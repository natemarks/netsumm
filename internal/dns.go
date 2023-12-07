package internal

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog"
)

// DNSLookup is a struct that contains information about a DNS lookup
type DNSLookup struct {
	LocalIP   string
	Server    string
	QueryHost string
}

// Measure performs a DNS lookup test for the given hostname
func (d DNSLookup) Measure(iterations int, mainLog *zerolog.Logger) PollSet {
	var result PollSet
	for i := 1; i <= iterations; i++ {
		result = append(result, DigLookup(d.LocalIP, d.Server, d.QueryHost, mainLog))
		time.Sleep(1 * time.Second)
	}
	return result
}

// DNSLookupWithServer performs a DNS lookup test for the given hostname
func DNSLookupWithServer(localIP, server, queryHost string, mainLog *zerolog.Logger) Poll {

	log := LoggerWithMapFields(mainLog, map[string]interface{}{
		"source":      localIP,
		"destination": server,
		"queryHost":   queryHost,
	})
	var poll = Poll{
		Type:        "DNS Lookup",
		Source:      localIP,
		Destination: server,
		StartTime:   time.Now(),
	}
	resolver := &net.Resolver{
		PreferGo: false,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial(network, fmt.Sprintf("%s:53", server))
		},
	}

	result, err := resolver.LookupIPAddr(context.TODO(), queryHost)
	if err != nil {
		log.Error().Err(err).Msg("dns lookup error")
	} else {

		log.Info().Msgf(
			"dns lookup successful: %v (%v ms)",
			result,
			DurationInMilliseconds(poll.StartTime, time.Now()))
	}
	poll.EndTime = time.Now()
	return poll
}

// DigLookup performs a DNS lookup test for the given hostname
func DigLookup(localIP, server, queryHost string, mainLog *zerolog.Logger) Poll {

	log := LoggerWithMapFields(mainLog, map[string]interface{}{
		"source":      localIP,
		"destination": server,
		"queryHost":   queryHost,
	})
	var poll = Poll{
		Type:        "DNS Lookup",
		Source:      localIP,
		Destination: server,
		StartTime:   time.Now(),
	}
	cmdInput := fmt.Sprintf("dig +short %s %s", server, queryHost)
	result := Execute(cmdInput)
	poll.EndTime = time.Now()
	duration := DurationInMilliseconds(poll.StartTime, poll.EndTime)

	if result.ExitCode != 0 {
		log.Error().Msgf("dig error (%v ms)", duration)
	} else {
		log.Info().Msgf(
			"dig successful (%v ms)",
			duration)
	}
	return poll
}
