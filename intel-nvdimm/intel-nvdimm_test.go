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
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_NewCollector(t *testing.T) {
	Convey("create new NvdimmCollector structure", t, func() {
		So(func() { NewCollector() }, ShouldNotPanic)
		result := NewCollector()
		So(result, ShouldNotBeNil)
	})
}

func Test_DiscoverDevices(t *testing.T) {
	Convey("discover devices", t, func() {
		col := NewCollector()
		result := col.DiscoverDevices()
		So(result, ShouldEqual, 0)
	})
}

func Test_NvdimmLabels(t *testing.T) {
	Convey("count nvmdimm metrics", t, func() {
		So(len(nvdimmLabels), ShouldEqual, 27)
	})
}

func Test_GetNvdimmMetrics(t *testing.T) {
	Convey("get nvdimm metrics", t, func() {
		col := NewCollector()
		result := col.DiscoverDevices()
		So(result, ShouldEqual, 0)
		ns := []plugin.Namespace{}
		for k, _ := range nvdimmLabels {
			ns = append(ns, plugin.NewNamespace("intel", "nvdimm", "nvdimm").AddDynamicElement("DimmUID", "Device UID").AddStaticElement(k))
		}
		metrics := col.GetNvdimmMetrics(ns)
		maxMetrics := len(nvdimmLabels) * int(col.Device_discovery_count)
		So(len(metrics), ShouldEqual, maxMetrics)
	})

}

func Test_CollectMetrics(t *testing.T) {
	Convey("collect metrics", t, func() {
		metrics := []plugin.Metric{}
	        nvdimm_ns := []plugin.Namespace{}
	        pool_ns := []plugin.Namespace{}
		namespace_ns := []plugin.Namespace{}
	        interleave_set_ns := []plugin.Namespace{}

		nc := NewCollector()
		So(nc, ShouldNotBeNil)

		nc.DiscoverDevices()
                nvmet := nc.GetNvdimmMetrics(nvdimm_ns)
                metrics = append(metrics, nvmet...)
		nc.DiscoveryPool()
                nvmet2 := nc.getPoolMetric(pool_ns)
                nvmet3 := nc.getInterleavesetMetric(interleave_set_ns)
                metrics = append(metrics, nvmet2...)
                metrics = append(metrics, nvmet3...)
		nc.DiscoveryNamespace()
                nvmet4 := nc.getNamespaceMetric(namespace_ns)
                metrics = append(metrics, nvmet4...)

		returned_metrics, _ := nc.CollectMetrics(metrics)
		So(returned_metrics, ShouldNotEqual, []plugin.Metric{})
	})

}

func Test_GetMetrics(t *testing.T) {
	nc := NewCollector()
	Convey("get collector", t, func () {
                So(nc, ShouldNotBeNil)
	})
	config_policy, _ := nc.GetConfigPolicy()
        Convey("get config policy", t, func () {
		So(config_policy, ShouldNotBeNil)
	})

	plugin_config := plugin.NewConfig()
	maxLabels := len(nvdimmLabels)+len(namespaceLabels)+len(poolLabels)+len(interleaveLabels)
	metricsTypes, _ := nc.GetMetricTypes(plugin_config)

	Convey("get metric types", t, func() {
		So(len(metricsTypes), ShouldEqual, maxLabels)
	})

	metrics, _ := nc.CollectMetrics(metricsTypes)
	maxMetrics := len(nvdimmLabels)*int(nc.Nvdimm.Device_discovery_count)+len(poolLabels)*int(nc.Pool.AmountPool)+len(namespaceLabels)*int(nc.Namespace.AmountNamespace)+len(interleaveLabels)
	Convey("get collect metrics", t, func() {
		So(len(metrics), ShouldEqual, maxMetrics)
	})
}
