package wrapper

import (
	"fmt"

	"github.com/perriea/tfwrapper/version"
)

func Help() {
	fmt.Printf("\033[1;31mtfwrapper v%s\n\033[0mCommands avalaible :\n", version.String())
	fmt.Print("-------------------------\n\n")
	Action("help", []string{})
}
