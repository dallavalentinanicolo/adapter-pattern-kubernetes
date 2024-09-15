package kubeClient

import (
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// InitializeKubeClient initializes a Kubernetes client.
// if kubeconfigPath is empty, it will attempt to use in-cluster configuration
// if in-cluster configuration fails, it will fall back to the default kubeconfig file
func InitializeKubeClient(kubeconfigPath string) (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error

	// If a kubeconfig path is provided, use it
	if kubeconfigPath != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			return nil, err
		}
	} else {
		// Attempt to use in-cluster config
		config, err = rest.InClusterConfig()
		if err != nil {
			// If in-cluster config fails, fall back to default kubeconfig path
			kubeconfigPath = os.Getenv("HOME") + "/.kube/config"
			config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
			if err != nil {
				return nil, err
			}
		}
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
