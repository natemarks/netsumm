package poll

import (
	"context"
	"log"
	"net"
	"time"
)

// Get preferred outbound ip of this machine
func GetSourceIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

// PerformDNSLookup performs a DNS lookup for the given hostname
func PerformDNSLookup(localIP, server, hostname string) Poll {
	var poll = Poll{
		Type:        "DNS Lookup",
		Source:      localIP,
		Destination: server,
		StartTime:   time.Now(),
	}
	resolver := &net.Resolver{
		PreferGo: false,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial(network, server)
		},
	}

	resolver.LookupIPAddr(context.Background(), hostname)

	poll.EndTime = time.Now()
	return poll
}
