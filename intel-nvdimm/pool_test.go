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
func Test_DiscoverPool(t *testing.T) {
	Convey("discover pool and get metrics", t, func() {
		pool := Pool{}
		pool.DiscoveryPool()
		So(int(pool.AmountPool), ShouldBeGreaterThan, 0)

		ns := []plugin.Namespace{}
		for k, _ := range poolLabels {
                        ns = append(ns, plugin.NewNamespace("intel", "nvdimm", "pool").AddDynamicElement("DimmUID", "Device UID").AddStaticElement(k))
                }
		So(len(ns), ShouldEqual, len(poolLabels))

		targetCount := len(poolLabels) * int(pool.AmountPool)
		So(len(pool.getPoolMetric(ns)), ShouldEqual, targetCount)

		unknown_metric := []plugin.Namespace{}
		unknown_metric = append(unknown_metric, plugin.NewNamespace("intel", "nvdimm", "pool").AddDynamicElement("DimmUID", "Device UID").AddStaticElement("unknown_metric"))
		So(pool.getPoolMetric(unknown_metric), ShouldContain, plugin.Metric{})
	})
}
