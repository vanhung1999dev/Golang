## 🔁 Preemptive vs Cooperative Scheduling

![](./image/2025-04-08_16-28.png)

- 🧠 Go used to rely on cooperative scheduling, but now uses preemptive scheduling (since Go 1.14), so it can preempt long-running goroutines even if they don’t yield.

## ⚙️ Common CPU Scheduling Algorithms

### 🔄 1. Round-Robin (RR)

Each thread gets a fixed time slice (quantum). <br>

After that, it's preempted and moved to the end of the queue. <br>

#### ✅ Simple, fair

❌ Context switching overhead can be high <br>

🧠 Go’s scheduler is heavily inspired by round-robin with enhancements like work-stealing. <br>

### 🎖️ 2. Priority Scheduling

Each thread/process is assigned a priority level. <br>

Higher priority runs first. <br>

✅ Good for real-time systems <br>

❌ Risk of starvation for lower-priority tasks (can be solved with aging) <br>

### ⌛ 3. Shortest Job First (SJF) / Shortest Remaining Time First (SRTF)

Runs the task with the least estimated time left. <br>

SRTF = Preemptive version of SJF. <br>

✅ Optimal average turnaround <br>
❌ Requires knowledge of job lengths — hard to predict <br>

### 🔃 4. Multilevel Queue

Processes are divided into multiple queues by type (e.g., I/O-bound vs CPU-bound). <br>

Each queue can use a different scheduling algorithm. <br>

## 🧠 What Go Uses

Go's scheduler is a hybrid: <br>

Based on M:N scheduling (many goroutines on few OS threads). <br>

Inspired by Round-Robin with:<br>

Work stealing<br>

Local run queues per processor<br>

Preemptive scheduling (since Go 1.14+)<br>

![](./image/ChatGPT%20Image%20Apr%208,%202025,%2004_31_47%20PM.png)
