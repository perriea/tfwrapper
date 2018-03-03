package wrapper

import (
	"testing"
)

// TestExecution : Execute command
func TestExcution(t *testing.T) {
	i := 0
	tables := []struct {
		args []string
	}{
		{[]string{"apply"}},
		{[]string{"^&*%&^"}},
		{[]string{"fmt"}},
		{[]string{""}},
	}

	for i < len(tables) {
		execution(tables[i].args)
		i++
	}
}
