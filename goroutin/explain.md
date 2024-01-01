## Go-Routine
- Main function is also routine ()
- Use keyword `go` to init new routine and it will run in `concurrent`
- If main routine die, all child-routine is also die

</br>

## Wait-Group
- A WaitGroup is initialized with a counter representing the number of goroutines to wait for.
- The Add() method increments the counter by the given value. This is called by each goroutine to indicate it is running.
- The main goroutine calls Add() to set the initial count, then launches worker goroutines.
- A WaitGroup is typically passed by a pointer to goroutines that need to be waited on.
- The Done() method decrements the counter by 1. Goroutines call this when finished.
- Each worker calls Done() when finished, decrementing the counter.
- The Wait() method blocks until the counter reaches 0, indicating all goroutines have finished.
Main calls Wait() to block until Done() brings counter to 0.


## Mutex
- A Mutex is a method used as a locking mechanism to ensure that only one Goroutine is accessing the critical section of code at any point of time. This is done to prevent race conditions from happening. Sync package contains the Mutex. Two methods defined on Mutex

- 1.Lock
- 2.Unlock


```
If one Goroutine already has the lock and if a new Goroutine is trying to get the lock, then the new Goroutine will be stopped until the mutex is unlocked. 
```