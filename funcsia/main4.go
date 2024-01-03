func fibonacci(n int) int {
	//print your code here
	a, b := 0, 1
	for i := 0; i < n-1; i++ {
		a, b = b, a+b
	}
	return b
}