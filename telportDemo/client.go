package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	wg               sync.WaitGroup = sync.WaitGroup{}
	sendCount        int            = 10000
	failSendCount    int            = 0
	successSendcount int            = 0
	sendTime         int            = 1
)

func main() {
	startTime := time.Now()
	for j := 0; j < sendCount; j++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			conn, err := net.Dial("tcp", "localhost:8088")
			if err != nil {
				failSendCount++
			} else {
				// fmt.Println("connect success")
				sender1(conn)
			}
		}(&wg)
	}
	wg.Wait()
	fmt.Println("successSendcount:", successSendcount, "failCount:", failSendCount, "failureRate:", float64(failSendCount)/float64(sendCount*sendTime)*100, "%  ////", time.Since(startTime))

}
func sender1(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("hi"))
	replay := make([]byte, 20)
	conn.Read(replay)
	// fmt.Println(string(replay))
	successSendcount++

}
