package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 2

	<-ch
	close(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
