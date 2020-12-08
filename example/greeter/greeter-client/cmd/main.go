package main

import (
	"context"
	"github.com/dotamixer/doom/example/greeter/greeter-client/idl/com/app/greeter/v1"
	"github.com/dotamixer/doom/pkg/resolver/consul"
	"log"
)

func main() {
	consul.SetDefaultHost("192.168.203.40:8500")

	client, err := greeter.NewGreeterServiceClientByDoom()
	if err != nil || client == nil {
		log.Println("Failed to new client. ", err)
		return
	}
	rsp, err := client.SayHello(context.TODO(), &greeter.SayHelloReq{
		Name: "EvanPan",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(rsp.Msg)
}
