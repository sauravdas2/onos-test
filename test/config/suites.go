// Copyright 2019-present Open Networking Foundation.
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

package config

import "github.com/onosproject/onos-test/pkg/onit"

// SmokeTestSuite is the primary onos-config test suite
type SmokeTestSuite struct {
	onit.TestSuite
}

// SetupTestSuite sets up the onos-config test suite
func (s *SmokeTestSuite) SetupTestSuite() {
	setup := s.Setup()
	setup.Topo().Nodes(2)
	setup.Config().Nodes(2)
	setup.SetupOrDie()
}

// CLITestSuite is the onos-config CLI test suite
type CLITestSuite struct {
	onit.TestSuite
}

// SetupTestSuite sets up the onos-config CLI test suite
func (s *CLITestSuite) SetupTestSuite() {
	setup := s.Setup()
	setup.Topo().Nodes(2)
	setup.Config().Nodes(2)
	setup.SetupOrDie()
}

// HATestSuite is the onos-config HA test suite
type HATestSuite struct {
	onit.TestSuite
}

// SetupTestSuite sets up the onos-config CLI test suite
func (s *HATestSuite) SetupTestSuite() {
	setup := s.Setup()
	setup.Topo().Nodes(2)
	setup.Config().Nodes(2)
	setup.SetupOrDie()
}
