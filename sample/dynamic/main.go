package main

import (
	"context"
	samplev1 "k8s-controller-sample/api/rollout_sample/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func init() {
	scheme := runtime.NewScheme()
	utilruntime.Must(samplev1.AddToScheme(scheme))
}

func main() {
	// 从kubeconfig中获取k8s配置
	restConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return
	}

	// 初始化dynamicClient客户端
	dynamicClient, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		return
	}

	// 使用dynamicClient获取unstructured对象
	obj, err := dynamicClient.Resource(samplev1.SchemeGroupVersionResource).Namespace("default").List(context.Background(), v1.ListOptions{})
	if err != nil {
		println(err.Error())
		return
	}

	// 将unstructured对象转化为Go结构体
	var list = &samplev1.RolloutSampleList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.UnstructuredContent(), &list)
	if err != nil {
		return
	}
	println(len(list.Items))
}
