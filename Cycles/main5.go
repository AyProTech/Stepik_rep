package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	for i := 1; i <= a; i++ {
		if i%b == 0 && i%c != 0 {
			fmt.Println(i)
			break
		}
	}
}
