/*
Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.

Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// custom to while-loop
	iterator := 0
	for iterator < 10{
		fmt.Println("iterator", iterator);
		iterator++;
	}

	// custom infinity loop
	j := 0
	for {
		// break loop
		if (j < 10) {
			fmt.Println("break infinity loop");
			break;
		}
		j++
	}

	/*
		if expression; condition {}
	*/
	if v := 1; v < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("positive")
	}

	// switch expression, value case condition
	switch num := 1; num {
	case 1: 
		fmt.Println("number", 1)
	case 2:
		fmt.Println("number", 2)
	default: 
		fmt.Println("no value")
	}

	// run same if else from top to bottom
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println(" less than 12")
	case t.Hour() > 12:
		fmt.Println(" greater than 12")
	default: 
		fmt.Println(" no value")
	}
	
}
