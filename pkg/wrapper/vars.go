package wrapper

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const path = "terraform.tfvars"

func existVarsConfig() bool {

	info, err = os.Stat(path)
	FatalError(err)

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
	// open file using READ & WRITE permission
	file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
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

func readVarsFile(path string) error {
	// re-open file
	file, err = os.OpenFile(path, os.O_RDWR, 0644)
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
