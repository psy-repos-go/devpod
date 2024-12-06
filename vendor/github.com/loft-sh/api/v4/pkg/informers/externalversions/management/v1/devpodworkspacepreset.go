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

// DevPodWorkspacePresetInformer provides access to a shared informer and lister for
// DevPodWorkspacePresets.
type DevPodWorkspacePresetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DevPodWorkspacePresetLister
}

type devPodWorkspacePresetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDevPodWorkspacePresetInformer constructs a new informer for DevPodWorkspacePreset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDevPodWorkspacePresetInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDevPodWorkspacePresetInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDevPodWorkspacePresetInformer constructs a new informer for DevPodWorkspacePreset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDevPodWorkspacePresetInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ManagementV1().DevPodWorkspacePresets().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ManagementV1().DevPodWorkspacePresets().Watch(context.TODO(), options)
			},
		},
		&managementv1.DevPodWorkspacePreset{},
		resyncPeriod,
		indexers,
	)
}

func (f *devPodWorkspacePresetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDevPodWorkspacePresetInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *devPodWorkspacePresetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&managementv1.DevPodWorkspacePreset{}, f.defaultInformer)
}

func (f *devPodWorkspacePresetInformer) Lister() v1.DevPodWorkspacePresetLister {
	return v1.NewDevPodWorkspacePresetLister(f.Informer().GetIndexer())
}