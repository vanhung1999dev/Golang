# Go Slice Internals Deep Dive
---

## 1. Overview

Slices are one of the most important Go data structures.

They look simple:

```go
s := []int{1,2,3}
```

But internally a slice is a lightweight descriptor over an array.

Understanding slices deeply means knowing:

* pointer / len / cap layout
* append growth behavior
* reallocation rules
* sharing underlying arrays
* copy semantics
* memory leaks from subslices
* performance tuning with preallocation

---

## 2. What a Slice Really Is

A slice is not the array itself.

Conceptually:

```go
type slice struct {
    ptr *T
    len int
    cap int
}
```

Where:

* `ptr` points to first visible element in backing array
* `len` = number of usable elements
* `cap` = elements available before reallocation

---

## 3. Example Layout

```go
arr := [5]int{10,20,30,40,50}
s := arr[1:4]
```

Slice view:

```text
array: [10,20,30,40,50]
         ^
         ptr -> 20
len = 3   // [20,30,40]
cap = 4   // [20,30,40,50]
```

Capacity counts from pointer position to end of backing array.

---

## 4. Why Slice Is Cheap to Pass

Passing slice to function copies only header:

```text
ptr + len + cap
```

Not entire underlying array.

So this is cheap:

```go
func f(s []int)
```

---

## 5. But Elements Are Shared

```go
func f(s []int) {
    s[0] = 999
}
```

Caller sees modification if same backing array.

Because header copied, data shared.

---

## 6. len vs cap

```go
s := make([]int, 3, 8)
```

Then:

```text
len = 3
cap = 8
```

You can access indexes:

```text
0..2
```

Append may grow into reserved capacity without realloc.

---

## 7. Append Fast Path

```go
s = append(s, x)
```

If:

```text
len < cap
```

Runtime writes new element into existing backing array.

Then:

```text
len++
```

No allocation.

Very fast.

---

## 8. Append Slow Path

If:

```text
len == cap
```

Need growth:

1. allocate new larger array
2. copy old elements
3. append new element
4. return new slice header

---

## 9. Why append Returns Slice

Because backing array may change.

Correct usage:

```go
s = append(s, 10)
```

Wrong to ignore returned slice when growth possible.

---

## 10. Growth Algorithm (Conceptual)

Historically small slices often grow around 2x.
Larger slices grow more conservatively (~1.25x range).

Purpose:

* reduce realloc frequency for small slices
* reduce memory waste for large slices

Exact runtime details may evolve by Go version.

---

## 11. Typical Growth Example

Start capacity 2:

```text
2 -> 4 -> 8 -> 16 -> ...
```

Larger sizes may become:

```text
1024 -> 1280 -> 1600 -> ...
```

Approximate, not strict contract.

---

## 12. Why Not Always 2x?

If huge slice always doubled:

```text
1 GB -> 2 GB -> 4 GB
```

Massive waste and memory pressure.

So larger slices grow slower.

---

## 13. Copy Cost During Growth

Reallocation copies existing elements.

If large struct elements:

```go
[]BigStruct
```

Growth can be expensive.

Prefer preallocation when size known.

---

## 14. Preallocation Pattern

```go
s := make([]int, 0, 100000)
```

Then many appends avoid repeated reallocations.

Excellent production optimization.

---

## 15. Nil Slice vs Empty Slice

```go
var a []int      // nil
b := []int{}     // empty non-nil
```

Both may have len=0.

But:

```text
a == nil true
b == nil false
```

Useful in APIs / JSON semantics.

---

## 16. Subslice Sharing Danger

```go
big := make([]byte, 1_000_000)
small := big[:10]
```

`small` still references huge backing array.

GC cannot free big array while `small` exists.

Potential memory retention bug.

---

## 17. Fix Memory Retention

Copy needed bytes:

```go
small2 := append([]byte(nil), small...)
```

Now independent backing array.

---

## 18. Full Slice Expression to Limit Capacity

```go
s := arr[1:3:3]
```

Meaning:

```text
low=1 high=3 max=3
```

Now capacity restricted.

Useful to prevent append modifying later elements of original array.

---

## 19. Example of Hidden Sharing Bug

```go
a := []int{1,2,3,4}
b := a[:2]
b = append(b, 99)
```

If capacity allows, append may overwrite `a[2]`.

Unexpected for beginners.

---

## 20. Copy Function

```go
n := copy(dst, src)
```

Copies min(len(dst), len(src)).

Use for safe independent slice creation.

---

## 21. Range Semantics

```go
for _, v := range s
```

`v` is copy of element.

Modifying `v` does not mutate slice element for non-pointer values.

---

## 22. Slice of Struct vs Slice of Pointer

### []Struct

* better locality
* fewer heap pointers
* better GC often

### []*Struct

* easier mutation sharing
* more pointer chasing
* more GC scanning

Choose carefully.

---

## 23. Concurrency Warning

Appending to same slice from multiple goroutines without synchronization is unsafe.

Because header/data may race.

Use mutex/channel ownership.

---

## 24. Performance Rules

### Good

```go
make([]T, 0, n)
append in loop
```

### Risky

```go
append unknown millions repeatedly without capacity hint
```

### Bad

Repeated tiny appends to huge struct slices under pressure.

---

## 25. Slice Header on 64-bit Systems

Usually:

```text
ptr = 8 bytes
len = 8 bytes
cap = 8 bytes
Total = 24 bytes
```

Approximate platform dependent.

---

## 26. Interview Question: Why append Sometimes Changes Original Slice?

Because if capacity remains, append writes into shared backing array.
No reallocation occurs.

If growth reallocates, new backing array may detach.

---

## 27. Decision Table

| Scenario                     | Best Practice           |
| ---------------------------- | ----------------------- |
| Known result size            | Preallocate cap         |
| Need independent copy        | use copy / append clone |
| Avoid parent overwrite       | full slice expr         |
| Huge source keep small piece | clone subslice          |
| Read-mostly data             | []Struct often good     |

---

## 28. Final Mental Model

Slice = window onto array.

```text
header points to storage
len = visible size
cap = future growth room
```

Append may:

```text
reuse same array
or allocate new one
```

---

## 29. Senior-Level Summary

To use slices well, understand:

* header semantics
* shared backing arrays
* growth reallocations
* copy costs
* memory retention bugs
* preallocation strategies
* locality and GC tradeoffs

These matter heavily in high-throughput Go systems.
