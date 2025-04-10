## 🔸 1. Slices = Fat Pointers (Kind of like Interfaces)

A slice in Go is internally: <br>

```
type slice struct {
    ptr    *T // pointer to underlying array
    len    int
    cap    int
}

```

So when you do: <br>

```
var s []int // zero value
fmt.Println(s == nil) // ✅ true

```

Because all fields are zeroed: <br>

- ptr = nil
- len = 0
- cap = 0

But if you do: <br>

```
s := make([]int, 0)
fmt.Println(s == nil) // ❌ false
```

- ptr ≠ nil (allocated empty array)
- len = 0
- cap ≥ 0

So even though it looks empty, it’s not nil. <br>
