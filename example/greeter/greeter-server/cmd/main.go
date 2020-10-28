package main

import (
	"github.com/dotamixer/doom/bootstrap"
	"github.com/dotamixer/doom/example/greeter/greeter-server/api/rpc"
	"github.com/dotamixer/doom/example/greeter/greeter-server/idl/com/app/greeter/v1"
	"github.com/dotamixer/doom/pkg/di"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

func initDI(c *dig.Container) {
	di.MustContainerProvide(c, rpc.NewHandler)
}

func main() {

	srv := bootstrap.NewServer(
		bootstrap.WithServiceName("greeter"))

	// 初始化业务侧的DI
	initDI(srv.Container)

	fn := func(grpcSrv *grpc.Server) {
		di.MustContainerInvoke(srv.Container, func(h greeter.GreeterServiceServer) {
			greeter.RegisterGreeterServiceServer(grpcSrv, h)
		})
	}

	srv.Init(bootstrap.WithRegisterGrpcCallback(fn))

	srv.Run()
}
