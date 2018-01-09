// This file was automatically generated by informer-gen

package v1alpha1

import (
	machine_v1alpha1 "github.com/gardener/node-controller-manager/pkg/apis/machine/v1alpha1"
	clientset "github.com/gardener/node-controller-manager/pkg/client/clientset"
	internalinterfaces "github.com/gardener/node-controller-manager/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gardener/node-controller-manager/pkg/client/listers/machine/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// MachineTemplateInformer provides access to a shared informer and lister for
// MachineTemplates.
type MachineTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MachineTemplateLister
}

type machineTemplateInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewMachineTemplateInformer constructs a new informer for MachineTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineTemplateInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.MachineV1alpha1().MachineTemplates(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.MachineV1alpha1().MachineTemplates(namespace).Watch(options)
			},
		},
		&machine_v1alpha1.MachineTemplate{},
		resyncPeriod,
		indexers,
	)
}

func defaultMachineTemplateInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewMachineTemplateInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *machineTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&machine_v1alpha1.MachineTemplate{}, defaultMachineTemplateInformer)
}

func (f *machineTemplateInformer) Lister() v1alpha1.MachineTemplateLister {
	return v1alpha1.NewMachineTemplateLister(f.Informer().GetIndexer())
}
