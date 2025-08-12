# Go Zero Values, Pointers, and Optional Fields — A Complete Guide

This guide covers:

- How Go initializes **zero values**
- How **pointers** and **dereferencing** work
- Why **nil pointers panic**
- How **struct zero values** behave
- How to handle **optional fields** (like in TypeScript)

---

## 1. Zero Values in Go

When you declare a variable without assigning a value, Go gives it the **zero value** for its type.

| Type kind | Example type     | Zero value      |
| --------- | ---------------- | --------------- |
| Boolean   | `bool`           | `false`         |
| Number    | `int`, `float64` | `0`, `0.0`      |
| String    | `string`         | `""` (empty)    |
| Struct    | `struct{X int}`  | all fields zero |
| Pointer   | `*int`           | `nil`           |
| Slice     | `[]int`          | `nil`           |
| Map       | `map[string]int` | `nil`           |
| Channel   | `chan int`       | `nil`           |
| Interface | `interface{}`    | `nil`           |
| Function  | `func()`         | `nil`           |

Example:

```go
package main

import "fmt"

type Sub struct {
	A int
	B string
}

func main() {
	var b bool        // false
	var n int         // 0
	var s string      // ""
	var sub Sub       // {0 ""}
	var p *int        // nil
	var sl []int      // nil
	var m map[string]int // nil
	var i interface{} // nil

	fmt.Println(b, n, s, sub, p, sl, m, i)
}
```

---

## 2. Struct Zero Values

Structs are **never nil** when stored as values — even if not initialized, Go allocates an empty struct.

```go
type Sub struct {
	Value string
}

func main() {
	var s Sub
	fmt.Println(s.Value) // "" — zero value for string
}
```

If the field is **not a pointer**, you can't tell if it was “missing” in JSON — missing and empty look the same.

---

## 3. Pointers

A **pointer** stores the **memory address** of a value.

- `&x` — address of `x`
- `*p` — value stored at the address `p` points to (**dereference**)

Example:

```go
x := 42
p := &x // p is a pointer to x
fmt.Println(p)  // address like 0xc0000180a8
fmt.Println(*p) // 42 — dereference
```

---

### Pointer Visual Diagram

```
x:  42         (value in memory)
p:  0xc0000120 (pointer storing address of x)

p ----> [ 42 ]
```

- `p` is the map (address).
- `*p` is the treasure at that location.

---

## 4. Dereferencing and Nil Pointers

If a pointer is `nil`, it points to nothing. Dereferencing it causes a **panic**.

```go
var p *int // nil
fmt.Println(p)  // <nil>

// ❌ panic: invalid memory address or nil pointer dereference
fmt.Println(*p)
```

**Safe usage:**

```go
if p != nil {
    fmt.Println(*p)
} else {
    fmt.Println("p is nil")
}
```

---

## 5. Interface Zero Values

An interface holds:

- A **type** (dynamic type)
- A **value** (data)

Zero value of an interface = both are nil.

```go
var i interface{}
fmt.Println(i) // <nil>

// ❌ panic: interface conversion: <nil> is nil, not string
fmt.Println(i.(string))
```

---

## 6. When Panics Happen

You **do not panic** just by reading zero values.  
Panics happen when you use a nil value as if it were initialized.

| Action                    | Zero value? | Panic? |
| ------------------------- | ----------- | ------ |
| Read zero bool/int/string | ✅          | ❌     |
| Access struct zero field  | ✅          | ❌     |
| Dereference nil pointer   | ✅          | ⚠ Yes  |
| Write to nil map          | ✅          | ⚠ Yes  |
| Append to nil slice       | ✅          | ❌ No  |
| Index nil slice           | ✅          | ⚠ Yes  |
| Type assert nil interface | ✅          | ⚠ Yes  |

---

## 7. Optional Fields (TypeScript vs Go)

TypeScript:

```ts
interface Config {
  isMandatory?: boolean;
}
```

- If missing → `undefined`

Go equivalent with **pointer**:

```go
type Config struct {
	IsMandatory *bool `json:"isMandatory,omitempty"`
}
```

- Missing in JSON → `nil`
- Provided → `true` or `false`
- Must nil-check before use.

Example:

```go
if c.IsMandatory != nil && *c.IsMandatory {
    fmt.Println("Mandatory")
}
```

For optional structs:

```go
type Metadata struct {
	Value string
}

type Config struct {
	Metadata *Metadata `json:"metadata,omitempty"`
}
```

- Missing → `nil`
- Present but empty → `&Metadata{}`
- Must nil-check before field access.

---

## 8. Summary

- **Zero values**: Every type in Go has a safe zero value.
- **Pointers**: Store addresses; dereference to get the value.
- **Nil pointers**: Dereferencing → panic.
- **Structs**: Non-pointer structs never nil; all fields zero value.
- **Optional fields**: Use pointers (`*T`, `*Struct`, `*[]T`) to detect missing data from JSON.
- **Rule of thumb**: Always nil-check before dereferencing.
