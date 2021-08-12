package cmd

import (
	"testing"
)

func TestBashCompletion(t *testing.T) {
	rootCmd := rootCmd
	rootCmd.AddCommand(completionCmd)

	output, err := executeCommand(rootCmd, "completion", "bash")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	checkStringContains(t, output, "bash completion for semtag")
}

func TestZshCompletion(t *testing.T) {
	rootCmd.AddCommand(completionCmd)

	output, err := executeCommand(rootCmd, "completion", "zsh")
	if output == "" {
		t.Error("Unexpected empty output")
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	checkStringContains(t, output, "zsh completion for semtag")
}
