package rpc

import (
	"context"
	"fmt"
	"github.com/dotamixer/doom/example/greeter/greeter-server/idl/com/app/greeter/v1"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	//app app.Service
}

func (h *Handler) SayHello(context context.Context, req *greeter.SayHelloReq) (*greeter.SayHelloRsp, error) {
	rsp := &greeter.SayHelloRsp{
		Msg: fmt.Sprintf("Hello, %s", req.Name),
	}
	logrus.Debug("Hello")
	return rsp, nil
}

func NewHandler() greeter.GreeterServiceServer {
	return &Handler{}
}
