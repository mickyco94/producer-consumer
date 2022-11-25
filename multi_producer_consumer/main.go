package main

import (
	"fmt"
	"sync"
	"time"
)

var input = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}

var producerCount = len(input)
var consumerCount = 5

func main() {
	jobs := make(chan int)
	producerWg := sync.WaitGroup{}
	consumerWg := sync.WaitGroup{}

	for i := 0; i < len(input); i++ {
		producerWg.Add(1)
		go produce(i, jobs, &producerWg)
	}
	go consume(jobs, &consumerWg)

	producerWg.Wait()
	close(jobs)
	consumerWg.Wait()
}

func produce(idx int, chn chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, v := range input[idx] {
		chn <- v
		time.Sleep(250 * time.Millisecond)
	}
}

func consume(chn <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range chn {
		fmt.Printf("Consumer received: %v\n", v)
		time.Sleep(500 * time.Millisecond)
	}
}
