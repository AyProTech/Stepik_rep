package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	sum := 0
	for a <= b {
		sum = sum + a
		a++
	}
	fmt.Println(sum)
}
