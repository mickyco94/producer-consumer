package main

import (
	"fmt"
	"time"
)

var input = []int{1, 2, 3, 4, 5, 6}

var consumerCount = 3

func main() {
	jobs := make(chan int)
	done := make(chan struct{})

	go produce(jobs)

	for i := 0; i < consumerCount; i++ {
		go consume(i, jobs, done)
	}

	<-done
}

func produce(chn chan int) {
	for _, v := range input {
		chn <- v
	}
	close(chn)
}

func consume(idx int, chn chan int, done chan struct{}) {
	for i := range chn {
		fmt.Printf("Consumer %v received: %v\n", idx, i)
		time.Sleep(5 * time.Second)
	}
	done <- struct{}{}
}
