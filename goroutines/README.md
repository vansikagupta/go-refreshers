# Concurrent programming
 is a paradigm in which the program is a composition of several autonoumous/independent processes/activities.

Goroutines and channels enable concurrent programming in which values are passed between independent activities, goroutines.

**Intuition behind Concurrency**\
* Some calculation occuring in background and spinner at the foreground.
* Networking is a natural domain for concurrency. In the Client Server model, the Server must serve independent clients concurrently.

# Goroutines

Concurrently executing activity is a goroutine in go. 
Similar to threads.

At the start of any program, there is only one goroutine which calls the main function.
It is called the **main goroutine**.\
New goroutines are created when go <func> statement is executed. The <func> gets executed/called in the newly created goroutine.\
There is no programmatic way to stop one goroutine from another.

See Example at spinner/main.go\
The program is a composition of two independent activities, spinning and fibonacci computation. 
Each is written as a separate function but both make progress concurrently.

At this point, we feel the need for our goroutines to be able to interact with one another and synchronize there work.

# Channels
* Channels are communication mechanism by which goroutines can send value to one another.
* A channel has an element type and can work with values of that particular type only. Example: chan int is an channel of integer type values.

We create channels with the built-in make func.\
    ch := make(chan int)\
    make creates the underlying data structure and returns only the **reference**. So channels references to the data structure created by make. Default value would be nil as with other reference types.\
When passing channels to functions, we are passing copy of the reference and thus both will be pointing to the same structure.

* Channels support two operations send and receive.
send: ch <- value\
receive: x := <- ch OR <- ch is also valid, we are discarding the received value.\
***Mind the direction of the arrow***
* Another operation on channels is to close them, to be done by the sender only, using the built-in function close(ch).
* Sending on a closed channel will cause panic.
* Once closed, sender cannot send any more values. Receiver can receive till the last sent value is retrieved and will get default element type value(0 for int) after the last value is fetched.
* Channels created with a non-zero integer capacity argument in make are **buffered channels**.\
ch := make(chan int, 3) // buffered channel\
ch := make(chan int, 0) OR make(chan int) // unbuffered channel

## Unbuffered Channels
**Send and receive are blocking operations.** The sender goroutine will be blocked until receiver receives the value. Receiving goroutine is blocked and can continue only after the receiving goroutines accepts the value sent.

This causes both goroutines to synchronize. Unbuffered channels are also known as synchronous channels.

When using the channel for synchronization alone, usually channel with element type empty struct{} is preferred.

## Pipelines
Sometimes we may want putput from one goroutine to be input to another. We can use channels as such a pipeline.
See 

## References
* Superb explaination: https://notes.shichao.io/gopl/ch8/
