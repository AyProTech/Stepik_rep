package main

import "fmt"

func main() {

	var a uint
	fmt.Scan(&a)
	a = a % 100
	a = a / 10
	fmt.Println(a)
}
