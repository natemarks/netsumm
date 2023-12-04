package config

import "github.com/natemarks/netsumm/poll"

type Target struct {
	Type        string `json:"type"`
	Destination string `json:"destination"`
	Payload     string `json:"payload"`
}

type Config struct {
	Targets []Target `json:"pollerConfig"`
	Source  string   `json:"source"`
}

// GetConfig returns a Config struct
func GetConfig() (config Config) {
	return Config{
		Targets: []Target{
			Target{
				Type:        "DNS Lookup",
				Destination: "8.8.8.8",
				Payload:     "www.google.com",
			},
			Target{
				Type:        "DNS Lookup",
				Destination: "8.8.8.8",
				Payload:     "www.cnn.com",
			},
			Target{
				Type:        "DNS Lookup",
				Destination: "8.8.8.8",
				Payload:     "www.microsoft.com",
			},
		},
		Source: poll.GetSourceIP(),
	}
}
