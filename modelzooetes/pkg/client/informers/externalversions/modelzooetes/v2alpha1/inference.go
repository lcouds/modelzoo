/*
Copyright 2019-2023 TensorChord Inc.

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	time "time"

	modelzooetesv2alpha1 "github.com/lcouds/modelzoo/modelzooetes/pkg/apis/modelzooetes/v2alpha1"
	versioned "github.com/lcouds/modelzoo/modelzooetes/pkg/client/clientset/versioned"
	internalinterfaces "github.com/lcouds/modelzoo/modelzooetes/pkg/client/informers/externalversions/internalinterfaces"
	v2alpha1 "github.com/lcouds/modelzoo/modelzooetes/pkg/client/listers/modelzooetes/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// InferenceInformer provides access to a shared informer and lister for
// Inferences.
type InferenceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2alpha1.InferenceLister
}

type inferenceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewInferenceInformer constructs a new informer for Inference type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInferenceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInferenceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredInferenceInformer constructs a new informer for Inference type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInferenceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TensorchordV2alpha1().Inferences(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TensorchordV2alpha1().Inferences(namespace).Watch(context.TODO(), options)
			},
		},
		&modelzooetesv2alpha1.Inference{},
		resyncPeriod,
		indexers,
	)
}

func (f *inferenceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInferenceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *inferenceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&modelzooetesv2alpha1.Inference{}, f.defaultInformer)
}

func (f *inferenceInformer) Lister() v2alpha1.InferenceLister {
	return v2alpha1.NewInferenceLister(f.Informer().GetIndexer())
}
