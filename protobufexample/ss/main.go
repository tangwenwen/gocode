package main

import (
	"demo/protobufexample"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	p1 := example.Person{
		Id:   *proto.Int32(1),
		Name: *proto.String("tww1"),
	}
	p2 := example.Person{
		Id:   2,
		Name: "tww2",
	}
	all_p := example.AllPerson{
		Per: []*example.Person{&p1, &p2},
	}
	// 序列化
	data, err := proto.Marshal(&all_p)
	if err != nil {
		log.Fatal("mashal data error:", err)
	}
	fmt.Println(data)
	// 反序列化
	var target example.AllPerson
	err = proto.Unmarshal(data, &target)
	fmt.Print(target.Per[0].Name)
}
