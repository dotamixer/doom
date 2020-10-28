package main

import (
	"fmt"
	"github.com/dotamixer/doom/tool/doom-protoc/internal/app"
	"github.com/dotamixer/doom/tool/doom-protoc/internal/flags"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	App   *app.App
	major = 2
	minor = 1
	patch = 0
)

var (
	RootCmd = &cobra.Command{
		Use:  "doom-protoc",
		Args: cobra.NoArgs,
	}
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "生成配置文件",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			App.Config()
		},
	}
)

var (
	fmtCmd = &cobra.Command{
		Use:   "fmt",
		Short: "格式化proto",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !cmd.HasAvailableFlags() {
				return errors.New("need flag")
			}
			App.Format()
			return nil
		},
	}
)

func init() {
	fmtCmd.PersistentFlags().StringSliceVarP(&flags.SrcFiles, "files", "f",
		[]string{}, "相对于导入目录下的源文件")
	fmtCmd.PersistentFlags().StringSliceVarP(&flags.SrcDirectories, "directories", "d",
		[]string{}, "相对于导入目录下源文件目录")
}

var (
	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "编译proto",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			App.Gen()
		},
	}
)

func init() {
	generateCmd.PersistentFlags().StringSliceVarP(&flags.SrcFiles, "files", "f",
		[]string{}, "相对于导入目录下的源文件")
	generateCmd.PersistentFlags().StringSliceVarP(&flags.SrcDirectories, "directories", "d",
		[]string{}, "相对于导入目录下源文件目录")
}

func init() {
	App, _ = app.InitApp()
}

func Run() {
	RootCmd.AddCommand(generateCmd)
	RootCmd.AddCommand(fmtCmd)
	RootCmd.AddCommand(configCmd)

	RootCmd.Version = fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	_ = RootCmd.Execute()
}

func main() {
	Run()
}
