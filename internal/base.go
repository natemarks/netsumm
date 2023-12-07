package internal

import (
	"github.com/rs/zerolog"
)

// Remote is an interface for a remote poller
type Remote interface {
	Measure(iterations int, mainLog *zerolog.Logger) PollSet
}

// LoggerWithMapFields adds fields to a logger
func LoggerWithMapFields(mainLog *zerolog.Logger, fields map[string]interface{}) zerolog.Logger {
	var log zerolog.Logger = *mainLog
	for key, value := range fields {
		log = log.With().Interface(key, value).Logger()
	}
	return log
}
