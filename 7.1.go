package main

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
	CanWrite() bool
}
type file struct {
}

func (D *file) WriteData(data interface{}) error {
	fmt.Println("writerdata:", data)
	return nil
}
func (D *file) CanWrite() bool {
	return true
}

func main() {
	f := new(file)
	var writer DataWriter
	writer = f
	writer.WriteData("data")
	fmt.Println(writer.CanWrite())
}
