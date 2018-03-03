package wrapper

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// readConfigYAML : Read config
func readConfigYAML() (Configuration, error) {
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
			dir, err = homedir.Dir()
			if err != nil {
				return Configuration{}, err
			}

			// Split path & generate good file
			folder = strings.Split(dir, "/")
			for k := i; k < len(folder); k++ {
				if k == i {
					config = fmt.Sprintf("%s_", folder[k])
				} else if k == (len(folder) - 1) {
					config = fmt.Sprintf("%s%s", config, folder[k])
				} else {
					config = fmt.Sprintf("%s%s_", config, folder[k])
				}
			}

			// Read file
			viper.SetConfigFile(config)
			viper.AddConfigPath(path)
			viper.SetConfigType("yaml")
			if err = viper.ReadInConfig(); err != nil {
				return Configuration{}, err
			}

			if err = viper.Unmarshal(&configuration); err != nil {
				return Configuration{}, err
			}

			return configuration, nil
		}
		i++
	}

	return Configuration{}, err
}

// existVarsConfig :
func existVarsConfig() bool {

	info, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		return false
	}

	fileDate, err := strconv.Atoi(info.ModTime().Format("20060102150405"))
	FatalError(err)

	fileNow, err := strconv.Atoi(time.Now().Format("20060102150405"))
	FatalError(err)

	if (fileNow - fileDate) > durationSess {
		return false
	}

	return true
}

// writeVarsConfig func : Write config file
func writeVarsConfig() error {
	// open file using READ & CREATE permission
	file, err = os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
	if Error(err) {
		return err
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(fmt.Sprintf("aws_region = \"%s\"\naws_access_key = \"%s\"\naws_secret_key = \"%s\"\naws_token = \"%s\"", "eu-west-1", os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), os.Getenv("AWS_SESSION_TOKEN")))
	if Error(err) {
		return err
	}

	// save changes
	err = file.Sync()
	if Error(err) {
		return err
	}
	return nil
}

// readConfigHCL config (terraform.tf)
func readConfigHCL() (string, error) {
	_, err = os.Stat(terraformVersionFile)
	if os.IsNotExist(err) {
		return "", err
	}

	dir, err = homedir.Dir()
	viper.SetConfigFile(terraformVersionFile)
	viper.AddConfigPath(dir)
	viper.SetConfigType("hcl")
	if err = viper.ReadInConfig(); err != nil {
		return "", err
	}

	if err = viper.Unmarshal(&tfConfiguration); err != nil {
		return "", err
	}

	if len(tfConfiguration.Terraform) > 0 {
		if tfConfiguration.Terraform[0].Version != "" {
			return tfConfiguration.Terraform[0].Version, nil
		}
		return "", nil
	}

	return "", nil
}
