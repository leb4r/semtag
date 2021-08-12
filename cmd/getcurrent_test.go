package cmd

import (
	"testing"
)

func TestGetCurrentCommmand(t *testing.T) {
	rootCmd.AddCommand(getCurrentCmd)

	output, err := executeCommandWithTempRepository(rootCmd, "get")
	if output == "" {
		t.Error("Unexpected empty output")
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	checkStringContains(t, output, "Current final version: ")
	checkStringContains(t, output, "Last tagged version: ")
}
