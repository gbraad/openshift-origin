// This file was automatically generated by informer-gen

package v1

import (
	quota_v1 "github.com/openshift/origin/pkg/quota/apis/quota/v1"
	clientset "github.com/openshift/origin/pkg/quota/generated/clientset"
	internalinterfaces "github.com/openshift/origin/pkg/quota/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/origin/pkg/quota/generated/listers/quota/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterResourceQuotaInformer provides access to a shared informer and lister for
// ClusterResourceQuotas.
type ClusterResourceQuotaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterResourceQuotaLister
}

type clusterResourceQuotaInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newClusterResourceQuotaInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.QuotaV1().ClusterResourceQuotas().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.QuotaV1().ClusterResourceQuotas().Watch(options)
			},
		},
		&quota_v1.ClusterResourceQuota{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *clusterResourceQuotaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&quota_v1.ClusterResourceQuota{}, newClusterResourceQuotaInformer)
}

func (f *clusterResourceQuotaInformer) Lister() v1.ClusterResourceQuotaLister {
	return v1.NewClusterResourceQuotaLister(f.Informer().GetIndexer())
}
