package main

import "fmt"

func main() {
	// 编译器将把123的类型推断为内置类型int。
	var x interface{} = 123
	// 情形一：
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true
	n = x.(int)
	fmt.Println(n) // 123
	// 情形二：
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false
	// 情形三：
	a = x.(float64) // 将产生一个恐慌
	fmt.Println(a)
}
