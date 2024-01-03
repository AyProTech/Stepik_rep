package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	for i, j := range a {
		if i%2 != 0 {
			fmt.Print(string(j))
		}

	}

}