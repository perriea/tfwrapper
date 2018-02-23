package wrapper

import (
	"fmt"
	"os"

	"github.com/perriea/tfwrapper/version"
)

func DisplayVersion() {
	fmt.Printf("\033[1;31mtfwrapper v%s\033[0m\n\n", version.String())
	test, release := version.LastVersion()
	if !test && release != nil {
		fmt.Printf("\033[1;31mYour version is out of date ! The latest version is %s.\nYou can update by downloading from Github (%s).\033[0m\n\n", *release.TagName, *release.HTMLURL)
	}
}

func Help() {
	Action([]string{})
	os.Exit(0)
}
