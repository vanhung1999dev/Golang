package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var channel = make(chan int)

func main() {
	wg.Add(2)
	go func(ch <-chan int) {
		number := <-ch
		fmt.Printf("receive data from go-routine 2 %v\n", number)
		wg.Done()
	}(channel)

	go func(ch chan<- int) {
		number := 10
		ch <- number
		wg.Done()
	}(channel)
	wg.Wait()
}
