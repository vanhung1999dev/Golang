## 1. Unnecessary Allocations (especially in hot paths)

```
for i := 0; i < n; i++ {
    s := fmt.Sprintf("value: %d", i) // allocates each time
}

```

- fmt.Sprintf allocates a new string every loop.
- Using []string with repeated append grows memory

### ✅ Fix:

- Reuse buffers (strings.Builder, sync.Pool)
- Pre-allocate slices with make([]T, 0, cap)
- Use strconv.Itoa over fmt.Sprintf when you can

## 2. Excessive Goroutines

```
for i := 0; i < 1000000; i++ {
    go doSomething(i) // system overload
}

```

- Each goroutine adds ~2KB stack space
- Scheduler overhead grows fast

### ✅ Fix:

- Use bounded worker pools (e.g. 100 workers)
- Use buffered channels and sync.WaitGroup

### 3. Incorrect Slice Growth

```
s := []int{}
for i := 0; i < 100000; i++ {
    s = append(s, i) // reallocates often
}

```

### ✅ Fix:

```
s := make([]int, 0, 100000)

```

## 4. Heavy Use of interface{} or reflect

❌ What happens: <br>

- interface{} disables inlining and adds boxing
- reflect introduces type checks, allocations, and indirection

✅ Fix: <br>

- Use concrete types or generics
- Avoid reflect in core logic

## 5. Boxing of Values in interface{}

```
var x interface{} = 42 // allocates

```

This escapes to the heap since interface{} holds a boxed copy. <br>

✅ Fix: <br>

- Avoid interface{} when concrete types suffice
- Watch for escape analysis (go build -gcflags="-m")

## 6. Copying Large Structs

```
type Big struct {
    A [1024]byte
}
func process(b Big) { ... } // passed by value

```

### ✅ Fix:

- Pass by pointer: func process(b \*Big) <br>

## 7. Memory Leaks via Goroutines / Channels

### ❌ What happens:

- Goroutines waiting on unreceived channels never exit.
- Buffered channels that aren't drained = memory retention.

### ✅ Fix:

- Always cancel with context.Context
- Drain channels and close them properly

## 8. Unoptimized JSON/encoding

### ❌ What happens:

- Using json.Marshal with deeply nested structs → allocations

### ✅ Fix:

- Reuse buffers with json.NewEncoder(buf)
- Use easyjson, go-json, segmentio/encoding, etc. for zero-alloc

## 9. Using defer in Hot Loops

### ❌ What happens:

```
for i := 0; i < 1000000; i++ {
    defer closeFile(f) // O(n) stack buildup
}

```

### ✅ Fix:

- Use defer outside loops or manually manage cleanup when perf matters.

## 10. Map Lookup with Non-Comparable Keys

### ❌ What happens:

- Using large struct keys or []byte → expensive comparison

### ✅ Fix:

- Use strings or small primitives for keys
- If using struct keys, keep them small and comparable

## 11. Overusing Mutexes

### ❌ What happens:

- Frequent contention = goroutine suspension

### ✅ Fix:

- Minimize lock hold time
- Use lock-free patterns (channels, atomics) when possible

## 12. Misusing Channels

### ❌ What happens:

- Channels used for logging, coordination, batching → adds latency

### ✅ Fix:

- Use channels sparingly
- Use slices, queues, or sync.Cond for high-throughput cases

## 13. String Concatenation in Loops

```
s := ""
for _, part := range parts {
    s += part // new allocation each time
}

```

### ✅ Fix:

Use strings.Builder for efficient string concat.

## 14. Large Heap Objects Not Reused

### ❌ What happens:

- Large slices/structs are repeatedly created and dropped

### ✅ Fix:

- Use sync.Pool for big reusable objects (e.g. buffers)

# 🔍 Bonus Tools to Help You Detect These

go build -gcflags="-m" => Shows escape analysis <br>
pprof => CPU/memory profiling <br>
benchstat => Compare benchmark results <br>
go test -bench . => Run benchmarks <br>
go tool trace => View goroutines, blocking, syscalls <br>
