// Code generated by informer-gen. DO NOT EDIT.

// This file was automatically generated by informer-gen

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/cluster/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=cluster.k8s.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("awsmachineclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cluster().V1alpha1().AWSMachineClasses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("azuremachineclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cluster().V1alpha1().AzureMachineClasses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cluster().V1alpha1().Clusters().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("gcpmachineclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cluster().V1alpha1().GCPMachineClasses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("machines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cluster().V1alpha1().Machines().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("machineclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cluster().V1alpha1().MachineClasses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("machinedeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cluster().V1alpha1().MachineDeployments().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("machinesets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cluster().V1alpha1().MachineSets().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("openstackmachineclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cluster().V1alpha1().OpenStackMachineClasses().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
