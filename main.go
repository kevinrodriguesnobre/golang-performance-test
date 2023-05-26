package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to perform a CPU-intensive task
func performCPUIntensiveTask() {
	sum := 0
	for i := 0; i < 500000000; i++ {
		// Perform some calculations
		sum += rand.Intn(100)
	}
}

func main() {
	// Start the timer
	start := time.Now()

	// Perform the CPU-intensive task
	performCPUIntensiveTask()

	// Calculate the elapsed time
	elapsed := time.Since(start)

	// Print the elapsed time
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
