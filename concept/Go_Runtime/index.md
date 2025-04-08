## ⚙️ What the Go Runtime Does for Goroutines:

### 1. Scheduling

- Maps many G (goroutines) → fewer M (OS threads)

- Uses the G-M-P model

- Handles work stealing, preemption, and fairness

### 2. Stack Management

- Each goroutine starts with a tiny stack (~2 KB).

- Go runtime grows/shrinks stacks dynamically to save memory.

- No manual stack allocation needed.

### 3. Garbage Collection

- Automatically frees unused memory

- Pauses goroutines as needed (very briefly)

- Tuned for low-latency concurrent GC

### 4. Preemption

- Goroutines can be stopped mid-execution (at safe points)

- Allows fairness — prevents long-running goroutines from hogging the CPU

### 5. Timers, Network Polling, Channels

- Handles: time.After, select

- Non-blocking I/O (polling using epoll/kqueue under the hood)

- Blocking/unblocking of goroutines on channels
