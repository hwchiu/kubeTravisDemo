package main

import (
	"testing"

	"github.com/linkernetworks/kubeconfig"
	"github.com/stretchr/testify/assert"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func TestCreatePod(t *testing.T) {
	podName := "test"
	config, err := kubeconfig.Load("")
	assert.NoError(t, err)

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err)

	err = createPod(clientset, podName)
	assert.NoError(t, err)

	pod, err := clientset.CoreV1().Pods("default").Get(podName, metav1.GetOptions{})
	assert.NotNil(t, pod)
	assert.NoError(t, err)

	err = clientset.CoreV1().Pods("default").Delete(podName, &metav1.DeleteOptions{})
	assert.NoError(t, err)

}
