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

package v1

import (
	clustermetav1 "github.com/onosproject/onos-test/pkg/onit/cluster/meta/v1"
)

type Pods interface {
	Get(name string) (*Pod, error)
	List() ([]*Pod, error)
}

func NewPods(objects clustermetav1.ObjectsClient) Pods {
	return &pods{
		ObjectsClient: objects,
	}
}

type pods struct {
	clustermetav1.ObjectsClient
}

func (c *pods) Get(name string) (*Pod, error) {
	object, err := c.ObjectsClient.Get(name, PodResource)
	if err != nil {
		return nil, err
	}
	return NewPod(object), nil
}

func (c *pods) List() ([]*Pod, error) {
	objects, err := c.ObjectsClient.List(PodResource)
	if err != nil {
		return nil, err
	}
	pods := make([]*Pod, len(objects))
	for i, object := range objects {
		pods[i] = NewPod(object)
	}
	return pods, nil
}
