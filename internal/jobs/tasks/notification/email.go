package notification

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Payload struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func DecodeToTask(msg string, task interface{}) (err error) {
	decodedstg, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return
	}
	msgBytes := []byte(decodedstg)
	err = json.Unmarshal(msgBytes, task)
	if err != nil {
		return
	}
	return
}

func SendMail(b64payload string) (bool, error) {

	payload := Payload{}
	DecodeToTask(b64payload, &payload)

	from := "<your-google-email>@gmail.com"
	pass := "<your-google-email-app-password>"
	to := payload.Email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + payload.Subject + "\n\n" +
		payload.Body

	fmt.Println(msg, pass)

	// if err != nil {
	// 	log.Printf("smtp error: %s", err)
	// 	return false, nil
	// }

	return true, nil
}
