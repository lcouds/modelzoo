/*
Copyright 2023 TensorChord Inc.

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/lcouds/modelzoo/ingress-operator/pkg/apis/modelzetes/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InferenceIngressLister helps list InferenceIngresses.
// All objects returned here must be treated as read-only.
type InferenceIngressLister interface {
	// List lists all InferenceIngresses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.InferenceIngress, err error)
	// InferenceIngresses returns an object that can list and get InferenceIngresses.
	InferenceIngresses(namespace string) InferenceIngressNamespaceLister
	InferenceIngressListerExpansion
}

// inferenceIngressLister implements the InferenceIngressLister interface.
type inferenceIngressLister struct {
	indexer cache.Indexer
}

// NewInferenceIngressLister returns a new InferenceIngressLister.
func NewInferenceIngressLister(indexer cache.Indexer) InferenceIngressLister {
	return &inferenceIngressLister{indexer: indexer}
}

// List lists all InferenceIngresses in the indexer.
func (s *inferenceIngressLister) List(selector labels.Selector) (ret []*v1.InferenceIngress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.InferenceIngress))
	})
	return ret, err
}

// InferenceIngresses returns an object that can list and get InferenceIngresses.
func (s *inferenceIngressLister) InferenceIngresses(namespace string) InferenceIngressNamespaceLister {
	return inferenceIngressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InferenceIngressNamespaceLister helps list and get InferenceIngresses.
// All objects returned here must be treated as read-only.
type InferenceIngressNamespaceLister interface {
	// List lists all InferenceIngresses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.InferenceIngress, err error)
	// Get retrieves the InferenceIngress from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.InferenceIngress, error)
	InferenceIngressNamespaceListerExpansion
}

// inferenceIngressNamespaceLister implements the InferenceIngressNamespaceLister
// interface.
type inferenceIngressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InferenceIngresses in the indexer for a given namespace.
func (s inferenceIngressNamespaceLister) List(selector labels.Selector) (ret []*v1.InferenceIngress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.InferenceIngress))
	})
	return ret, err
}

// Get retrieves the InferenceIngress from the indexer for a given namespace and name.
func (s inferenceIngressNamespaceLister) Get(name string) (*v1.InferenceIngress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("inferenceingress"), name)
	}
	return obj.(*v1.InferenceIngress), nil
}
