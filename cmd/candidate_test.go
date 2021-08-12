package cmd

import "testing"

func TestCandidateCommmand(t *testing.T) {
	rootCmd.AddCommand(betaCmd)

	_, err := executeCommandWithTempRepository(rootCmd, "beta")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
