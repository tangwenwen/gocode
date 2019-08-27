package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	i := 0
	for input.Scan() {
		i++
		if i > 5 {
			break
		}
		counts[input.Text()]++

	}
	for k, v := range counts {
		fmt.Printf("%s\t%d\t", k, v)
	}
}
