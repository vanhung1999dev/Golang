package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var mutex = sync.Mutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)    // add to routine
		mutex.Lock() // lock to read counter variable
		go printCounter()
		mutex.Lock() // lock to incr counter variable
		go incrCounter()
	}
	wg.Wait() // await until two thread completed
}

func printCounter() {
	fmt.Println(counter)
	mutex.Unlock()
	wg.Done() // notification to a wait group and -1
}

func incrCounter() {
	counter++
	mutex.Unlock()
	wg.Done()
}
