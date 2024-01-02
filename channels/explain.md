## Channel 

- Channels are the pipes that connect concurrent goroutines. You can send values into channels from one 
goroutine and receive those values into another goroutine.
- **By default sends and receives block until both the sender and receiver are ready**
```
channelName <- value // send data to channel
myVar := <- channelName // receive data from channel
```

</br>

## SendOnly and ReceiveOnly
- `chan <-` // send only
- `<- chan` // Receive only
- Note: `Data was send is shallow copy, not by reference` 