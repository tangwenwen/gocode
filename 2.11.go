//package main
//
//import (
//	"fmt"
//)
//func swap( a, b *int){
//	fmt.Println(*a,*b)
//	fmt.Println(a,b)
//	var t = *a
//	*a = *b
//	*b =t
//	fmt.Println(*a,*b)
//	fmt.Println(a,b)
//
//}
//func main(){
//	var s , t = 1 , 2
//	fmt.Printf("before:%p,%p\n",&s,&t)
//	swap(&s,&t)
//	fmt.Printf("after:%p,%p\n",&s,&t)
//	fmt.Println(s , t)
//}
package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "process mode")

func main() {
	flag.Parse()
	fmt.Print(*mode)
}
