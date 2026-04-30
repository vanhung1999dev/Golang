# Go Memory Safety & Unsafe Deep Dive

## unsafe.Pointer, Memory Safety, When (and When NOT) to Use It (Senior / FAANG Level)

---

## 1. Overview

Go is designed as a memory-safe language.

Normally Go protects you from:

* invalid pointer arithmetic
* use-after-free bugs
* arbitrary memory reinterpretation
* buffer overflows from raw pointer math
* many classes of undefined behavior common in C/C++

The `unsafe` package is the escape hatch.

It allows low-level memory manipulation outside many normal Go guarantees.

Senior engineers should understand:

* what `unsafe.Pointer` is
* how it differs from typed pointers
  n- why it exists
* real use cases
* GC interaction risks
* portability risks
* when not to use it

---

## 2. What Is `unsafe`?

Standard package:

```go
import "unsafe"
```

It exposes primitives such as:

```go
unsafe.Pointer
unsafe.Sizeof(x)
unsafe.Alignof(x)
unsafe.Offsetof(x.f)
unsafe.Slice(...)
unsafe.String(...)
```

Used by runtime, high-performance libraries, systems code.

---

## 3. What Is `unsafe.Pointer`?

Generic raw pointer type.

Conceptually similar to:

```text
void* in C
```

It can hold pointer to any object.

```go
var x int = 10
p := unsafe.Pointer(&x)
```

---

## 4. Why Normal Pointers Are Safer

Typed pointers:

```go
var p *int
```

Compiler knows:

* pointee type
* alignment expectations
* valid dereference type
* some escape/liveness reasoning

With `unsafe.Pointer`, you bypass much of that type safety.

---

## 5. Common Conversion Pattern

```go
x := int64(5)
p := unsafe.Pointer(&x)
q := (*int64)(p)
fmt.Println(*q)
```

Convert raw pointer back to typed pointer.

---

## 6. Why `unsafe` Exists

Some low-level tasks need capabilities normal Go forbids.

Examples:

* zero-copy conversions
* binary protocol parsing
* memory-mapped files
* calling OS / syscalls
* implementing runtime internals
* highly tuned data structures
* interoperability with C

---

## 7. `unsafe.Sizeof`

```go
var x int64
n := unsafe.Sizeof(x)
```

Returns size in bytes of type/value.

Useful for layout analysis.

---

## 8. `unsafe.Alignof`

```go
unsafe.Alignof(x)
```

Returns alignment requirement.

Useful for struct tuning.

---

## 9. `unsafe.Offsetof`

```go
type User struct {
    ID int64
    Age int32
}

unsafe.Offsetof(User{}.Age)
```

Returns byte offset of field inside struct.

---

## 10. Pointer Arithmetic Pattern

Go forbids direct arithmetic on pointers.

But common unsafe pattern uses `uintptr`:

```go
p := unsafe.Pointer(&arr[0])
next := unsafe.Pointer(uintptr(p) + unsafe.Sizeof(arr[0]))
```

This computes adjacent element address.

Dangerous if misused.

---

## 11. Why `uintptr` Is Dangerous

`uintptr` is integer, not GC-tracked pointer.

If you convert pointer -> uintptr and keep it:

```text
GC may move / reclaim assumptions may break
liveness may be lost
```

Rule: do conversion in one expression when possible.

---

## 12. GC Interaction (Very Important)

Garbage collector tracks real pointers.

If you hide pointer in integer arithmetic carelessly:

```go
u := uintptr(unsafe.Pointer(p))
```

GC may not treat `u` as pointer reference.

This can create subtle bugs.

---

## 13. Correct Short-Lived Pattern

```go
ptr := (*T)(unsafe.Pointer(uintptr(base) + off))
```

Use immediately.
Avoid storing uintptr addresses long-term.

---

## 14. Reinterpreting Memory

Example:

```go
var x uint32 = 0x3f800000
f := *(*float32)(unsafe.Pointer(&x))
```

Same bytes interpreted as different type.

Powerful but risky.

---

## 15. Zero-Copy String / Byte Tricks

Historically libraries used unsafe to convert:

```text
[]byte <-> string
```

without allocation.

Danger:

* strings expected immutable
* modifying backing bytes can violate assumptions

Modern helpers exist but still require caution.

---

## 16. `unsafe.Slice`

Build slice from pointer + length.

```go
s := unsafe.Slice(ptr, n)
```

Useful when interfacing with foreign memory.

Must guarantee memory validity.

---

## 17. `unsafe.String`

Can build string view over bytes/memory.

Must ensure underlying bytes remain valid and unchanged as required.

---

## 18. Memory Safety Risks

Using unsafe can introduce:

* out-of-bounds reads/writes
* misaligned access
* use-after-free style bugs via foreign memory
* stale pointers
* GC visibility bugs
* corrupted structs
* portability issues
* data races still possible

---

## 19. Portability Risks

Assumptions may differ across:

* 32-bit vs 64-bit
* architecture alignment rules
* future Go runtime changes
* endianness in some contexts

Unsafe code may break across platforms.

---

## 20. When To Use `unsafe`

Reasonable cases:

### Performance-Proven Hot Path

After profiling shows allocation/copy bottleneck.

### Systems / Runtime Adjacent Code

Drivers, mmap, syscall wrappers.

### Serialization Libraries

Highly optimized codecs.

### Interop

C memory or OS APIs.

### Data Layout Introspection

Sizeof/Offsetof/Alignof.

---

## 21. When NOT To Use `unsafe`

### To look clever

Never.

### Premature optimization

Use profiler first.

### Ordinary business logic

Use normal Go.

### If maintainers cannot reason about it

n
Complexity cost too high.

### If safe alternative exists with acceptable performance

Prefer safe code.

---

## 22. Channel / Map / Slice Internals via Unsafe?

Possible to inspect runtime internals, but dangerous and version fragile.

Do not depend on private runtime layout in production unless absolutely necessary.

---

## 23. Example Good Use Case

Read struct field offsets for optimized binary encoder:

```go
unsafe.Offsetof(T{}.Field)
```

No dangerous mutation.
Reasonable use.

---

## 24. Example Bad Use Case

Force-casting unrelated structs with different layouts:

```go
(*B)(unsafe.Pointer(&a))
```

May silently corrupt logic.

---

## 25. Performance Reality

Unsafe may remove:

* copies
* allocations
* bounds checks in some patterns
* reflection overhead

But can also worsen:

* maintainability
* debugging time
* future upgrades
* correctness incidents

---

## 26. Safer Alternatives First

Consider before unsafe:

* better algorithms
* preallocation
* pooling carefully
* generics
* bytes.Buffer / strings.Builder
* sync.Pool
* standard library APIs

---

## 27. Review Checklist for Unsafe Code

* Is benchmark gain real?
* Is memory lifetime correct?
* Is alignment valid?
* Any race conditions?
* Works across architectures?
* Clear comments added?
* Tests + fuzzing added?
* Can future maintainers understand it?

---

## 28. Senior Interview Answer

> I use `unsafe` sparingly and only when profiling justifies it. It can remove copies or enable low-level interop, but it bypasses Go's memory safety guarantees. Main risks are GC visibility issues, invalid aliasing assumptions, portability problems, and maintainability cost.

---

## 29. Decision Table

| Scenario                        | Recommendation          |
| ------------------------------- | ----------------------- |
| Need field offsets              | unsafe.Offsetof         |
| Need size/alignment             | unsafe.Sizeof / Alignof |
| Hot zero-copy proven bottleneck | Maybe unsafe with tests |
| Business CRUD service           | Avoid unsafe            |
| Runtime/syscall/mmap work       | Reasonable              |
| Unsure if needed                | Do not use              |

---

## 30. Final Mental Model

`unsafe` trades:

```text
safety + portability + simplicity
for
control + potential speed + low-level access
```

---

## 31. Senior-Level Summary

Use normal Go by default.
Use `unsafe` only intentionally.

Understand:

* pointer semantics
* GC interaction
* lifetime rules
* alignment
* architecture assumptions
* benchmarked benefit vs risk

In strong engineering teams, unsafe code should be rare, reviewed, tested, and justified.
