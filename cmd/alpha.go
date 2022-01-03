package cmd

import (
	"github.com/spf13/cobra"
)

var alphaCmd = &cobra.Command{
	Use:   "alpha",
	Short: "Tags the current ref as a alpha version, the tag will contain all the commits from the last final version.",
	PreRun: func(cmd *cobra.Command, args []string) {
		initGit()
	},
	RunE: alphaAction,
}

func init() {
	rootCmd.AddCommand(alphaCmd)
}

func alphaAction(cmd *cobra.Command, args []string) error {
	// perform alpha bump
	v, err := bumpVersion(repository.LastVersion, Scope, "alpha", Metadata)
	if err != nil {
		return err
	}

	// create tag
	if err := tagAction(repository, v.String(), Output, Force); err != nil {
		return err
	}
	return nil
}
