package cmd

import (
	"fmt"
	"os"

	"github.com/coreos/go-semver/semver"
	repo "github.com/leb4r/semtag/pkg/git"
	"github.com/leb4r/semtag/pkg/utils"
	"github.com/leb4r/semtag/pkg/version"
)

var (
	repository *repo.Repo
)

func tagAction(repository *repo.Repo, tag string, dryrun bool, force bool) error {
	// get status of worktree
	// exit if --force is not set and worktree contains changes
	if status := repository.Status; len(status) > 0 && !force {
		fmt.Printf("\nThe following changes were found in the worktree:\n\n%s\n--force was not declared. Tag was not created.\n", status)
		os.Exit(1)
	}

	// override the tag to be created if -v flag is set
	if Version != "" {
		tag = Version
	}

	if dryrun {
		fmt.Printf("To be tagged: %s", tag)
	} else {
		if err := repository.CreateTag(tag); err != nil {
			return err
		}
	}

	return nil
}

func bumpVersion(v *version.Version, scope string, preRelease string, metadata string) (*version.Version, error) {
	newVersion := v
	err := newVersion.Bump(scope)
	utils.CheckIfError(err)

	// set pre-release and metadata
	newVersion.Semver.PreRelease = semver.PreRelease(preRelease)
	newVersion.Semver.Metadata = metadata

	return v, nil
}
