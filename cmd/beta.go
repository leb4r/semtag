package cmd

import (
	"github.com/spf13/cobra"
)

var betaCmd = &cobra.Command{
	Use:   "beta",
	Short: "Tags the current ref as a beta version, the tag will contain all the commits from the last final version.",
	PreRun: func(cmd *cobra.Command, args []string) {
		initGit()
	},
	RunE: betaAction,
}

func init() {
	rootCmd.AddCommand(betaCmd)
}

func betaAction(cmd *cobra.Command, args []string) error {
	// perform beta bump
	v, err := bumpVersion(repository.LastVersion, Scope, "beta", Metadata)
	if err != nil {
		return err
	}

	// create tag
	if err := tagAction(repository, v.String(), Output, Force); err != nil {
		return err
	}
	return nil
}
