package internal

import (
	"log"
	"net"
)

// GetSourceIP returns the IP address of the source
func GetSourceIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

// Remote is an interface for a remote poller
type Remote interface {
	Measure(iterations int) PollSet
}
