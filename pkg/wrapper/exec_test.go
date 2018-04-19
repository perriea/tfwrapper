package wrapper

import (
	"testing"
)

// TestSwitchVersion : switch version of Terraform
func TestSwitchVersion(t *testing.T) {
	var (
		i int
	)

	i = 0
	tables := []struct {
		version string
		result  bool
	}{
		{"", false},
		{"0", false},
		{"0.0.1", false},
		{"0.7.0", true},
		{"0.9.1", true},
		{"0.10.9", false},
		{"^&%&^", false},
	}

	for i < len(tables) {
		yamlProvider.Provider.Version.Terraform = tables[i].version
		err = switchVersion()
		if err != nil && !tables[i].result {
			t.Error(err)
		}
		i++
	}
}

// TestExecution : Execute command
func TestExecCmd(t *testing.T) {
	var i int

	i = 0
	tables := []struct {
		args []string
	}{
		{[]string{"plan"}},
		{[]string{"^&*%&^"}},
		{[]string{"fmt"}},
		{[]string{""}},
	}

	for i < len(tables) {
		ExecCmd(tables[i].args, false, true)

		i++
	}
}
