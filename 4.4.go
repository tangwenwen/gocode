package main

import "fmt"

func main() {
	// for index,v :=range "asdasd"{
	// 	fmt.Printf("%d,%d\t",index,v)
	// }

	c := make(chan int)
	go func() {
		c <- 2
		c <- 3
		c <- 1
	}()
	go func() {
		c <- 5
		c <- 6
		c <- 7
	}()
	for v := range c {
		fmt.Println(v)
	}

}
