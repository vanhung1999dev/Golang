# Internal Structure of a Slice

```
type sliceHeader struct {
    Data uintptr // Pointer to the underlying array
    Len  int     // Current length of the slice
    Cap  int     // Capacity of the slice
}

```

- **Data** → A pointer to the first element of the underlying array.

- **Len** → The number of elements currently in the slice.

- **Cap** → The number of elements the slice can hold before needing to grow.

## Example: Creating a Slice

```
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5} // Underlying array
    slice := arr[1:4]            // Slice referencing part of `arr`

    fmt.Println(slice)           // Output: [2 3 4]
    fmt.Println(len(slice))      // Output: 3
    fmt.Println(cap(slice))      // Output: 4
}

```

- slice := arr[1:4] creates a slice [2,3,4] referencing arr.

- The length is 3 (number of elements in the slice).

- The capacity is 4 (from arr[1] to the end of arr).

## What Happens When a Slice Grows Beyond Its Capacity?

### Growth Strategy

When appending to a slice exceeds its capacity, Go **allocates a new larger array**and copies the old elements into it. This avoids excessive memory usage while still allowing growth <br>

### Example: Slice Growth

```
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3}
    fmt.Println(len(slice), cap(slice)) // Output: 3 3

    slice = append(slice, 4) // Exceeds capacity, triggers reallocation
    fmt.Println(len(slice), cap(slice)) // Output: 4 6
}

```

## Growth Pattern

- Go **doubles** the slice’s capacity when needed.

- If the **old capacity < 1024**, it **doubles** (cap = cap \* 2).

- If the **old capacity >= 1024**, it **grows by 25%** (cap = cap + cap/4).

## Example

```
s := make([]int, 3, 3)
fmt.Println(cap(s)) // 3

s = append(s, 10) // Triggers reallocation
fmt.Println(cap(s)) // 6 (growth factor: x2)

```

# Edge Cases to Watch Out For

## . Slice Re_slicing (Changing len, Keeping cap)

```
arr := []int{1, 2, 3, 4, 5}
slice := arr[1:3] // [2, 3]
slice = slice[:4] // Expands the slice (within capacity)

fmt.Println(slice) // Output: [2 3 4 5]

```

- You can increase the **length** of a slice **up to its capacity**.

- Beyond capacity → **Causes runtime panic**.

## Modifying a Slice Modifies the Underlying Array

```
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // [2,3,4]
slice[0] = 100

fmt.Println(arr) // Output: [1 100 3 4 5]

```

- Changes **reflect** in the original array.

## Appending May Break Reference Sharing

```
arr := [3]int{1, 2, 3}
slice1 := arr[:2] // [1,2]
slice2 := append(slice1, 4) // Exceeds arr capacity

fmt.Println(arr)   // Output: [1, 2, 3] (Unchanged)
fmt.Println(slice2) // Output: [1, 2, 4] (New array)

```

- append() **creates a new array** if capacity is exceeded.

## Nil vs. Empty Slice

```
var s1 []int // nil slice
s2 := []int{} // empty slice

fmt.Println(s1 == nil) // true
fmt.Println(s2 == nil) // false

```

- nil slices have len = 0, cap = 0, and no backing array.

- Empty slices have len = 0, cap > 0 and an allocated array.
