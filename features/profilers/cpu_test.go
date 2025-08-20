package main

import (
	"testing"
	"time"
)

// TestCPUProfileDemo creates a clear CPU hotspot for profiling.
// Run this test with the CPU profiler enabled in GoLand.
func TestCPUProfileDemo(t *testing.T) {
	iters := 20_000_000

	// CPU hotspot
	total := CPUWork(iters)
	t.Logf("CPU result: %.3f", total)

	// Short pause so the profiler captures results before test exits
	time.Sleep(1 * time.Second)
}
