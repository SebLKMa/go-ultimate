package main

import (
	"context"
	"fmt"
	"time"
)

func waitUntil(ctx context.Context, until time.Time) {
	timer := time.NewTimer(time.Until(until))
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Println("Task executed at:", time.Now())
		return
	case <-ctx.Done():
		return
	}
}

func main() {
	// our context, for now we use context.Background()
	ctx := context.Background()

	// when we want to wait till
	until, _ := time.Parse(time.RFC3339, "2025-10-02T02:30:00+00:00")

	// and now we wait
	waitUntil(ctx, until)

	// Do what ever we want....
	/*
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			fmt.Println("Task executed at:", time.Now())
		}
	*/
}
