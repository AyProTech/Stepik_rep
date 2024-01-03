package main

import "fmt"

func main() {
	var v, a, b, c int
	fmt.Scan(&v)
	a = v % 10
	b = (v / 10) % 10
	c = v / 100
	if a != b && a != c && b != c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
