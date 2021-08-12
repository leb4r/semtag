package cmd

import (
	"testing"
)

func TestGetFinalCommmand(t *testing.T) {
	rootCmd.AddCommand(getFinalCmd)

	output, err := executeCommandWithTempRepository(rootCmd, "getlast")
	if output == "" {
		t.Error("Unexpected empty output")
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	checkStringContains(t, output, "Last tagged version: ")
}
