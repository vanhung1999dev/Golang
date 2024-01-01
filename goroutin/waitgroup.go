package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2) // add to routine
	go countAnimalToSleep("fork")
	go countAnimalToSleep("dog")
	wg.Wait() // await until two thread completed
}

func countAnimalToSleep(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, i)
	}
	wg.Done() // notification to a wait group and -1
}
