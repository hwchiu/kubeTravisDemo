package main

import (
	"fmt"
	"log"

	"github.com/linkernetworks/kubeconfig"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createPod(clientset *kubernetes.Clientset, name string) error {
	_, err := clientset.CoreV1().Pods("default").Create(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  name,
					Image: "nginx",
				},
			},
		},
	})
	return err
}

func main() {
	fmt.Println("vim-go")

	config, err := kubeconfig.Load("")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	pods, _ := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
	for _, v := range pods.Items {
		fmt.Println(v.Name)
	}
}
