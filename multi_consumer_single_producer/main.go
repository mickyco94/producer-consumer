package main

import (
	"fmt"
	"time"
)

var input = []int{1, 2, 3, 4, 5, 6}

var consumerCount = 5

func main() {
	jobs := make(chan int)
	done := make(chan struct{})

	go producer(jobs)

	for i := 0; i < consumerCount; i++ {
		go consumer(jobs, done)
	}

	<-done
}

func producer(chn chan int) {
	for _, v := range input {
		chn <- v
	}
	close(chn)
}

func consumer(chn chan int, done chan struct{}) {
	for i := range chn {
		fmt.Printf("Consumer received: %v\n", i)
		time.Sleep(5 * time.Second)
	}
	done <- struct{}{}
}
