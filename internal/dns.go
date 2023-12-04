package internal

import (
	"context"
	"fmt"
	"net"
	"time"
)

// DNSLookup is a struct that contains information about a DNS lookup
type DNSLookup struct {
	LocalIP   string
	Server    string
	QueryHost string
}

// Measure performs a DNS lookup test for the given hostname
func (d DNSLookup) Measure(iterations int) PollSet {
	var result PollSet
	for i := 1; i <= iterations; i++ {
		result = append(result, PerformDNSLookup(d.LocalIP, d.Server, d.QueryHost))
		time.Sleep(1 * time.Second)
	}
	return result
}

// PerformDNSLookup performs a DNS lookup for the given hostname
func PerformDNSLookup(localIP, server, queryHost string) Poll {

	var poll = Poll{
		Type:        "DNS Lookup",
		Source:      localIP,
		Destination: server,
		StartTime:   time.Now(),
	}
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial(network, fmt.Sprintf("%s:53", server))
		},
	}

	resolver.LookupIPAddr(context.Background(), queryHost)
	poll.EndTime = time.Now()
	return poll
}
