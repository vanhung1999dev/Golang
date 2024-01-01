## Go-Routine
- Main function is also routine ()
- Use keyword `go` to init new routine and it will run in `concurrent`
- If main routine die, all child-routine is also die

## Wait-Group
- A WaitGroup is initialized with a counter representing the number of goroutines to wait for.
- The Add() method increments the counter by the given value. This is called by each goroutine to indicate it is running.
- The main goroutine calls Add() to set the initial count, then launches worker goroutines.
- A WaitGroup is typically passed by a pointer to goroutines that need to be waited on.
- The Done() method decrements the counter by 1. Goroutines call this when finished.
- Each worker calls Done() when finished, decrementing the counter.
- The Wait() method blocks until the counter reaches 0, indicating all goroutines have finished.
Main calls Wait() to block until Done() brings counter to 0.
