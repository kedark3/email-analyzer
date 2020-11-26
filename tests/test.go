// File to read Sample Emails from the directory and show fomatted output for each email msg
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type formattedEmail struct {
	To        string `json:"To"`
	From      string `json:"From"`
	Date      string `json:"Date"`
	Subject   string `json:"Subject"`
	MessageID string `json:"Message-ID"`
}

func main() {

	// Read from an URL
	url := "http://localhost:8080/api/v1/emails"
	// Test emails directory
	dir := "./sampleEmails/"

	// Read files from the directory
	files, err := ioutil.ReadDir(dir)
	checkError(err)

	// For each file in the directory, make post request using getJSON function
	for _, file := range files {
		fileContents, err := os.Open(dir + file.Name())
		checkError(err)
		fmt.Println(getJSON(url, fileContents))
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getJSON(url string, fileContents io.Reader) string {
	var fmtEmail formattedEmail
	// Make Post request
	response, err := http.Post(url, "text/plain", fileContents)
	checkError(err)
	// Make sure to close the body
	defer response.Body.Close()

	// Read the body
	bytes, err := ioutil.ReadAll(response.Body)
	checkError(err)
	// Unmarshal data from body into fmtEmail struct object
	err = json.Unmarshal(bytes, &fmtEmail)
	checkError(err)
	// Marshal it for the purpose of pretty printing with identation
	bytesJSON, err := json.MarshalIndent(fmtEmail, "", "    ")
	checkError(err)

	return string(bytesJSON)
}
