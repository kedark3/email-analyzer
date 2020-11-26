package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"      // to log erros
	"net/http" // provides necessary functionality for rest api
	"regexp"
	"strings"
)

// FormattedEmail struct defines 5 JSON fields we need to extract from the Email
type FormattedEmail struct {
	To        string `json:"To"`
	From      string `json:"From"`
	Date      string `json:"Date"`
	Subject   string `json:"Subject"`
	MessageID string `json:"Message-ID"`
}

// This function will take the fieldName and extract it from the Email
func extractFieldFromEmail(fieldName, emailText string) string {
	// regex is compiled of the form "\nTo:.*"
	re := regexp.MustCompile("\n" + fieldName + ".*")
	// if we matched 1 field, then process it
	if matched := re.FindAllString(emailText, 1); len(matched) == 1 {
		// Split on the field name and then choose 1st element from split and trim surrounding spaces
		return strings.TrimSpace(strings.Split(string(matched[0]), fieldName)[1])
	}
	return ""
}

// This function takes the emailText and returns struct of FormattedEmail type by parsing
// email contents and error object
func parseEmails(emailText string) (FormattedEmail, error) {

	if len(emailText) == 0 {
		err := errors.New("Empty emailText received, cannot parse it")
		return FormattedEmail{}, err
	}

	fmtEmail := FormattedEmail{
		To:        extractFieldFromEmail("To:", emailText),
		From:      extractFieldFromEmail("From:", emailText),
		Date:      extractFieldFromEmail("Date:", emailText),
		Subject:   extractFieldFromEmail("Subject:", emailText),
		MessageID: extractFieldFromEmail("Message-ID:", emailText),
	}
	return fmtEmail, nil
}

// Defined function to handle requests for API V1
func emailsV1(w http.ResponseWriter, r *http.Request) {
	// Set content type for all the responses to JSON
	w.Header().Set("Content-type", "application/json")

	switch r.Method {

	case "POST":
		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		// Request body is passed on to parseEmails
		fmtEmail, err := parseEmails(string(body))
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "Please use HTTP POST method with text/plain formatted email"}`))
			break
		}
		// return JSON formatted object
		json.NewEncoder(w).Encode(fmtEmail)
	default:
		// If any other kind of request arrives, return a message and StatusNotFound
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Please use HTTP POST method with text/plain formatted email"}`))
	}
}

func main() {
	// HandleFunc enables us to call correct function
	http.HandleFunc("/api/v1/emails", emailsV1)
	// ListenAndServe starts serving on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
