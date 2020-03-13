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

package v1beta1

import (
	clustermetav1 "github.com/onosproject/onos-test/pkg/onit/api/meta/v1"
)

type DeploymentsClient interface {
	Deployments() DeploymentsReader
}

func NewDeploymentsClient(objects clustermetav1.ObjectsClient) DeploymentsClient {
	return &deploymentsClient{
		ObjectsClient: objects,
	}
}

type deploymentsClient struct {
	clustermetav1.ObjectsClient
}

func (c *deploymentsClient) Deployments() DeploymentsReader {
	return NewDeploymentsReader(c.ObjectsClient)
}