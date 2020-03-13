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
	clustermetav1 "github.com/onosproject/onos-test/pkg/onit/api/meta/v1"
)

type NodesClient interface {
	Nodes() NodesReader
}

func NewNodesClient(objects clustermetav1.ObjectsClient) NodesClient {
	return &nodesClient{
		ObjectsClient: objects,
	}
}

type nodesClient struct {
	clustermetav1.ObjectsClient
}

func (c *nodesClient) Nodes() NodesReader {
	return NewNodesReader(c.ObjectsClient)
}