package main

import "fmt"

const (
	USD int = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Print(RMB, symbol[RMB])
	array := [4][2]int{}
	fmt.Println(array[:2])
	a := make([]int, 0, 4)
	a = append([]int{1, 2}, a...)
	a = append(append(a[:1], []int{3}...), a[1:]...)
	fmt.Println(a)

}
