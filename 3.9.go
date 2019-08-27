package main

import "fmt"

func main() {
	var mapList map[string]int
	mapList = map[string]int{"k1": 1, "k2": 2}

	map2 := map[string]int{"k1": 3, "k2": 4}
	mapList = map2
	fmt.Print(mapList["k1"])
	fmt.Print(mapList["k2"])
	var scen []int
	delete(map2, "k2")
	for _, v := range map2 {
		scen = append(scen, v)
	}
	fmt.Println(scen)

}
