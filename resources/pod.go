package pod

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// PodData holds information about a Kubernetes pod, including its name, namespace, and status.
type PodData struct {
	Name      string
	Namespace string
	Status    string
}

// GetPendingPod retrieves a list of pods in the cluster that are in the "Pending" phase and returns them
// It returns a slice of PodData representing the pending pods, and an error if any occurs during fetching
func GetPendingPod(clientset *kubernetes.Clientset) ([]PodData, error) {
	// List all pods across all namespaces
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("Error fetching pods: %v", err)
	}

	var pendingPods []PodData
	// Filter pods that are in the "Pending" phase
	for _, pod := range pods.Items {
		if pod.Status.Phase == "Pending" {
			pendingPods = append(pendingPods, PodData{
				Name:      pod.Name,
				Namespace: pod.Namespace,
				Status:    string(pod.Status.Phase),
			})
		}
	}
	// Log the total number of pods and the number of pending pods
	return pendingPods, nil
}
