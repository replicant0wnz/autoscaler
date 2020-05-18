// Copyright 2018 Drone.IO Inc
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package mocks

//go:generate mockgen -package=mocks -destination=mock_engine.go   github.com/replicant0wnz/autoscaler Engine
//go:generate mockgen -package=mocks -destination=mock_server.go   github.com/replicant0wnz/autoscaler ServerStore
//go:generate mockgen -package=mocks -destination=mock_provider.go github.com/replicant0wnz/autoscaler Provider
//go:generate mockgen -package=mocks -destination=mock_metrics.go  github.com/replicant0wnz/autoscaler/metrics Collector
//go:generate mockgen -package=mocks -destination=mock_drone.go    github.com/drone/drone-go/drone Client
//go:generate mockgen -package=mocks -destination=mock_docker.go   docker.io/go-docker APIClient
