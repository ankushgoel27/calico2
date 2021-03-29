// Copyright (c) 2021 Tigera, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v3

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	projectcalicov3 "github.com/projectcalico/apiserver/pkg/apis/projectcalico/v3"
	clientset "github.com/projectcalico/apiserver/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/projectcalico/apiserver/pkg/client/informers_generated/externalversions/internalinterfaces"
	v3 "github.com/projectcalico/apiserver/pkg/client/listers_generated/projectcalico/v3"
)

// KubeControllersConfigurationInformer provides access to a shared informer and lister for
// KubeControllersConfigurations.
type KubeControllersConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v3.KubeControllersConfigurationLister
}

type kubeControllersConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewKubeControllersConfigurationInformer constructs a new informer for KubeControllersConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubeControllersConfigurationInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubeControllersConfigurationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredKubeControllersConfigurationInformer constructs a new informer for KubeControllersConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubeControllersConfigurationInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoV3().KubeControllersConfigurations().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoV3().KubeControllersConfigurations().Watch(context.TODO(), options)
			},
		},
		&projectcalicov3.KubeControllersConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubeControllersConfigurationInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubeControllersConfigurationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubeControllersConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&projectcalicov3.KubeControllersConfiguration{}, f.defaultInformer)
}

func (f *kubeControllersConfigurationInformer) Lister() v3.KubeControllersConfigurationLister {
	return v3.NewKubeControllersConfigurationLister(f.Informer().GetIndexer())
}
