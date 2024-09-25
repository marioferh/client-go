// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterMonitoringLister helps list ClusterMonitorings.
// All objects returned here must be treated as read-only.
type ClusterMonitoringLister interface {
	// List lists all ClusterMonitorings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterMonitoring, err error)
	// Get retrieves the ClusterMonitoring from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterMonitoring, error)
	ClusterMonitoringListerExpansion
}

// clusterMonitoringLister implements the ClusterMonitoringLister interface.
type clusterMonitoringLister struct {
	indexer cache.Indexer
}

// NewClusterMonitoringLister returns a new ClusterMonitoringLister.
func NewClusterMonitoringLister(indexer cache.Indexer) ClusterMonitoringLister {
	return &clusterMonitoringLister{indexer: indexer}
}

// List lists all ClusterMonitorings in the indexer.
func (s *clusterMonitoringLister) List(selector labels.Selector) (ret []*v1.ClusterMonitoring, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterMonitoring))
	})
	return ret, err
}

// Get retrieves the ClusterMonitoring from the index for a given name.
func (s *clusterMonitoringLister) Get(name string) (*v1.ClusterMonitoring, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clustermonitoring"), name)
	}
	return obj.(*v1.ClusterMonitoring), nil
}
