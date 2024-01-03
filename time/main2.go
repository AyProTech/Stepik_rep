package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	a := bufio.NewScanner(os.Stdin)
	a.Scan()

	firstTime, err := time.Parse("2006-01-02 15:04:05", a.Text())
	if err != nil {
		panic(err)

	}

	if firstTime.Hour() > 13 {
		newTime := firstTime.AddDate(0, 0, 1)

		fmt.Println(newTime.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println(a.Text())
	}
}
