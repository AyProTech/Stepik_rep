package main

import "fmt"

func main() {

	var a, b, c uint
	fmt.Scan(&a)
	b = a / 30
	c = (a % 30) * 2

	fmt.Println("It is", b, "hours", c, "minutes.")
}
