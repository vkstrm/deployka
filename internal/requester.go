package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// MakeRequest : Will POST the passed MessageBody to the API
func MakeRequest(url string, body MessageBody) []ResponseBody {
	response, err := http.Post(url, "application/json", bytes.NewBuffer(marshal(body)))
	errorCheck(err)
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Printf("Response returned with status %d", response.StatusCode)
		panic(response.StatusCode)
	}

	return grabResponseBody(response.Body)
}

func marshal(body MessageBody) []byte {
	bytes, err := json.Marshal(body)
	errorCheck(err)

	return bytes
}

func grabResponseBody(body io.Reader) []ResponseBody {
	bytes, err := ioutil.ReadAll(body)
	errorCheck(err)

	return unmarshal(bytes)
}

func unmarshal(bytes []byte) []ResponseBody {
	var resBody []ResponseBody
	err := json.Unmarshal(bytes, &resBody)
	errorCheck(err)

	return resBody
}
