/*
Copyright 2019-2023 TensorChord Inc.

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"net/http"

	v2alpha1 "github.com/lcouds/modelzoo/modelzooetes/pkg/apis/modelzooetes/v2alpha1"
	"github.com/lcouds/modelzoo/modelzooetes/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type TensorchordV2alpha1Interface interface {
	RESTClient() rest.Interface
	InferencesGetter
}

// TensorchordV2alpha1Client is used to interact with features provided by the tensorchord.ai group.
type TensorchordV2alpha1Client struct {
	restClient rest.Interface
}

func (c *TensorchordV2alpha1Client) Inferences(namespace string) InferenceInterface {
	return newInferences(c, namespace)
}

// NewForConfig creates a new TensorchordV2alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*TensorchordV2alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new TensorchordV2alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*TensorchordV2alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &TensorchordV2alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new TensorchordV2alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TensorchordV2alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TensorchordV2alpha1Client for the given RESTClient.
func New(c rest.Interface) *TensorchordV2alpha1Client {
	return &TensorchordV2alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v2alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TensorchordV2alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
