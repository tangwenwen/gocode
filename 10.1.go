package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x).Elem()
	v.SetFloat(2.33)
	fmt.Println(x)
}
