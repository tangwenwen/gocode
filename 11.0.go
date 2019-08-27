package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	service := "192.168.1.114:7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError1(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError1(err)
	var parm string
	fmt.Scanf("%s", &parm)
	_, err = conn.Write([]byte(parm))
	if err == nil {
		resv := make([]byte, 128)
		_, err = conn.Read([]byte(resv))
		fmt.Println(string(resv))
	}
	os.Exit(0)
}
func checkError1(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
