package wrapper

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var (
	file *os.File
	info os.FileInfo
)

const path = "terraform.tfvars"

func existConfig() bool {

	info, err = os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	fileDate, err := strconv.Atoi(info.ModTime().Format("20060102150405"))
	if isError(err) {
		return false
	}

	fileNow, err := strconv.Atoi(time.Now().Format("20060102150405"))
	if isError(err) {
		return false
	}

	if (fileNow - fileDate) > 3000 {
		return false
	}
	return true
}

func writeConfig() {
	// open file using READ & WRITE permission
	file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(fmt.Sprintf("aws_region = \"%s\"\naws_access_key = \"%s\"\naws_secret_key = \"%s\"\naws_token = \"%s\"", "eu-west-1", os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), os.Getenv("AWS_SESSION_TOKEN")))
	if isError(err) {
		return
	}

	// save changes
	err = file.Sync()
	if isError(err) {
		return
	}
}

func readFile(path string) {
	// re-open file
	file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
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
			isError(err)
			break
		}
	}

	fmt.Println("==> done reading from file")
	fmt.Println(string(text))
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
