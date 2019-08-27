// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net"
// 	"os"
// )
// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
// 		os.Exit(1)
// 	}
// 	service := os.Args[1]
// 	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
// 	checkError(err)
// 	conn, err := net.DialTCP("tcp", nil, tcpAddr)
// 	checkError(err)
// 	var n int
// 	n, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
// 	fmt.Println(n)
// 	checkError(err)
// 	result, err := ioutil.ReadAll(conn)
// 	checkError(err)
// 	fmt.Println(string(result))
// 	os.Exit(0)
// }
// func checkError(err error) {
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
// 		os.Exit(1)
// 	}
// }

package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
func handleClient(conn net.Conn) {
	// defer 	conn.Close()
	// starttime := time.Now()
	// time.Sleep(5*time.Second)
	// duertiem :=time.Since(starttime)
	// duertiemstr := duertiem.String()
	// conn.Write([]byte(duertiemstr))
	conn.SetReadDeadline(time.Now().Add(10 * time.Second)) // 设置2分钟超时
	request := make([]byte, 128)                           // 退出前关闭连接
	for {
		read_len, err := conn.Read(request)

		if read_len == 0 {
			fmt.Println("客户端已关闭连接")
			break // 客户端已关闭连接
		} else if daytime := strconv.FormatInt(time.Now().Unix(), 10); strings.TrimSpace(string(request[:read_len])) == "timestrap" {
			fmt.Println("untime")
			conn.Write([]byte(daytime))
		} else {
			fmt.Println("time")
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
		fmt.Println("请求数据:", string(request))
		if err != nil {
			fmt.Println("error:", err)
			break
		}
		// request = make([]byte, 128) // 清除上次读取的内容
	}
}
