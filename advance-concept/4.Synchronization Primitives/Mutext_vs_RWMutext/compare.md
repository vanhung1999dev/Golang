# Go Mutex vs RWMutex Deep Dive

## Performance, Contention, Starvation Mode, Internals (Senior / FAANG Level)

---

## 1. Overview

In Go, `sync.Mutex` and `sync.RWMutex` are core synchronization primitives.

Understanding them deeply means knowing:

* When `Mutex` is faster than `RWMutex`
* How contention affects throughput
* What starvation mode solves
* How lock behavior interacts with CPU caches and scheduler
* Real-world alternatives beyond locks

---

## 2. sync.Mutex

Allows only one goroutine to hold the lock at a time.

Best for:

* Short critical sections
* Mixed read/write workloads
* Moderate contention
* Simpler correctness

Often the fastest real-world default.

---

## 3. Mutex Internal Model

Conceptually:

```go
type Mutex struct {
    state int32
    sema  uint32
}
```

`state` stores bits such as:

* locked
  n- woken
* starving
* waiter count

---

## 4. Fast Path Lock()

```go
mu.Lock()
```

Runtime first tries atomic CAS:

```text
unlocked -> locked
```

If successful:

* no parking
* no scheduler involvement
* very fast

---

## 5. Slow Path Under Contention

If already locked:

1. Spin briefly
2. Retry CAS
3. If still fails:

   * enqueue waiter
   * park goroutine
4. Unlock wakes waiter later

---

## 6. Why Spin?

If lock will be released soon, spinning is cheaper than parking/unparking.

Good for tiny critical sections.

Bad if lock is held for long time.

---

## 7. Starvation Mode

Go Mutex has two modes.

### Normal Mode

Optimized for throughput.

New arriving goroutine may steal lock before old waiter wakes.

Pros:

* high throughput
* low handoff latency

Cons:

* old waiters may starve

### Starvation Mode

If waiter waits too long, runtime switches mode.

Unlock directly hands lock to oldest waiter.

Pros:

* fairness
* bounded waiting

Cons:

* lower throughput

---

## 8. Contention Behavior

### Low Contention

CAS succeeds often.

Mutex extremely fast.

### Moderate Contention

Spin helps.

### Heavy Contention

Many goroutines fighting same lock causes:

* cache line bouncing
* park/unpark overhead
* scheduler churn
* starvation mode activation

Throughput may collapse.

---

## 9. Cache Line Bouncing

Lock state lives in shared memory.

Many CPU cores modifying same cache line creates expensive coherence traffic.

Often the real bottleneck under contention.

---

## 10. sync.RWMutex

Allows:

* many readers simultaneously
* one writer exclusively

Good for read-heavy workloads.

---

## 11. RWMutex Internal Concept

Contains logic for:

* reader count
* waiting writers
* writer exclusion
* semaphores

More complex than Mutex.

---

## 12. Read Lock Path

```go
rw.RLock()
```

Usually:

* atomically increment reader count
* proceed if no writer pending

Unlock:

```go
rw.RUnlock()
```

* decrement reader count
* wake waiting writer if last reader exits

---

## 13. Write Lock Path

```go
rw.Lock()
```

Writer must:

1. block new readers
2. wait active readers to drain
3. acquire exclusive ownership

---

## 14. Why RWMutex Can Be Slower

Many engineers assume RWMutex is always faster.

False.

Reasons:

* extra atomic ops
* reader bookkeeping
* shared counter contention
* writer pauses readers

---

## 15. When Mutex Beats RWMutex

Tiny critical section:

```go
value := x
```

RWMutex overhead may exceed benefit.

Mutex can win even in read-heavy workloads.

---

## 16. When RWMutex Wins

Read-heavy workload with meaningful read time:

```go
rw.RLock()
scan structure
rw.RUnlock()
```

Multiple readers can overlap.

---

## 17. Writer Fairness in RWMutex

Classic RW locks can starve writers.

Go avoids this by blocking new readers when writer is waiting.

Flow:

1. existing readers finish
2. writer runs
3. readers resume

---

## 18. Side Effect of Writer Arrival

When writer appears:

* new readers stop
* read latency spikes possible
* throughput may dip temporarily

Important for low-latency systems.

---

## 19. Benchmark Truth

No universal winner.

Depends on:

* read/write ratio
* critical section duration
* core count
* contention level
* latency goals

Always benchmark real workload.

---

## 20. Better Alternatives Sometimes

### Atomic Snapshot / atomic.Pointer

Great for config/state snapshots.

### Sharded Mutex

Split hot lock into many buckets.

### Copy-on-write

Rare writes, many reads.

### sync.Map

Useful for specific read-mostly patterns.

---

## 21. Scheduler Interaction

Blocked lock waiters are parked goroutines.

Unlock may call wake path:

```text
goready(waiter)
```

Then scheduler puts goroutine back into run queue.

---

## 22. Practical Decision Table

| Scenario                        | Best Choice    |
| ------------------------------- | -------------- |
| Short critical section          | Mutex          |
| Mixed reads/writes              | Mutex          |
| 99% reads, expensive reads      | RWMutex        |
| Rare writes, immutable snapshot | atomic.Pointer |
| Hot shared map                  | Sharded Mutex  |
| Unknown workload                | Benchmark      |

---

## 23. Common Production Mistakes

### Use RWMutex everywhere

Wrong.

### Hold lock during I/O

```go
mu.Lock()
db.Query()
mu.Unlock()
```

Terrible for latency.

### One global lock for everything

Creates bottleneck.

### Trust microbenchmarks only

Real contention differs.

---

## 24. Final Mental Model

### Mutex

```text
One lane road
Cheap toll gate
```

### RWMutex

```text
Many reader lanes
Writer temporarily closes highway
More traffic control overhead
```

---

## 25. Senior-Level Summary

Use `Mutex` by default.

Use `RWMutex` only when profiling proves benefit.

Real performance depends on:

* critical section size
* contention level
* read/write ratio
* hardware topology
* fairness requirements
* p99 latency goals
