package main

import (
	"context"
	"encoding/json"
	samplev1 "k8s-controller-sample/api/rollout_sample/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	runtimeSchema "k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	scheme := runtime.NewScheme()
	utilruntime.Must(samplev1.AddToScheme(scheme))

	restConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return
	}

	dynamicClient, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		return
	}

	unstructuredObj, err := dynamicClient.Resource(runtimeSchema.GroupVersionResource{
		Group:    samplev1.GroupName,
		Version:  samplev1.Version,
		Resource: "rolloutsamples",
	}).Namespace("default").List(context.Background(), v1.ListOptions{})

	b, _ := json.Marshal(unstructuredObj)
	println(string(b))

	var list = &samplev1.RolloutSampleList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredObj.UnstructuredContent(), &list)
	if err != nil {
		return
	}
	println(len(list.Items))
	//if err != nil {
	//	if k8serrors.IsNotFound(err) {
	//		return nil, nil
	//	}
	//	return nil, err
	//}
	//
	//var rolloutTicketRun *pipelinev1.RolloutTicketRun
	//if err != nil {
	//	return nil, err
	//}
	//if err = json.Unmarshal(b, &rolloutTicketRun);
}
