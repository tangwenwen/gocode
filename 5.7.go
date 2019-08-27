// package main
//
// import (
// 	"fmt"
// )
//
// // 遍历切片的每个元素, 通过给定函数进行元素访问
// func visit(list []int, f func(int)) {
//
// 	for _, v := range list {
// 		f(v)
// 	}
// }
//
// func main() {
//
// 	// 使用匿名函数打印切片内容
// 	visit([]int{1, 2, 3, 4}, func(v int) {
// 		if v%2 == 0 {
// 			fmt.Println(v)
// 		}
// 	})
// }
package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "", "skill to perform")

func main() {

	flag.Parse()

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok1 := skill[*skillParam]; ok1 {
		fmt.Println(ok1)
		f()
	} else {
		fmt.Println("skill not found")
	}

}
