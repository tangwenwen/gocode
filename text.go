package main

import "fmt"

type Phoneer interface {
	call()
}

type apple struct {
	name string
}
type huawei struct {
	name string
}
type position struct {
	x float64
	y float64
}

func (pp apple) call() {
	fmt.Println("apple call", pp.name)
}
func (pp huawei) call() {
	fmt.Println("huawei call", pp.name)
}
func (this *position) call() {
	fmt.Println("x:", this.x, ";", "y:", this.y)
}

func main() {
	var p Phoneer
	appl := apple{"woshi apple"}
	p = appl
	p.call()
	huawe := huawei{"woshi huawei"}
	p = huawe
	p.call()
	pos := position{1.5, 2.5}
	p = &pos
	p.call()
}

// package main
//
// type nillinterfacer interface {
// 	option()
// }
//
// func main() {
// 	var o nillinterfacer
// 	o.option()
// }
