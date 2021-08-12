package version

import (
	"errors"
	"strings"

	"github.com/coreos/go-semver/semver"
)

// Version wraps semver and includes flag for leading `v`
type Version struct {
	LeadingV bool
	Semver   *semver.Version
}

// New creates a new version from a string
func New(version string) *Version {
	return &Version{
		LeadingV: strings.HasPrefix(version, "v"),
		Semver:   semver.New(strings.TrimPrefix(version, "v")),
	}
}

// String returns semantic version string including leading `v`
// if necessary
func (v Version) String() string {
	if v.LeadingV {
		return "v" + v.Semver.String()
	}
	return v.Semver.String()
}

// Bump the version string according to scope
func (v Version) Bump(scope string) error {
	switch scope {
	case "patch":
		v.Semver.BumpPatch()
	case "minor":
		v.Semver.BumpMinor()
	case "major":
		v.Semver.BumpMajor()
	default:
		return errors.New("scope must be one of: patch, minor, or major")
	}
	return nil
}
