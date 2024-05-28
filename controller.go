package main

import (
	"context"
	"fmt"
	clientset "k8s-controller-sample/api/generated/clientset/versioned"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"time"
)

type Controller struct {
	kubeClient   kubernetes.Interface
	sampleClient clientset.Interface
	workqueue    workqueue.RateLimiter
}

func NewController(ctx context.Context,
	kubeClient kubernetes.Interface,
	sampleClient clientset.Interface,
) *Controller {

	return &Controller{
		kubeClient:   kubeClient,
		sampleClient: sampleClient,
		workqueue:    workqueue.NewMaxOfRateLimiter(),
	}
}

func (c *Controller) Run(ctx context.Context, workers int) error {
	defer c.workqueue.ShutDown()
	logger := klog.FromContext(ctx)

	// Start the informer factories to begin populating the informer caches
	logger.Info("Starting Foo controller")

	// Wait for the caches to be synced before starting workers
	logger.Info("Waiting for informer caches to sync")

	if ok := cache.WaitForCacheSync(ctx.Done(), c.deploymentsSynced, c.foosSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	logger.Info("Starting workers", "count", workers)
	// Launch two workers to process Foo resources
	for i := 0; i < workers; i++ {
		go wait.UntilWithContext(ctx, c.runWorker, time.Second)
	}

	logger.Info("Started workers")
	<-ctx.Done()
	logger.Info("Shutting down workers")

	return nil
}
