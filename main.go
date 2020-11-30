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

// This function takes the emailText and returns struct of FormattedEmail type by parsing
// email contents and error object
func parseEmails(emailText string) (FormattedEmail, error) {

	if len(emailText) == 0 {
		err := errors.New("Empty emailText received, cannot parse it")
		return FormattedEmail{}, err
	}

	var to, from, date, subject, messageID string
	emailTextSliced := strings.Split(emailText, "\n")
	for index, line := range emailTextSliced {
		switch {
		case strings.HasPrefix(line, "To:") && to == "":
			to = strings.TrimSpace(strings.Split(line, "To:")[1])
		case strings.HasPrefix(line, "From:") && from == "":
			from = strings.TrimSpace(strings.Split(line, "From:")[1])
		case strings.HasPrefix(line, "Message-ID:") && messageID == "":
			messageID = strings.TrimSpace(strings.Split(line, "Message-ID:")[1])
		case strings.HasPrefix(line, "Subject:") && subject == "":
			subject = strings.TrimSpace(strings.Split(line, "Subject:")[1])
			// Subject could be multiple lines, hence make sure to iterate through next
			// couple of indices to check if we need to add those to `subject`
			// when we find a line beginning with `(?m)^[A-Za-z-]*:` we know we are on the next
			// field in the email header and hence our loop ends there.
			// TODO: this has a slight risk if subject is the last line in the header before body begins
			re2 := regexp.MustCompile(`(?m)^[A-Za-z-]*:`)
			for !re2.MatchString(emailTextSliced[index+1]) {
				subject += emailTextSliced[index+1]
				index++
			}

		case strings.HasPrefix(line, "Date:") && date == "":
			date = strings.TrimSpace(strings.Split(line, "Date:")[1])
		}
	}

	fmtEmail := FormattedEmail{
		To:        to,
		From:      from,
		Date:      date,
		Subject:   subject,
		MessageID: messageID,
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
