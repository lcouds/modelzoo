/*
Copyright 2019-2023 TensorChord Inc.

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v2alpha1

import (
	v2alpha1 "github.com/lcouds/modelzoo/modelzooetes/pkg/apis/modelzetes/v2alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InferenceLister helps list Inferences.
// All objects returned here must be treated as read-only.
type InferenceLister interface {
	// List lists all Inferences in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2alpha1.Inference, err error)
	// Inferences returns an object that can list and get Inferences.
	Inferences(namespace string) InferenceNamespaceLister
	InferenceListerExpansion
}

// inferenceLister implements the InferenceLister interface.
type inferenceLister struct {
	indexer cache.Indexer
}

// NewInferenceLister returns a new InferenceLister.
func NewInferenceLister(indexer cache.Indexer) InferenceLister {
	return &inferenceLister{indexer: indexer}
}

// List lists all Inferences in the indexer.
func (s *inferenceLister) List(selector labels.Selector) (ret []*v2alpha1.Inference, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.Inference))
	})
	return ret, err
}

// Inferences returns an object that can list and get Inferences.
func (s *inferenceLister) Inferences(namespace string) InferenceNamespaceLister {
	return inferenceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InferenceNamespaceLister helps list and get Inferences.
// All objects returned here must be treated as read-only.
type InferenceNamespaceLister interface {
	// List lists all Inferences in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2alpha1.Inference, err error)
	// Get retrieves the Inference from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2alpha1.Inference, error)
	InferenceNamespaceListerExpansion
}

// inferenceNamespaceLister implements the InferenceNamespaceLister
// interface.
type inferenceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Inferences in the indexer for a given namespace.
func (s inferenceNamespaceLister) List(selector labels.Selector) (ret []*v2alpha1.Inference, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.Inference))
	})
	return ret, err
}

// Get retrieves the Inference from the indexer for a given namespace and name.
func (s inferenceNamespaceLister) Get(name string) (*v2alpha1.Inference, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2alpha1.Resource("inference"), name)
	}
	return obj.(*v2alpha1.Inference), nil
}
