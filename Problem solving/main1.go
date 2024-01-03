package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	sum := 0
	for a != 0 {
		sum = sum + a%10
		a = a / 10
	}
	fmt.Println(sum)

}
