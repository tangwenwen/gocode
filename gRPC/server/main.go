package main

import (
	"context"
	"demo/gRPC/DataProto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strings"
)

const (
	HOST string = "localhost"
	PORT string = "1111"
)

type FormData struct {
}

func (fd *FormData) DoFormat(ctx context.Context, in *DataProto.Data) (out *DataProto.Data, err error) {
	str := in.Text
	out = &DataProto.Data{Text: strings.ToUpper(str)}
	return out, nil
}
func main() {
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("faile listen at: " + HOST + ":" + PORT)
	} else {
		fmt.Println("Demo server is listening at: " + HOST + ":" + PORT)
	}
	rpcServer := grpc.NewServer()
	DataProto.RegisterFormDataServer(rpcServer, &FormData{})
	reflection.Register(rpcServer)
	if err = rpcServer.Serve(listener); err != nil {
		fmt.Println("faile serve at: " + HOST + ":" + PORT)
	}

}
