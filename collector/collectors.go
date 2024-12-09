/*
Copyright 2022 Koor Technologies, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package collector

import (
	"context"

	"github.com/ceph/go-ceph/rados"
	rgwadmin "github.com/ceph/go-ceph/rgw/admin"
	"github.com/prometheus/client_golang/prometheus"
)

const Namespace = "ceph"

type Client struct {
	Name string

	RGWAdminAPI *rgwadmin.API
	Rados       *rados.Conn
}

type Collector interface {
	Update(context.Context, *Client, chan<- prometheus.Metric) error
}

type NewCollectorFunc func() (Collector, error)

var Factories map[string]NewCollectorFunc = map[string]NewCollectorFunc{}
