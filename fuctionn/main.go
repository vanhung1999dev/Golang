package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	const total = sum(1, 1)
	fmt.Println("total", total)
}
