package main

import "fmt"

func main() {
	var n, m int
	sum := 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&m)
		if m > 9 && m < 100 && m%8 == 0 {
			sum = sum + m
		}
	}
	fmt.Println(sum)
}
