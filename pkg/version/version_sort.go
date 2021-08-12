package version

import (
	"sort"
)

// Versions is a slice of Version
type Versions []*Version

func (s Versions) Len() int {
	return len(s)
}

func (s Versions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Versions) Less(i, j int) bool {
	return s[i].Semver.LessThan(*s[j].Semver)
}

// Sort sorts the given slice of Version
func Sort(versions []*Version) {
	sort.Sort(Versions(versions))
}
