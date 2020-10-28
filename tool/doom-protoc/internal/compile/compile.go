package compile

import (
	"fmt"
	"github.com/dotamixer/doom/tool/doom-protoc/internal/config"
	"github.com/dotamixer/doom/tool/doom-protoc/internal/proto"
	"github.com/dotamixer/doom/tool/doom-protoc/internal/wkt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Compiler struct {
	config *config.Config
}

type metaCmd struct {
	file           string
	args           []string
	extraModifiers string
}

func NewCompiler(cfg *config.Config) *Compiler {
	return &Compiler{
		config: cfg,
	}
}

func (c *Compiler) Compile(desc proto.DescriptorSource, deleteDirectory bool) (err error) {

	for _, itr := range c.config.Generate.Plugins {
		metaCmds := c.generateCmd(desc, itr.Type)
		//插件
		arg := fmt.Sprintf("--%s_out=", itr.Name)
		outputPath := os.ExpandEnv(c.config.Generate.Output)
		if deleteDirectory {
			err = os.RemoveAll(outputPath)
			if err != nil {
				log.Fatalf("remove path failed ", outputPath)
			}
		}

		err = os.MkdirAll(outputPath, os.ModePerm)
		if err != nil {
			log.Fatalf("MkdirAll failed [%v]", err)
		}
		log.Println("out path ", outputPath)
		//参数
		arg = arg + itr.Flags
		var optArg string
		for _, cmd := range metaCmds {
			optArg = ""
			if len(cmd.extraModifiers) > 0 {
				optArg = arg + "," + cmd.extraModifiers + ":" + outputPath
			} else {
				optArg = arg + ":" + outputPath
			}
			//输出路径
			tmp := append(cmd.args, optArg)
			tmp = append(tmp, cmd.file)
			tmpCmd := exec.Command("protoc", tmp...)
			log.Println(tmpCmd.String())
			out, err := tmpCmd.CombinedOutput()
			if err != nil {
				log.Println("compile Error:", err)
				log.Println(string(out))
			}
		}
	}
	return nil
}

func (e *Compiler) generateCmd(desc proto.DescriptorSource, typ string) []*metaCmd {
	var (
		M   string
		ret []*metaCmd
	)
	for name, fileDesc := range desc.Files() {
		M = ""
		if _, ok := wkt.Filenames[name]; ok {
			continue
		}
		log.Println("compile file ", name)
		//生成命令
		fs := fileDesc.GetDependencies()
		var ms []string
		for _, fd := range fs {
			dependName := fd.GetName()
			opt := fd.GetFileOptions().GoPackage
			if opt == nil {
				str := fd.GetPackage()
				opt = &str
			}
			//处理指定包名的情况
			tmps := strings.Split(*opt, ";")

			m := fmt.Sprintf("M%s=%s/%s/%s", dependName, e.config.Module, e.config.Generate.Modifier, tmps[0])
			ms = append(ms, m)

			if typ == "go" {
				if m, ok := wkt.FilenameToGoModifierMap[dependName]; ok {
					ms = append(ms, "M"+dependName+"="+m)
				}
			} else {
				if m, ok := wkt.FilenameToGogoModifierMap[dependName]; ok {
					ms = append(ms, "M"+dependName+"="+m)
				}
			}

		}
		if len(ms) > 0 {
			M = strings.Join(ms, ",")
		}

		var args []string
		//-I 包含路径
		for _, itr := range e.config.Includes {
			args = append(args, "-I"+itr)
		}
		cmd := metaCmd{
			file:           filepath.Join(e.config.ImportPath, name),
			args:           args,
			extraModifiers: M,
		}
		ret = append(ret, &cmd)
	}
	return ret
}
