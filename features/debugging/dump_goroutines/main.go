package main

import (
	"fmt"
	"time"
)

// Worker simulates some work in an infinite loop
func worker(id int, stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Printf("Worker %d: stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

// Printer listens to messages and prints them
func printer(messages <-chan string) {
	for msg := range messages {
		fmt.Printf("Printer: %s\n", msg)
	}
}

func main() {
	stop := make(chan struct{})
	message := make(chan string)

	// Start multiple worker goroutines
	for i := 1; i <= 3; i++ {
		go worker(i, stop)
	}

	// Start a goroutine that produces messages
	go func() {
		for i := 1; i <= 5; i++ {
			msg := fmt.Sprintf("message %d", i)
			message <- msg
			time.Sleep(500 * time.Millisecond)
		}
		close(message)
	}()

	// Start a printer goroutine
	go printer(message)

	// Let everything run for a while
	time.Sleep(5 * time.Second)

	// Stop workers gracefully
	close(stop)

	// Give them a moment to exit before the main exits
	fmt.Println("Main: exiting")
}
