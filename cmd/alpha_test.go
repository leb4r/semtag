package cmd

import "testing"

func TestAlphaCommmand(t *testing.T) {
	rootCmd.AddCommand(alphaCmd)

	_, err := executeCommandWithTempRepository(rootCmd, "alpha")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
