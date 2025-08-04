# 📘 Golang `struct` Deep Dive: Declaration, Behavior, Value vs Pointer Embedding

---

## 🧱 1. What is a `struct` in Go?

A `struct` is a composite data type in Go used to group related fields together.

### Declaration

```go
type User struct {
    Name  string
    Age   int
    Email string
}
```

### Initialization

```go
u1 := User{"Alice", 30, "alice@example.com"}         // full
u2 := User{Name: "Bob"}                              // partial
u3 := new(User)                                      // returns *User with zero-values
```

### Access and Update

```go
fmt.Println(u1.Name) // "Alice"
u1.Age = 31
```

---

## 🧩 2. Attaching Methods to Structs

```go
func (u User) Greet() string {
    return "Hello, " + u.Name
}

func (u *User) Birthday() {
    u.Age++
}
```

- **Value receiver**: operates on a copy
- **Pointer receiver**: can modify the original

---

## 🧬 3. Struct Embedding (Composition in Go)

Go doesn’t support classical inheritance. Instead, it allows struct embedding to reuse fields and methods.

---

### A. Embedding as Value (non-pointer)

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return "I am " + a.Name
}

type Dog struct {
    Animal  // embedded by value
    Breed   string
}

func main() {
    d := Dog{
        Animal: Animal{Name: "Buddy"},
        Breed:  "Golden Retriever",
    }

    fmt.Println(d.Speak())      // ✅ inherits method
    d.Animal.Name = "Charlie"
    fmt.Println(d.Speak())      // I am Charlie
}
```

- `Animal` is copied into `Dog` as a value.
- Any external change to another `Animal` will **not** affect the embedded one.

---

### B. Embedding as Pointer

```go
type Cat struct {
    *Animal // pointer embedded
    Color   string
}

func main() {
    a := &Animal{Name: "Mimi"}
    c := Cat{
        Animal: a,
        Color:  "White",
    }

    fmt.Println(c.Speak())      // I am Mimi
    a.Name = "Luna"
    fmt.Println(c.Speak())      // I am Luna
}
```

- `Cat` holds a **pointer** to `Animal`.
- Changes to `a.Name` are reflected in `c`.

---

## 🤯 Value vs Pointer Embedding Comparison

| Feature         | Value Embedding            | Pointer Embedding           |
| --------------- | -------------------------- | --------------------------- |
| Type            | Struct directly            | Pointer to struct           |
| Behavior        | Copies the value           | Shares the same memory      |
| Changes Reflect | No (isolated)              | Yes (mutual change)         |
| Memory Use      | Slightly more per instance | Efficient for large structs |
| Use Cases       | Immutable/simple data      | Shared or mutable state     |

---

## 🔧 When to Use Pointer Embedding

- You want to **share state** across multiple structs
- Struct is **large**, and copying would be inefficient
- You want **modifications** to propagate

---

## 🚫 When NOT to Use Pointer Embedding

- The embedded struct is **small and simple**
- You want **structs to be independent**
- You don't want to worry about **nil pointers**

## Structs are never nil unless you use a pointer to struct.

```
type MyStruct struct {
    Name string
}

var s MyStruct
fmt.Println(s == nil) // ❌ compile error

```

You can’t compare a struct to nil — it’s a value type. <br>

```
var ps *MyStruct = nil
fmt.Println(ps == nil) // ✅ true

```

So for structs: <br>

- You only get nil when you're dealing with pointers
- A value struct (even with zero fields) is never nil

---

## ✅ Summary

- `struct` is Go’s way of grouping data logically
- Methods can be attached to structs with value or pointer receivers
- Embedding lets one struct reuse another’s fields and methods
- Value embedding → **copy** behavior (safe, isolated)
- Pointer embedding → **shared** behavior (efficient but careful)

---
