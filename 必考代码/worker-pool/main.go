package main

import (
	"fmt"
	"sync"
)

// 只启动三个go routine, 并发打印0-10000

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := range ch {
				fmt.Println(j)
			}

			wg.Done()
		}()
	}

	for k := 0; k < 10000; k++ {
		ch <- k
	}

	close(ch)
	wg.Wait()

}
