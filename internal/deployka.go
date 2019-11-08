package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var baseURL = ""
var user = ""

// FetchPipes : Get the pipelines from the API
func FetchPipes() {
	res := MakeRequest(baseURL, FetchPipesMessage())
	fmt.Print("Current status of the pipes:\n\n")
	PrintPipes(res)
}

// BlockPipes : Block the passed pipes
func BlockPipes(pipes []string) {
	res := MakeRequest(baseURL, BlockPipeMessage(user, pipes))
	fmt.Printf("Now blocking:\n")
	PrintPipes(res)
}

// UnblockPipes : Unblock the passed pipes
func UnblockPipes(pipes []string) {
	res := MakeRequest(baseURL, UnblockPipeMessage(user, pipes))
	fmt.Printf("Unblocked:\n")
	PrintPipes(res)
}

// CheckConfig : Check if the config file is present and load the values
func CheckConfig() {
	exists := configFileExists()
	if !exists {
		fmt.Println("Config file is missing. Try \"deployka config\".")
		os.Exit(0)
	}

	user, baseURL = getConfigValues()
}

// Config : Configure the application
func Config() {
	exists := configFileExists()
	reader := bufio.NewReader(os.Stdin)
	var username string
	var url string
	if exists {
		username, url = getConfigValues()
		fmt.Printf("Configuration exists\n\tUsername: %s\n\tURL: %s\n", username, url)
		fmt.Print("Make new? [y/N] ")
		answer, _ := reader.ReadString('\n')
		if answer != "y\n" {
			fmt.Println("Old configuration kept.")
			return
		}
	}

	fmt.Print("Username: ")
	username, _ = reader.ReadString('\n')

	fmt.Print("URL: ")
	url, _ = reader.ReadString('\n')

	username = strings.TrimSpace(username)
	url = strings.TrimSpace(url)

	err := writeToConfig(username, url)
	if err != nil {
		fmt.Printf("Writig configuration failed\n %v \n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("\nNew configuration\nUsername: %s\nURL: %s\n", username, url)
}
