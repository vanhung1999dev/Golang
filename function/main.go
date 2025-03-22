package main

import "fmt"

func main() {
	fmt.Println("sum", add(1,2))
	sum, desc := add2(1,2)
	
	fmt.Println(desc, sum)

	fmt.Println("sum3", add3(1,2))

	anonyms := func(x int) int {
		return x
	}
	fmt.Println("anonyms", anonyms(2))
}

func add(num1 int, num2 int) int {
	return num1 + num2;
}

func add2(num1, num2 int) (int, string) {
	return num1 + num2, "sum"
}

// not prefer this way, not readable => implicit return
func add3(num1, num2 int) (sum int) { // implicit int sum = 0, and auto return sum;
	sum = num1 + num2
	return;
}