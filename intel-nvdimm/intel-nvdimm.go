/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2017 Intel Corporation

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

package nvdimm

import (
	"fmt"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"strconv"
)

// #cgo LDFLAGS: -L/lib64 -lixpdimm
// #include <nvm_management.h>
// #include <nvm_types.h>
// #include <nvm_context.h>
import "C"

type label struct {
	description string
	unit        string
}

const (
	PluginName = "nvdimm"
	Version    = 1

	nsVendor = "intel"
	nsClass  = "nvdimm"
	nsType   = "smart"
	devname  = "device"
)

func init() {
}

type NvdimmCollector struct {
	Nvdimm
	Pool
	Namespace
}

func NewCollector() *NvdimmCollector {
	return &NvdimmCollector{}
}

func (nc *NvdimmCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	nvdimm_ns := []plugin.Namespace{}
	pool_ns := []plugin.Namespace{}
	namespace_ns := []plugin.Namespace{}
	interleave_set_ns := []plugin.Namespace{}

	metrics := []plugin.Metric{}
	// TODO: Handle errors here

	// Insert metrics that will be collected
	for _, metric := range mts {
		metric_name := metric.Namespace
		switch metric_name[2].Value { // Divide metrics by category
		case "nvdimm":
			nvdimm_ns = append(nvdimm_ns, metric_name)
		case "pool":
			pool_ns = append(pool_ns, metric_name)
		case "namespace":
			namespace_ns = append(namespace_ns, metric_name)
		case "interleave_set":
			interleave_set_ns = append(interleave_set_ns, metric_name)
		}
	}

	if len(nvdimm_ns) > 0 {
	nc.DiscoverDevices()
	    nvmet := nc.GetNvdimmMetrics(nvdimm_ns)
	    metrics = append(metrics, nvmet...)
	}
	if len(pool_ns) > 0 || len(interleave_set_ns) > 0 {
	    nc.DiscoveryPool()
	    nvmet2 := nc.getPoolMetric(pool_ns)
	    nvmet3 := nc.getInterleavesetMetric(interleave_set_ns)
	    metrics = append(metrics, nvmet2...)
	    metrics = append(metrics, nvmet3...)
	}
	if len(pool_ns) > 0 || len(interleave_set_ns) > 0 {
		nc.DiscoveryNamespace()
	    nvmet4 := nc.getNamespaceMetric(namespace_ns)
	metrics = append(metrics, nvmet4...)
	}
	return metrics, nil
}

func (NvdimmCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()

	return *policy, nil
}

func (NvdimmCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	for name, label := range nvdimmLabels {
		metric := plugin.Metric{
			Namespace:   plugin.NewNamespace("intel", "nvdimm", "nvdimm").AddDynamicElement("DimmUID", "Device UID").AddStaticElement(name),
			Description: "dynamic NVDIMM metric: " + name,
			Unit:        label.unit,
			Version:     1,
		}
		metrics = append(metrics, metric)
	}

	for name, label := range poolLabels {
		metric := plugin.Metric{
			Namespace:   plugin.NewNamespace("intel", "nvdimm", "pool").AddDynamicElement("PoolID", "Pool UID").AddStaticElement(name),
			Description: "dynamic Pool metric: " + name,
			Unit:        label.unit,
			Version:     1,
		}
		metrics = append(metrics, metric)
	}

	for name, label := range interleaveLabels {
		metric := plugin.Metric{
			Namespace:   plugin.NewNamespace("intel", "nvdimm", "interleave_set").AddDynamicElement("InterleaveID", "Interleave UID").AddStaticElement(name),
			Description: "dynamic Interleave_set metric: " + name,
			Unit:        label.unit,
			Version:     1,
		}
		metrics = append(metrics, metric)
	}

	for name, label := range namespaceLabels {
		metric := plugin.Metric{
			Namespace:   plugin.NewNamespace("intel", "nvdimm", "namespace").AddDynamicElement("NamespaceID", "Namespace UID").AddStaticElement(name),
			Description: "dynamic Namespace metric: " + name,
			Unit:        label.unit,
			Version:     1,
		}
		metrics = append(metrics, metric)
	}

	return metrics, nil
}
