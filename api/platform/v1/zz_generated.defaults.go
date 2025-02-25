// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by defaulter-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	applicationv1 "tkestack.io/tke/api/application/v1"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&CSIOperator{}, func(obj interface{}) { SetObjectDefaults_CSIOperator(obj.(*CSIOperator)) })
	scheme.AddTypeDefaultingFunc(&CSIOperatorList{}, func(obj interface{}) { SetObjectDefaults_CSIOperatorList(obj.(*CSIOperatorList)) })
	scheme.AddTypeDefaultingFunc(&Cluster{}, func(obj interface{}) { SetObjectDefaults_Cluster(obj.(*Cluster)) })
	scheme.AddTypeDefaultingFunc(&ClusterList{}, func(obj interface{}) { SetObjectDefaults_ClusterList(obj.(*ClusterList)) })
	scheme.AddTypeDefaultingFunc(&ConfigMap{}, func(obj interface{}) { SetObjectDefaults_ConfigMap(obj.(*ConfigMap)) })
	scheme.AddTypeDefaultingFunc(&ConfigMapList{}, func(obj interface{}) { SetObjectDefaults_ConfigMapList(obj.(*ConfigMapList)) })
	scheme.AddTypeDefaultingFunc(&CronHPA{}, func(obj interface{}) { SetObjectDefaults_CronHPA(obj.(*CronHPA)) })
	scheme.AddTypeDefaultingFunc(&CronHPAList{}, func(obj interface{}) { SetObjectDefaults_CronHPAList(obj.(*CronHPAList)) })
	scheme.AddTypeDefaultingFunc(&Helm{}, func(obj interface{}) { SetObjectDefaults_Helm(obj.(*Helm)) })
	scheme.AddTypeDefaultingFunc(&HelmList{}, func(obj interface{}) { SetObjectDefaults_HelmList(obj.(*HelmList)) })
	scheme.AddTypeDefaultingFunc(&IPAM{}, func(obj interface{}) { SetObjectDefaults_IPAM(obj.(*IPAM)) })
	scheme.AddTypeDefaultingFunc(&IPAMList{}, func(obj interface{}) { SetObjectDefaults_IPAMList(obj.(*IPAMList)) })
	scheme.AddTypeDefaultingFunc(&LBCF{}, func(obj interface{}) { SetObjectDefaults_LBCF(obj.(*LBCF)) })
	scheme.AddTypeDefaultingFunc(&LBCFList{}, func(obj interface{}) { SetObjectDefaults_LBCFList(obj.(*LBCFList)) })
	scheme.AddTypeDefaultingFunc(&LogCollector{}, func(obj interface{}) { SetObjectDefaults_LogCollector(obj.(*LogCollector)) })
	scheme.AddTypeDefaultingFunc(&LogCollectorList{}, func(obj interface{}) { SetObjectDefaults_LogCollectorList(obj.(*LogCollectorList)) })
	scheme.AddTypeDefaultingFunc(&Machine{}, func(obj interface{}) { SetObjectDefaults_Machine(obj.(*Machine)) })
	scheme.AddTypeDefaultingFunc(&MachineList{}, func(obj interface{}) { SetObjectDefaults_MachineList(obj.(*MachineList)) })
	scheme.AddTypeDefaultingFunc(&PersistentEvent{}, func(obj interface{}) { SetObjectDefaults_PersistentEvent(obj.(*PersistentEvent)) })
	scheme.AddTypeDefaultingFunc(&PersistentEventList{}, func(obj interface{}) { SetObjectDefaults_PersistentEventList(obj.(*PersistentEventList)) })
	scheme.AddTypeDefaultingFunc(&Prometheus{}, func(obj interface{}) { SetObjectDefaults_Prometheus(obj.(*Prometheus)) })
	scheme.AddTypeDefaultingFunc(&PrometheusList{}, func(obj interface{}) { SetObjectDefaults_PrometheusList(obj.(*PrometheusList)) })
	scheme.AddTypeDefaultingFunc(&TappController{}, func(obj interface{}) { SetObjectDefaults_TappController(obj.(*TappController)) })
	scheme.AddTypeDefaultingFunc(&TappControllerList{}, func(obj interface{}) { SetObjectDefaults_TappControllerList(obj.(*TappControllerList)) })
	scheme.AddTypeDefaultingFunc(&VolumeDecorator{}, func(obj interface{}) { SetObjectDefaults_VolumeDecorator(obj.(*VolumeDecorator)) })
	scheme.AddTypeDefaultingFunc(&VolumeDecoratorList{}, func(obj interface{}) { SetObjectDefaults_VolumeDecoratorList(obj.(*VolumeDecoratorList)) })
	return nil
}

func SetObjectDefaults_CSIOperator(in *CSIOperator) {
	SetDefaults_CSIOperatorStatus(&in.Status)
}

func SetObjectDefaults_CSIOperatorList(in *CSIOperatorList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_CSIOperator(a)
	}
}

func SetObjectDefaults_Cluster(in *Cluster) {
	SetDefaults_ClusterSpec(&in.Spec)
	for i := range in.Spec.BootstrapApps {
		a := &in.Spec.BootstrapApps[i]
		applicationv1.SetDefaults_AppSpec(&a.App.Spec)
	}
	SetDefaults_ClusterStatus(&in.Status)
}

func SetObjectDefaults_ClusterList(in *ClusterList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Cluster(a)
	}
}

func SetObjectDefaults_ConfigMap(in *ConfigMap) {
	SetDefaults_ConfigMap(in)
}

func SetObjectDefaults_ConfigMapList(in *ConfigMapList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ConfigMap(a)
	}
}

func SetObjectDefaults_CronHPA(in *CronHPA) {
	SetDefaults_CronHPAStatus(&in.Status)
}

func SetObjectDefaults_CronHPAList(in *CronHPAList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_CronHPA(a)
	}
}

func SetObjectDefaults_Helm(in *Helm) {
	SetDefaults_HelmStatus(&in.Status)
}

func SetObjectDefaults_HelmList(in *HelmList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Helm(a)
	}
}

func SetObjectDefaults_IPAM(in *IPAM) {
	SetDefaults_IPAMStatus(&in.Status)
}

func SetObjectDefaults_IPAMList(in *IPAMList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_IPAM(a)
	}
}

func SetObjectDefaults_LBCF(in *LBCF) {
	SetDefaults_LBCFStatus(&in.Status)
}

func SetObjectDefaults_LBCFList(in *LBCFList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_LBCF(a)
	}
}

func SetObjectDefaults_LogCollector(in *LogCollector) {
	SetDefaults_LogCollectorStatus(&in.Status)
}

func SetObjectDefaults_LogCollectorList(in *LogCollectorList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_LogCollector(a)
	}
}

func SetObjectDefaults_Machine(in *Machine) {
	SetDefaults_MachineStatus(&in.Status)
}

func SetObjectDefaults_MachineList(in *MachineList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Machine(a)
	}
}

func SetObjectDefaults_PersistentEvent(in *PersistentEvent) {
	SetDefaults_PersistentEventStatus(&in.Status)
}

func SetObjectDefaults_PersistentEventList(in *PersistentEventList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PersistentEvent(a)
	}
}

func SetObjectDefaults_Prometheus(in *Prometheus) {
	SetDefaults_PrometheusStatus(&in.Status)
}

func SetObjectDefaults_PrometheusList(in *PrometheusList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Prometheus(a)
	}
}

func SetObjectDefaults_TappController(in *TappController) {
	SetDefaults_TappControllerStatus(&in.Status)
}

func SetObjectDefaults_TappControllerList(in *TappControllerList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_TappController(a)
	}
}

func SetObjectDefaults_VolumeDecorator(in *VolumeDecorator) {
	SetDefaults_VolumeDecoratorStatus(&in.Status)
}

func SetObjectDefaults_VolumeDecoratorList(in *VolumeDecoratorList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_VolumeDecorator(a)
	}
}
