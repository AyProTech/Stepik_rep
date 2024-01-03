package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a%10 == 1 && a != 11 {
		fmt.Println(a, "korova")
	} else if a%10 == 2 && a != 12 {
		fmt.Println(a, "korovy")
	} else if a%10 == 3 && a != 13 {
		fmt.Println(a, "korovy")
	} else if a%10 == 4 && a != 14 {
		fmt.Println(a, "korovy")
	} else {
		fmt.Println(a, "korov")
	}
}
