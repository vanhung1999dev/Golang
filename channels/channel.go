package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)                 // add to routine
	channel := make(chan int) // declare channel by keyword `chan`
	go func() {               // go routine 1 will block until it receives data from go routine 2
		number := <-channel // receive data from channel
		fmt.Printf("receive data from go-routine 2 %v\n", number)
		data := 12
		channel <- data // continue send data to channel
		wg.Done()
	}()
	go func() { // go routine 2 will send data to channel
		number := 10
		channel <- number // send data to channel
		data := <-channel
		fmt.Printf("receive data from go-routine 1 %v\n", data)
		wg.Done()
	}()
	wg.Wait() // await until two thread completed
}
