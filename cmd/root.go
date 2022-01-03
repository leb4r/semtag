package cmd

import (
	repo "github.com/leb4r/semtag/pkg/git"
	"github.com/leb4r/semtag/pkg/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// CfgFile is path to config file
	CfgFile string
	// Scope represents what kind of bump to perform
	Scope string
	// Force tag application
	Force bool
	// Output only
	Output bool
	// Version represents the specific tag
	Version string
	// Metadata represents suffix to append to tag
	Metadata string
)

var rootCmd = &cobra.Command{
	Use:   "semtag",
	Short: "Tag your repository according to Semantic Versioning",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&Output, "output", "o", false, "Output the version only, shows the bumped version, but doesn't perform tag.")
	rootCmd.PersistentFlags().BoolVarP(&Force, "force", "f", false, "Forces tagging, even if there are un-staged or un-commited changes.")

	rootCmd.PersistentFlags().StringVarP(&CfgFile, "config", "c", "", "Specifies which config file to use")
	rootCmd.PersistentFlags().StringVarP(&Metadata, "metadata", "m", "", "Specifies the metadata (+BUILD) for the version.")
	rootCmd.PersistentFlags().StringVarP(&Version, "version", "v", "", `Specifies manually the version to be tagged, must be a valid semantic version
				 in the format X.Y.Z where X, Y and Z are positive integers.`)
	rootCmd.PersistentFlags().StringVarP(&Scope, "scope", "s", "patch",
		`The scope that must be increased, can be major, minor or patch.
		The resulting version will match X.Y.Z(-PRERELEASE)(+BUILD)
		where X, Y and Z are positive integers, PRERELEASE is an optional
		string composed of alphanumeric characters describing if the build is
		a release candidate, alpha or beta version, with a number.
		BUILD is also an optional string composed of alphanumeric
		characters and hyphens.
		Setting the scope as 'auto', the script will chose the scope between
		'minor' and 'patch', depending on the amount of lines added (<10% will
		choose patch).`)

	// bind flags to config
	if err := viper.BindPFlag("version", rootCmd.PersistentFlags().Lookup("version")); err != nil {
		utils.ThrowError(err)
	}
	if err := viper.BindPFlag("metadata", rootCmd.PersistentFlags().Lookup("metadata")); err != nil {
		utils.ThrowError(err)
	}
	if err := viper.BindPFlag("scope", rootCmd.PersistentFlags().Lookup("scope")); err != nil {
		utils.ThrowError(err)
	}
}

// initConfig reads in config file
func initConfig() {
	if CfgFile != "" {
		// use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// find home directory.
		home, err := homedir.Dir()
		utils.CheckIfError(err)

		viper.SetConfigName(".semtag")
		viper.SetConfigType("yaml")

		// search for config in $HOME and current directory
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// config file not found; nothing to do here
		} else {
			// config file was found but another error was produced; same
		}
	}
}

func initGit() {
	repository = repo.New(".")
}
