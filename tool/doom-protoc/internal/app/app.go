package app

import (
	"github.com/dotamixer/doom/tool/doom-protoc/internal/compile"
	"github.com/dotamixer/doom/tool/doom-protoc/internal/config"
	"github.com/dotamixer/doom/tool/doom-protoc/internal/flags"
	"github.com/dotamixer/doom/tool/doom-protoc/internal/format"
	"github.com/dotamixer/doom/tool/doom-protoc/internal/proto"
	"log"
	"path/filepath"
	"strings"
)

type App struct {
	config    *config.Config
	compiler  *compile.Compiler
	formatter *format.Formatter
}

func NewApp(config *config.Config, compiler *compile.Compiler, formatter *format.Formatter) *App {
	return &App{
		config:    config,
		compiler:  compiler,
		formatter: formatter,
	}
}

func (a *App) Format() {
	var (
		err error

		absFiles []string
	)
	err = a.config.Load()
	if err != nil {
		log.Fatal(err)
	}

	absFiles = a.specialFile()
	if len(absFiles) > 0 {
		a.formatter.Format(absFiles)
	}
}

func (a *App) Gen() {
	var (
		err             error
		deleteDirectory bool
	)
	err = a.config.Load()
	if err != nil {
		log.Fatal(err)
	}
	protos := a.specialFile()
	//不是编译指定文件时，对编译output 目录进行删除后重建
	if len(protos) == 0 {
		protos = a.config.Protos
		deleteDirectory = true
	}
	descSource, err := proto.DescriptorSourceFromProtoFiles(a.config.Includes, protos...)
	if err != nil {
		log.Fatalf("Failed to process proto source files. %v", err)
	}

	err = a.compiler.Compile(descSource, deleteDirectory)
	if err != nil {
		log.Fatalf("compile error %v", err)
	}

	return
}

func (a *App) Config() {
	err := a.config.Output()
	log.Fatal(err)
	return
}

func (a *App) specialFile() []string {
	var (
		absFiles []string
	)
	//文件参数
	sourceFiles := map[string]struct{}{}
	for _, itr := range flags.SrcFiles {
		sourceFiles[itr] = struct{}{}
	}
	for _, itr := range a.config.Protos {
		if _, ok := sourceFiles[itr]; ok {
			absPath := filepath.Join(a.config.ImportPath, itr)
			absPath = filepath.ToSlash(absPath)
			absFiles = append(absFiles, absPath)
		}
	}
	//目录参数
	for _, itr := range a.config.Protos {
		for _, dir := range flags.SrcDirectories {
			if strings.Contains(itr, dir) {
				absPath := filepath.Join(a.config.ImportPath, itr)
				absPath = filepath.ToSlash(absPath)
				absFiles = append(absFiles, absPath)
			}
		}
	}

	return absFiles
}
