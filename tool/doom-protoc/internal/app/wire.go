//+build wireinject

package app

import (
	"github.com/dotamixer/doom/tool/doom-protoc/internal/compile"
	"github.com/dotamixer/doom/tool/doom-protoc/internal/config"
	"github.com/dotamixer/doom/tool/doom-protoc/internal/format"
	"github.com/google/wire"
)

func InitApp() (*App, error) {
	panic(wire.Build(config.NewConfig, compile.NewCompiler, format.NewFormatter, NewApp))
}
