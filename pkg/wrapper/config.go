package wrapper

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/perriea/tfwrapper/pkg/app"
)

type Configuration struct {
	Aws       Aws       `yaml:"aws"`
	Terraform Terraform `yaml:"terraform"`
}

type Aws struct {
	General     General     `yaml:"general"`
	Credentials Credentials `yaml:"credentials"`
}

type Terraform struct {
	Vars Vars `yaml:"vars"`
}

type Vars struct {
	AwsAccount string `yaml:"aws_account"`
	ClientName string `yaml:"client_name"`
}

type General struct {
	Account string `yaml:"account"`
	Region  string `yaml:"region"`
}

type Credentials struct {
	Profile string `yaml:"profile"`
	Role    string `yaml:"role"`
}

var (
	configuration Configuration
	dir           string
	subfolder     []string
	folder        []string
	config        string
)

// ReadConfig : test
func readConfig() (bool, Configuration) {
	i := 0
	path := ""

	// Read in five subdirectories
	for i < 5 {
		// Check if folder exist
		path = fmt.Sprintf("%s%s", strings.Join(subfolder, ""), "conf")
		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			subfolder = append(subfolder, "../")
		} else {
			// Get current directory
			dir, err := os.Getwd()
			app.Error(err)

			// Split path & generate good file
			folder = strings.Split(dir, "/")
			for k := (i + 2); k < len(folder); k++ {
				if k == (i + 2) {
					config = fmt.Sprintf("%s_", folder[k])
				} else if k == (len(folder) - 1) {
					config = fmt.Sprintf("%s%s", config, folder[k])
				} else {
					config = fmt.Sprintf("%s%s_", config, folder[k])
				}
			}

			// Read file
			viper.SetConfigName(config)
			viper.AddConfigPath(path)
			if err := viper.ReadInConfig(); err != nil {
				fmt.Printf("%s", err)
				return false, configuration
			}
			err = viper.Unmarshal(&configuration)
			if err != nil {
				fmt.Printf("%s", err)
				return false, configuration
			}

			return true, configuration
		}
		i++
	}

	return false, configuration
}
