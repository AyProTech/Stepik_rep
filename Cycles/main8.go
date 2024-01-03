package main

import "fmt"

func main() {
	var a, b, c, d, m int
	fmt.Scan(&a, &b)
	a1 := a
	a = 0
	for a1 != 0 {
		a = a*10 + a1%10
		a1 = a1 / 10
	}

	for a != 0 {
		c = a % 10
		m = b
		for m != 0 {
			d = m % 10
			if c == d {
				fmt.Print(d, " ")
			}
			m = m / 10
		}
		a = a / 10
	}
}
