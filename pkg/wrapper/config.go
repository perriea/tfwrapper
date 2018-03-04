package wrapper

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// readYAMLConfig : Read config
func readYAMLConfig() (YAMLConfig, error) {
	var (
		i         int
		path      string
		dir       string
		config    string
		subfolder []string
		folder    []string
	)

	i = 0
	config = ""
	path = ""

	// Read in five subdirectories
	for i < maxRotate {
		// Check if folder exist
		path = fmt.Sprintf("%s%s", strings.Join(subfolder, ""), "conf")
		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			subfolder = append(subfolder, "../")
		}
		i++
	}

	dir, err = os.Getwd()
	if err != nil {
		return YAMLConfig{}, err
	}

	folder = strings.Split(dir, "/")
	for k := (len(folder) - len(subfolder)); k < len(folder); k++ {
		if (len(folder) - 1) == k {
			config = fmt.Sprintf("%s%s", config, folder[k])
		} else {
			config = fmt.Sprintf("%s%s_", config, folder[k])
		}
	}

	// Read file
	viper.SetConfigName(config)
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	if err = viper.ReadInConfig(); err != nil {
		return YAMLConfig{}, err
	}

	if err = viper.Unmarshal(&yamlProvider); err != nil {
		return YAMLConfig{}, err
	}

	return yamlProvider, nil
}

func validConfigAuth() bool {
	var (
		fileDate int
		fileNow  int
	)

	// File exist or not
	info, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		return false
	}

	// Convert string (hour) to int
	fileDate, err = strconv.Atoi(info.ModTime().Format("20060102150405"))
	FatalError(err)
	fileNow, err = strconv.Atoi(time.Now().Format("20060102150405"))
	FatalError(err)
	if (fileNow - fileDate) > durationSess {
		return false
	}

	return true
}

// writeVarsConfig func : Write config file
func writeAuthConfig(provider string) error {
	var (
		config string
	)

	// open file using READ & CREATE permission
	file, err = os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
	if Error(err) {
		return err
	}
	defer file.Close()

	switch provider {
	case "aws":
		config = fmt.Sprintf("aws_region = \"%s\"\naws_access_key = \"%s\"\naws_secret_key = \"%s\"\naws_token = \"%s\"\nenv = \"%s\"", yamlProvider.Provider.General.Region, os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), os.Getenv("AWS_SESSION_TOKEN"), yamlProvider.Provider.General.Env)
	case "gcp":
		// impossible interpolation: https://github.com/hashicorp/terraform/issues/10059
		config = fmt.Sprintf("gcp_credentials = \"%s\"\ngcp_project = \"%s\"\ngcp_region = \"%s\"\nenv = \"%s\"", yamlProvider.Provider.Credentials.Profile, yamlProvider.Provider.General.Project, yamlProvider.Provider.General.Region, yamlProvider.Provider.General.Env)
	}

	// write some text line-by-line to file
	_, err = file.WriteString(config)
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
