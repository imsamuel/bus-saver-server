package main

import (
	"net/smtp"
	"os"
)

// Message defines the contents of the message sent from the website's contact
// form.
type Message struct {
	Name string
	Email string
	Message string
}

var (
	to = "samuelywp@gmail.com" // receiver email addr
	client = os.Getenv("CLIENT") // client email addr - does the sending
	password = os.Getenv("PASSWORD") // password of client email account
)

// Constructs an email message using the values from the website's contact form
// and sends it to the receiver email addr.
func sendEmail(m Message) error {
	subject := m.Name + " sent you a new message"
	body := "To: " + to + "\r\nSubject: " + subject + "\r\n\r\n" + "Name: " + m.Name + "\r\n\r\n" + "Email: " + m.Email + "\r\n\r\n" + "Message: " + m.Message
	auth := smtp.PlainAuth("", client, password, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, client, []string{to}, []byte(body))
	return err
}
