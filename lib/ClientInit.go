package lib

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

var K8sClient *kubernetes.Clientset

func init() {
	config := &rest.Config{
		Host: "http://127.0.0.1:8009",
	}
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	K8sClient = c
}
