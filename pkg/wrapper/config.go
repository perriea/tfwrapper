package wrapper

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// readYAMLConfig : Read config
func readYAMLConfig() (YAMLConfig, error) {
	var (
		i                 int
		path, dir, config string
		folder, subfolder []string
	)

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

	data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.yml", path, config))
	if err != nil {
		return YAMLConfig{}, err
	}

	if err = yaml.Unmarshal(data, &yamlProvider); err != nil {
		return YAMLConfig{}, err
	}

	return yamlProvider, nil
}

func validConfigAuth() bool {
	var (
		info os.FileInfo
		age  int64
	)

	// File exist or not
	info, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		return false
	}

	age = int64(time.Since(info.ModTime()).Seconds())
	if age > durationSess {
		return false
	}

	return true
}

// writeVarsConfig func : Write config file
func writeAuthConfig(provider string) error {
	var (
		config string
		file   *os.File
	)

	// open file using READ & CREATE permission
	file, err = os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	config = fmt.Sprintf("region = \"%s\"\nenv = \"%s\"\naws_access_key = \"%s\"\naws_secret_key = \"%s\"\naws_token = \"%s\"\ngcp_credentials = \"%s\"\ngcp_project = \"%s\"",
		yamlProvider.Terraform.General.Region, yamlProvider.Terraform.General.Env, os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), os.Getenv("AWS_SESSION_TOKEN"), yamlProvider.Terraform.Credentials.Profile, yamlProvider.Terraform.General.Project)

	// write some text line-by-line to file
	_, err = file.WriteString(config)
	if err != nil {
		return err
	}

	return file.Sync()
}
