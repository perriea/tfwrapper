package wrapper

const (
	// Terraform
	binary     = "terraform"
	configFile = "terraform.tfvars"

	// Wrapper
	terraformDefaultVersion = "0.10.8"
	maxRotate               = 5

	// AWS STS duration session max
	durationSess = 900
)

var (
	// Configuration wrapper
	yamlProvider YAMLConfig

	profile *string
	err     error
)

// YAMLConfig Config YAML
type YAMLConfig struct {
	Cloud    string   `yaml:"cloud"`
	Provider provider `yaml:"provider"`
}

type provider struct {
	General     general     `yaml:"general"`
	Credentials credentials `yaml:"credentials"`
	Version     versions    `yaml:"versions"`
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

// Version : Config HCL Version
type versions struct {
	Terraform string `yaml:"terraform"`
	Provider  string `yaml:"provider"`
}
