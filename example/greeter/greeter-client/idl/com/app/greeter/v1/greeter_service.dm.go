package greeter

import (
	"google.golang.org/grpc"
)

var (
	ServiceName = "com.app.greeter.v1"
)

func NewClient() (client GreeterServiceClient, err error) {
	var (
		conn *grpc.ClientConn
	)
	conn, err = grpc.Dial("consul://default/com.app.greeter.v1",
		grpc.WithInsecure())
	if err != nil {
		return
	}

	client = NewGreeterServiceClient(conn)

	return
}
