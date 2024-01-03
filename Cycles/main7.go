package main

import "fmt"

func main() {
	var a, b, c int
	sum := 0
	fmt.Scan(&a, &b, &c)
	for a < c {
		if a <= c {
			sum++
		}
		a = a + a*b/100
	}
	fmt.Println(sum)
}
