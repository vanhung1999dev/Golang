# Go Memory Layout & Data Structures Deep Dive

## Struct Alignment, Padding, Cache Line Effects, False Sharing (Senior / FAANG Level)

---

## 1. Overview

High-performance Go systems are not only about algorithms. They are also about how data is laid out in memory.

Understanding memory layout helps you optimize:

* CPU cache efficiency
* memory footprint
* latency under load
* multicore scalability
* garbage collector pressure

Senior engineers should understand:

* struct alignment
  n- padding waste
* field ordering
* cache lines
* false sharing
* data-oriented design

---

## 2. Why Memory Layout Matters

Two structs with same logical fields can perform very differently.

Because CPUs access memory in cache lines, not individual fields.

Poor layout causes:

* extra memory usage
* more cache misses
* lower throughput
* contention between cores

---

## 3. Alignment Basics

Most CPUs prefer values aligned to natural boundaries.

Examples (typical 64-bit systems):

| Type    | Size | Preferred Alignment |
| ------- | ---: | ------------------: |
| bool    |    1 |                   1 |
| int32   |    4 |                   4 |
| int64   |    8 |                   8 |
| pointer |    8 |                   8 |
| float64 |    8 |                   8 |

If a field requires 8-byte alignment, compiler may insert padding bytes.

---

## 4. Struct Padding Example

```go
type Bad struct {
    A bool   // 1 byte
    B int64  // 8 bytes
    C bool   // 1 byte
}
```

Compiler may place padding between fields.

Conceptual layout:

```text
A _ _ _ _ _ _ _ B B B B B B B B C _ _ _ _ _ _ _
```

This struct can be much larger than expected.

---

## 5. Better Field Ordering

```go
type Good struct {
    B int64
    A bool
    C bool
}
```

Now large aligned field first, smaller fields later.

Often reduces total size.

---

## 6. Why Size Reduction Matters

If one struct saves 8 bytes and you store 10 million objects:

```text
80 MB saved
```

Also improves cache density.

More objects fit in L1/L2/L3 cache.

---

## 7. How to Inspect Size

```go
unsafe.Sizeof(v)
unsafe.Alignof(v)
unsafe.Offsetof(v.Field)
```

Useful for profiling memory layout.

---

## 8. Nested Struct Alignment

```go
type Inner struct {
    X int64
}

type Outer struct {
    A byte
    I Inner
}
```

`Inner` may force alignment padding inside `Outer`.

Nested structs inherit alignment costs.

---

## 9. Arrays vs Slices Memory Layout

### Array

Contiguous fixed-size memory block.

### Slice

Header contains:

* pointer to backing array
* length
* capacity

Elements still stored contiguously in backing array.

---

## 10. Why Contiguous Data Is Fast

Sequential scan of slice:

```go
for i := range arr { ... }
```

CPU prefetcher can load upcoming cache lines efficiently.

Pointer-chasing linked structures often slower.

---

## 11. Cache Line Basics

Typical cache line size:

```text
64 bytes
```

When CPU loads one address, it often loads whole cache line.

So neighboring fields may be fetched “for free”.

---

## 12. Spatial Locality

Fields accessed together should be near each other.

Example:

```go
type User struct {
    ID   int64
    Age  int32
    Flag bool
}
```

Frequently-read hot fields grouped together can improve locality.

---

## 13. Temporal Locality

Recently accessed memory tends to be accessed again soon.

Cache rewards repeated hot data access.

---

## 14. False Sharing (Very Important)

False sharing happens when:

* two goroutines on different CPU cores update different variables
* variables happen to live on same cache line

Even though logically independent, hardware treats line as shared.

This causes cache invalidation traffic.

---

## 15. Example False Sharing

```go
type Counters struct {
    A int64
    B int64
}
```

Core 1 updates `A`, Core 2 updates `B` repeatedly.

If both inside same 64-byte cache line:

```text
cache line ping-pongs between cores
```

Throughput collapses.

---

## 16. Why It Is Called False Sharing

They are not sharing variables logically.

They only share physical cache line.

---

## 17. Fix False Sharing with Padding

```go
type Counters struct {
    A int64
    _ [56]byte
    B int64
}
```

Now `A` and `B` likely on different cache lines.

---

## 18. Per-Worker Sharding Pattern

Instead of one global counter:

```text
counter[NumCPU]
```

Each worker updates own slot.
Aggregate later.

Reduces atomic contention and false sharing.

---

## 19. Hot vs Cold Fields

Split frequently-used and rarely-used fields.

```go
type RequestHot struct {
   ID int64
   Status int32
}

type RequestCold struct {
   Debug string
   Metadata map[string]string
}
```

Keep hot path compact.

---

## 20. Pointer Fields and GC Cost

Pointers increase garbage collector scanning work.

A struct with many pointers may cost more than packed scalar fields.

Sometimes flattening data reduces GC overhead.

---

## 21. Array of Structs vs Struct of Arrays

### Array of Structs (AoS)

```go
[]Point{{X,Y},{X,Y}}
```

Good when accessing full object.

### Struct of Arrays (SoA)

```go
X []float64
Y []float64
```

Good for vectorized scans of one field.

Often better cache behavior in analytics workloads.

---

## 22. Map Memory Notes

Go maps use buckets and overflow structures.

Great for lookup flexibility, but less cache-friendly than flat slices.

For tiny key spaces, slices/arrays may outperform maps.

---

## 23. Channel / Mutex Layout Note

Even synchronization primitives can suffer false sharing if many hot locks are adjacent in arrays.

Example:

```go
locks := make([]sync.Mutex, 1024)
```

Neighbor locks may share cache lines.

---

## 24. Benchmarking Example

Two counters:

### Bad

```go
struct { A int64; B int64 }
```

### Better

```go
struct { A int64; pad[56]; B int64 }
```

Under multicore write load, padded version may be dramatically faster.

---

## 25. Common Production Mistakes

### Ignore struct size in huge slices

Millions of rows magnify waste.

### Random field ordering

n
Creates padding.

### Global hot counters

Cause contention.

### Overusing pointer-rich graphs

Bad locality + GC cost.

---

## 26. Practical Optimization Workflow

1. Measure memory usage
2. Inspect hot structs
3. Reorder fields
4. Benchmark again
5. Check p99 latency / CPU usage
6. Avoid premature micro-optimization

---

## 27. Senior Interview Answer

If asked about false sharing:

> False sharing occurs when independent variables reside on the same cache line and are updated by different cores, causing cache invalidations. Even without lock contention, performance can degrade heavily. We mitigate it with padding, sharding, or redesigning data layout.

---

## 28. Decision Table

| Scenario                | Best Practice                  |
| ----------------------- | ------------------------------ |
| Large slice of structs  | Optimize field order           |
| Hot counters            | Shard or pad                   |
| Hot + cold fields mixed | Split structs                  |
| Scan one numeric column | Consider SoA                   |
| Pointer-heavy objects   | Reduce indirection if possible |

---

## 29. Final Mental Model

Memory performance depends on:

```text
what data is touched
how often
which core touches it
whether it fits cache
whether cores fight over cache lines
```

---

## 30. Senior-Level Summary

Good Go performance often comes from better data layout, not clever syntax.

Understand:

* alignment
* padding
* cache lines
* locality
* false sharing
* GC scanning cost
* data-oriented design

These are major differentiators for high-scale systems.
