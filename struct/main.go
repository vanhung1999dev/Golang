package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Person struct {
	name string
	age int
}


// struct tag
// Struct tags are used by many serialization packages; the built-in encoding/json , encoding/xml , and many other external packages such as yaml use them. Struct tags allow you to add extra information about a field. They have a known format: key:”value”
type Car struct {
	name string `filed: name`
}


// should capital first character if use from outside package
type Bike struct {
	Name string
}

func main() {
	fmt.Println(Vertex{1, 2})

	p := Person{"hung", 25};
	p1 := Person{} // init default value for name and age
	p2 := Person{name: "bang"} // init default age = 0;
	p3 := Person{age: 10}
	g := &p;

	g.name = "crush"

	fmt.Println("name", p.name)
	fmt.Println("p1", p1)
	fmt.Println("p2", p2)
	fmt.Println("p3", p3)
	fmt.Println("g pointer", g)
}
