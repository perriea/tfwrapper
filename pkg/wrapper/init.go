package wrapper

import (
	"github.com/perriea/tfversion/terraform"
)

// Files
const terraformVersionFile = "terraform.tf"
const configFile = "terraform.tfvars"
const binary = "terraform"

// Wrapper
const terraformDefaultVersion = "0.10.8"
const maxRotate = 5

// AWS STS
const durationSess = 900

var (
	profile *string

	// Configuration wrapper
	yamlProvider YAMLConfig
	hclTerraform HCLConfig

	err error
)

// YAMLConfig Config YAML
type YAMLConfig struct {
	Terraform string   `yaml:"terraform"`
	Cloud     string   `yaml:"cloud"`
	Provider  provider `yaml:"provider"`
}

// Amazon Web Service
type provider struct {
	General     general     `yaml:"general"`
	Credentials credentials `yaml:"credentials"`
}

type general struct {
	Region  string `yaml:"region"`
	Env     string `yaml:"env"`
	Project string `yaml:"project"` // only google
	Account string `yaml:"account"`
}

type credentials struct {
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
