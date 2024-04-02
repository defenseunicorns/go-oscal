package logging

import (
	"fmt"
	"os"

	"github.com/defenseunicorns/go-oscal/src/pkg/files"
)

// OpenLogFile opens the provided input file or returns an error if the file does not exist.
func OpenLogFile(path string) (file *os.File, err error) {
	err = files.CreateFileDirs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file path: %s", err)
	}
	file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %s", err)
	}
	return file, nil
}
