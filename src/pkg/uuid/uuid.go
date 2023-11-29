package uuid

import (
	"github.com/google/uuid"
)

// Wrapper functionality for a UUID generator
// Useful for future modification and standardization across projects

func NewUUID() string {
	return uuid.New().String()
}

func NewUUIDWithSource(source string) string {
	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(source)).String()
}
