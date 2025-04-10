## 🧱 What is a "fat pointer"?

A fat pointer is a data structure that holds more than just a memory address. <br>

In Go, certain types like slices and interfaces are implemented using fat pointers because they need to track more than just the data’s address — such as size, capacity, or type. <br>

## 🥗 Fat Pointer in Slices

A slice in Go is implemented like this (conceptually): <br>

```
type SliceHeader struct {
    Data uintptr  // Pointer to the underlying array
    Len  int      // Current length
    Cap  int      // Capacity
}

```

✅ So a slice is a fat pointer: it contains a pointer + metadata (len, cap). <br>

When you pass a slice around, you're passing this **small struct by value — which contains a pointer to the data, not the data itself**. So it's lightweight but still carries the essential info. <br>
