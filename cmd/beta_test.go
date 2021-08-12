package cmd

import "testing"

func TestBetaCommmand(t *testing.T) {
	rootCmd.AddCommand(betaCmd)

	_, err := executeCommandWithTempRepository(rootCmd, "final")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
