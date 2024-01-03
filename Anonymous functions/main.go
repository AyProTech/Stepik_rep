// fn
fn := func(x uint) uint {
	j := int(x)
	a := 0
	cnt := 0
	b := ""
	for j != 0 {
		a = j % 10
		if a%2 == 0 && a != 0 {
			b += strconv.Itoa(a)
			cnt++
		}
		j = j / 10
	}
	d := 0
	if cnt != 0 {
		c, _ := strconv.Atoi(b)

		for c != 0 {
			d = d*10 + c%10
			c /= 10
		}

	} else {
		return uint(100)
	}
	return uint(d)

}