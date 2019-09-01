# go-mirco
---
### Tool installation

（1）First install the protocol tool
    
```
  go get github.com/micro/protoc-gen-micro
```

 (2) Two tools that must be installed
     
> * [proto](https://github.com/google/protobuf)
   
> * [protoc-gen-go](https://github.com/golang/protobuf)

### achieve

We then create a mirco-me folder under $GOPATH/src/demo to create the server and client and protos folders. Create server_main.go in the server folder and create client_main.go and protos folders in the client folder. Create a hello.proto file

#### add the following to the protos/hello.proto file

```
   syntax = "proto3";

   service Greeter {
      rpc Hello(HelloRequest) returns (HelloResponse) {}
   }

   message HelloRequest {
      string name = 1;
   }

   message HelloResponse {
      string greeting = 2;
   }

```
#### add the following to /server/server_main.go

```
   package main
import (
"context"
"fmt"

micro "github.com/micro/go-micro"
proto "../proto"

)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
rsp.Greeting = "Hello " + req.Name
return nil
}

func main() {
// Create a new service. Optionally include some options here.
service := micro.NewService(
micro.Name("greeter"),
)

// Init will parse the command line flags.
service.Init()

// Register handler
proto.RegisterGreeterHandler(service.Server(), new(Greeter))

// Run the server
if err := service.Run(); err != nil {
    fmt.Println(err)
}


}

```

#### add the following to client/client_main.go

```
    package main

import (
"context"
"fmt"

micro "github.com/micro/go-micro"
proto "../proto"


)

func main() {
// Create a new service. Optionally include some options here.
service := micro.NewService(micro.Name("greeter.client"))
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

```

Open the terminal and switch to the protos folder and execute the following command.

```
  protoc --proto_path=. --micro_out=. --go_out=. hello.proto
```

Two files hello.pb.go and hello.micro.go will be created in this folder.

Run the server_main.go file under the server. The terminal displays the following:

```
   2019/08/31 05:22:35 Transport [http] Listening on [::]:65170
   2019/08/31 05:22:35 Broker [http] Connected to [::]:65171
   2019/08/31 05:22:35 Registry [mdns] Registering node: greeter-4c82f3f5-d3d6-4753-891f-9b34969a6161
```

Run the client client_main.go file terminal display content is as follows:

```
   Hello John

   Process finished with exit code 0
```
