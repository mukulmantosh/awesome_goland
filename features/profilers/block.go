package main

import (
	"sync"
	"time"
)

// MutexBlocking Example 1: Goroutines contending on a single mutex
func MutexBlocking(n int) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(10 * time.Millisecond) // hold the lock for a bit
			mu.Unlock()
		}()
	}

	wg.Wait()
}

// ChannelBlocking Example 2: Goroutines blocking on a channel
func ChannelBlocking(n int) {
	ch := make(chan int, 1) // very small buffer
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()

			// Try sending — will block if a channel is full
			ch <- id
			time.Sleep(2 * time.Millisecond)
			<-ch // receive back to free space
		}(i)
	}

	wg.Wait()
}
