package main

import "fmt"

func main() {
	var n, b int
	sum := 0
	fmt.Scan(&n)
	a := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&b)

		a = append(a, b)
	}
	for i := 0; i < n; i++ {
		if a[i] > 0 {
			sum++
		}
	}
	fmt.Println(sum)
	// здесь должен быть вашкод
}
