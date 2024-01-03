package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	for a != 0 {
		b = a % 10
		a = a / 10
	}
	fmt.Println(b)
}
