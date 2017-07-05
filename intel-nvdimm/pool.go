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
import "C"

var poolLabels = map[string]label{
	"capacity": label{
		description: "",
		unit:        "",
	},
	"free_capacity": label{
		description: "",
		unit:        "",
	},
	"socket_id": label{
		description: "",
		unit:        "",
	},
	"nvdimm_count": label{
		description: "",
		unit:        "",
	},
	"interleave_set_count": label{
		description: "",
		unit:        "",
	},

	"pool_health": label{
		description: "",
		unit:        "",
	},
}

type Pool struct {
	AmountPool C.int
	ArrayPools []C.struct_pool
}

//Fill structs for Pool and Interleave_set
func (p *Pool) DiscoveryPool() {
	//C.nvm_create_context()
	p.AmountPool = C.nvm_get_pool_count()

	if p.AmountPool < 0 {
		fmt.Printf("Error: not found pool\n")
	} else {
		p.ArrayPools = make([]C.struct_pool, p.AmountPool) //Allocate memory on array of pools
		arrayPools_ptr := (*C.struct_pool)(unsafe.Pointer(&p.ArrayPools[0]))
		C.nvm_get_pools(arrayPools_ptr, C.NVM_UINT8(p.AmountPool))
	}
}

//Main function for getting metrics from Pool
func (p *Pool) getPoolMetric(nss []plugin.Namespace) []plugin.Metric {

	amountPool := int(p.AmountPool)
	metric := plugin.Metric{}
	metrics := []plugin.Metric{}

	for _, ns := range nss {
		metricName := ns.Element(len(ns) - 1).Value
		//For all uid
		if ns[3].Value == "*" {
			for i := 0; i < amountPool; i++ {

				newNS := make([]plugin.NamespaceElement, len(ns))
				copy(newNS, ns)
				newNS[3].Value = C.GoString(&p.ArrayPools[i].pool_uid[0])

				metric = p.getPoolValueOfProperty(i, metricName, newNS)
				fmt.Println(metric.Namespace)
				fmt.Println(metric.Data)
				metrics = append(metrics, metric)
			}

		} else { //For specific uid
			newNS := make([]plugin.NamespaceElement, len(ns))
			copy(newNS, ns)

			//Check where in ArrayPools is requested UID
			for i := 0; i < amountPool; i++ {
				if ns[3].Value == C.GoString(&p.ArrayPools[i].pool_uid[0]) {
					metric = p.getPoolValueOfProperty(i, metricName, newNS)
					fmt.Println(metric.Namespace)
					fmt.Println(metric.Data)
					metrics = append(metrics, metric)
				}
			}
		}
	}
	return metrics
}

func (p *Pool) getPoolValueOfProperty(i int, metricName string, ns []plugin.NamespaceElement) plugin.Metric {

	var v uint
	var v16I C.NVM_INT16
	var v16 C.NVM_UINT16
	var v64 C.NVM_UINT64

	switch metricName {
	case "capacity": //Metrics for Pool
		v64 = p.ArrayPools[i].capacity
		v = uint(v64)
		metric := plugin.Metric{
			Namespace: ns,
			Data:      v,
		}
		return metric
	case "free_capacity":
		v64 = p.ArrayPools[i].free_capacity
		v = uint(v64)
		metric := plugin.Metric{
			Namespace: ns,
			Data:      v,
		}
		return metric
	case "nvdimm_count":
		v16 = p.ArrayPools[i].dimm_count
		v = uint(v16)
		metric := plugin.Metric{
			Namespace: ns,
			Data:      v,
		}
		return metric
	case "interleave_set_count":
		v16 = p.ArrayPools[i].ilset_count
		v = uint(v16)
		metric := plugin.Metric{
			Namespace: ns,
			Data:      v,
		}
		return metric
	case "socket_id":
		v16I = p.ArrayPools[i].socket_id
		v = uint(v16I)
		metric := plugin.Metric{
			Namespace: ns,
			Data:      v,
		}
		return metric
	case "pool_health":
		health := PoolHealthString(int(p.ArrayPools[i].health)).String()
		metric := plugin.Metric{
			Namespace: ns,
			Data:      health,
		}
		return metric
	default:
		fmt.Println("No exist metric")
		metric := plugin.Metric{}
		return metric
	}
}
