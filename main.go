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

package main

import (
	"github.com/intelsdi-x/snap-plugin-collector-intel-nvdimm/intel-nvdimm"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

func main() {
	//nv := nvdimm.NvdimmCollector{}
	//                            METADATA            COLLECTOR    SOMETHING

	plugin.StartCollector(nvdimm.NewCollector(), nvdimm.PluginName, nvdimm.Version)

	//var namespaces []plugin.Namespace
	//var collectedMetrics []plugin.Metric

	//	nvdimmCollector := nvdimm.NewCollector()
	//	nvdimmCollector.DiscoverDevices()
	/*	namespaces = append(namespaces, plugin.NewNamespace("intel", "nvdimm", "nvdimm", "*", "device_capacity"))
			namespaces = append(namespaces, plugin.NewNamespace("intel", "nvdimm", "nvdimm", "*", "channel_pos"))
			namespaces = append(namespaces, plugin.NewNamespace("intel", "nvdimm", "nvdimm", "*", "spare_capacity"))
		    namespaces = append(namespaces, plugin.NewNamespace("intel", "nvdimm", "nvdimm", "*", "wear_level"))
		    namespaces = append(namespaces, plugin.NewNamespace("intel", "nvdimm", "nvdimm", "*", "media_errors_uncorrectable"))
		    namespaces = append(namespaces, plugin.NewNamespace("intel", "nvdimm", "nvdimm", "*", "media_errors_corrected"))*/
	/*    namespaces = append(namespaces, plugin.NewNamespace("intel", "nvdimm", "nvdimm", "*", "volatile_capacity"))
	      namespaces = append(namespaces, plugin.NewNamespace("intel", "nvdimm", "nvdimm", "*", "persistent_capacity"))
	      namespaces = append(namespaces, plugin.NewNamespace("intel", "nvdimm", "nvdimm", "*", "unconfigured_capacity"))
	      namespaces = append(namespaces, plugin.NewNamespace("intel", "nvdimm", "nvdimm", "*", "inaccessible_capacity"))
	      namespaces = append(namespaces, plugin.NewNamespace("intel", "nvdimm", "nvdimm", "*", "reserved_capacity"))
	*/
	//    collectedMetrics := nvdimmCollector.GetNvdimmMetrics(namespaces)
	//    fmt.Printf("CollectedMetrics size = %d \n", len(collectedMetrics))
	//    fmt.Printf("%s\n", collectedMetrics[0].Namespace.String())

	/*	for _, met := range collectedMetrics {
	    fmt.Print(met.Namespace.String(), " ")
	    fmt.Println(met.Data)
	}*/
}
