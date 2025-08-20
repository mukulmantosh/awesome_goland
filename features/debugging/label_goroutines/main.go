package main

import (
	"context"
	"fmt"
	"runtime/pprof"
	"time"
)

func main() {
	ctx := context.Background()

	// Worker goroutine with label
	go pprof.Do(ctx, pprof.Labels("goroutine", "worker"), func(ctx context.Context) {
		for {
			fmt.Println("Worker: processing task...")
			time.Sleep(1 * time.Second) // simulate work
		}
	})

	// Printer goroutine with label
	go pprof.Do(ctx, pprof.Labels("goroutine", "printer"), func(ctx context.Context) {
		for {
			fmt.Println("Printer: logging status...")
			time.Sleep(2 * time.Second) // simulate slower work
		}
	})

	// Keep the program alive so goroutines run
	time.Sleep(8 * time.Second)
	fmt.Println("Main: exiting")
}
