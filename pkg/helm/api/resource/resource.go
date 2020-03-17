// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"github.com/onosproject/onos-test/pkg/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"time"
)

// Type is a resource type
type Type struct {
	Kind Kind
	Name string
}

// Kind is a resource kind
type Kind struct {
	Group   string
	Version string
	Kind    string
}

// Waiter is an interface for resources that support waiting for readiness
type Waiter interface {
	Wait(time.Duration) error
}

// Client is a resource client
type Client kubernetes.Client

// Filter is a resource filter
type Filter func(kind metav1.GroupVersionKind, meta metav1.ObjectMeta) (bool, error)

// NoFilter is a filter that accepts all resources
var NoFilter Filter = func(kind metav1.GroupVersionKind, meta metav1.ObjectMeta) (bool, error) {
	return true, nil
}

// NewUIDFilter returns a new filter for the given owner UIDs
func NewUIDFilter(uids ...types.UID) Filter {
	return func(kind metav1.GroupVersionKind, meta metav1.ObjectMeta) (bool, error) {
		for _, owner := range meta.OwnerReferences {
			for _, uid := range uids {
				if owner.UID == uid {
					return true, nil
				}
			}
		}
		return false, nil
	}
}

// NewResource creates a new resource
func NewResource(meta metav1.ObjectMeta, kind Kind, client Client) *Resource {
	return &Resource{
		Client:    client,
		Kind:      kind,
		Namespace: meta.Namespace,
		Name:      meta.Name,
		UID:       meta.UID,
	}
}

// Resource is a Kubernetes resource
type Resource struct {
	Client
	Kind      Kind
	Namespace string
	Name      string
	UID       types.UID
}