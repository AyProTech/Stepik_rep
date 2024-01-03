func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(c)
	}
	// package main уже объявлен, нужные импорты уже есть
}
