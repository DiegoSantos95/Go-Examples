package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done(): // Check if the context has been canceled
			fmt.Printf("Worker %d: Context canceled\n", id)
			return
		case t := <-ticker.C:
			// Simulate some work
			fmt.Printf("Worker %d: Some Work completed at %s\n", id, t.String())
		}
	}
}

func main() {
	// Create a parent context
	parentCtx := context.Background()

	// Create a context with cancellation
	ctx, cancel := context.WithCancel(parentCtx)

	wg := &sync.WaitGroup{}

	// Launch some goroutines with the context
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(ctx, wg, i)
	}

	// Simulate some time passing
	time.Sleep(3 * time.Second)

	// Cancel the context to stop the goroutines
	cancel()

	// Wait for the goroutines to finish
	wg.Wait()

	fmt.Println("Main: All workers have finished")
}
