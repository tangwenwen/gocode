// package main
//
// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )
//
// var wg sync.WaitGroup
//
// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }
// func main() {
// 	court := make(chan int)
// 	wg.Add(2)
// 	go player("aaa", court)
// 	go player("bbb", court)
// 	court <- 1
// 	wg.Wait()
// }
// func player(name string, court chan int) {
// 	defer wg.Done()
// 	for {
// 		ball, ok := <-court
// 		if !ok {
// 			fmt.Printf("Player %s Won\n", name)
// 			return
// 		}
// 		n := rand.Intn(100)
// 		if n%2 == 0 {
// 			fmt.Println(n)
// 			fmt.Printf("Player %s Missed\n", name)
// 			close(court)
// 			return
// 		}
// 		fmt.Printf("Player %s Hit %d\n", name, ball)
// 		ball++
// 		court <- ball
// 	}
// }
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)
	go runner(baton)
	baton <- 1

	wg.Wait()
}
func runner(baton chan int) {
	var newrunner int
	runno := <-baton
	fmt.Println("running no:", runno)
	if runno != 4 {
		newrunner = runno + 1
		go runner(baton)
	}
	if runno == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runno)
		wg.Done()
		return
	}
	time.Sleep(time.Second)
	baton <- newrunner

}
