package main

import (
	"bytes"
	"demo/cstest/conf"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type sendObj struct {
	Code int
	Data sendData
}

type sendData struct {
	User_name  string  `json:"user_name"`  //用户名
	Lng        float64 `json:"lng"`        //经度
	Lat        float64 `json:"lat"`        //纬度
	Time       int64   `json:"time"`       //时间
	Access_key string  `json:"access_key"` //登录key
	Gps_type   string  `json:"gps_type"`   //gps类型
}

var (
	logfile, _ = os.OpenFile("C:/Users/123/Desktop/untitled1/src/demo/tcpServerDemo/log/log.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	logger     = log.New(logfile, "", log.Llongfile)
)

func main() {
	runtime.GOMAXPROCS(4)
	logger.SetFlags(log.Ldate | log.Ltime)

	//建立socket的MQ连接
	amqpconn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "failed to connect to RabbitMQ")
	defer amqpconn.Close()

	//建立管道
	ch, err := amqpconn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()
	//建立队列来存储/转发信息
	q, err := ch.QueueDeclare(
		"info_queue",
		true,
		false,
		false,
		false,
		nil)
	failOnError(err, "failed to declare q queue")

	add, err := net.ResolveTCPAddr("tcp4", conf.GetConfig().ServerConf.SADDRESS)
	checkError(err)

	listenner, err := net.ListenTCP("tcp", add)
	checkError(err)
	logger.Println("starting listening...")
	var wg sync.WaitGroup
	wg.Add(1)
	go acceptClient(listenner, q, ch)
	wg.Wait()

}
func acceptClient(listenner net.Listener, q amqp.Queue, ch *amqp.Channel) {
	for {
		defer listenner.Close()
		conn, err := listenner.Accept()
		if err != nil {
			logger.Println("err3", err)
		}
		tcp := conn.(*net.TCPConn)
		go haddleConnection(tcp, q, ch)
	}

}
func haddleConnection(conn *net.TCPConn, q amqp.Queue, ch *amqp.Channel) {

	defer conn.Close()
	readChan := make(chan []byte) //读数据管道
	sendChan := make(chan []byte) //发数据管道
	stopChan := make(chan bool)   //是否停止管道

	go readConn(conn, readChan, stopChan)
	go sendByMQ(sendChan, stopChan, q, ch)
	// go sendByHttp(sendChan, stopChan)
	for {
		select {
		case readByte := <-readChan:
			sendChan <- serialize(readByte)
			// sendByMQ(sendChan, stopChan)
		case stop := <-stopChan:
			if stop {
				break

			}
		}
	}
}

func sendByMQ(sendChan <-chan []byte, stopChan chan<- bool, q amqp.Queue, ch *amqp.Channel) {
	err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         <-sendChan,
		})
	failOnError(err, "failed to publish a message")
	stopChan <- true
}

func readConn(conn net.Conn, readChan chan<- []byte, stopChan chan<- bool) {
	for {
		conn.SetReadDeadline(time.Now().Add(time.Microsecond * 100))
		data := make([]byte, 64)
		if n, err := conn.Read(data); err != nil || n == 0 {
			logger.Println("read error;", err, data)
			break
		}
		buf := make([]byte, 11)
		data1 := []byte{16, 0, 0, 0, 3, 0, 0, 0, 50, 48, 48}
		copy(buf, data1)
		_, err := conn.Write(buf)
		checkError(err)
		logger.Println(conn.RemoteAddr().String(), "Reply code sent!")
		fmt.Println(conn.RemoteAddr().String(), "Receive successful ,Disconnect")
		readChan <- data
	}
	stopChan <- true
}
func sendByHttp(sendChan <-chan []byte, stopChan chan<- bool) {
	data := <-sendChan
	url := "http://tangwenwen.top:7999/work_sheet"
	request, err := http.NewRequest("POST", url, bytes.NewReader(data))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "*/*")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		logger.Println(err.Error())
		return
	}
	logger.Println(url, "reply status code:", resp.Status)
	stopChan <- true
}

func serialize(data []byte) []byte {
	lngStartIndex := bytes.IndexRune(data, rune(':'))
	userNameStartIndex := bytes.IndexRune(data, rune(';'))
	latStartIndex := bytes.LastIndex(data, []byte(","))
	gps_type := string(data[userNameStartIndex-1 : userNameStartIndex])
	username := string(data[userNameStartIndex+1 : lngStartIndex])
	lng := string(data[lngStartIndex+1 : latStartIndex])
	lat := data[latStartIndex+1:]
	latEnd := bytes.IndexByte(lat, 0)
	lngFloat, err := strconv.ParseFloat(lng, 64)
	latFloat, err := strconv.ParseFloat(string(lat[:latEnd]), 64)
	instance := sendObj{
		Code: 10002,
		Data: sendData{
			User_name:  username,
			Lng:        lngFloat,
			Lat:        latFloat,
			Time:       time.Now().Unix(),
			Access_key: "431be92ca9dd41b49fac982fdee63eca",
			Gps_type:   gps_type,
		},
	}
	jsonByte, err := json.Marshal(instance)
	if err != nil {
		logger.Println("Marshal json error:", err)
	}
	return jsonByte
}

func checkError(err error) {
	if err != nil {
		logger.Println(os.Stderr, "Fatal error:", err.Error())
		os.Exit(1)
	}
}
func failOnError(err error, msg string) {
	if err != nil {
		logger.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}
