func adding(x, y string) int64 {

	a := []rune(x)
	b := []rune(y)
	n := ""
	m := ""
	for _, i := range a {
		if unicode.IsDigit(i) {
			n += string(i)
		}
	}
	for _, i := range b {
		if unicode.IsDigit(i) {
			m += string(i)
		}
	}
	cnt, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	cnt1, err1 := strconv.Atoi(m)
	if err != nil {
		panic(err1)
	}

	return int64(cnt + cnt1)

}