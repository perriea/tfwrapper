package wrapper

import (
	"fmt"
)

// FatalError function
func FatalError(err error) {
	if err != nil {
		fmt.Println("Crash ! Please open an issue (https://github.com/perriea/tfwrapper/issues):")
		panic(err)
	}
}

// Error function
func Error(err error) bool {
	if err != nil {
		return true
	}
	return false
}
