func removeDuplicates(inputStream chan string, outputStream chan string) {
	defer close(outputStream)
	var str string
	for {
		v, ok := <-inputStream
		if !ok {
			break
		}
		if v != str {
			outputStream <- v
			str = v
		}
	}
}