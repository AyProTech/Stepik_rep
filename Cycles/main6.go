package main

import "fmt"

func main() {
	var i int

	for {
		fmt.Scan(&i)
		if i < 10 {
			continue
		} else if i > 9 && i <= 100 {
			fmt.Println(i)
		} else {
			break
		}

	}

}
