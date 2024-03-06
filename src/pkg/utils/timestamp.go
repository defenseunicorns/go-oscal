package utils

import "time"

func GetTimestamp() time.Time {
	// Defaults RFC3339 when marshalling and unmarshalling
	return time.Now()
}
