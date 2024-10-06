package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(10 * time.Second): // Simulate long-running work
		fmt.Println("Task Completed")
	case <-ctx.Done(): // Check if context is canceled
		fmt.Println("Task Canceled:", ctx.Err())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Get the context from the request
	go longRunningTask(ctx)

	fmt.Fprintln(w, "Task Started 1")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
