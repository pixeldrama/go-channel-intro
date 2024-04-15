package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{"task1", "task2", "task3"}

	ch := make(chan string, 3)

	for _, task := range tasks {
		ch <- task
	}

	go worker(ch)

	//time.Sleep(10 * time.Second)
}

// G2
func worker(ch chan string) {
	for {
		data := <-ch
		fmt.Println("Process " + data)
		time.Sleep(time.Second)
	}
}

/*func main() {
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
}*/

/* Signal example */
/*func doWork(ch chan struct{}) {
	// Do some work...
	time.Sleep(2 * time.Second)

	// Send a signal on the channel
	close(ch)
}

func main() {
	// Create a signalling channel with a zero-sized struct
	doneCh := make(chan struct{})

	// Start a worker that will signal completion when it's done
	go doWork(doneCh)

	// Wait for the completion signal
	<-doneCh
	fmt.Println("Work completed!")
}*/
