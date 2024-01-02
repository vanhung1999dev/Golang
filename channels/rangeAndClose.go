package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var capacityOfChannel = 2
var channel = make(chan int, capacityOfChannel)

func main() {
	wg.Add(2)
	go func(ch <-chan int) {
		for data := range ch { // get each data in channel
			fmt.Printf("receive data from go-routine 2 %v\n", data)
		}
		wg.Done()
	}(channel)

	go func(ch chan<- int) {
		number1 := 10
		number2 := 11
		ch <- number1
		ch <- number2
		close(ch) // noti that closed channel and does not send data, stop receive data
		wg.Done()
	}(channel)
	wg.Wait()
}
