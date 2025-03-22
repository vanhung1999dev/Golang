/*
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
 Each closure is bound to its own sum variable.
*/

package main

import "fmt"


// func adder() return function and this func also return int
// func can access to the sum outside of it'scope
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder() // now pos = func(x) int{}
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), // call func(x)
			neg(-2*i),
		)
	}
}
