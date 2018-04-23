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
	Cloud     string    `yaml:"cloud"`
	Terraform terraform `yaml:"terraform"`
}

type terraform struct {
	Version     string      `yaml:"version"`
	Provider    string      `yaml:"provider"`
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
