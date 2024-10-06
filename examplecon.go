package main

import (
	"fmt"
	"time"
)

func task1(c chan string) {
	start := time.Now()
	fmt.Println("Task 1 started")
	time.Sleep(3 * time.Second) // Simulate task 1 taking 3 seconds
	duration := time.Since(start)
	c <- fmt.Sprintf("Task 1 finished in %v", duration)
}

func task2(c chan string) {
	start := time.Now()
	fmt.Println("Task 2 started")
	time.Sleep(2 * time.Second) // Simulate task 2 taking 2 seconds
	duration := time.Since(start)
	c <- fmt.Sprintf("Task 2 finished in %v", duration)
}

func main() {
	// Capture the start time for the main routine
	mainStart := time.Now()

	// Create channels to receive results from the tasks
	task1Chan := make(chan string)
	task2Chan := make(chan string)

	// Start both tasks concurrently
	go task1(task1Chan)
	go task2(task2Chan)

	// Use select to process whichever finishes first
	completed := 0
	for completed < 2 {
		select {
		case result1 := <-task1Chan:
			fmt.Println(result1)
			completed++
		case result2 := <-task2Chan:
			fmt.Println(result2)
			completed++
		}
	}

	// Total time taken for the main routine to complete
	totalDuration := time.Since(mainStart)
	fmt.Printf("Total time taken: %v\n", totalDuration)
}
