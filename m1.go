package main

import (
	"fmt"
	"sync"
)

func printNums(wg1 *sync.WaitGroup, ch1 chan int, ch2 chan int) {
	defer wg1.Done()
	for i := 0; i <= 10; i++ {
		fmt.Println("111")
		fmt.Printf("%d", <-ch1)
		fmt.Println("222")
		ch2 <- i
	}
}

func printchars(wg1 *sync.WaitGroup, ch1 chan int, ch2 chan int) {
	defer wg1.Done()
	for i := 0; i <= 10; i++ {
		fmt.Println("333")
		ch1 <- i
		fmt.Println("444")
		fmt.Printf("%c", 'a'+<-ch2)
	}
}

func main() {
	var wg1 sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg1.Add(2)
	go printNums(&wg1, ch1, ch2)
	go printchars(&wg1, ch1, ch2)
	wg1.Wait()
}

// 输出：0a1b2c3d4e5f6g7h8i9j10k
// printNums 和printchars交替输出
