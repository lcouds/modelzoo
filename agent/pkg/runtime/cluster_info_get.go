package runtime

import (
	"strings"

	"github.com/lcouds/modelzoo/agent/api/types"
	"github.com/lcouds/modelzoo/agent/pkg/k8s"
	"github.com/lcouds/modelzoo/agent/pkg/version"
)

func (r generalRuntime) GetClusterInfo(cluster *types.ManagedCluster) error {
	info, err := k8s.GetKubernetesVersion(r.kubeClient)
	if err != nil {
		return err
	}
	cluster.KubernetesVersion = info.GitVersion
	cluster.Platform = info.Platform

	v := version.GetVersion()
	cluster.Version = v.Version

	resources, err := r.ListServerResource()
	if err != nil {
		return err
	}
	cluster.ServerResources = strings.Join(resources, ";")
	return nil
}
