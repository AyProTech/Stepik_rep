package main

import "fmt"

func main() {
	var n, b int
	sum := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		if b == 0 {
			sum++
		}
	}
	fmt.Println(sum)

}
