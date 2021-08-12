package cmd

import (
	"testing"
)

func TestGetLastCommmand(t *testing.T) {
	rootCmd.AddCommand(getFinalCmd)

	output, err := executeCommandWithTempRepository(rootCmd, "getfinal")
	if output == "" {
		t.Error("Unexpected empty output")
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	checkStringContains(t, output, "Final tagged version: ")
}
