# Some Facts about Golang Channel

## Use Case Communication between goroutines

The amazing `goroutines`[^1] needs something for data exchange and this is where Golang Channels come into play. They
are more or less behave like FIFO circular queue.

```go
go func (){
	time.sleep(time.Minute)
	print("world")
}

print("hello ")

```
### Writing and Reading with Channels

```go
ch := make(chan int, 1)
ch <- 3   // sends 3 to the channel
v := <-ch // receives 3 from the channel and assigns it to v
```

### Channel Types and there Use Cases:

- Buffered Channels (Asynchronous Communication): `ch := make(chan int, 3)`
- Unbuffered Channels (Synchronous Communication): `ch := make(chan string)`
- Signalling Channels (Async Channels with Zero-sized Elements): `ch:= make(chan struct{})`

### Select Keyword

```go
package main

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

### Facts:

- Channels are blocking for reading and writing
- Can be directed `recvCh := make(<-chan int)`
- Do not share memory (data is copied to the channel struct)
- Have to be closed or are blocking forever
- Deadlocks possible, but could be detected

## Resources
- [Exploring the Depths of Golang Channels: A Comprehensive Guide](https://medium.com/@ravikumar19997/exploring-the-depths-of-golang-channels-a-comprehensive-guide-53e1a97cafe6)
- [Golang: channels implementation](https://dmitryvorobev.blogspot.com/2016/08/golang-channels-implementation.html)
- [Default Selection](https://go.dev/tour/concurrency/6)

[^1]: Powerpoint: Java Virtual Threads and goroutines (Campus Day)
