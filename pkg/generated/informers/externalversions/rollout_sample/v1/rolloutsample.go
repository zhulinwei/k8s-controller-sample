/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	rolloutsamplev1 "k8s-controller-sample/api/rollout_sample/v1"
	versioned "k8s-controller-sample/pkg/generated/clientset/versioned"
	internalinterfaces "k8s-controller-sample/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "k8s-controller-sample/pkg/generated/listers/rollout_sample/v1"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RolloutSampleInformer provides access to a shared informer and lister for
// RolloutSamples.
type RolloutSampleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.RolloutSampleLister
}

type rolloutSampleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRolloutSampleInformer constructs a new informer for RolloutSample type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRolloutSampleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRolloutSampleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRolloutSampleInformer constructs a new informer for RolloutSample type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRolloutSampleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ZlwV1().RolloutSamples(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ZlwV1().RolloutSamples(namespace).Watch(context.TODO(), options)
			},
		},
		&rolloutsamplev1.RolloutSample{},
		resyncPeriod,
		indexers,
	)
}

func (f *rolloutSampleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRolloutSampleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *rolloutSampleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rolloutsamplev1.RolloutSample{}, f.defaultInformer)
}

func (f *rolloutSampleInformer) Lister() v1.RolloutSampleLister {
	return v1.NewRolloutSampleLister(f.Informer().GetIndexer())
}
