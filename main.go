package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
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

func filePerformance() {
	fileSize := 1024 * 1024 * 100 // 100MB
	filePath := "testfile.txt"

	// Generate a test file with random data
	generateTestFile(filePath, fileSize)

	// Measure file reading performance
	readStartTime := time.Now()
	readFile(filePath)
	readEndTime := time.Now()

	readDuration := readEndTime.Sub(readStartTime)
	fmt.Printf("Read time: %s\n", readDuration)

	// Measure file writing performance
	writeStartTime := time.Now()
	writeFile(filePath)
	writeEndTime := time.Now()

	writeDuration := writeEndTime.Sub(writeStartTime)
	fmt.Printf("Write time: %s\n", writeDuration)

	// Clean up the test file
	os.Remove(filePath)
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

	filePerformance()
}

func generateTestFile(filePath string, fileSize int) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating test file:", err)
		return
	}
	defer file.Close()

	data := make([]byte, fileSize)
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	for i := 0; i < fileSize; i++ {
		data[i] = byte(rand.Intn(256)) // Generate a random byte
	}

	if _, err := file.Write(data); err != nil {
		fmt.Println("Error writing test file data:", err)
		return
	}
}

func readFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}

func writeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	data := []byte("Hello, World!")

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}