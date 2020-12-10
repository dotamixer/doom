package pprof

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime/pprof"
	"time"
)

type Handler struct {
	opts *Options
	ctx context.Context
	cancel context.CancelFunc
}

func NewHandler() *Handler {

	h := &Handler{}

	return h
}

func (h *Handler) Start(opts *Options) {
	h.opts = opts

	logrus.Infof("start pprof... opt:[%+v]", *opts)

	h.ctx, h.cancel = context.WithCancel(context.TODO())

	go h.cpuProfile()
	go h.heapProfile()

	go func() {
		ticker := time.NewTicker(opts.Frequency)

		for {
			select {
			case <-h.ctx.Done():
				return
			case <-ticker.C:
				go h.cpuProfile()
				go h.heapProfile()
			}
		}
	}()
}

func (h *Handler) Stop() {
	if h.cancel != nil {
		h.cancel()
		h.cancel = nil
	}
}

//取基础文件名
func getFileName() (filename string) {
	now := time.Now()
	filename = now.Format("2006_01_02_15_04")
	return
}


// 生成 CPU 报告
func (h *Handler) cpuProfile() {

	pprof.StopCPUProfile()

	filename := fmt.Sprintf("cpu_%s.prof", getFileName())
	fp := filepath.Join(h.opts.Path, filename)

	f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	err = pprof.StartCPUProfile(f)
	if err != nil {
		return
	}
	defer pprof.StopCPUProfile()

	time.Sleep(h.opts.Frequency)

}

// 生成堆内存报告
func (h *Handler) heapProfile() {
	filename := fmt.Sprintf("heap_%s.prof", getFileName())
	fp := filepath.Join(h.opts.Path, filename)

	f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	pprof.WriteHeapProfile(f)
}

