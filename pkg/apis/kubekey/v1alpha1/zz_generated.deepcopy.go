// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	in.Sources.DeepCopyInto(&out.Sources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CalicoCfg) DeepCopyInto(out *CalicoCfg) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CalicoCfg.
func (in *CalicoCfg) DeepCopy() *CalicoCfg {
	if in == nil {
		return nil
	}
	out := new(CalicoCfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CephRBD) DeepCopyInto(out *CephRBD) {
	*out = *in
	if in.Monitors != nil {
		in, out := &in.Monitors, &out.Monitors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CephRBD.
func (in *CephRBD) DeepCopy() *CephRBD {
	if in == nil {
		return nil
	}
	out := new(CephRBD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chart) DeepCopyInto(out *Chart) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chart.
func (in *Chart) DeepCopy() *Chart {
	if in == nil {
		return nil
	}
	out := new(Chart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]HostCfg, len(*in))
		copy(*out, *in)
	}
	in.RoleGroups.DeepCopyInto(&out.RoleGroups)
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	in.Kubernetes.DeepCopyInto(&out.Kubernetes)
	out.Network = in.Network
	in.Registry.DeepCopyInto(&out.Registry)
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]Addon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.KubeSphere = in.KubeSphere
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneEndpoint) DeepCopyInto(out *ControlPlaneEndpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneEndpoint.
func (in *ControlPlaneEndpoint) DeepCopy() *ControlPlaneEndpoint {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalEtcd) DeepCopyInto(out *ExternalEtcd) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalEtcd.
func (in *ExternalEtcd) DeepCopy() *ExternalEtcd {
	if in == nil {
		return nil
	}
	out := new(ExternalEtcd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlusterFS) DeepCopyInto(out *GlusterFS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlusterFS.
func (in *GlusterFS) DeepCopy() *GlusterFS {
	if in == nil {
		return nil
	}
	out := new(GlusterFS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostCfg) DeepCopyInto(out *HostCfg) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostCfg.
func (in *HostCfg) DeepCopy() *HostCfg {
	if in == nil {
		return nil
	}
	out := new(HostCfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostGroups) DeepCopyInto(out *HostGroups) {
	*out = *in
	if in.All != nil {
		in, out := &in.All, &out.All
		*out = make([]HostCfg, len(*in))
		copy(*out, *in)
	}
	if in.Etcd != nil {
		in, out := &in.Etcd, &out.Etcd
		*out = make([]HostCfg, len(*in))
		copy(*out, *in)
	}
	if in.Master != nil {
		in, out := &in.Master, &out.Master
		*out = make([]HostCfg, len(*in))
		copy(*out, *in)
	}
	if in.Worker != nil {
		in, out := &in.Worker, &out.Worker
		*out = make([]HostCfg, len(*in))
		copy(*out, *in)
	}
	if in.K8s != nil {
		in, out := &in.K8s, &out.K8s
		*out = make([]HostCfg, len(*in))
		copy(*out, *in)
	}
	if in.Client != nil {
		in, out := &in.Client, &out.Client
		*out = make([]HostCfg, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostGroups.
func (in *HostGroups) DeepCopy() *HostGroups {
	if in == nil {
		return nil
	}
	out := new(HostGroups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSphere) DeepCopyInto(out *KubeSphere) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSphere.
func (in *KubeSphere) DeepCopy() *KubeSphere {
	if in == nil {
		return nil
	}
	out := new(KubeSphere)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kubernetes) DeepCopyInto(out *Kubernetes) {
	*out = *in
	if in.ApiserverCertExtraSans != nil {
		in, out := &in.ApiserverCertExtraSans, &out.ApiserverCertExtraSans
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kubernetes.
func (in *Kubernetes) DeepCopy() *Kubernetes {
	if in == nil {
		return nil
	}
	out := new(Kubernetes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolume) DeepCopyInto(out *LocalVolume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolume.
func (in *LocalVolume) DeepCopy() *LocalVolume {
	if in == nil {
		return nil
	}
	out := new(LocalVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig) DeepCopyInto(out *NetworkConfig) {
	*out = *in
	out.Calico = in.Calico
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfig.
func (in *NetworkConfig) DeepCopy() *NetworkConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NfsClient) DeepCopyInto(out *NfsClient) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NfsClient.
func (in *NfsClient) DeepCopy() *NfsClient {
	if in == nil {
		return nil
	}
	out := new(NfsClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryConfig) DeepCopyInto(out *RegistryConfig) {
	*out = *in
	if in.RegistryMirrors != nil {
		in, out := &in.RegistryMirrors, &out.RegistryMirrors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InsecureRegistries != nil {
		in, out := &in.InsecureRegistries, &out.InsecureRegistries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryConfig.
func (in *RegistryConfig) DeepCopy() *RegistryConfig {
	if in == nil {
		return nil
	}
	out := new(RegistryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleGroups) DeepCopyInto(out *RoleGroups) {
	*out = *in
	if in.Etcd != nil {
		in, out := &in.Etcd, &out.Etcd
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Master != nil {
		in, out := &in.Master, &out.Master
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Worker != nil {
		in, out := &in.Worker, &out.Worker
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleGroups.
func (in *RoleGroups) DeepCopy() *RoleGroups {
	if in == nil {
		return nil
	}
	out := new(RoleGroups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sources) DeepCopyInto(out *Sources) {
	*out = *in
	out.Chart = in.Chart
	in.Yaml.DeepCopyInto(&out.Yaml)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sources.
func (in *Sources) DeepCopy() *Sources {
	if in == nil {
		return nil
	}
	out := new(Sources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	out.LocalVolume = in.LocalVolume
	out.NfsClient = in.NfsClient
	out.GlusterFS = in.GlusterFS
	in.CephRBD.DeepCopyInto(&out.CephRBD)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Yaml) DeepCopyInto(out *Yaml) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Yaml.
func (in *Yaml) DeepCopy() *Yaml {
	if in == nil {
		return nil
	}
	out := new(Yaml)
	in.DeepCopyInto(out)
	return out
}
