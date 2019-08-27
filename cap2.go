// package main
//
// import "fmt"
//
// type User struct {
// 	id   int
// 	name string
// }
//
// func (user User) print() {
// 	fmt.Println("id:", user.id, "name:", user.name)
// }
//
// func main() {
// 	u := User{
// 		id:   100,
// 		name: "tww",
// 	}
// 	aaa := u.print
// 	u.id, u.name = 101, "qqq"
// 	u.print()
// 	aaa()
//
// }

// package main
//
// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )
//
// func Newtest() chan int {
// 	c := make(chan int)
// 	rand.Seed(time.Now().UnixNano())
// 	go func() {
// 		time.Sleep(time.Second)
// 		c <- rand.Int()
// 	}()
// 	return c
// }
// func main() {
// 	tack := Newtest()
// 	fmt.Println( <-tack)
//
// }

// package main
//
// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )
//
// func main() {
// 	var wg sync.WaitGroup
// 	runtime.GOMAXPROCS(4)
// 	wg.Add(3)
// 	c := make(chan int,1)
// 	for i := 0; i < 3; i++ {
// 		go func(id int) {
// 			defer wg.Done()
// 			c <- 1
// 			for j := 0; j < 3; j++ {
// 				fmt.Println(id, j)
// 			}
//
// 		}(i)<-c
// 	}
// 	wg.Wait()
// }

package main

/*
#include <stdio.h>
#include <stdlib.h>
void hello() {
printf("Hello, World!\n");
}
*/
import "C"

func main() {
	C.hello()
}
