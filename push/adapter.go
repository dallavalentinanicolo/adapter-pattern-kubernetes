package push

// Adapter integrates the push notification system with the existing notifier systems
type Adapter struct {
	PushService PushNotification
}

// SendNotification is an implementation method that adapts the interface
// It sends a notification message using the underlying push notification service
func (a Adapter) SendNotification(message string) {
	a.PushService.Alert("Send notification", message, []string{
		"sre-team",
	})
}
