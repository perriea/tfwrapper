package version

import (
	"context"

	"github.com/google/go-github/github"
)

var (
	ctx    context.Context
	client *github.Client
)

// Version number that is being run at the moment.
const Version = "0.0.1"

// Repository Git
const Repository = "https://github.com/perriea/tfwrapper"

// Prerelease marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
var Prerelease = ""

func init() {
	client = github.NewClient(nil)
	ctx = context.Background()
}
