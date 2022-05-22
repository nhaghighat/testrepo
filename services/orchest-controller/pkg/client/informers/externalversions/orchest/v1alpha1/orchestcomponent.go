/*
Copyright 2022 The orchest Authors.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	orchestv1alpha1 "github.com/orchest/orchest/services/orchest-controller/pkg/apis/orchest/v1alpha1"
	versioned "github.com/orchest/orchest/services/orchest-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/orchest/orchest/services/orchest-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/orchest/orchest/services/orchest-controller/pkg/client/listers/orchest/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OrchestComponentInformer provides access to a shared informer and lister for
// OrchestComponents.
type OrchestComponentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.OrchestComponentLister
}

type orchestComponentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewOrchestComponentInformer constructs a new informer for OrchestComponent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOrchestComponentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOrchestComponentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredOrchestComponentInformer constructs a new informer for OrchestComponent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOrchestComponentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OrchestV1alpha1().OrchestComponents(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OrchestV1alpha1().OrchestComponents(namespace).Watch(context.TODO(), options)
			},
		},
		&orchestv1alpha1.OrchestComponent{},
		resyncPeriod,
		indexers,
	)
}

func (f *orchestComponentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOrchestComponentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *orchestComponentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&orchestv1alpha1.OrchestComponent{}, f.defaultInformer)
}

func (f *orchestComponentInformer) Lister() v1alpha1.OrchestComponentLister {
	return v1alpha1.NewOrchestComponentLister(f.Informer().GetIndexer())
}
