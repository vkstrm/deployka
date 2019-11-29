package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// MakeRequest : Will POST the passed MessageBody to the API
func MakeRequest(url string, key string, body MessageBody) ([]ResponseBody, error) {
	b := marshal(body)
	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	request.Header.Set("x-api-key", key)
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)

	if err != nil {
		log.Println("API request failed. Check the URL.")
		return nil, fmt.Errorf("failed to fetch status: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Printf("Response returned with status %d", response.StatusCode)
		return nil, fmt.Errorf("POST %s returned %d", url, response.StatusCode)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Failed to read response body")
		return nil, fmt.Errorf("failed to read response: %v", err)
	}
	defer response.Body.Close()

	return unmarshal(bytes)

}

// Marshal the message body to bytes
func marshal(body MessageBody) []byte {
	bytes, err := json.Marshal(body)

	if err != nil {
		log.Println("Failed to marshal")
		panic(err)
	}

	return bytes
}



// Unmarshal the bytes to a response body
func unmarshal(bytes []byte) ([]ResponseBody, error) {
	var resBody []ResponseBody
	err := json.Unmarshal(bytes, &resBody)

	if err != nil {
		log.Println("Failed to unmarshal")
		return nil, fmt.Errorf("failed to unmarshal: %v", err)
	}

	return resBody, nil
}
