package main

import (
	"flag"
	"fmt"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/rootxrishabh/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error %s while building config from hostOS\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil{
			fmt.Printf("Error %s building config from cluster\n", err.Error())
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil{
		fmt.Printf("Error %s while building clientset\n", err.Error())
	}

	var c Controller
	c.processItem()
}

git commit -am "Commit message"
