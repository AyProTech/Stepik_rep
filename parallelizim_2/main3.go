func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	done := make(chan int)
	go func() {
		defer close(done)

		select {
		case a := <-firstChan:
			done <- (a * a)
		case a := <-secondChan:
			done <- (a * 3)
		case <-stopChan:
			return
		}
	}()
	return done
}