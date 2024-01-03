package main

import "fmt"

func main() {
	var n int
	a := 1
	fmt.Scan(&n)
	for i := 0; a <= n; i++ {
		fmt.Print(a, " ")
		a = a * 2
	}

}
