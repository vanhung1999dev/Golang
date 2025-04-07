# Escape Analysis in Go

Escape analysis is a compiler optimization technique in Go that determines whether a variable can be allocated on the stack or must escape to the heap. <br>

## When Does a Variable Escape to the Heap?

### Returning a Reference to a Local Variable:

```
func foo() *int {
    x := 42
    return &x // x escapes because it must survive function return
}
```

- Since x is returned, it escapes to the heap to avoid being deallocated.

### Interface Conversion (Dynamic Types):

```
func printVal(i interface{}) {
    fmt.Println(i) // i escapes to heap
}

func main() {
    x := 42
    printVal(x) // x escapes because it’s stored in an interface{}
}
```

- Storing x inside interface{} forces it onto the heap.

### Slice of Structs (or Large Structs in General):

```
type Data struct {
    values [1024]int
}

func foo() *Data {
    d := Data{}
    return &d // Heap allocation due to large struct
}
```

- Large structs are often allocated on the heap to avoid excessive stack usage.

### Closures Capturing Variables:

```
func closure() func() int {
    x := 10
    return func() int { return x } // x escapes to heap
}
```

- x is used inside a closure, so it must persist beyond closure()'s execution.

### Heap Allocation for Structs with Methods that Require Pointers:

```
type Person struct {
    name string
}

func (p *Person) SetName(newName string) {
    p.name = newName
}

func main() {
    p := &Person{"John"} // Escapes to heap
}
```

- Since p is used with a pointer receiver, it’s allocated on the heap.

# How to Check Escape Analysis in Go? Use the -gcflags="-m" compiler flag:

```
go run -gcflags="-m" main.go

```

OUTPUT: <br>

```
main.go:10:6: moved to heap: x

```
