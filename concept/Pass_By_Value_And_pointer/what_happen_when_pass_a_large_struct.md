## When you pass a large struct by value in Go, the entire struct is copied. This can have a few important implications:

## 🔁 1. Copying Overhead

- The struct's entire contents are copied to the new variable or function parameter.
- If the struct is large, this copying can become expensive in terms of performance (CPU cycles and memory bandwidth).
- Copying is shallow for value types but copies pointers as-is (not the data they point to).

```
type LargeStruct struct {
    A [1000000]int
}

func doSomething(s LargeStruct) {
    // s is a full copy of the argument
}

```

Calling doSomething(someLargeStruct) copies all 1 million integers. <br>

## 🧠 2. Separate Memory

- The original and the copy are completely independent.
- Changes made to the struct inside the function do not affect the original struct.

```
func modify(s LargeStruct) {
    s.A[0] = 100 // Only changes the copy
}

```

## 📦 3. Escape Analysis & Stack vs Heap

- If a struct is large and passed around, the compiler may promote it to the heap to avoid excessive stack usage.
- Heap allocations are slower and involve GC pressure (garbage collection), impacting performance.

# 💡 Best Practices

- If your struct is large (e.g., > a few KB), pass by pointer (\*Struct) instead of by value.

  - This avoids copying.
  - Still allows for modification (unless you make the receiver method or function parameter read-only by convention).

- Even if the struct is small, but you need to modify it, use a pointer.

## 🔬 Extra Tip – Profiling

```
go build -gcflags="-m"

```
