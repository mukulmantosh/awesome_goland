package main

import "math"

// CPUWork burns CPU cycles with math operations.
func CPUWork(iters int) float64 {
	sum := 0.0
	for i := 0; i < iters; i++ {
		x := float64(i % 1000)
		sum += math.Sqrt(x) * math.Cos(x*0.001)
	}
	return sum
}
