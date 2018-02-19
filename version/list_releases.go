package version

import (
	"context"

	"github.com/google/go-github/github"
)

var (
	ctx    context.Context
	client *github.Client
)

func init() {
	client = github.NewClient(nil)
	ctx = context.Background()
}

// LastVersion : Check last version of package
func LastVersion() (bool, *github.RepositoryRelease) {
	releases, _, err := client.Repositories.ListReleases(ctx, "perriea", "tfwrapper", nil)
	if err != nil {
		panic(err)
	}

	if *releases[0].TagName == String() {
		return true, releases[0]
	}
	return false, releases[0]
}