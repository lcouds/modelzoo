package runtime

import (
	"context"
	"fmt"

	ingressclientset "github.com/lcouds/modelzoo/ingress-operator/pkg/client/clientset/versioned"
	inferenceclientset "github.com/lcouds/modelzoo/modelzooetes/pkg/client/clientset/versioned"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/lcouds/modelzoo/agent/errdefs"
)

func (r generalRuntime) InferenceDelete(ctx context.Context, namespace,
	inferenceName, ingressNamespace, event string) error {

	if r.eventEnabled {
		err := r.eventRecorder.CreateDeploymentEvent(namespace, inferenceName, event, "")
		if err != nil {
			return err
		}
	}

	getOpts := metav1.GetOptions{}

	// This makes sure we don't delete non-labelled deployments
	_, err := r.inferenceClient.TensorchordV2alpha1().
		Inferences(namespace).
		Get(context.TODO(), inferenceName, getOpts)

	if err != nil {
		if k8serrors.IsNotFound(err) {
			return errdefs.NotFound(err)
		} else {
			return errdefs.System(err)
		}
	}

	if err := deleteInference(ctx, namespace, r.inferenceClient,
		r.ingressClient, ingressNamespace,
		inferenceName, r.ingressEnabled); err != nil {
		return err
	}
	return nil
}

func deleteInference(ctx context.Context,
	namespace string,
	clientset inferenceclientset.Interface,
	ingressClient ingressclientset.Interface,
	baseNamespace string,
	inferenceName string, ingressEnabled bool) error {
	foregroundPolicy := metav1.DeletePropagationForeground
	opts := &metav1.DeleteOptions{PropagationPolicy: &foregroundPolicy}

	if deployErr := clientset.TensorchordV2alpha1().Inferences(namespace).
		Delete(ctx, inferenceName, *opts); deployErr != nil {

		if k8serrors.IsNotFound(deployErr) {
			return errdefs.NotFound(deployErr)
		} else {
			return errdefs.System(deployErr)
		}
	}

	if ingressEnabled && ingressClient != nil {
		name := fmt.Sprintf("%s.%s", inferenceName, namespace)
		if err := ingressClient.TensorchordV1().InferenceIngresses(baseNamespace).Delete(ctx, name, *opts); err != nil {
			if k8serrors.IsNotFound(err) {
				return errdefs.NotFound(err)
			} else {
				return errdefs.System(err)
			}
		}
	}

	return nil
}
