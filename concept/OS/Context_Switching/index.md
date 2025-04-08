# ✅ 2. Context Switching

Context switching is the process where the CPU stops executing one thread/process and switches to another. <br>

Why? <br>

- To give other threads a chance to run (time slicing).
- When a thread blocks (e.g., waiting for I/O).
- For multitasking and responsiveness.

## 🧠 How OS Switches Between Threads/Processes

The OS scheduler handles this in steps: <br>

**1.Saves the current context:** <br>

- CPU registers (PC, stack pointer, general-purpose registers)
- Program counter (where the thread left off)
- Thread state (ready, running, waiting)

**2.Restores the next thread's context:** <br>

- Loads its saved registers and memory state
- Updates CPU to continue from where the thread was paused

**3.Updates the scheduler's data:** <br>

- Marks the old thread as ready/waiting
- Marks the new one as running

# ⚠️ Cost of Context Switching

![](./image/2025-04-08_16-20.png)

## ⏱️ Estimated cost:

- Switching threads of the same process: ~1–2 microseconds
- Switching between processes: more costly due to memory map and cache disruption

# 🚀 Why It Matters for Go

- OS threads → high context switch cost
- Goroutines → low-cost, user-space context switching

  - Go saves only minimal state (e.g., stack pointer)
  - No syscall or OS involvement

- **That’s why Go can have millions of goroutines, while threads are expensive.**
  ![](./image/ChatGPT%20Image%20Apr%208,%202025,%2004_22_40%20PM.png)
