package main

import (
	"demo/mygRPC"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	HOST string = "localhost"
	PORT string = "1111"
)

type Greeter struct{}

func (G *Greeter) SayHi(ctx context.Context, in *DataPro.SendMessage) (*DataPro.ReplyMessage, error) {
	if in.Name == "tww" {
		return &DataPro.ReplyMessage{Message: "hi"}, nil
	} else {
		return &DataPro.ReplyMessage{Message: "gun"}, nil
	}
}
func main() {
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("err", err)
	}
	service := grpc.NewServer()
	DataPro.RegisterGreeterServer(service, &Greeter{})
	reflection.Register(service)
	err = service.Serve(listener)
	if err != nil {
		fmt.Println("err", err)
	}

}
