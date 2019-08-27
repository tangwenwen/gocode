package main

import (
	"context"
	"demo/gRPC/DataProto"
	"fmt"
	"google.golang.org/grpc"
)

const ADDRESS string = "localhost:1111"

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		fmt.Println("cannot connect:" + ADDRESS)
	}
	defer conn.Close()
	client := DataProto.NewFormDataClient(conn)
	resq, err := client.DoFormat(context.Background(), &DataProto.Data{Text: "qqqqqqqqqqqqq"})
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(resq.Text)
}
