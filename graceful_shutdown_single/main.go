package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var input = []int{1, 2, 3, 4, 5, 6}

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	jobs := make(chan int)
	done := make(chan struct{})

	go produce(jobs)
	go consume(jobs, done, sig)

	<-done
}

func produce(chn chan<- int) {
	for _, v := range input {
		chn <- v
		time.Sleep(1 * time.Second)
	}
	close(chn)
}

func consume(chn chan int, done chan<- struct{}, sig <-chan os.Signal) {
	for {
		select {
		case <-sig:
			fmt.Printf("Shutting down\n")
			done <- struct{}{}
			return
		case v, open := <-chn:
			if !open {
				done <- struct{}{}
				return
			}
			fmt.Printf("Consumer received: %v\n", v)
			time.Sleep(250 * time.Millisecond)
		}
	}
}
