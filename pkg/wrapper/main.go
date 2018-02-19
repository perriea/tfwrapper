package wrapper

import "os"

var (
	profile *string

	config string
	dir    string

	data []string
	// Read config
	subfolder []string
	folder    []string

	// File info
	file *os.File
	info os.FileInfo

	// Configuration wrapper
	configuration Configuration

	err error
)

// General Vars
type General struct {
	Account string `yaml:"account"`
	Region  string `yaml:"region"`
}

// Configuration file wrapper
type Configuration struct {
	Aws       Aws       `yaml:"aws"`
	Terraform Terraform `yaml:"terraform"`
}

// Aws Config
type Aws struct {
	General     General     `yaml:"general"`
	Credentials Credentials `yaml:"credentials"`
}

// Terraform Vars
type Terraform struct {
	Vars Vars `yaml:"vars"`
}

// Vars Clients
type Vars struct {
	AwsAccount string `yaml:"aws_account"`
	ClientName string `yaml:"client_name"`
}

// Credentials AWS
type Credentials struct {
	Profile string `yaml:"profile"`
	Role    string `yaml:"role"`
}
