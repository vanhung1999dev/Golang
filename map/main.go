/*
A map maps keys to values.

The zero value of a map is nil. A nil map has no keys, nor can keys be added.

The make function returns a map of the given type, initialized and ready for use.
*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}


/*
name map[type of key]value, default value => nil and not ready to use,
solution1: m = map[string]Vertex{}
solution2: m = make(map[string]Vertex)
*/

/*
Insert or update an element in map m:

m[key] = elem
Retrieve an element:

elem = m[key]
Delete an element:

delete(m, key)
Test that a key is present with a two-value assignment:

elem, ok = m[key]
If key is in m, ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type.
*/
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	freq := make(map[string]int)

	for i := 0; i<20; i++ {
		freq["last_index"] = i
	}

	fmt.Println("freq", freq)

	dict := make(map[string]int)
	dict["harry potter"] = 1
	dict["seven sky"] = 2

	_, isOk := dict["comic"]
	fmt.Println("Is comic exist",isOk)

	for key, val := range dict {
		fmt.Println(key, val)
	}
}