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
	fmt.Println(a[3])
	// здесь должен быть ваш код
}
