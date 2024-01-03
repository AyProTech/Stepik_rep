package main

import (
	"bufio"
	"fmt"
	"os"

	"unicode"
)

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	b := []rune(text)
	a := len(b)
	if unicode.IsUpper(b[0]) && (b[a-1]) == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}

}
