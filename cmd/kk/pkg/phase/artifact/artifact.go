/*
 Copyright 2021 The KubeSphere Authors.

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

package artifact

import (
	"errors"

	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/artifact"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/common"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/core/module"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/core/pipeline"
)

func NewArtifactImportPipeline(runtime *common.KubeRuntime) error {

	m := []module.Module{
		&artifact.UnArchiveModule{},
	}

	p := pipeline.Pipeline{
		Name:    "ArtifactImportPipeline",
		Modules: m,
		Runtime: runtime,
	}
	if err := p.Start(); err != nil {
		return err
	}
	return nil
}

func ArtifactImport(args common.Argument) error {
	var loaderType string

	loaderType = common.AllInOne

	runtime, err := common.NewKubeRuntime(loaderType, args)
	if err != nil {
		return err
	}
	switch runtime.Cluster.Kubernetes.Type {
	case common.Kubernetes:
		if err := NewArtifactImportPipeline(runtime); err != nil {
			return err
		}
	default:
		return errors.New("unsupported cluster kubernetes type")
	}

	return nil
}
