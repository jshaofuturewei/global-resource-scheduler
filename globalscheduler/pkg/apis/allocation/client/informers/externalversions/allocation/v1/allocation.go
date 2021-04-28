/*
Copyright 2020 Authors of Arktos.

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
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "k8s.io/kubernetes/globalscheduler/pkg/apis/allocation/client/clientset/versioned"
	internalinterfaces "k8s.io/kubernetes/globalscheduler/pkg/apis/allocation/client/informers/externalversions/internalinterfaces"
	v1 "k8s.io/kubernetes/globalscheduler/pkg/apis/allocation/client/listers/allocation/v1"
	allocationv1 "k8s.io/kubernetes/globalscheduler/pkg/apis/allocation/v1"
)

// AllocationInformer provides access to a shared informer and lister for
// Allocations.
type AllocationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.AllocationLister
}

type allocationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
	tenant           string
}

// NewAllocationInformer constructs a new informer for Allocation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAllocationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAllocationInformer(client, namespace, resyncPeriod, indexers, nil)
}

func NewAllocationInformerWithMultiTenancy(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tenant string) cache.SharedIndexInformer {
	return NewFilteredAllocationInformerWithMultiTenancy(client, namespace, resyncPeriod, indexers, nil, tenant)
}

// NewFilteredAllocationInformer constructs a new informer for Allocation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAllocationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return NewFilteredAllocationInformerWithMultiTenancy(client, namespace, resyncPeriod, indexers, tweakListOptions, "system")
}

func NewFilteredAllocationInformerWithMultiTenancy(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc, tenant string) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AllocationV1().AllocationsWithMultiTenancy(namespace, tenant).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) watch.AggregatedWatchInterface {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AllocationV1().AllocationsWithMultiTenancy(namespace, tenant).Watch(options)
			},
		},
		&allocationv1.Allocation{},
		resyncPeriod,
		indexers,
	)
}

func (f *allocationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAllocationInformerWithMultiTenancy(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions, f.tenant)
}

func (f *allocationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&allocationv1.Allocation{}, f.defaultInformer)
}

func (f *allocationInformer) Lister() v1.AllocationLister {
	return v1.NewAllocationLister(f.Informer().GetIndexer())
}