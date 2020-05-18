// Copyright 2018 Drone.IO Inc
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package digitalocean

import (
	"context"
	"strconv"
	"testing"

	"github.com/digitalocean/godo"
	"github.com/replicant0wnz/autoscaler"

	"github.com/golang/mock/gomock"
	"github.com/h2non/gock"
)

func TestDestroy(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.digitalocean.com").
		Get("/v2/droplets/3164494").
		Reply(200).
		BodyString(respDropletCreate)

	gock.New("https://api.digitalocean.com").
		Delete("/v2/droplets/3164494").
		Reply(204)

	mockContext := context.TODO()
	mockInstance := &autoscaler.Instance{
		ID: "3164494",
	}

	p := New(
		WithSSHKey("58:8e:30:66:fc:e2:ff:ad:4f:6f:02:4b:af:28:0d:c7"),
		WithToken("77e027c7447f468068a7d4fea41e7149a75a94088082c66fcf555de3977f69d3"),
	)

	err := p.Destroy(mockContext, mockInstance)
	if err != nil {
		t.Error(err)
	}

	if !gock.IsDone() {
		t.Errorf("Expected http requests not detected")
	}
}

func TestDestroyDeleteError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	defer gock.Off()

	gock.New("https://api.digitalocean.com").
		Get("/v2/droplets/3164494").
		Reply(200).
		BodyString(respDropletCreate)

	gock.New("https://api.digitalocean.com").
		Delete("/v2/droplets/3164494").
		Reply(500)

	mockContext := context.TODO()
	mockInstance := &autoscaler.Instance{
		ID: "3164494",
	}

	p := New(
		WithSSHKey("58:8e:30:66:fc:e2:ff:ad:4f:6f:02:4b:af:28:0d:c7"),
		WithToken("77e027c7447f468068a7d4fea41e7149a75a94088082c66fcf555de3977f69d3"),
	)

	err := p.Destroy(mockContext, mockInstance)
	if err == nil {
		t.Errorf("Expect error returned from digital ocean")
	} else if _, ok := err.(*godo.ErrorResponse); !ok {
		t.Errorf("Expect ErrorResponse digital ocean")
	}

	if !gock.IsDone() {
		t.Errorf("Expected http requests not detected")
	}
}

func TestDestroyFindError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	defer gock.Off()

	gock.New("https://api.digitalocean.com").
		Get("/v2/droplets/3164494").
		Reply(500)

	mockContext := context.TODO()
	mockInstance := &autoscaler.Instance{
		ID: "3164494",
	}

	p := New(
		WithSSHKey("58:8e:30:66:fc:e2:ff:ad:4f:6f:02:4b:af:28:0d:c7"),
		WithToken("77e027c7447f468068a7d4fea41e7149a75a94088082c66fcf555de3977f69d3"),
	)

	err := p.Destroy(mockContext, mockInstance)
	if err == nil {
		t.Errorf("Expect error returned from digital ocean")
	} else if _, ok := err.(*godo.ErrorResponse); !ok {
		t.Errorf("Expect ErrorResponse digital ocean")
	}

	if !gock.IsDone() {
		t.Errorf("Expected http requests not detected")
	}
}

func TestDestroyNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	defer gock.Off()

	gock.New("https://api.digitalocean.com").
		Get("/v2/droplets/3164494").
		Reply(404)

	mockContext := context.TODO()
	mockInstance := &autoscaler.Instance{
		ID: "3164494",
	}

	p := New(
		WithSSHKey("58:8e:30:66:fc:e2:ff:ad:4f:6f:02:4b:af:28:0d:c7"),
		WithToken("77e027c7447f468068a7d4fea41e7149a75a94088082c66fcf555de3977f69d3"),
	)

	err := p.Destroy(mockContext, mockInstance)
	if err == nil {
		t.Errorf("Expect error returned from digital ocean")
	} else if err != autoscaler.ErrInstanceNotFound {
		t.Errorf("Expect ErrInstanceNotFound")
	}

	if !gock.IsDone() {
		t.Errorf("Expected http requests not detected")
	}
}

func TestDestroyInvalidInput(t *testing.T) {
	i := &autoscaler.Instance{}
	p := provider{}
	err := p.Destroy(context.TODO(), i)
	if _, ok := err.(*strconv.NumError); !ok {
		t.Errorf("Expected invalid or missing ID error")
	}
}
