/*
 * Copyright 2019 THL A29 Limited, a Tencent company.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package images

import (
	"reflect"
	"sort"

	"tkestack.io/tke/pkg/spec"
	"tkestack.io/tke/pkg/util/containerregistry"
)

type Components struct {
	ETCD               containerregistry.Image
	CoreDNS            containerregistry.Image
	Pause              containerregistry.Image
	NvidiaDevicePlugin containerregistry.Image
	Keepalived         containerregistry.Image

	GPUManager        containerregistry.Image
	Busybox           containerregistry.Image
	GPUQuotaAdmission containerregistry.Image

	MetricsServer  containerregistry.Image
	AddonResizer   containerregistry.Image
	Cilium         containerregistry.Image
	CiliumOperator containerregistry.Image
	Ipamd          containerregistry.Image
	Masq           containerregistry.Image
	CiliumRouter   containerregistry.Image
}

func (c Components) Get(name string) *containerregistry.Image {
	v := reflect.ValueOf(c)
	for i := 0; i < v.NumField(); i++ {
		v, _ := v.Field(i).Interface().(containerregistry.Image)
		if v.Name == name {
			return &v
		}
	}
	return nil
}

var KubecomponetNames = append(KubeNodeImages, "kube-apiserver", "kube-controller-manager", "kube-scheduler", "kube-proxy")
var KubeNodeImages = []string{"kube-proxy"}

var components = Components{
	ETCD:               containerregistry.Image{Name: "etcd", Tag: "v3.4.7"},
	CoreDNS:            containerregistry.Image{Name: "coredns", Tag: "1.7.0"},
	Pause:              containerregistry.Image{Name: "pause", Tag: "3.2"},
	NvidiaDevicePlugin: containerregistry.Image{Name: "nvidia-device-plugin", Tag: "1.0.0-beta4"},
	Keepalived:         containerregistry.Image{Name: "keepalived", Tag: "2.0.16-r0"},

	GPUManager:        containerregistry.Image{Name: "gpu-manager", Tag: "v1.1.5"},
	Busybox:           containerregistry.Image{Name: "busybox", Tag: "1.31.1"},
	GPUQuotaAdmission: containerregistry.Image{Name: "gpu-quota-admission", Tag: "v1.0.0"},

	MetricsServer: containerregistry.Image{Name: "metrics-server", Tag: "v0.3.6"},
	AddonResizer:  containerregistry.Image{Name: "addon-resizer", Tag: "1.8.11"},

	Cilium:         containerregistry.Image{Name: "cilium", Tag: "v1.9.5"},
	CiliumOperator: containerregistry.Image{Name: "cilium-operator-generic", Tag: "v1.9.5"},
	Ipamd:          containerregistry.Image{Name: "tke-eni-ipamd", Tag: "v3.2.6"},
	Masq:           containerregistry.Image{Name: "ip-masq-agent", Tag: "v1.0.0"},
	CiliumRouter:   containerregistry.Image{Name: "cilium-router", Tag: "v0.1.0"},
}

func List() []string {
	var items []string

	for _, version := range spec.K8sVersionsWithV {
		for _, name := range KubecomponetNames {
			items = append(items, containerregistry.Image{Name: name, Tag: version}.BaseName())
		}
	}

	v := reflect.ValueOf(components)
	for i := 0; i < v.NumField(); i++ {
		v, _ := v.Field(i).Interface().(containerregistry.Image)
		items = append(items, v.BaseName())
	}
	sort.Strings(items)

	return items
}

func Get() Components {
	return components
}

func ListKubernetesImageFullNamesWithVersion(version string, images []string) []string {
	var items []string
	for _, name := range images {
		items = append(items, containerregistry.Image{Name: name, Tag: "v" + version}.FullName())
	}

	items = append(items, components.ETCD.FullName())
	items = append(items, components.Pause.FullName())
	items = append(items, components.CoreDNS.FullName())

	return items
}
