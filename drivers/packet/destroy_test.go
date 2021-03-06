// Copyright 2018 Drone.IO Inc
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package packet

import (
	"context"
	"reflect"
	"testing"

	"github.com/replicant0wnz/autoscaler"
	"github.com/h2non/gock"
	"github.com/packethost/packngo"
)

func TestDestroyError(t *testing.T) {
	defer gock.Off()

	gock.New(baseURL).
		Delete(getDevice + "/" + instanceID).
		Reply(400)

	err := prov.Destroy(context.Background(), &autoscaler.Instance{ID: instanceID})

	if err == nil {
		t.Errorf("Expect error when deleting a device")
	} else if _, ok := err.(*packngo.ErrorResponse); !ok {
		t.Errorf("expected: %s , got: %s ", reflect.TypeOf(&packngo.ErrorResponse{}), reflect.TypeOf(err))
	}
}
