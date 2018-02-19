package wrapper

import (
	"fmt"
	"os"
	"strings"

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
