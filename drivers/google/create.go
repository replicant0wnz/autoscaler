// Copyright 2018 Drone.IO Inc
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package google

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/replicant0wnz/autoscaler"
	"github.com/replicant0wnz/autoscaler/logger"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
)

func (p *provider) Create(ctx context.Context, opts autoscaler.InstanceCreateOpts) (*autoscaler.Instance, error) {
	p.init.Do(func() {
		p.setup(ctx)
	})

	buf := new(bytes.Buffer)
	err := p.userdata.Execute(buf, &opts)
	if err != nil {
		return nil, err
	}

	name := strings.ToLower(opts.Name)

	logger := logger.FromContext(ctx).
		WithField("zone", p.zone).
		WithField("image", p.image).
		WithField("size", p.size).
		WithField("name", opts.Name)

	logger.Debugln("instance insert")

	in := &compute.Instance{
		Name:           name,
		Zone:           fmt.Sprintf("projects/%s/zones/%s", p.project, p.zone),
		MinCpuPlatform: "Automatic",
		MachineType:    fmt.Sprintf("projects/%s/zones/%s/machineTypes/%s", p.project, p.zone, p.size),
		Metadata: &compute.Metadata{
			Items: []*compute.MetadataItems{
				{
					Key:   "user-data",
					Value: googleapi.String(buf.String()),
				},
			},
		},
		Tags: &compute.Tags{
			Items: p.tags,
		},
		Disks: []*compute.AttachedDisk{
			{
				Type:       "PERSISTENT",
				Boot:       true,
				Mode:       "READ_WRITE",
				AutoDelete: true,
				DeviceName: name,
				InitializeParams: &compute.AttachedDiskInitializeParams{
					SourceImage: fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s", p.image),
					DiskType:    fmt.Sprintf("projects/%s/zones/%s/diskTypes/%s", p.project, p.zone, p.diskType),
					DiskSizeGb:  p.diskSize,
				},
			},
		},
		CanIpForward: false,
		NetworkInterfaces: []*compute.NetworkInterface{
			{
				Network: p.network,
				Subnetwork: p.subnetwork,
				AccessConfigs: []*compute.AccessConfig{
					{
						Name: "External NAT",
						Type: "ONE_TO_ONE_NAT",
					},
				},
			},
		},
		Labels: p.labels,
		Scheduling: &compute.Scheduling{
			Preemptible:       false,
			OnHostMaintenance: "MIGRATE",
			AutomaticRestart:  googleapi.Bool(true),
		},
		DeletionProtection: false,
		ServiceAccounts: []*compute.ServiceAccount{
			{
				Scopes: p.scopes,
				Email:  "default",
			},
		},
	}

	op, err := p.service.Instances.Insert(p.project, p.zone, in).Context(ctx).Do()
	if err != nil {
		logger.WithError(err).
			Errorln("instance insert failed")
		return nil, err
	}

	logger.Debugln("pending instance insert operation")

	err = p.waitZoneOperation(ctx, op.Name)
	if err != nil {
		logger.WithError(err).
			Errorln("instance insert operation failed")
		return nil, err
	}

	logger.Debugln("instance insert operation complete")

	resp, err := p.service.Instances.Get(p.project, p.zone, name).Do()
	if err != nil {
		logger.WithError(err).
			Errorln("cannot get instance details")
		return nil, err
	}

	instance := &autoscaler.Instance{
		Provider: autoscaler.ProviderGoogle,
		ID:       name,
		Name:     opts.Name,
		Image:    p.image,
		Region:   p.zone,
		Size:     p.size,
		Address:  resp.NetworkInterfaces[0].NetworkIP,
	}

	logger.
		WithField("name", instance.Name).
		WithField("ip", instance.Address).
		Debugln("instance inserted")

	return instance, nil
}
