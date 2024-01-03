func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	ans := make(chan int)
	go func() {
		defer close(ans)
		var sum int

		for {
			select {
			case a := <-arguments:
				sum += a
			case <-done:
				ans <- sum
				return
			}
		}
	}()
	return ans
}




