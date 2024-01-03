// package main уже объявлен.
func main() {
	// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
	// Другие использовать нельзя.

	sum := 0
	file := bufio.NewScanner(os.Stdin)

	for file.Scan() {
		i, _ := strconv.Atoi(file.Text())
		sum += i

	}
	io.WriteString(os.Stdout, strconv.Itoa(sum))

}