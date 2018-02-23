package wrapper

import (
	"testing"
)

// TestExecution : Execute command
func TestExcution(t *testing.T) {
	i := 0
	tables := []struct {
		args   []string
		result bool
	}{
		{[]string{"apply"}, false},
		{[]string{"^&*%&^"}, false},
		{[]string{"fmt"}, true},
		{[]string{""}, false},
	}

	for i < len(tables) {
		err = execution(tables[i].args)
		if err != nil && tables[i].result {
			t.Fatal(err.Error())
		}
		i++
	}
}
