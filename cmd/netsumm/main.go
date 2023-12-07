package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/natemarks/netsumm/version"

	"github.com/natemarks/netsumm/config"
	"github.com/natemarks/netsumm/internal"
	"github.com/rs/zerolog"
)

func getPoller(localIP string, target config.Target, mainLog *zerolog.Logger) internal.Remote {
	switch target.Type {
	case "DNS Lookup":
		return internal.DNSLookup{
			LocalIP:   localIP,
			Server:    target.Destination,
			QueryHost: target.Data,
		}
	case "TCP Connection":
		return internal.TCPConnection{
			LocalIP:  localIP,
			RemoteIP: target.Destination,
			Port:     target.Data,
		}
	case "Ping":
		return internal.Ping{
			LocalIP: localIP,
			Server:  target.Destination,
			Timeout: target.Data,
		}
	default:
		panic(fmt.Errorf("unknown target type: %s", target.Type))
	}
}

func worker(config config.Config, target config.Target, wg *sync.WaitGroup, ch chan<- internal.PollSet, mainLog *zerolog.Logger) {
	defer wg.Done()

	remote := getPoller(config.Source, target, mainLog)
	result := remote.Measure(config.Iterations, mainLog)
	ch <- result
}

func main() {
	config := config.GetConfigFromFile()
	hostname, _ := os.Hostname()
	config.Source = hostname
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stderr).With().Str("version", version.Version).Timestamp().Logger()

	// Create a channel for CustomStruct
	ch := make(chan internal.PollSet)

	// Create a WaitGroup to wait for the goroutines to finish
	var wg sync.WaitGroup

	for _, target := range config.Targets {
		wg.Add(1)
		go worker(config, target, &wg, ch, &logger)
	}

	// Start a goroutine to close the channel after all workers are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Receive and process results from the channel
	for result := range ch {
		summary := internal.GetSummary(result)
		fmt.Println(summary)
	}

	fmt.Println("All workers have finished.")
}
