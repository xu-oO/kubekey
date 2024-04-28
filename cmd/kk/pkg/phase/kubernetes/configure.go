/*
 Copyright 2022 The KubeSphere Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package kubernetes

import (
	"github.com/pkg/errors"

	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/addons"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/bootstrap/precheck"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/certs"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/common"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/core/module"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/core/pipeline"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/filesystem"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/kubernetes"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/plugins"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/plugins/network"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/plugins/storage"
)

func NewCreateConfigureKubernetesPipeline(runtime *common.KubeRuntime) error {
	skipLocalStorage := true
	if runtime.Arg.DeployLocalStorage != nil {
		skipLocalStorage = !*runtime.Arg.DeployLocalStorage
	} else if runtime.Cluster.KubeSphere.Enabled {
		skipLocalStorage = false
	}
	m := []module.Module{
		&precheck.NodePreCheckModule{},
		&kubernetes.StatusModule{},
		&network.DeployNetworkPluginModule{},
		&kubernetes.ConfigureKubernetesModule{},
		&filesystem.ChownModule{},
		&certs.AutoRenewCertsModule{Skip: !runtime.Cluster.Kubernetes.EnableAutoRenewCerts()},
		&kubernetes.SaveKubeConfigModule{},
		&plugins.DeployPluginsModule{},
		&addons.AddonsModule{},
		&storage.DeployLocalVolumeModule{Skip: skipLocalStorage},
	}

	p := pipeline.Pipeline{
		Name:    "CreateConfigureKubernetesPipeline",
		Modules: m,
		Runtime: runtime,
	}
	if err := p.Start(); err != nil {
		return err
	}
	return nil
}

func CreateConfigureKubernetes(args common.Argument) error {
	var loaderType string
	if args.FilePath != "" {
		loaderType = common.File
	} else {
		loaderType = common.AllInOne
	}

	runtime, err := common.NewKubeRuntime(loaderType, args)
	if err != nil {
		return err
	}

	switch runtime.Cluster.Kubernetes.Type {
	case common.Kubernetes:
		if err := NewCreateConfigureKubernetesPipeline(runtime); err != nil {
			return err
		}
	default:
		return errors.New("unsupported cluster kubernetes type")
	}

	return nil
}
