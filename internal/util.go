package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const configPath = ".deployka-config"

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func getConfigFile() (bool, *os.File) {
	file, err := os.Open(configPath)
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, file
}

// Deletes and recreates the file if it exists
// TODO: It'd be nice if it could be done better
func writeToConfig(username string, url string) error {
	if configFileExists() {
		deleteConfigFile()
	}
	file := createConfigFile()
	defer file.Close()
	_, err := file.WriteString(fmt.Sprintf("username=%v", username))
	_, err = file.WriteString(fmt.Sprintf("url=%s", url))
	return err
}

func getConfigValues(file *os.File) (string, string) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		panic(err)
	}
	username := strings.Split(scanner.Text(), "=")[1]
	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		panic(err)
	}
	url := strings.Split(scanner.Text(), "=")[1]
	return username, url
}

// Only create if it doesn't exist
func createConfigFile() *os.File {
	file, err := os.Create(configPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return file
}

func deleteConfigFile() {
	err := os.Remove(configPath)
	if err != nil {
		fmt.Printf("Error when removing file\n%s", err.Error())
		os.Exit(1)
	}
}

func configFileExists() bool {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
