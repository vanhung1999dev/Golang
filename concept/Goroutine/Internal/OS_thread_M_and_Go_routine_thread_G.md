## 🔧 1. Goroutines (G) – Lightweight User Threads

- A goroutine is a lightweight, user-space thread managed by the Go runtime, not the OS.
- Created using go func() {...} — it’s cheap (few KB of stack) and fast to start.
- Thousands (even millions) of goroutines can exist at once.

### 🧠 Think of G as a task waiting to be executed.

Each G holds: <br>

- A function to run
- Its stack
- Metadata (e.g., status)

## 🧵 2. Machine (M) – OS Threads

- An M is a real OS thread (like those created by pthread in C).
- M executes a G (goroutine), but it can only do so through a P (Processor).

Each M: <br>

- Is scheduled by the OS kernel
- Can be reused across multiple goroutines
- Can block (e.g., on system calls or I/O)

### 🧠 Think of M as the worker doing the actual execution.

## 🛠️ 3. How G and M Work Together (with P)

Go’s scheduler uses the G–M–P model: <br>

```
G = Goroutine (task)
M = OS Thread (execution engine)
P = Processor (scheduler context)

```

### 🔄 The Flow:

- G is ready to run → it gets added to a P’s run queue.
- M (OS thread) picks up P, pulls G from it, and executes it.
- When G yields or finishes:
  - M gets next G from the P's queue.
  - If no work is left, M can steal from other Ps or sleep.

## 🧱 Visual Analogy

- G = Tasks in a task queue
- M = Workers
- P = Workstation with a queue of tasks and tools

Only if a worker (M) has a workstation (P) can it work. Each workstation has a queue of tasks (Gs). If the worker blocks (e.g., waiting for I/O), the workstation is reassigned to another free worker. <br>

## Blocking Example

If a goroutine does a blocking syscall: <br>

- Its M blocks too (since the syscall is in kernel space).
- To avoid wasting CPU, Go detaches the P from that M.
- Another idle M is assigned that P to continue running Gs.

Together, they allow Go to efficiently multiplex thousands of goroutines over a small pool of OS threads, leading to low memory use, fast scheduling, and high performance. <br>
