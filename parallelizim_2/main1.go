done := make(chan struct{})
go func(d chan struct{}) {
	work()
	close(done)
}(done)
<-done
