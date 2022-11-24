package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

var input = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}

var producerCount = len(input)
var consumerCount = 5

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt)
		<-sig
		cancel()
	}()

	jobs := make(chan int)
	producerWg := sync.WaitGroup{}
	consumerWg := sync.WaitGroup{}

	for i := 0; i < len(input); i++ {
		producerWg.Add(1)
		go produce(ctx, i, jobs, &producerWg)
	}
	for i := 0; i < consumerCount; i++ {
		consumerWg.Add(1)
		go consume(ctx, i, jobs, &consumerWg)

	}

	producerWg.Wait()
	close(jobs)
	consumerWg.Wait()
}

func produce(ctx context.Context, idx int, chn chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, v := range input[idx] {
		select {
		case <-ctx.Done():
			fmt.Printf("Exiting producer: %v\n", idx)
			return
		default:
			chn <- v
			time.Sleep(5 * time.Second)
		}
	}
}

func consume(ctx context.Context, idx int, chn <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Exiting consumer: %v\n", idx)
			return
		case v, open := <-chn:
			if !open {
				return
			}
			fmt.Printf("Consumer received: %v\n", v)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
