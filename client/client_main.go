package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	proto "../proto"
	"github.com/micro/go-micro"
)


func main() {
	opition:=func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:2379"}
	}
	etcdRegister:= etcdv3.NewRegistry(opition)

		// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter.client"),
		micro.Registry(etcdRegister),
		)
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}