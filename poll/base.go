package poll

import (
	"fmt"
	"time"
)

type Poll struct {
	Type        string    `json:"type"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
}

type Summary struct {
	Type        string `json:"type"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	PollCount   int    `json:"pollCount"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
	Avg         int    `json:"avg"`
}

func (s Summary) String() string {
	return fmt.Sprintf("Type: %s\nSource: %s\nDestination: %s\nPollCount: %d\nMin: %d\nMax: %d\nAvg: %d\n", s.Type, s.Source, s.Destination, s.PollCount, s.Min, s.Max, s.Avg)
}

func DurationInMilliseconds(start, end time.Time) int {
	return int(end.Sub(start).Milliseconds())
}

// GetSummary returns a Summary struct
func GetSummary(polls []Poll) Summary {
	var summary = Summary{
		Type:        polls[0].Type,
		Source:      polls[0].Source,
		Destination: polls[0].Destination,
		PollCount:   len(polls),
		Min:         DurationInMilliseconds(polls[0].StartTime, polls[0].EndTime),
		Max:         DurationInMilliseconds(polls[0].StartTime, polls[0].EndTime),
		Avg:         DurationInMilliseconds(polls[0].StartTime, polls[0].EndTime),
	}
	durationSum := 0
	for _, poll := range polls {
		duration := DurationInMilliseconds(poll.StartTime, poll.EndTime)
		if duration < summary.Min {
			summary.Min = duration
		}
		if duration > summary.Max {
			summary.Max = duration
		}
		durationSum += duration
	}
	summary.Avg = durationSum / len(polls)
	return summary
}
