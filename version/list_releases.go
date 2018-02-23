package version

import (
	"github.com/google/go-github/github"
)

// LastVersion : Check last version of package
func LastVersion() (bool, *github.RepositoryRelease) {
	releases, _, err := client.Repositories.ListReleases(ctx, "perriea", "tfwrapper", nil)
	if err != nil {
		return true, nil
	}

	if *releases[0].TagName == String() {
		return true, releases[0]
	}
	return false, releases[0]
}
