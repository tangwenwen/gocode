package main

import (
	"fmt"
	"github.com/pkufergus/goroutinepool"
	"time"
)

func main() {

	p := goroutinepool.NewRoutinePool(10)
	for i := 0; i < 100; i++ {
		count := i
		p.AddJob(func() {
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("%d\n", count)
		})

	}
	p.WaitAll()
}
