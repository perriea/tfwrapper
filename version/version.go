package version

import (
	"fmt"
)

// String returns the complete version string, including prerelease
func String() string {
	if prerelease != "" {
		return fmt.Sprintf("%s-%s", Version, prerelease)
	}
	return Version
}
