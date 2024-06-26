package runtime

import (
	"context"
	"fmt"

	ingressv1 "github.com/lcouds/modelzoo/ingress-operator/pkg/apis/modelzooetes/v1"
	"github.com/lcouds/modelzoo/modelzooetes/pkg/apis/modelzooetes/v2alpha1"
	"github.com/lcouds/modelzoo/modelzooetes/pkg/consts"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/lcouds/modelzoo/agent/api/types"
	"github.com/lcouds/modelzoo/agent/errdefs"
	"github.com/lcouds/modelzoo/agent/pkg/config"
	localconsts "github.com/lcouds/modelzoo/agent/pkg/consts"
)

func (r generalRuntime) InferenceCreate(ctx context.Context,
	req types.InferenceDeployment, cfg config.IngressConfig, event string, serverPort int) error {

	namespace := req.Spec.Namespace

	if r.eventEnabled {
		err := r.eventRecorder.CreateDeploymentEvent(namespace, req.Spec.Name, event, "")
		if err != nil {
			return err
		}
	}

	inf, err := makeInference(req)
	if err != nil {
		return err
	}

	// Create the ingress
	// TODO(gaocegege): Check if the domain is already used.
	if r.ingressEnabled {
		name := req.Spec.Labels[localconsts.LabelName]

		if r.ingressAnyIPToDomain {
			// Get the service with type=loadbalancer.
			svcs, err := r.kubeClient.CoreV1().Services("").List(ctx, metav1.ListOptions{})
			if err != nil {
				return errdefs.System(fmt.Errorf("failed to list services: %v", err))
			}

			if len(svcs.Items) == 0 {
				return errdefs.System(fmt.Errorf("no service with type=LoadBalancer"))
			}
			var externalIP string
			for _, s := range svcs.Items {
				if s.Spec.Type == v1.ServiceTypeLoadBalancer {
					if len(s.Status.LoadBalancer.Ingress) == 0 {
						continue
					}
					externalIP = s.Status.LoadBalancer.Ingress[0].IP
					break
				}
			}
			// Set the domain to
			ingressDomain := fmt.Sprintf("%s.%s", externalIP, localconsts.Domain)
			cfg.Domain = ingressDomain
		}

		domain, err := makeDomain(name, cfg.Domain)
		if err != nil {
			return errdefs.InvalidParameter(err)
		}

		// Set the domain.
		// Create the inference with the ingress domain.
		if inf.Spec.Annotations == nil {
			inf.Spec.Annotations = make(map[string]string)
		}
		if cfg.TLSEnabled {
			inf.Spec.Annotations[AnnotationDomain] = fmt.Sprintf("https://%s", domain)
		} else {
			inf.Spec.Annotations[AnnotationDomain] = fmt.Sprintf("http://%s", domain)
		}

		_, err = r.inferenceClient.TensorchordV2alpha1().
			Inferences(namespace).Create(
			ctx, inf, metav1.CreateOptions{})
		if err != nil {
			if k8serrors.IsAlreadyExists(err) {
				return errdefs.Conflict(err)
			} else {
				return errdefs.System(err)
			}
		}

		cfg.Domain = domain
		ingress, err := makeIngress(req, cfg)
		if err != nil {
			return err
		}

		_, err = r.ingressClient.TensorchordV1().
			InferenceIngresses(cfg.Namespace).
			Create(ctx, ingress, metav1.CreateOptions{})
		if err != nil {
			if k8serrors.IsAlreadyExists(err) {
				return errdefs.Conflict(err)
			} else {
				return errdefs.System(err)
			}
		}
	} else {
		// Set the gateway kubernetes service domain.
		domain := fmt.Sprintf("gateway.default:%d/api/v1/%s/%s/", serverPort, string(req.Spec.Framework), req.Spec.Name)
		if inf.Spec.Annotations == nil {
			inf.Spec.Annotations = make(map[string]string)
		}
		if cfg.TLSEnabled {
			inf.Spec.Annotations[AnnotationDomain] = fmt.Sprintf("https://%s", domain)
		} else {
			inf.Spec.Annotations[AnnotationDomain] = fmt.Sprintf("http://%s", domain)
		}
		_, err = r.inferenceClient.TensorchordV2alpha1().
			Inferences(namespace).Create(
			ctx, inf, metav1.CreateOptions{})
		if err != nil {
			if k8serrors.IsAlreadyExists(err) {
				return errdefs.Conflict(err)
			} else {
				return errdefs.System(err)
			}
		}
	}
	return nil
}

func makeInference(request types.InferenceDeployment) (*v2alpha1.Inference, error) {
	is := &v2alpha1.Inference{
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.Spec.Name,
			Namespace: request.Spec.Namespace,
			Labels: map[string]string{
				consts.LabelInferenceName: request.Spec.Name,
			},
		},
		Spec: v2alpha1.InferenceSpec{
			Name:             request.Spec.Name,
			Image:            request.Spec.Image,
			Framework:        v2alpha1.Framework(request.Spec.Framework),
			Port:             request.Spec.Port,
			Command:          request.Spec.Command,
			Args:             request.Spec.Args,
			EnvVars:          request.Spec.EnvVars,
			Secrets:          request.Spec.Secrets,
			Constraints:      request.Spec.Constraints,
			Labels:           request.Spec.Labels,
			Annotations:      request.Spec.Annotations,
			HTTPProbePath:    request.Spec.HTTPProbePath,
			ModelBasePath:    request.Spec.ModelBasePath,
			SchedulerName:    request.Spec.SchedulerName,
			RuntimeClassName: request.Spec.RuntimeClassName,
		},
	}

	if request.Spec.Scaling != nil {
		is.Spec.Scaling = &v2alpha1.ScalingConfig{
			MinReplicas:     request.Spec.Scaling.MinReplicas,
			MaxReplicas:     request.Spec.Scaling.MaxReplicas,
			TargetLoad:      request.Spec.Scaling.TargetLoad,
			ZeroDuration:    request.Spec.Scaling.ZeroDuration,
			StartupDuration: request.Spec.Scaling.StartupDuration,
		}
		if request.Spec.Scaling.Type != nil {
			buf := v2alpha1.ScalingType(*request.Spec.Scaling.Type)
			is.Spec.Scaling.Type = &buf
		}
	}

	if request.Spec.Models != nil {
		var models []v2alpha1.ModelConfig
		for _, modelConfig := range request.Spec.Models {
			models = append(models, v2alpha1.ModelConfig{
				Name:    modelConfig.Name,
				Image:   modelConfig.Image,
				Command: modelConfig.Command,
			})
		}
		is.Spec.Models = models
	}

	if request.Spec.Volumes != nil {
		var volumes []v2alpha1.VolumeConfig
		for _, volumeConfig := range request.Spec.Volumes {
			volumes = append(volumes, v2alpha1.VolumeConfig{
				Name:      volumeConfig.Name,
				MountPath: volumeConfig.MountPath,
				SubPath:   volumeConfig.SubPath,
				NFS: &corev1.NFSVolumeSource{
					Server:   volumeConfig.NFS.Server,
					Path:     volumeConfig.NFS.Path,
					ReadOnly: volumeConfig.NFS.ReadOnly,
				},
			})
		}
		is.Spec.Volumes = volumes
	}

	//rr, err := createResources(request)
	//if err != nil {
	//	return nil, errdefs.InvalidParameter(err)
	//}

	//is.Spec.Resources = &rr
	is.Spec.Resources = request.Spec.Resources
	return is, nil
}

func makeIngress(request types.InferenceDeployment, cfg config.IngressConfig) (*ingressv1.InferenceIngress, error) {
	labels := map[string]string{
		consts.LabelInferenceName:      request.Spec.Name,
		consts.LabelInferenceNamespace: request.Spec.Namespace,
	}

	if request.Spec.Labels == nil {
		return nil, errdefs.InvalidParameter(fmt.Errorf("labels is required"))
	}

	// Avoiding Ingress name conflicts
	name := fmt.Sprintf("%s.%s", request.Spec.Name, request.Spec.Namespace)

	ingress := &ingressv1.InferenceIngress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: cfg.Namespace,
			Labels:    labels,
		},
		Spec: ingressv1.InferenceIngressSpec{
			Domain:        cfg.Domain,
			Framework:     string(request.Spec.Framework),
			IngressType:   "nginx",
			BypassGateway: false,
			Function:      name,
			TLS: &ingressv1.InferenceIngressTLS{
				Enabled: cfg.TLSEnabled,
			},
		},
	}

	return ingress, nil
}
