package push

import "fmt"

// PushNotification represents a legacy customer system for push notifications,
// which is not used by current notification systems such as OneSignal
type PushNotification struct {
	PodPending string
}

// Alert sends an alert notification to a list of persons with the given title and message
// It uses the PodPending field to include relevant information in the notification
func (p *PushNotification) Alert(title string, message string, persons []string) {
	for _, person := range persons {
		fmt.Printf("Hi %s, something went wrong as there are %s pods pending. Title: %s, Body: %s\n", person, p.PodPending, title, message)
	}
}
