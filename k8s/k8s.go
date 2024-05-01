package k8s

import (
	"context"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Get_Pods(namespace string) []string {
	userHomeDir, _ := os.UserHomeDir()
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
	kubeConfig, _ := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	clientset, _ := kubernetes.NewForConfig(kubeConfig)
	pods, _ := clientset.CoreV1().Pods("kube-system").List(context.Background(), v1.ListOptions{})
	var pods_list []string
	for _, pod := range pods.Items {
		pods_list = append(pods_list, pod.Name)
	}
	return pods_list
}
