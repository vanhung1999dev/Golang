package main

import (
	"fmt"
	"math/cmplx"
)

var bool1, bool2 bool // same type, in Go, it always init default value for each type, in this case, value = 0 for boolean type
var sum int
var desc string

// variable := value, but it only work in function

/*
Type in GO
1: Basic => int, string, bool
2: Aggregate => array, struct
3: Reference => pointers, slice, func, channel, maps
4: Interface 
*/

/*
bool

string

int  int8  int16  int32  int64 exp: int8 => 2^8 => range from [-127, 127] 
uint uint8 uint16 uint32 uint64 uintptr => only for positive value

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
*/

// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe), %T => type, %v => value, %q => string with quote, default when print will empty


// Go need to explicit convert type 

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i int
	var s string
	var b bool
	var f float64

	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %q\n", s, s)
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", f, f)

	// convert int => float64
	i32 := 10;
	f64:= float64(i32);

	fmt.Printf("Type: %T Value: %v\n", i32, i32)
	fmt.Printf("Type: %T Value: %v\n", f64, f64)

	const OS_VERSION = "linux" // type will be refer in compiler and cannot re-assign
}