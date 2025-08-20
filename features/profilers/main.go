package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Starting...")

	// CPU work
	total := 0.0
	for i := 0; i < 50_000_000; i++ {
		total += math.Sqrt(float64(i))
	}
	fmt.Println("CPU result:", total)

	// Memory allocations
	var data [][]byte
	for i := 0; i < 100; i++ {
		buf := make([]byte, 1024*1024) // 1 MB
		data = append(data, buf)
	}
	fmt.Println("Allocated ~100 MB")

	// Keep alive briefly so profiler can attach
	time.Sleep(3 * time.Second)
	fmt.Println("Done.")
}
