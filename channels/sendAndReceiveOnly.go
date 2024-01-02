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
		number := <-ch
		fmt.Printf("receive data from go-routine 2 %v\n", number)
		wg.Done()
	}(channel)

	go func(ch chan<- int) {
		number1 := 10
		number2 := 11
		ch <- number1
		ch <- number2
		wg.Done()
	}(channel)
	wg.Wait()
}
