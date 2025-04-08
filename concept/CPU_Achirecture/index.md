# 🔸 1. CPU Cores

Each core can independently execute tasks. <br>

A single-core CPU can do one thing at a time (unless it switches via context switching). <br>

A multi-core CPU (e.g., 4-core) can run multiple threads truly in parallel — one per core. <br>

## ✅ Impact on Scheduling:

More cores = more real parallelism. <br>

Go scheduler tries to run goroutines on available cores using logical "Processors" (P in Go's scheduler). <br>

# 🔸 2. CPU Cache

L1, L2, L3 caches are small, fast memory close to the core. <br>

L1: fastest, per-core. <br>

L2: bigger, slower, often shared between cores. <br>

L3: even larger, shared across all cores. <br>

## ✅ Impact on Scheduling:

Cache locality matters: keeping data near the core that needs it boosts performance. <br>

Go's work-stealing scheduler tries to let goroutines run on the same core to take advantage of hot caches. <br>

Thread migration (moving a goroutine to a different thread or core) can cause cache misses → slowdowns. <br>

# 3. Hardware Threads (Hyper-Threading)

Intel’s Hyper-Threading or AMD’s SMT (Simultaneous Multithreading): each core can run 2 threads at once (logical threads). <br>

Ex: A 4-core CPU with hyper-threading appears as 8 logical processors. <br>

It doesn’t double performance, but helps hide latencies (like waiting for memory). <br>

## ✅ Impact on Scheduling:

Go’s GOMAXPROCS controls how many logical processors Go uses. <br>

More logical CPUs available → more goroutines can run "in parallel." <br>

Go runtime doesn't distinguish between real vs hyper-threaded cores — it just sees "CPU slots". <br>
