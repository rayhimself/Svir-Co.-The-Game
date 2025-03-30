package k8s

import (
	"context"
	"os"
	"path/filepath"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)


func GerCubeConfig() *restclient.Config {
	userHomeDir, _ := os.UserHomeDir()
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
	fmt.Println(kubeConfigPath)
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		fmt.Printf("error getting Kubernetes config: %v\n", err)
		os.Exit(1)
	}
	return kubeConfig
}

func GetDeploy(namespace string, kubeConfig *restclient.Config) *appsv1.DeploymentList {
	clientset, _ := kubernetes.NewForConfig(kubeConfig)
	deploys, _ := clientset.AppsV1().Deployments(namespace).List(context.Background(), v1.ListOptions{})
	return deploys
}

func DeleteDeploy (namespace, delpoyName string, kubeConfig *restclient.Config) {
	clientset, _ := kubernetes.NewForConfig(kubeConfig)
	deletePolicy := v1.DeletePropagationForeground
	clientset.AppsV1().Deployments(namespace).Delete(context.Background(), delpoyName, v1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}

func CreateDeploy (namespace, delpoyName, palnt string, kubeConfig *restclient.Config) {
	deployment := &appsv1.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name: delpoyName,
			Labels: map[string]string{
				"plant": palnt,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &v1.LabelSelector{
				MatchLabels: map[string]string{
					"plant": palnt,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Labels: map[string]string{
						"plant": palnt,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	clientset, _ := kubernetes.NewForConfig(kubeConfig)
	clientset.AppsV1().Deployments(namespace).Create(context.Background(), deployment, v1.CreateOptions{})

}