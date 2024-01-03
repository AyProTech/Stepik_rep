func minimumFromFour() int {
	//print your code here
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if a < b && a < c && a < d {
		return a
	} else if b < c && b < d {
		return b
	} else if c < d {
		return c
	} else {
		return d
	}

}