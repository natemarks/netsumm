package main

import (
	"fmt"
	"github.com/natemarks/netsumm/poll"
	"sync"
	"time"
)

func worker(sourceIP string, wg *sync.WaitGroup, ch chan<- []poll.Poll) {
	defer wg.Done()

	var result []poll.Poll
	for i := 1; i <= 20; i++ {
		result = append(result, poll.PerformDNSLookup(sourceIP, "8.8.8.8", "google.com"))
		time.Sleep(1 * time.Second)
	}
	ch <- result
}

func main() {
	sourceIP := poll.GetSourceIP()
	// Create a channel for CustomStruct
	ch := make(chan []poll.Poll)

	// Create a WaitGroup to wait for the goroutines to finish
	var wg sync.WaitGroup

	// Start 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(sourceIP, &wg, ch)
	}

	// Start a goroutine to close the channel after all workers are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Receive and process results from the channel
	for result := range ch {
		summary := poll.GetSummary(result)
		fmt.Println(summary)
	}

	fmt.Println("All workers have finished.")
}
