# Using time ticker

Based on  
[A Better Scheduling in Go](https://stephenafamo.com/blog/posts/better-scheduling-in-go)  
[Schedule task at specific time](https://stephenafamo.com/blog/posts/how-to-schedule-task-at-specific-time-in-go)  

## Using goroutines to run time.NewTicker

### example4.go
```
ticker := time.NewTicker(500 * time.Millisecond):
This line creates a new Ticker object that will send a time.Time value on its C channel every 500 milliseconds.

defer ticker.Stop():
It is crucial to call Stop() on the ticker when it's no longer needed to release associated resources. defer ensures this happens when the main function exits.

done := make(chan bool):
A channel done is created to provide a mechanism to signal the goroutine to terminate gracefully.

go func() { ... }():
This launches an anonymous function as a new goroutine.

for { select { ... } }:
Inside the goroutine, an infinite for loop combined with a select statement is used to listen for events.
    case <-done:: If a value is received on the done channel, it indicates that the goroutine should stop. A message is printed, and return exits the goroutine.
    case t := <-ticker.C:: When the ticker fires, the current time t is received from ticker.C, and a message is printed. 

time.Sleep(3 * time.Second):
In the main goroutine, this simulates other work being done while the ticker goroutine runs in the background.

done <- true:
After 3 seconds, a true value is sent on the done channel, signaling the ticker goroutine to stop.

time.Sleep(100 * time.Millisecond):
A small delay is added to allow the goroutine to process the done signal and exit before the main function potentially terminates.
```
