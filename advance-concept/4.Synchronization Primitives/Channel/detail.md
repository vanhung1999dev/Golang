# What a Channel Really Is

When you write:

```go
ch := make(chan int, 3)
```

Go runtime creates an internal structure (simplified):

```
channel object:
- buffer memory
- capacity = 3
- count = current items
- send index
- receive index
- lock
- send wait queue
- receive wait queue
- closed flag
```

Think:

```
Mailbox + waiting lists + mutex + state
```

# Internal Shape (Simplified)
```go
hchan {
    buf      []memory
    cap      int
    count    int
    sendx    int
    recvx    int
    closed   bool
    lock

    sendq    waiting senders
    recvq    waiting receivers
}
```

# Two Kinds of Channels

## Unbuffered

```go
ch := make(chan int)
```

Capacity = 0

No storage.

Send must meet receiver immediately.

## Buffered

```go
ch := make(chan int, 3)
```

Has storage.

```
[slot0][slot1][slot2]
```

# SEND Operation Internals

When you do:

```go
ch <- 10
```

Runtime roughly does:

```
1. lock channel
2. if closed => panic
3. if waiting receiver exists:
      hand value directly to receiver
      wake receiver
4. else if buffer has space:
      store into buffer
5. else:
      park sender in sendq
6. unlock
```

## Example: Buffered Send

```go
ch := make(chan int, 2)

ch <- 1
ch <- 2
```

Buffer:

```
[1][2]
count=2 full
```

Next:

```go
ch <- 3
```

Now:

```
buffer full
no space
sender goroutine parks
added to sendq
```

# What “Park” Means

Current goroutine pauses.

Not dead.

Not error.

Stored by runtime as waiting goroutine.

Gwaiting

Later runtime wakes it.

# RECEIVE Operation Internals

When:

```go
x := <-ch
```

Runtime roughly:

```
1. lock channel
2. if buffer has data:
      read item
3. else if waiting sender exists:
      receive directly from sender
      wake sender
4. else if closed:
      return zero value, ok=false
5. else:
      park receiver in recvq
6. unlock
```

## Example Receive from Full Buffer
```go
ch := make(chan int, 2)
ch <- 1
ch <- 2

x := <-ch
```

Now:

```
x = 1
buffer becomes [2]
space freed
```

If sender was blocked waiting to send 3, runtime wakes sender:

```
buffer becomes [2][3]
```

# Direct Handoff (Very Important)

If unbuffered channel:

```go
ch := make(chan int)
```

Sender:

```go
ch <- 10
```

Receiver:

```go
x := <-ch
```

No buffer used.

**Runtime copies value directly sender -> receiver.**

# FIFO Wait Queues

Channel keeps queues:

```
sendq = blocked senders
recvq = blocked receivers
```

Usually oldest waiter gets served first.

This helps fairness.

# CLOSE Operation Internals

When:

```go
close(ch)
```

Runtime:

```
1. lock channel
2. if already closed => panic
3. mark closed=true
4. wake all waiting receivers
5. wake all waiting senders (they panic when resumed)
6. unlock
```

# After Close Behavior

## Receive After Close

If buffered values remain:

**still receive remaining values first**

Then once empty:
```
v, ok := <-ch
```

Returns:

```
zero value, false
```

Example:

```
chan int => 0,false
chan string => "",false
```

# Send After Close

```go
ch <- 5
```

Always:

```
panic: send on closed channel
```

# Why Receivers Wake on Close

Example:

```go
x := <-ch
```

```
If nobody will ever send again, receiver would wait forever

So close wakes them.
```

They get:

```
zero value + ok=false
```

# RANGE Internals
```go
for v := range ch {
    ...
}
```

Compiler turns into roughly:

```
for {
    v, ok := <-ch
    if !ok {
        break
    }
    ...
}
```

```
So range stops only when:

channel closed AND empty
```

### Important

If channel never closes:

```go
for v := range ch
```

can wait forever.

# SELECT Internals

Example:

```go
select {
case x := <-ch1:
case ch2 <- 5:
case <-done:
}
```

Runtime roughly:

```
1. check all cases
2. if one ready -> run it
3. if multiple ready -> choose pseudo-randomly
4. if none ready:
      if default exists -> run default
      else park goroutine waiting on all cases
```

### How Waiting on Many Channels Works

If no case ready:

**Runtime registers goroutine on each channel’s wait queue.**

When one becomes ready:

```
wake goroutine
remove from others
continue selected case
```