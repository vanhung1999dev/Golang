package main

import (
	"fmt"
	"time"
)

func main() {
	go countAnimalToSleep("fork")
	countAnimalToSleep("dog")
}

func countAnimalToSleep(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}
