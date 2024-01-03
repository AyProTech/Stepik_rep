package main

import "fmt"

func main() {
	var v int
	fmt.Scan(&v)
	if v > 0 {
		fmt.Println("Число положительное")
	} else if v < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}

}
