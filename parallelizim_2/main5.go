package main

// пакет используется для проверки выполнения условия задачи, не удаляйте его

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		for i := range in1 {
			out <- i * 2
		}
	}()

}
