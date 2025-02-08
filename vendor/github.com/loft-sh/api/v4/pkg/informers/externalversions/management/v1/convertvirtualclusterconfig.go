// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	managementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	versioned "github.com/loft-sh/api/v4/pkg/clientset/versioned"
	internalinterfaces "github.com/loft-sh/api/v4/pkg/informers/externalversions/internalinterfaces"
	v1 "github.com/loft-sh/api/v4/pkg/listers/management/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConvertVirtualClusterConfigInformer provides access to a shared informer and lister for
// ConvertVirtualClusterConfigs.
type ConvertVirtualClusterConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ConvertVirtualClusterConfigLister
}

type convertVirtualClusterConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewConvertVirtualClusterConfigInformer constructs a new informer for ConvertVirtualClusterConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConvertVirtualClusterConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConvertVirtualClusterConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredConvertVirtualClusterConfigInformer constructs a new informer for ConvertVirtualClusterConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConvertVirtualClusterConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ManagementV1().ConvertVirtualClusterConfigs().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ManagementV1().ConvertVirtualClusterConfigs().Watch(context.TODO(), options)
			},
		},
		&managementv1.ConvertVirtualClusterConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *convertVirtualClusterConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConvertVirtualClusterConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *convertVirtualClusterConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&managementv1.ConvertVirtualClusterConfig{}, f.defaultInformer)
}

func (f *convertVirtualClusterConfigInformer) Lister() v1.ConvertVirtualClusterConfigLister {
	return v1.NewConvertVirtualClusterConfigLister(f.Informer().GetIndexer())
}
