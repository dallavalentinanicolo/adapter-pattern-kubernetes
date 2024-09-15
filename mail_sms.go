package main

import (
	"fmt"
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
}

// notification system
type Sms struct {
	phoneNumber string
}

// Implement the SendNotification method for the Mail struct
func (m Mail) SendNotification(message string) {
	fmt.Printf("\nI'm going to send a notification via email to: %s\nMessage:\n%s\n", m.emailAddress, message)
}

// Implement the SendNotification method for the Sms struct
func (s Sms) SendNotification(message string) {
	fmt.Printf("\nI'm going to send a notification via SMS to the phone: %s\nMessage:\n%s\n", s.phoneNumber, message)
}
