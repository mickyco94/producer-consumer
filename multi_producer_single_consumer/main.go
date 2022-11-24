package main

import (
	"fmt"
	"sync"
)

var input = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}

func main() {
	jobs := make(chan int)
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	for i := 0; i < len(input); i++ {
		wg.Add(1)
		go produce(i, jobs, &wg)
	}
	go consume(jobs, done)

	wg.Wait()
	close(jobs)
	<-done
}

func produce(idx int, chn chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, v := range input[idx] {
		chn <- v
	}
}

func consume(chn <-chan int, done chan<- struct{}) {
	for v := range chn {
		fmt.Printf("Consumer received: %v\n", v)
	}
	done <- struct{}{}
}
