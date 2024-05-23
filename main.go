package main

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	master         string
	kubeConfigPath string
)

func init() {
	flag.StringVar(&master, "master", "", "master url")
	flag.StringVar(&kubeConfigPath, "kubeconfig", "", "absolute path to the kubeconfig file")
}

func main() {
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return
	}
	println(clientset.LegacyPrefix)
}
