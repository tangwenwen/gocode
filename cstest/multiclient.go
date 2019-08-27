package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var (
	wg               sync.WaitGroup = sync.WaitGroup{}
	sendCount        int            = 50
	failSendCount    int            = 0
	successSendcount int            = 0
	sendTime         int            = 60
)

func main() {
	startTime := time.Now()
	for i := 1; i <= sendTime; i++ {
		for j := 0; j < sendCount; j++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				conn, err := net.Dial("tcp", "192.168.1.114:3002")
				if err != nil {
					fmt.Fprintf(os.Stderr, "Fatal error2: %s\n", err.Error())
					failSendCount++
				} else {
					fmt.Println("connect success")
					sender1(conn)
				}
			}(&wg)
		}
		time.Sleep(1e9)
		fmt.Println(i)
	}
	wg.Wait()
	fmt.Println("successSendcount:", successSendcount, "failCount:", failSendCount, "failureRate:", float64(failSendCount)/float64(sendCount*sendTime)*100, "%  ////", time.Since(startTime))

}
func sender1(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 49)
	data := []byte{15, 49, 0, 0, 0, 1, 0, 0,
		0, 9, 40, 0, 0, 0, 48, 59, 49, 48, 54,
		52, 55, 54, 51, 50, 50, 50, 57, 49, 50,
		58, 49, 50, 49, 46, 49, 56, 48, 53, 52,
		44, 51, 49, 46, 49, 53, 51, 51, 56, 51}
	copy(buf, data)
	_, err := conn.Write(buf)
	if err != nil {
		fmt.Println(buf)
		return
	}
	successSendcount++
	fmt.Println("send over")
}
