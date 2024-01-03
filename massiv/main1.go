package main

import "fmt"

func main() {
	var workArray [10]uint8
	var a, b int
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}
	for i := 3; i > 0; i-- {
		fmt.Scan(&a, &b)
		workArray[a], workArray[b] = workArray[b], workArray[a]
	}
	for _, i := range workArray {
		fmt.Print(i, " ")
	}
}
