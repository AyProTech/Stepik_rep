package main

import "fmt"

func main() {
	var n, b, min int
	ls := []int{}
	sum := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		ls = append(ls, b)
	}
	min = ls[0]
	for i := 0; i < n; i++ {

		if ls[i] < min {
			min = ls[i]
		}
	}
	for i := 0; i < n; i++ {
		if ls[i] == min {
			sum++
		}
	}
	fmt.Println(sum)
}
