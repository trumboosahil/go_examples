package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func longRunningTask(ctx context.Context) error {
	// Simulate a long-running task
	select {
	case <-time.After(5 * time.Second): // Simulate work
		fmt.Println("Task Completed")
		return nil // Task completed successfully
	case <-ctx.Done(): // Context canceled
		fmt.Println("Task Canceled:", ctx.Err())
		return ctx.Err() // Return cancellation error
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Create a channel to signal when the task is done
	done := make(chan error, 1)

	go func() {
		// Simulate a long-running task
		done <- longRunningTask(ctx)
	}()

	select {
	case err := <-done:
		// Handle task completion or cancellation
		if err != nil {
			if ctx.Err() == context.Canceled {
				fmt.Fprintln(w, "Request was canceled by the client")
			} else {
				http.Error(w, "Error occurred while processing the task", http.StatusInternalServerError)
			}
			return
		}

		// Task completed successfully
		fmt.Fprintln(w, "Task completed successfully")
	case <-ctx.Done():
		// Ensure the response is sent when the request is canceled
		w.WriteHeader(http.StatusRequestTimeout)
		fmt.Fprintln(w, "Request was canceled by the client")
		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
