package main

import (
	"fmt"
	kubeClient "go-adapter-pattern/clientk8s"
	metricsProm "go-adapter-pattern/prometheus"
	"go-adapter-pattern/push"
	pod "go-adapter-pattern/resources"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"k8s.io/client-go/kubernetes"
)

// displayPendingPods fetches the pending pods and displays them using an HTML template
func displayPendingPods(clientset *kubernetes.Clientset, w http.ResponseWriter) {
	// Fetch the list of pending pods using the pod package.
	pendingPods, err := pod.GetPendingPod(clientset)
	if err != nil {
		http.Error(w, "Error fetching pending pods", http.StatusInternalServerError)
		log.Printf("Error fetching pending pods: %v", err)
		return
	}

	// Update the Prometheus metric for pending pods
	metricsProm.UpdatePodPendingMetric(len(pendingPods))

	// Load the HTML template from the file
	tmplPath := filepath.Join("templates", "k8s_resource.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		log.Printf("Error loading template: %v", err)
		return
	}

	// Render the template with the list of pending pods
	err = tmpl.Execute(w, pendingPods)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error rendering template: %v", err)
	}
}

const (
	enableSms              = true  // simulate by display terminal message due to cost
	enablePushNotification = false // via Telegram
	enableMail             = false // via Gmail SMTP
	kubeconfigPath         = ""
)

var (
	sms Sms = Sms{phoneNumber: "012345678901011121314151617"}
)

func main() {
	// Specify the kubeconfig path (usually ~/.kube/config for local environments)
	clientset, err := kubeClient.InitializeKubeClient(kubeconfigPath)
	if err != nil {
		log.Fatalf("Failed to initialize Kubernetes client: %v", err)
	}

	// Expose the Prometheus metrics on /metrics endpoint
	go metricsProm.ExposeMetrics()

	// Initialize the previous pending pod count to monitor changes
	var previousPendingPodCount int

	// Start a goroutine to check pending pods every minute
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-ticker.C:
				// Fetch the list of pending pods
				pendingPods, err := pod.GetPendingPod(clientset)
				if err != nil {
					log.Printf("Error fetching pending pods: %v", err)
					continue
				}

				currentPendingPodCount := len(pendingPods)

				// Update the Prometheus metric for pending pods
				metricsProm.UpdatePodPendingMetric(currentPendingPodCount)

				// Check if the pending pod count has changed
				if currentPendingPodCount != previousPendingPodCount {
					// Prepare a notification message based on the number of pending pods
					var message string
					var singolarOrPlural string
					if currentPendingPodCount > 0 {
						if currentPendingPodCount > 1 {
							singolarOrPlural = "are"
						} else {
							singolarOrPlural = "is"
						}
						message = fmt.Sprintf("Hey SRE, there %s %d pending pods in your cluster.", singolarOrPlural, currentPendingPodCount)
					} else if currentPendingPodCount == 0 && previousPendingPodCount > 0 {
						message = "Good news! All pending pods have been resolved."
					}

					// Send notifications based on the flags
					sendNotifications(message, currentPendingPodCount)
				}

				// Update the previous pending pod count
				previousPendingPodCount = currentPendingPodCount
			}
		}
	}()

	// Set up the HTTP server with routing using an anonymous function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			displayPendingPods(clientset, w)
		case "/check":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200 OK"))
		default:
			http.Error(w, "404 page not found", http.StatusNotFound)
		}
	})

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// sendNotifications sends notifications based on the enabled systems
func sendNotifications(message string, currentPendingPodCount int) {
	// Create a slice of Notifier interfaces to send notifications
	var notifiers []Notifier

	// Add SMS, Push Notification, and Email based on the flags
	if enableSms {
		notifiers = append(notifiers, sms)
	}

	if enablePushNotification {
		if telegramBothToken, telegramChatId := os.Getenv("TELEGRAM_BOT_TOKEN"), os.Getenv("TELEGRAM_CHAT_ID"); telegramBothToken != "" && telegramChatId != "" {
			pushService := push.PushNotification{
				PodPending: strconv.Itoa(currentPendingPodCount),
				Token:      telegramBothToken,
				ChatID:     telegramChatId,
			}
			notificatPush := push.Adapter{PushService: pushService}
			notifiers = append(notifiers, notificatPush)
		} else {
			log.Fatal("ERROR, please export TELEGRAM_BOT_TOKEN and TELEGRAM_BOT_TOKEN")
			os.Exit(1)
		}
	}

	if enableMail {
		if emailAddressSender, passwordApp := os.Getenv("EMAIL_SENDER"), os.Getenv("GMAIL_APP_PASSWORD"); emailAddressSender != "" && passwordApp != "" {
			mailNotifier := Mail{
				emailAddress: emailAddressSender,
				password:     passwordApp,
				smtpHost:     "smtp.gmail.com",
				smtpPort:     "587",
				to:           []string{os.Getenv("EMAIL_TO")},
			}
			// Append the mailNotifier inside the if block
			notifiers = append(notifiers, mailNotifier)
		} else {
			log.Fatal("ERROR: EMAIL_SENDER or GMAIL_APP_PASSWORD is not set.")
			os.Exit(1)
		}
	}

	// Send notifications via all enabled notifiers
	for _, notifier := range notifiers {
		notifier.SendNotification(message)
	}
}
