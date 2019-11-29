package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	neturl "net/url"
	"os"
	"os/user"
	"path"
	"regexp"
	"strings"
)

type configuration struct {
	Username string
	URL      string
}

// Initialize the configuration file
func initConfigFile() error {
	basepath := getBasePath()

	if !pathExists(basepath) {
		err := os.Mkdir(basepath, 0777)

		if err != nil {
			return fmt.Errorf("failed to create configuration directory: %v", err)
		}
	}

	cpath := path.Join(basepath, "config")

	err := createFile(cpath)
	if err != nil {
		return fmt.Errorf("failed to create configuration file: %v", err)
	}

	return nil
}

// Get home/.deployka
func getBasePath() string {
	const configDir = ".deployka"

	user, err := user.Current()

	if err != nil {
		log.Println("Couldn't get user.")
		panic(err)
	}

	return path.Join(user.HomeDir, configDir)
}

// Get path/.deployka/config
func getConfigFilePath() string {
	return path.Join(getBasePath(), "config")
}

// Write the configuration to the file at path
// Deletes and recreates the file if it exists
func writeToFile(path string, username string, url string) error {
	if pathExists(path) {
		deleteFile(path)
		err := createFile(path)

		if err != nil {
			log.Println("Failed to change configuration.")
			return fmt.Errorf("failed to overwrite existing configuration file: %v", err)
		}
	}

	info := configuration{
		Username: username,
		URL:      url,
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		log.Printf("Marshal error: %v", err.Error())
		return fmt.Errorf("failed during serialization of configuration file: %v", err)
	}

	err = ioutil.WriteFile(path, bytes, 0777)
	if err != nil {
		return fmt.Errorf("failed to write configuration file: %v", err)
	}
	return nil
}

// Get the configuration values from the file at path
func getConfigValues(path string) (string, string) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Printf("Read file error: %v", err.Error())
	}

	var info *configuration
	err = json.Unmarshal(data, &info)

	if err != nil {
		log.Printf("Unmarshal error: %v", err.Error())
	}

	return info.Username, info.URL
}

// Create a file at path
func createFile(filepath string) error {
	_, err := os.Create(filepath)
	return err
}

// Delete file at path
func deleteFile(path string) {
	err := os.Remove(path)

	if err != nil {
		log.Printf("Error when removing file\n%s", err.Error())
		os.Exit(1)
	}
}

// Check if file path has anything
func pathExists(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func parseAPIKey(key string) (string, error) {
	key = strings.TrimSpace(key)
	re := regexp.MustCompile(`^[a-zA-Z0-9]{20,50}$`)

	if !re.MatchString(key) {
		return "", fmt.Errorf("invalid API key format: %s", key)
	}

	return key, nil
}

func parseURL(u string) (string, error) {
	url, err := neturl.Parse(strings.TrimSpace(u))
	if err != nil {
		return "", err
	}

	if url.Scheme != "https" {
		return "", fmt.Errorf("missing or wrong scheme in URL: %s", u)
	}

	return url.String(), nil
}