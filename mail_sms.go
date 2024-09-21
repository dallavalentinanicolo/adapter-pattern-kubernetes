package main

import (
	"fmt"
	"log"
	"net/smtp"
)

// Create a common model for all derived structs.
// All structs that decide to implement the Notifier interface
// must implement the SendNotification method.

// This model is used for basic notification systems like
// email and SMS.
type Notifier interface {
	SendNotification(message string)
}

// AWS SNS Email notification system
type Mail struct {
	emailAddress string
	password     string
	smtpHost     string
	smtpPort     string
	to           []string
}

// notification system
type Sms struct {
	phoneNumber string
}

// Implement the SendNotification method for the Mail struct
func (m Mail) SendNotification(message string) {
	// Subject and Body of the email
	subject := "Subject: Alerting Pod pending on your cluster\n"
	body := message
	fullMessage := []byte(subject + "\n" + body)

	// Set up authentication information
	auth := smtp.PlainAuth("", m.emailAddress, m.password, m.smtpHost)

	// Send the email
	err := smtp.SendMail(m.smtpHost+":"+m.smtpPort, auth, m.emailAddress, m.to, fullMessage)
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	fmt.Printf("\nEmail notification sent to: %v\nMessage:\n%s\n", m.to, message)
}

// Implement the SendNotification method for the Sms struct
func (s Sms) SendNotification(message string) {
	fmt.Printf("\nI'm going to send a notification via SMS to the phone: %s\nMessage:\n%s\n", s.phoneNumber, message)
}
