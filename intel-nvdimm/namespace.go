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
	"unsafe"
)

// #cgo LDFLAGS: -L/lib64 -lixpdimm
// #include <nvm_management.h>
// #include <nvm_types.h>
// #include <nvm_context.h>
// void nvmGetNamespaceDetails(struct namespace_details *p_namespace, struct namespace_discovery *p_discovery) {
//	char *something = &p_discovery->namespace_uid[0];
//	//fprintf(stderr, "ID: %s\n", something);
//	nvm_get_namespace_details(something, p_namespace);
// }
// int getNamespaceType(struct namespace_details ns_details) {
//     return (int)ns_details.type;
// }
import "C"

var namespaceLabels = map[string]label{
	"type": label{
		description: "",
		unit:        "",
	},
	"block_size": label{
		description: "",
		unit:        "",
	},
	"block_count": label{
		description: "",
		unit:        "",
	},
	"health": label{
		description: "",
		unit:        "",
	},
	"enable_state": label{
		description: "",
		unit:        "",
	},
	"btt": label{
		description: "",
		unit:        "",
	},
}

type Namespace struct {
	AmountNamespace    C.int
	NamespaceDiscovery []C.struct_namespace_discovery
	NamespaceDetails   []C.struct_namespace_details
}

func (p *Namespace) DiscoveryNamespace() {
	//C.nvm_create_context()
	p.AmountNamespace = C.nvm_get_namespace_count()

	if p.AmountNamespace <= 0 {
		fmt.Printf("Error: not found namespace\n")
	} else {
		p.NamespaceDiscovery = make([]C.struct_namespace_discovery, p.AmountNamespace)
		arrayDiscovery_ptr := (*C.struct_namespace_discovery)(unsafe.Pointer(&p.NamespaceDiscovery[0]))
		C.nvm_get_namespaces(arrayDiscovery_ptr, C.NVM_UINT8(p.AmountNamespace))

		p.NamespaceDetails = make([]C.struct_namespace_details, p.AmountNamespace)
		arrayDetails := (*C.struct_namespace_details)(unsafe.Pointer(&p.NamespaceDetails[0]))
		C.nvmGetNamespaceDetails(arrayDetails, arrayDiscovery_ptr)
	}
}

func (nse *Namespace) getNamespaceMetric(nss []plugin.Namespace) []plugin.Metric {

	aN := int(nse.AmountNamespace)
	metric := plugin.Metric{}
	metrics := []plugin.Metric{}

	for _, ns := range nss {
		metricName := ns.Element(len(ns) - 1).Value
		//For all uid
		if ns[3].Value == "*" {
			for i := 0; i < aN; i++ {

				newNS := make([]plugin.NamespaceElement, len(ns))
				copy(newNS, ns)
				newNS[3].Value = C.GoString(&nse.NamespaceDetails[i].pool_uid[0])

				metric = getValueOfPropertyNamespace(nse, i, metricName, newNS)
				fmt.Println(metric.Namespace)
				fmt.Println(metric.Data)
				metrics = append(metrics, metric)
			}

		} else { //For specific uid
			newNS := make([]plugin.NamespaceElement, len(ns))
			copy(newNS, ns)

			//Check where in ArrayPools is requested UID
			for i := 0; i < aN; i++ {
				if ns[3].Value == C.GoString(&nse.NamespaceDetails[i].pool_uid[0]) {
					metric = getValueOfPropertyNamespace(nse, i, metricName, newNS)
					fmt.Println(metric.Namespace)
					fmt.Println(metric.Data)
					metrics = append(metrics, metric)
				}
			}
		}

	}
	return metrics
}

func getValueOfPropertyNamespace(nse *Namespace, i int, metricName string, ns []plugin.NamespaceElement) plugin.Metric {
	var v uint
	var v32 C.NVM_UINT32
	var v64 C.NVM_UINT64

	switch metricName {
	case "type":
		v := TypeNamespace(int(C.getNamespaceType(nse.NamespaceDetails[i]))).String()
		metric := plugin.Metric{
			Namespace: ns,
			Data:      v,
		}
		return metric
	case "block_size":
		v32 = nse.NamespaceDetails[i].block_size
		v = uint(v32)
		metric := plugin.Metric{
			Namespace: ns,
			Data:      v,
		}
		return metric
	case "block_count":
		v64 = nse.NamespaceDetails[i].block_count
		v = uint(v64)
		metric := plugin.Metric{
			Namespace: ns,
			Data:      v,
		}
		return metric
	case "health":
		health := HealthNamespace(int(nse.NamespaceDetails[i].health)).String()
		metric := plugin.Metric{
			Namespace: ns,
			Data:      health,
		}
		return metric
	case "enable_state":
		state := StateNamespace(int(nse.NamespaceDetails[i].enabled)).String()
		metric := plugin.Metric{
			Namespace: ns,
			Data:      state,
		}
		return metric
	case "btt":
		value := uint8(nse.NamespaceDetails[i].btt)
		metric := plugin.Metric{
			Namespace: ns,
			Data:      value,
		}
		return metric
	default:
		fmt.Println("No exists metric")
		metric := plugin.Metric{}
		return metric
	}
}
