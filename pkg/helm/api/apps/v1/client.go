// Code generated by onit-generate. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/onos-test/pkg/helm/api/resource"
)

type Client interface {
	DaemonSetsClient
	DeploymentsClient
	ReplicaSetsClient
	StatefulSetsClient
}

func NewClient(resources resource.Client, filter resource.Filter) Client {
	return &client{
		Client:             resources,
		DaemonSetsClient:   NewDaemonSetsClient(resources, filter),
		DeploymentsClient:  NewDeploymentsClient(resources, filter),
		ReplicaSetsClient:  NewReplicaSetsClient(resources, filter),
		StatefulSetsClient: NewStatefulSetsClient(resources, filter),
	}
}

type client struct {
	resource.Client
	DaemonSetsClient
	DeploymentsClient
	ReplicaSetsClient
	StatefulSetsClient
}