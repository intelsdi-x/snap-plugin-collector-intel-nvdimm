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
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	log "github.com/sirupsen/logrus"
)

// #cgo LDFLAGS: -L/lib64 -lixpdimm
// #include <nvm_management.h>
// #include <nvm_types.h>
import "C"

var interleaveLabels = map[string]label{
	"size": label{
		description: "",
		unit:        "",
	},
	"available_size": label{
		description: "",
		unit:        "",
	},
	"dimm_count": label{
		description: "",
		unit:        "",
	},
	"mirrored": label{
		description: "",
		unit:        "",
	},
	"interleave_set_health": label{
		description: "",
		unit:        "",
	},
	"encryption_status": label{
		description: "",
		unit:        "",
	},
}

//Main function for getting metrics from Pool
func (p *Pool) getInterleavesetMetric(nss []plugin.Namespace) []plugin.Metric {
	metric := plugin.Metric{}
	metrics := []plugin.Metric{}

	for _, ns := range nss {
		metricName := ns.Element(len(ns) - 1).Value
		if ns[3].Value == "*" {
			for i, array := range p.ArrayPools {
				newNS := plugin.CopyNamespace(ns)
				newNS[3].Value = C.GoString(&array.pool_uid[0])

				metric = p.getInterleavesetValueOfProperty(i, metricName, newNS)
				metrics = append(metrics, metric)
			}
		} else { //For specific uid
			newNS := plugin.CopyNamespace(ns)
			//Check where in ArrayPools is requested UID
			for i, pool := range p.ArrayPools {
				if ns[3].Value == C.GoString(&pool.pool_uid[0]) {
					metric = p.getInterleavesetValueOfProperty(i, metricName, newNS)
					metrics = append(metrics, metric)
				}
			}
		}
	}
	return metrics
}

func (p *Pool) getInterleavesetValueOfProperty(i int, metricName string, ns []plugin.NamespaceElement) plugin.Metric {
	var v uint
	var v64 C.NVM_UINT64

	switch metricName {
	case "size": //Metrics for Interleave_set
		v64 = p.ArrayPools[i].ilsets[0].size
		v = uint(v64)
		metric := plugin.Metric{
			Namespace: ns,
			Data:      v,
		}
		return metric
	case "available_size":
		v64 = p.ArrayPools[i].ilsets[0].available_size
		v = uint(v64)
		metric := plugin.Metric{
			Namespace: ns,
			Data:      v,
		}
		return metric
	case "dimm_count":
		v := uint8(p.ArrayPools[i].ilsets[0].dimm_count)
		metric := plugin.Metric{
			Namespace: ns,
			Data:      v,
		}
		return metric
	case "mirrored":
		v := uint8(p.ArrayPools[i].ilsets[0].mirrored)
		metric := plugin.Metric{
			Namespace: ns,
			Data:      v,
		}
		return metric
	case "interleave_set_health":
		health := InterleaveSetHealth(int(p.ArrayPools[i].ilsets[0].health)).String()
		metric := plugin.Metric{
			Namespace: ns,
			Data:      health,
		}
		return metric
	case "encryption_status":
		vEnum := EncryptionString(int(p.ArrayPools[i].ilsets[0].encryption)).String()
		metric := plugin.Metric{
			Namespace: ns,
			Data:      vEnum,
		}
		return metric

	default:
		log.Debug("No exist metric")
		metric := plugin.Metric{}
		return metric
	}
}
