package push

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// PushNotification represents a legacy customer system for push notifications,
// which is not used by current notification systems such as OneSignal
type PushNotification struct {
	PodPending string
	Token      string
	ChatID     string
	Message    string
}

// sendTelegramNotification sends a message to a Telegram chat via the Telegram Bot API
func (p *PushNotification) sendTelegramNotification(message string) error {

	telegramAPI := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", p.Token)

	// Create the data for the message (chat ID and message text)
	data := url.Values{}
	data.Set("chat_id", p.ChatID)
	data.Set("text", message)

	// Make a POST request to Telegram's API
	resp, err := http.Post(
		telegramAPI,
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error sending message, status code: %d", resp.StatusCode)
	}

	return nil
}

// Alert sends an alert notification to a list of persons with the given title and message
// It uses the PodPending field to include relevant information in the notification
func (p *PushNotification) Alert(title string, message string, persons []string) {
	notificationMessage := fmt.Sprintf("%s: %s\nPods Pending: %s", title, message, p.PodPending)
	for _, person := range persons {
		fmt.Printf("Notifying %s: %s\n", person, notificationMessage)
		// Send the message via Telegram
		err := p.sendTelegramNotification(notificationMessage)
		if err != nil {
			fmt.Errorf("Failed to send Telegram notification: %v\n", err)
			os.Exit(1)
		}
	}
}
