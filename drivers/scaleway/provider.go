// Copyright 2018 Drone.IO Inc
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package scaleway

import (
	"sync"
	"text/template"

	"github.com/replicant0wnz/autoscaler/drivers/internal/userdata"
	"github.com/scaleway/scaleway-sdk-go/scw"

	"github.com/replicant0wnz/autoscaler"
)

// provider implements a Scaleway provider.
type provider struct {
	init sync.Once

	accessKey     string
	secretKey     string
	orgID         string
	securityGroup string
	dynamicIP     bool
	zone          scw.Zone // fr-par-1 or nl-ams-1
	size          string
	image         string
	tags          []string
	userdata      *template.Template

	client *scw.Client
}

// New returns a new Scaleway provider.
func New(opts ...Option) (autoscaler.Provider, error) {
	p := new(provider)
	for _, opt := range opts {
		err := opt(p)
		if err != nil {
			return nil, err
		}
	}

	if p.zone == "" {
		p.zone = scw.ZoneFrPar1
	}
	if p.size == "" {
		p.size = "dev1-l"
	}
	if p.image == "" {
		p.image = "ubuntu-bionic"
	}
	if p.userdata == nil {
		p.userdata = userdata.T
	}

	return p, nil
}
