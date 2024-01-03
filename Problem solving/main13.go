package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 1
	a, b := 1, 1
	for i := 1; a != n; i++ {
		a, b = b, a+b
		sum += 1
		if b > n {
			break
		}
	}
	if a == n {
		fmt.Println(sum)
	} else {
		fmt.Println(-1)
	}

}
