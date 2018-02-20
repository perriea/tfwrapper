package wrapper

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// ReadConfig AWS || GCP || Azure
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
			dir, err = os.Getwd()
			Error(err)

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

func existVarsConfig() bool {

	info, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		return false
	}

	fileDate, err := strconv.Atoi(info.ModTime().Format("20060102150405"))
	FatalError(err)

	fileNow, err := strconv.Atoi(time.Now().Format("20060102150405"))
	FatalError(err)

	if (fileNow - fileDate) > 3000 {
		return false
	}
	return true
}

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

func readVarsFile(configFile string) error {
	// re-open file
	file, err = os.OpenFile(configFile, os.O_RDWR, 0644)
	if Error(err) {
		return err
	}
	defer file.Close()

	// read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// break if error occured
		if err != nil && err != io.EOF {
			return err
			//break
		}
	}

	fmt.Println("==> done reading from file")
	fmt.Println(string(text))
	return nil
}
