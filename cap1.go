// package main
//
// import "fmt"
//
// const (
// 	_        = iota
// 	KB int64 = 1 << (10 * iota)
// 	MB
// 	GB
// 	TB
// )
//
// func main() {
// 	s := "asd"
// 	fmt.Println(&s)
// 	s, y := "hello", 20
// 	fmt.Println(&s, y, s)
// 	fmt.Println(KB)
// }

package main

import "fmt"

func main() {
	str := "asd阿斯顿"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	for _, r := range str {
		fmt.Printf("%c", r)
	}
}
