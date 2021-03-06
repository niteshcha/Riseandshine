/***
Copyright 2019 Cisco Systems Inc. All rights reserved.

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
	time "time"

	acisnatv1 "github.com/noironetworks/aci-containers/pkg/snatlocalinfo/apis/aci.snat/v1"
	versioned "github.com/noironetworks/aci-containers/pkg/snatlocalinfo/clientset/versioned"
	internalinterfaces "github.com/noironetworks/aci-containers/pkg/snatlocalinfo/informers/externalversions/internalinterfaces"
	v1 "github.com/noironetworks/aci-containers/pkg/snatlocalinfo/listers/aci.snat/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SnatLocalInfoInformer provides access to a shared informer and lister for
// SnatLocalInfos.
type SnatLocalInfoInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SnatLocalInfoLister
}

type snatLocalInfoInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSnatLocalInfoInformer constructs a new informer for SnatLocalInfo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSnatLocalInfoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSnatLocalInfoInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSnatLocalInfoInformer constructs a new informer for SnatLocalInfo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSnatLocalInfoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AciV1().SnatLocalInfos(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AciV1().SnatLocalInfos(namespace).Watch(context.TODO(), options)
			},
		},
		&acisnatv1.SnatLocalInfo{},
		resyncPeriod,
		indexers,
	)
}

func (f *snatLocalInfoInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSnatLocalInfoInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *snatLocalInfoInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&acisnatv1.SnatLocalInfo{}, f.defaultInformer)
}

func (f *snatLocalInfoInformer) Lister() v1.SnatLocalInfoLister {
	return v1.NewSnatLocalInfoLister(f.Informer().GetIndexer())
}
