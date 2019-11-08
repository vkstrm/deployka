package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configPath = "deployka-config"

type configuration struct {
	Username string
	URL      string
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// Deletes and recreates the file if it exists
// TODO: It'd be nice if it could be done better
func writeToConfig(username string, url string) error {
	if configFileExists() {
		deleteConfigFile()
	}

	info := configuration{
		Username: username,
		URL:      url,
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		fmt.Printf("Marshal error: %v", err.Error())
	}

	err = ioutil.WriteFile(configPath, bytes, 0777)
	if err != nil {
		fmt.Printf("WriteFile error: %v", err.Error())
	}
	return err
}

func getConfigValues() (string, string) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Read file error: %v", err.Error())
	}
	var info *configuration
	err = json.Unmarshal(data, &info)
	if err != nil {
		fmt.Printf("Unmarshal error: %v", err.Error())
	}

	return info.Username, info.URL
}

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
