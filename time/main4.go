package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"time"
)

func main() {
	const now = 1589570165
	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	b := strings.Trim(a, "\n")

	b = strings.Replace(b, " мин. ", "m", 1)
	b = strings.Replace(b, " сек.", "s", 1)

	dur, err := time.ParseDuration(b)
	if err != nil {
		panic(err)
	}

	unixDur := time.Unix(now, 0).UTC()

	timeDur := unixDur.Add(dur)

	fmt.Println(timeDur.Format(time.UnixDate))

}
