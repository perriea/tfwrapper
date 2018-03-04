package wrapper

import (
	"os"

	"github.com/perriea/tfversion/terraform"
)

// Files
const terraformVersionFile = "terraform.tf"
const configFile = "terraform.tfvars"
const binary = "terraform"

// Wrapper
const terraformDefaultVersion = "0.10.8"
const maxRotate = 5

// AWS
const durationSess = 900

var (
	profile *string

	// File info
	file *os.File
	info os.FileInfo

	// Configuration wrapper
	yamlProvider YAMLConfig
	hclTerraform HCLConfig

	err error
)

// YAMLConfig Config YAML
type YAMLConfig struct {
	AWS aws `yaml:"aws"`
}

// Amazon Web Service
type aws struct {
	General     awsGeneral     `yaml:"general"`
	Credentials awsCredentials `yaml:"credentials"`
}

type awsGeneral struct {
	Account string `yaml:"account"`
	Region  string `yaml:"region"`
	Env     string `yaml:"env"`
}

type awsCredentials struct {
	Profile string `yaml:"profile"`
	Role    string `yaml:"role"`
}

// HCLConfig : Config HCL Version
type HCLConfig struct {
	Terraform []terraformVersion
}
type terraformVersion struct {
	Version string `mapstructure:"version"`
}

func init() {
	terraform.InitFolder()
}
