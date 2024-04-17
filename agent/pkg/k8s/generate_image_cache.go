package k8s

import (
	"time"

	kubefledged "github.com/lcouds/kube-fledged/pkg/apis/kubefledged/v1alpha3"
	"github.com/lcouds/modelzoo/agent/api/types"
	"github.com/lcouds/modelzoo/agent/pkg/consts"
	modelzooetes "github.com/lcouds/modelzoo/modelzooetes/pkg/apis/modelzooetes/v2alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func MakeImageCache(req types.ImageCache, inference *modelzooetes.Inference) *kubefledged.ImageCache {
	nodeSlector := map[string]string{
		consts.LabelServerResource: string(req.NodeSelector),
	}
	cache := &kubefledged.ImageCache{
		ObjectMeta: v1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(inference, schema.GroupVersionKind{
					Group:   modelzooetes.SchemeGroupVersion.Group,
					Version: modelzooetes.SchemeGroupVersion.Version,
					Kind:    modelzooetes.Kind,
				}),
			},
		},
		Spec: kubefledged.ImageCacheSpec{
			CacheSpec: []kubefledged.CacheSpecImages{
				{
					Images: []kubefledged.Image{
						{
							Name:           req.Image,
							ForceFullCache: req.ForceFullCache,
						},
					},
					NodeSelector: nodeSlector,
				},
			},
		},
		Status: kubefledged.ImageCacheStatus{
			StartTime: &metav1.Time{Time: time.Now()},
		},
	}
	return cache
}
