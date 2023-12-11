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

	t.Run("returns an error if the inputfile does not exist", func(t *testing.T) {
		UpgradeCmd.SetArgs([]string{"-f", "test.json"})
		err := UpgradeCmd.Execute()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("returns an error if no version is provided", func(t *testing.T) {
		UpgradeCmd.SetArgs([]string{"-f", "../testdata/catalog.json"})
		err := UpgradeCmd.Execute()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("returns an error if it fails to create a validator", func(t *testing.T) {
		invalidVersion := "1.0.0"
		UpgradeCmd.SetArgs([]string{"-f", "../testdata/catalog.json", "-v", invalidVersion})
		err := UpgradeCmd.Execute()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
