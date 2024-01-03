package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a string
	a, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	s := strings.ReplaceAll(a, " ", "")
	s = strings.ReplaceAll(s, ",", ".")
	si := strings.Split(s, ";")
	n1, _ := strconv.ParseFloat(si[0], 64)
	n2, _ := strconv.ParseFloat(si[1], 64)
	fmt.Printf("%.4f", n1/n2)

}
