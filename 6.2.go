// package main
//
// import "fmt"
//
// type Sss struct {
// 	A int
// 	B string
// 	C *int
// }
//
// func newSss(a int, b string, c *int) *Sss {
// 	return &Sss{
// 		A: a,
// 		B: b,
// 		C: c,
// 	}
// }
//
// var s = newSss(1, "asdasd", &www)
//
// var www int = 2
//
// func main() {
// 	fmt.Printf("%s", s.B)
// }

package main

import "fmt"

type people struct {
	name  string
	child *people
}

func main() {
	relation := &people{
		name: "pp",
		child: &people{
			name: "dd",
			child: &people{
				name: "ww",
			},
		},
	}
	for {
		fmt.Println(relation.name)

		if relation.child == nil {
			break
		}
		relation = relation.child
	}
}
