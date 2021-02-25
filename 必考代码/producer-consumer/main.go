package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(2)
	go consumer(ch, &wg)
	go producer(ch, &wg)

	wg.Wait()
}

func producer(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		out <- i
	}
}

func consumer(input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range input {
		fmt.Println("Received", i)
	}
}
