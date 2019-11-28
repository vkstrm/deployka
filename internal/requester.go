package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// MakeRequest : Will POST the passed MessageBody to the API
func MakeRequest(url string, key string, body MessageBody) []ResponseBody {
	b := marshal(body)
	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
	request.Header.Set("x-api-key", key)
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)

	if err != nil {
		log.Println("API request failed. Check the URL.")
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Printf("Response returned with status %d", response.StatusCode)
		panic(response.StatusCode)
	}

	return grabResponseBody(response.Body)
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

// Get the responsebody from the reader
func grabResponseBody(body io.Reader) []ResponseBody {
	bytes, err := ioutil.ReadAll(body)

	if err != nil {
		log.Println("Failed to read")
		panic(err)
	}

	return unmarshal(bytes)
}

// Unmarshal the bytes to a response body
func unmarshal(bytes []byte) []ResponseBody {
	var resBody []ResponseBody
	err := json.Unmarshal(bytes, &resBody)

	if err != nil {
		log.Println("Failed to unmarshal")
		panic(err)
	}

	return resBody
}
