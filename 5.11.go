package main

import (
	"fmt"
	"time"
)

var filbs [41]int

func main() {
	start := time.Now()
	result := 0
	for i := 0; i < 41; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	endtime := time.Now()
	du := endtime.Sub(start)
	fmt.Println(du)
}
func fibonacci(n int) (res int) {
	fmt.Println(filbs)
	if filbs[n] != 0 {
		res = filbs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	filbs[n] = res
	return
}
