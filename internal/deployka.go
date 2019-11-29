package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var baseURL = ""
var apikey = ""

// FetchPipes : Get the pipelines from the API
func FetchPipes() error {
	res, err := MakeRequest(baseURL, apikey, FetchPipesMessage())
	if err != nil {
		return fmt.Errorf("failed to fetch pipes: %v", err)
	}

	fmt.Print("Current status of the pipes:\n\n")
	PrintPipes(res)
	return nil
}

// BlockPipes : Block the passed pipes
func BlockPipes(pipes []string) error {
	res, err := MakeRequest(baseURL, apikey, BlockPipeMessage(pipes))
	if err != nil {
		return fmt.Errorf("failed to block pipes: %v", err)
	}

	if len(res) == 0 {
		fmt.Println("No pipes blocked")
	} else {
		fmt.Printf("Now blocking:\n")
		PrintPipes(res)
	}

	return nil
}

// UnblockPipes : Unblock the passed pipes
func UnblockPipes(pipes []string) error {
	res, err := MakeRequest(baseURL, apikey, UnblockPipeMessage(pipes))
	if err != nil {
		return fmt.Errorf("failed to unblock pipes: %v", err)
	}

	if len(res) == 0 {
		fmt.Println("No pipes unblocked")
	} else {
		fmt.Printf("Unblocked:\n")
		PrintPipes(res)
	}

	return nil
}

// CheckConfig : Check if the config file is present and load the values
func CheckConfig() {
	path := getConfigFilePath()
	exists := pathExists(path)

	if !exists {
		fmt.Println(`Config file is missing. Run "deployka config".`)
		os.Exit(0)
	}

	apikey, baseURL = getConfigValues(path)
}

// Config : Configure the application by setting apikey and Deployka API URL
func Config(input io.Reader) error {
	configPath := getConfigFilePath()
	exists := pathExists(configPath)
	reader := bufio.NewReader(input)
	var apikey string
	var url string

	if exists {
		apikey, url = getConfigValues(configPath)
		fmt.Printf("Configuration exists\n\tAPI key: %s\n\tURL: %s\n", apikey, url)
		fmt.Print("Make new? [y/N] ")
		answer, _ := reader.ReadString('\n')

		if answer != "y\n" {
			fmt.Println("Old configuration kept.")
			return nil
		}
	} else {
		err := initConfigFile()
		if err != nil {
			return err
		}
	}

	fmt.Print("API key: ")
	apikey, _ = reader.ReadString('\n')
	apikey, err := parseAPIKey(apikey)
	if err != nil {
		return err
	}

	fmt.Print("URL: ")
	url, _ = reader.ReadString('\n')
	url, err = parseURL(url)
	if err != nil {
		return err
	}

	err = writeToFile(configPath, apikey, url)

	if err != nil {
		log.Printf("Writig configuration failed\n %v \n", err.Error())
		return fmt.Errorf("failed to write configuration file: %v", err)
	}

	fmt.Printf("\nNew configuration\nAPI key: %s\nURL: %s\n", apikey, url)
	return nil
}
