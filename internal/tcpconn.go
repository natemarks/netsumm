package internal

import (
	"fmt"
	"net"
	"time"
)

// TCPConnection is a struct that represents a TCP connection
type TCPConnection struct {
	LocalIP  string
	RemoteIP string
	Port     string
}

// Measure performs a TCP connection test for the given hostname
func (t TCPConnection) Measure(iterations int) PollSet {
	var result PollSet
	for i := 1; i <= iterations; i++ {
		result = append(result, TimeTCPConnection(t.LocalIP, t.RemoteIP, t.Port))
		time.Sleep(1 * time.Second)
	}
	return result
}

// TimeTCPConnection performs a DNS lookup for the given hostname
func TimeTCPConnection(localIP, remoteIP, port string) Poll {

	var poll = Poll{
		Type:        "TCP Connection",
		Source:      localIP,
		Destination: remoteIP,
		StartTime:   time.Now(),
	}
	address := fmt.Sprintf("%s:%s", remoteIP, port)
	conn, _ := net.DialTimeout("tcp", address, 5*time.Second)

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	poll.EndTime = time.Now()
	return poll
}
