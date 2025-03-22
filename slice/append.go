/*
Appending to a slice
It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.

func append(s []T, vs ...T) []T
The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

NOTE: https://go.dev/blog/slices-intro

*/

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	// how append work, whenever the capacity is not enough, it will x2 current capacity
	var arr []int
	for i := 0; i<20 ; i++ {
		arr = append(arr, i)
		fmt.Printf("len=%d cap=%d %v\n", len(arr), cap(arr), arr)
	}

	// when you call append, Go will always check the capacity and re-allocated if need
	// so you can pre-define the capacity and length if it can predict
	const PREDICT_SIZE = 5
	arr1 := make([]int, 0, PREDICT_SIZE)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}