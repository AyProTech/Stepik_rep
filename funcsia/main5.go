func sumInt(a ...int) (int, int) {
	var b, s int
	for _, i := range a {
		b += 1
		s = s + i
	}

	return b, s
}