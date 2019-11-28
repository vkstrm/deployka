package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var baseURL = ""
var apikey = ""

// FetchPipes : Get the pipelines from the API
func FetchPipes() {
	res := MakeRequest(baseURL, apikey, FetchPipesMessage())
	fmt.Print("Current status of the pipes:\n\n")
	PrintPipes(res)
}

// BlockPipes : Block the passed pipes
func BlockPipes(pipes []string) {
	res := MakeRequest(baseURL, apikey, BlockPipeMessage(pipes))
	fmt.Printf("Now blocking:\n")
	PrintPipes(res)
}

// UnblockPipes : Unblock the passed pipes
func UnblockPipes(pipes []string) {
	res := MakeRequest(baseURL, apikey, UnblockPipeMessage(pipes))
	fmt.Printf("Unblocked:\n")
	PrintPipes(res)
}

// CheckConfig : Check if the config file is present and load the values
func CheckConfig() {
	path := getConfigFilePath()
	exists := pathExists(path)

	if !exists {
		fmt.Println("Config file is missing. Try \"deployka config\".")
		os.Exit(0)
	}

	apikey, baseURL = getConfigValues(path)
}

// Config : Configure the application by setting apikey and Deployka API URL
func Config() {
	configPath := getConfigFilePath()
	exists := pathExists(configPath)
	reader := bufio.NewReader(os.Stdin)
	var apikey string
	var url string

	if exists {
		apikey, url = getConfigValues(configPath)
		fmt.Printf("Configuration exists\n\tAPI key: %s\n\tURL: %s\n", apikey, url)
		fmt.Print("Make new? [y/N] ")
		answer, _ := reader.ReadString('\n')

		if answer != "y\n" {
			fmt.Println("Old configuration kept.")
			return
		}
	} else {
		initConfigFile()
	}

	fmt.Print("API key: ")
	apikey, _ = reader.ReadString('\n')

	fmt.Print("URL: ")
	url, _ = reader.ReadString('\n')

	apikey = strings.TrimSpace(apikey)
	url = strings.TrimSpace(url)

	err := writeToFile(configPath, apikey, url)

	if err != nil {
		log.Printf("Writig configuration failed\n %v \n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("\nNew configuration\nAPI key: %s\nURL: %s\n", apikey, url)
}
