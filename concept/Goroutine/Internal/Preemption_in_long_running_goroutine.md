# How does Go handle preemption in long-running goroutines?

Go originally used a cooperative scheduler, but now it supports preemption, especially to handle long-running goroutines that might block other work. Here's a breakdown of how Go handles preemption, why it matters, and how it works under the hood. <br>

## 🧠 What is Preemption?

- Preemption is when the Go runtime forcibly stops a goroutine to give time to others — even if the current goroutine hasn’t voluntarily yielded.
- This is crucial to prevent a goroutine from hogging the CPU, especially in tight loops or compute-heavy tasks.

## ⚙️ How Go Handles Preemption

### 📜 History:

- Before Go 1.14: Only cooperative scheduling — a goroutine had to reach a safe point (e.g., function call, channel op, select) to be paused.

- Since Go 1.14: Introduced asynchronous preemption — the runtime can interrupt long-running goroutines mid-execution.

## 🧩 Preemption Mechanism (Since Go 1.14+)

### 🔄 Overview:

- Go runtime timer loop checks if a goroutine is running too long.
- It sets a flag on the goroutine’s stack (g.preempt).
- It triggers a preemption signal (e.g., SIGPREEMPT).

- At the next safe point (or function prologue), the goroutine:
  - Sees the preempt flag.
  - Yields control voluntarily to let the scheduler run something else.

**🧠 Safe points = function calls, memory allocations, channel sends/receives, etc.** <br>

## 🔬 Example: Long-running loop

```
func busyLoop() {
	for {
		// some tight loop with no I/O or blocking
	}
}

```

Before Go 1.14, this would block the scheduler, starving other goroutines. <br>

After Go 1.14, the runtime: <br>

- Marks it for preemption.
- Triggers a signal.
- Yields execution when possible, letting other goroutines run.

## 🏃‍♂️ What Is a “Safe Point”?

A safe point is a specific location in the code where the Go runtime can safely stop, inspect, or preempt a goroutine. These points ensure that the runtime can manipulate the goroutine (e.g., for garbage collection or scheduling) without corrupting program state: <br>

- Stop a goroutine
- Inspect its stack
- Move it out of the CPU and run something else

Examples include: <br>

- Function prologues (start of a new function)
- Channel operations
- Select statements
- runtime.Gosched() (manual yield)

## 📍 Why Safe Points Are Needed

Stopping a goroutine at any arbitrary machine instruction can leave the program in an inconsistent or dangerous state (e.g., in the middle of an instruction, with stack/registers half-updated). Safe points solve this by letting the runtime only intervene at known-good points.

## 🧰 Developer Tooling

= runtime.Gosched() — manually yield

- GODEBUG=schedtrace=1000 — trace scheduler activity
- pprof and trace — see which goroutines are running too long

## 🧠 What Happens at a Safe Point?

At a safe point, the Go runtime can: <br>

- Preempt the goroutine (scheduler)
- Scan stack frames (GC)
- Move the goroutine to another M or P (scheduling)
- Safely walk the stack for profiling

## 🔄 Where Do Safe Points Happen?

### Function Entry (prologue):

- Go inserts checks at the beginning of many functions.
- Example: A tight loop like for {} is not preemptable, but for { doSomething() } is, because the function call creates a safe point.

### Channel operations:

- Send/receive on a channel yields a safe point.

### Select statements

### Memory allocations:

- Calls to new(), make(), or anything that grows the heap.

### runtime.Gosched():

- Manual yield — tells the scheduler “I'm at a safe point, switch me out if needed”.

### Syscalls and I/O operations

### Defer, panic, recover: these internally involve stack manipulations and are also safe points.

```
// ❌ Not preemptable — tight loop
func tightLoop() {
	for {
		// nothing here — no safe point
	}
}

// ✅ Preemptable — function call introduces safe point
func preemptableLoop() {
	for {
		helper() // function call = safe point
	}
}

```

## 📦 Compiler-Inserted Checks

- When compiling your Go program, the compiler may insert instructions (in function prologues) to:
- Check the g.preempt flag.
- If set, yield control.

## 🧵 How This Helps Go Scheduler

When Go wants to stop or switch a goroutine: <br>

- It sets the g.preempt flag.
- When the goroutine hits a safe point:

  - It sees the flag.
  - It yields to the scheduler.

- If it never reaches a safe point, it can’t be stopped — which is why tight loops without function calls can cause goroutine starvation.
