package main

import "fmt"

func main() {
	var n, b int
	fmt.Scan(&n)
	a := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&b)

		a = append(a, b)
	}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a[i], " ")
		}
	}

	// здесь должен быть вашкод
}
