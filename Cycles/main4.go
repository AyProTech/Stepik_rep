package main

import "fmt"

func main() {
	var max, i, sum int
	fmt.Scan(&i)
	for i != 0 {
		if i > max {
			max = i
			sum = 1
		} else if i == max {
			sum++
		}
		fmt.Scan(&i)
	}
	fmt.Println(sum)
}
