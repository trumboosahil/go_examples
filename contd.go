package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex         // Mutex to protect the shared resource
	cond := sync.NewCond(&mutex) // Condition variable bound to the mutex

	// Shared resource
	ready := false

	// Consumer Goroutine
	go func() {
		mutex.Lock() // Acquire the lock before checking the condition
		for !ready { // While the shared resource (ready) is not true
			fmt.Println("Consumer: Waiting for the producer to signal...")
			cond.Wait() // Wait until the condition is signaled
		}
		fmt.Println("Consumer: Received signal from producer! Proceeding...")
		mutex.Unlock() // Unlock the mutex after the work is done
	}()

	// Simulate a delay to let the consumer start waiting
	time.Sleep(2 * time.Second)

	// Producer Goroutine
	go func() {
		mutex.Lock() // Acquire the lock before modifying the shared resource
		fmt.Println("Producer: Working on something...")
		time.Sleep(2 * time.Second) // Simulate doing some work
		ready = true                // Set the shared resource to true
		fmt.Println("Producer: Signaling the consumer that work is ready!")
		cond.Signal()  // Signal the waiting goroutine
		mutex.Unlock() // Unlock the mutex after signaling
	}()

	// Wait a bit for both goroutines to complete
	time.Sleep(5 * time.Second)
	fmt.Println("Main: All tasks completed.")
}
