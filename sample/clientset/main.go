package main

import (
	"context"
	samplev1 "k8s-controller-sample/api/rollout_sample/v1"
	"k8s-controller-sample/pkg/generated/clientset/versioned"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/clientcmd"
)

func init() {
	scheme := runtime.NewScheme()
	utilruntime.Must(samplev1.AddToScheme(scheme))
}

func main() {
	restConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return
	}
	clientset, err := versioned.NewForConfig(restConfig)
	if err != nil {
		println(err.Error())
		return
	}
	list, err := clientset.ZlwV1().RolloutSamples("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		println(err.Error())
		return
	}
	println(len(list.Items))
}
