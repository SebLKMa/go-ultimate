package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a ticker that ticks every 2 seconds
	ticker := time.NewTicker(1 * time.Second)
	// Ensure the ticker is stopped when main exits to release resources
	defer ticker.Stop()

	// Create a channel to signal the goroutine to quit
	quit := make(chan struct{})

	// Start a goroutine to listen for ticker events
	go func() {
		for {
			select {
			case <-ticker.C: // Received a tick from the ticker
				fmt.Println("Tick1 at:", time.Now())
			case <-quit: // Received a quit signal
				fmt.Println("Ticker1 goroutine stopping.")
				return // Exit the goroutine
			}
		}
	}()

	// Create another ticker that ticks every 2 seconds
	ticker2 := time.NewTicker(2 * time.Second)
	// Ensure the ticker is stopped when main exits to release resources
	defer ticker2.Stop()

	// Start another goroutine to listen for ticker events
	go func() {
		for {
			select {
			case <-ticker2.C: // Received a tick from the ticker
				fmt.Println("Tick2 at:", time.Now())
			case <-quit: // Received a quit signal
				fmt.Println("Ticker2 goroutine stopping.")
				return // Exit the goroutine
			}
		}
	}()

	// Simulate some work in the main goroutine
	fmt.Println("Main goroutine working...")
	time.Sleep(60 * time.Second) // Let the ticker run for a few ticks

	// Signal the ticker goroutine to quit
	close(quit)
	fmt.Println("Main goroutine sent quit signal.")

	// Give the goroutine a moment to process the quit signal
	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine finished.")
}
