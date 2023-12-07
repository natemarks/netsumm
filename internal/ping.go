package internal

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
)

// Ping is a struct that contains information about a DNS lookup
type Ping struct {
	LocalIP string
	Server  string
	Timeout string // in seconds
}

// Measure performs a DNS lookup test for the given hostname
func (p Ping) Measure(iterations int, mainLog *zerolog.Logger) PollSet {
	var result PollSet
	for i := 1; i <= iterations; i++ {
		result = append(result, ExternalPing(p.LocalIP, p.Server, p.Timeout, mainLog))
		time.Sleep(1 * time.Second)
	}
	return result
}

// ExternalPing ping using the system ping command
func ExternalPing(localIP, server, timeout string, mainLog *zerolog.Logger) Poll {

	log := LoggerWithMapFields(mainLog, map[string]interface{}{
		"source":      localIP,
		"destination": server,
		"timeout":     timeout,
	})
	var poll = Poll{
		Type:        "Ping",
		Source:      localIP,
		Destination: server,
		StartTime:   time.Now(),
	}
	cmdInput := fmt.Sprintf("ping -c 1 -W %s %s", timeout, server)
	result := Execute(cmdInput)
	poll.EndTime = time.Now()
	duration := DurationInMilliseconds(poll.StartTime, poll.EndTime)
	if result.ExitCode != 0 {
		log.Error().Msgf("ping error (%v ms)", duration)
	} else {
		log.Info().Msgf(
			"ping successful (%v ms)",
			duration)
	}

	return poll
}
