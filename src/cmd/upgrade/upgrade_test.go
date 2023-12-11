package upgrade

import (
	"bytes"
	"log"
	"testing"
)

func TestUpgradeCmd(t *testing.T) {
	t.Parallel()
	log.SetOutput(new(bytes.Buffer))

	t.Run("returns an error if no inputfile is provided", func(t *testing.T) {
		UpgradeCmd.SetArgs([]string{"-f", ""})
		err := UpgradeCmd.Execute()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("returns an error if the inputfile is not a json or yaml file", func(t *testing.T) {
		UpgradeCmd.SetArgs([]string{"-f", "test.txt"})
		err := UpgradeCmd.Execute()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
