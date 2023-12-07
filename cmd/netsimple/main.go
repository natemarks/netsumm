package main

import (
	"fmt"
	"os"

	"github.com/natemarks/netsumm/version"

	"github.com/natemarks/netsumm/internal"
	"github.com/rs/zerolog"
)

func getDNSAndQuery() (string, string) {
	args := os.Args[1:] // Exclude the first argument (program name)

	var arg1, arg2 string

	if len(args) > 0 {
		arg1 = args[0]
	}

	if len(args) > 1 {
		arg2 = args[1]
	}

	return arg1, arg2
}
func main() {
	hostname, _ := os.Hostname()
	server, query := getDNSAndQuery()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stderr).With().Str("version", version.Version).Timestamp().Logger()

	dnsLookup := internal.DNSLookup{
		LocalIP:   hostname,
		Server:    server,
		QueryHost: query,
	}
	result := dnsLookup.Measure(20, &logger)
	summary := internal.GetSummary(result)
	fmt.Println(summary)
}
