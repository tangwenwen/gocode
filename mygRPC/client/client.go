package main

import (
	"demo/mygRPC"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1111", grpc.WithInsecure())
	if err != nil {
		fmt.Println("err", err)
	}
	defer conn.Close()
	t := DataPro.NewGreeterClient(conn)
	res := os.Args[1]

	rep, err := t.SayHi(context.Background(), &DataPro.SendMessage{Name: res})
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(rep.Message)
}
