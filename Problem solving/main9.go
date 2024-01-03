package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	sum1 := 0
	for n != 0 {
		sum = sum + n%10
		n = n / 10
	}

	for sum != 0 {
		sum1 = sum1 + sum%10
		sum = sum / 10
	}
	fmt.Println(sum1)
}
