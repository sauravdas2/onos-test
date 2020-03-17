// Code generated by onit-generate. DO NOT EDIT.

package v1beta1

import (
	"github.com/onosproject/onos-test/pkg/helm/api/resource"
)

type Client interface {
	DeploymentsClient
	StatefulSetsClient
}

func NewClient(resources resource.Client, filter resource.Filter) Client {
	return &client{
		Client:             resources,
		DeploymentsClient:  NewDeploymentsClient(resources, filter),
		StatefulSetsClient: NewStatefulSetsClient(resources, filter),
	}
}

type client struct {
	resource.Client
	DeploymentsClient
	StatefulSetsClient
}