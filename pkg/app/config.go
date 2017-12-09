package app

import (
	"bufio"
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	home string
	err  error
)

func init() {
	home, err = homedir.Dir()
	Error(err)
}

func ReadInput() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Stack Name: ")
	text, err := reader.ReadString('\n')
	Error(err)
	fmt.Println(text)
}

// CreateConfig : Create configuration if not existing
func CreateConfig() {

	viper.SetConfigName("tfwrapper")
	viper.AddConfigPath(home)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	} else {

	}

}
