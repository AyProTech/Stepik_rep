package main

import "fmt"

func main() {
	var a, b int
	x := 0
	y := 0
	fmt.Scan(&a)
	b = a % 1000
	a = a / 1000
	for i := 1; i <= 3; i++ {
		x = x + a%10
		y = y + b%10
		a = a / 10
		b = b / 10
	}
	if x == y {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
