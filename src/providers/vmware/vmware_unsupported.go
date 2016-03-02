// Copyright 2015 CoreOS, Inc.
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

// The vmware provider fetches a configuration from the VMware Guest Info
// interface.

// +build !amd64

package vmware

import (
	"errors"

	"github.com/coreos/ignition/config"
)

func (p provider) FetchConfig() (config.Config, error) {
	return config.Config{}, errors.New("vmware provider is not supported on this architecture")
}

func (p *provider) IsOnline() bool {
	return false
}
