package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // Simulate long-running work
		fmt.Println("Task Completed")
	case <-ctx.Done(): // Context canceled or deadline exceeded
		fmt.Println("Task Canceled:", ctx.Err())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second) // Set 3-second timeout
	defer cancel()                                                 // Ensure cancel is called to release resources

	go longRunningTask(ctx)

	fmt.Fprintln(w, "Task Started")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
