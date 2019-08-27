// 接口的类型转换

package main

import "fmt"

type FIyer interface {
	Fly()
}
type Walker interface {
	Walk()
}
type bird struct {
}

func (this *bird) Fly() {
	fmt.Println("bird:fly")
}
func (this *bird) Walk() {
	fmt.Println("bird:walk")
}

type pig struct {
}

func (this *pig) Walk() {
	fmt.Println("pig:walk")
}

func main() {
	// animals := map[string]interface{}{
	// 	"bird": new(bird),
	// 	"pig":  new(pig),
	// }
	// for name, obj := range animals {
	// 	f, isFlayer := obj.(FIyer)
	// 	w, isWalk := obj.(Walker)
	// 	fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlayer, isWalk)
	// 	if isFlayer {
	// 		f.Fly()
	// 	}
	// 	if isWalk {
	// 		w.Walk()
	// 	}
	//
	// }
	bb := new(pig)
	bb.Walk()
	bb.Fly()
}
