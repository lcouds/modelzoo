package runtime

import (
	"context"

	"github.com/lcouds/modelzoo/agent/api/types"
	"github.com/lcouds/modelzoo/agent/pkg/k8s"
	modelzetes "github.com/lcouds/modelzoo/modelzooetes/pkg/apis/modelzetes/v2alpha1"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r generalRuntime) ImageCacheCreate(ctx context.Context, req types.ImageCache, inference *modelzetes.Inference) error {
	imageCache := k8s.MakeImageCache(req, inference)
	logrus.Infof("%v", imageCache)

	if _, err := r.kubefledgedClient.KubefledgedV1alpha3().
		ImageCaches(req.Namespace).
		Create(ctx, imageCache, metav1.CreateOptions{}); err != nil {
		return err
	}
	return nil
}
