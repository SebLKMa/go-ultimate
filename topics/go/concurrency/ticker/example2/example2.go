package main

import (
	"fmt"
	"time"
)

func main() {
	// This will print the time every 5 seconds
	for theTime := range time.Tick(time.Second * 5) {
		fmt.Println("Task exected at " + theTime.Format("2006-01-02 15:04:05"))
	}
}
