package internal

import (
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog"
)

// TCPConnection is a struct that represents a TCP connection
type TCPConnection struct {
	LocalIP  string
	RemoteIP string
	Port     string
}

// Measure performs a TCP connection test for the given hostname
func (t TCPConnection) Measure(iterations int, mainLog *zerolog.Logger) PollSet {
	var result PollSet
	for i := 1; i <= iterations; i++ {
		result = append(result, TimeTCPConnection(t.LocalIP, t.RemoteIP, t.Port, mainLog))
		time.Sleep(1 * time.Second)
	}
	return result
}

// TimeTCPConnection performs a DNS lookup for the given hostname
func TimeTCPConnection(localIP, remoteIP, port string, mainLog *zerolog.Logger) Poll {
	log := LoggerWithMapFields(mainLog, map[string]interface{}{
		"source":      localIP,
		"destination": remoteIP,
		"port":        port,
	})
	var poll = Poll{
		Type:        "TCP Connection",
		Source:      localIP,
		Destination: remoteIP,
		StartTime:   time.Now(),
	}
	address := fmt.Sprintf("%s:%s", remoteIP, port)
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		mainLog.Error().Err(err).Msg("tcp connection error")
	} else {
		log.Info().Msgf("tcp connection successful")
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	poll.EndTime = time.Now()
	return poll
}
