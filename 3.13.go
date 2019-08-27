package main

import (
	"container/list"
	"fmt"
)

func main() {
	asd := list.New()
	asd.PushBack("1111")
	asd.PushFront("2222")
	element := asd.PushBack("fff")
	for i := asd.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value)
	}
	asd.InsertAfter("high", element)
	asd.InsertBefore("noon", element)
	asd.Remove(element)
	for i := asd.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value)
	}

}
