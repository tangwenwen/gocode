// package main
//
// import "fmt"
//
// type bag struct {
// 	item []int
// }
// func (b *bag) Insert (id int) {
// 	b.item=append(b.item, id)
// }
// func main()  {
// 	b :=new(bag)
// 	b.Insert(1024)
// 	fmt.Println(b.item)
// }

package main

import "fmt"

type pro struct {
	value int
}

func (s *pro) setv(v int) {
	s.value = v
}

func (g *pro) getv() int {
	return g.value
}
func main() {
	pppp := new(pro)
	pppp.setv(111)
	fmt.Println(pppp.getv())
}
