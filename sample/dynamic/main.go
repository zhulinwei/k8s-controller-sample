package main

import (
	"context"
	"encoding/json"
	samplev1 "k8s-controller-sample/api/rollout_sample/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func init () {
	scheme := runtime.NewScheme()
	utilruntime.Must(samplev1.AddToScheme(scheme))
}

func main () {
	restConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return
	}

	dynamicClient, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		return
	}

	unstructuredObj, err := dynamicClient.Resource(samplev1.SchemeGroupVersionResource).Namespace("default").List(context.Background(), v1.ListOptions{})

	b, _ := json.Marshal(unstructuredObj)
	println(string(b))

	var list = &samplev1.RolloutSampleList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredObj.UnstructuredContent(), &list)
	if err != nil {
		return
	}
	println(len(list.Items))
}
