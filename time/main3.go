package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"time"
)

func main() {
	a, err := bufio.NewReader(os.Stdin).ReadString('\n')

	b := strings.Split(a, ",")

	a1, err := time.Parse("02.01.2006 15:04:05", b[0][:19])
	if err != nil {
		panic(err)
	}
	b1, err := time.Parse("02.01.2006 15:04:05", b[1][:19])
	if err != nil {
		panic(err)
	}

	if a1.Before(b1) {
		fmt.Print(b1.Sub(a1))
		return
	}
	fmt.Print(a1.Sub(b1))

}
