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

package precheck

import (
	"github.com/pkg/errors"

	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/common"
	"github.com/xu-oO/kubekey/v3/cmd/kk/pkg/core/connector"
)

type KubeSphereExist struct {
	common.KubePrepare
}

func (k *KubeSphereExist) PreCheck(runtime connector.Runtime) (bool, error) {
	currentKsVersion, ok := k.PipelineCache.GetMustString(common.KubeSphereVersion)
	if !ok {
		return false, errors.New("get current KubeSphere version failed by pipeline cache")
	}
	if currentKsVersion != "" {
		return true, nil
	}
	return false, nil
}
