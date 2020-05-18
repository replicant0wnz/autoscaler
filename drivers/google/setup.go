// Copyright 2018 Drone.IO Inc
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package google

import (
	"context"
	"reflect"

	"github.com/replicant0wnz/autoscaler/logger"

	compute "google.golang.org/api/compute/v1"
)

func (p *provider) setup(ctx context.Context) error {
	if reflect.DeepEqual(p.tags, defaultTags) {
		return p.setupFirewall(ctx)
	}
	return nil
}

func (p *provider) setupFirewall(ctx context.Context) error {
	logger := logger.FromContext(ctx)

	logger.Debugln("finding default firewall rules")

	_, err := p.service.Firewalls.Get(p.project, "default-allow-docker").Context(ctx).Do()
	if err == nil {
		logger.Debugln("found default firewall rule")
		return nil
	}

	rule := &compute.Firewall{
		Allowed: []*compute.FirewallAllowed{
			{
				IPProtocol: "tcp",
				Ports:      []string{"2376"},
			},
		},
		Direction:    "INGRESS",
		Name:         "default-allow-docker",
		Network:      p.network,
		Priority:     1000,
		SourceRanges: []string{"0.0.0.0/0"},
		TargetTags:   []string{"allow-docker"},
	}

	op, err := p.service.Firewalls.Insert(p.project, rule).Context(ctx).Do()
	if err != nil {
		logger.WithError(err).
			Errorln("cannot create firewall operation")
		return err
	}

	err = p.waitGlobalOperation(ctx, op.Name)
	if err != nil {
		logger.WithError(err).
			Errorln("cannot create firewall rule")
	}

	return err
}
