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
    "strconv"
	"unsafe"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	
	log "github.com/sirupsen/logrus"
)

// #cgo LDFLAGS: -L/lib64 -lixpdimm
// #include <nvm_management.h>
// #include <nvm_types.h>
// int getSensorType(struct device_details dev_details, int i) {
//     return (int)dev_details.sensors[i].type;
// }
// NVM_UINT64 getSensorReading(struct device_details dev_details, int sensor_type) {
//     int i;
//     for (i = 0; i < NVM_MAX_DEVICE_SENSORS; i++) {
//         if (dev_details.sensors[i].type == sensor_type) {
//              return dev_details.sensors[i].reading;
//         }
//     }
// }
// int getSensorUnit(struct device_details dev_details, int sensor_type) {
//     int i;
//     for (i = 0; i < NVM_MAX_DEVICE_SENSORS; i++) {
//         if (dev_details.sensors[i].type == sensor_type) {
//              return (int)dev_details.sensors[i].units;
//         }
//     }
// }
import "C"

type Nvdimm struct {
	Memory_topology_count  C.int
	Device_discovery_count C.int
	Memory_topology        []C.struct_memory_topology
	Device_array           []C.struct_device_discovery
	Device_details         []C.struct_device_details
}

var nvdimmLabels = map[string]label{
	"interleave_set_id": label{
		description: "",
		unit:        "",
	},
	"pool_id": label{
		description: "",
		unit:        "",
	},
	"device_capacity": label{
		description: "",
		unit:        "",
	},
	"volatile_capacity": label{
		description: "",
		unit:        "",
	},
	"persistent_capacity": label{
		description: "",
		unit:        "",
	},
	"unconfigured_capacity": label{
		description: "",
		unit:        "",
	},
	"inaccessible_capacity": label{
		description: "",
		unit:        "",
	},
	"reserved_capacity": label{
		description: "",
		unit:        "",
	},
	"speed": label{
		description: "",
		unit:        "",
	},
	"power_management_enabled": label{
		description: "",
		unit:        "",
	},
	"power_limit": label{
		description: "",
		unit:        "",
	},
	"peak_power_budget": label{
		description: "",
		unit:        "",
	},
	"avg_power_budget": label{
		description: "",
		unit:        "",
	},
	"bytes_read": label{
		description: "",
		unit:        "",
	},
	"host_reads": label{
		description: "",
		unit:        "",
	},
	"bytes_written": label{
		description: "",
		unit:        "",
	},
	"host_writes": label{
		description: "",
		unit:        "",
	},
	"spare_capacity": label{
		description: "",
		unit:        "",
	},
	"wear_level": label{
		description: "",
		unit:        "",
	},
	"power_cycles": label{
		description: "",
		unit:        "",
	},
	"unsafe_shutdowns": label{
		description: "",
		unit:        "",
	},
	"media_errors_uncorrectable": label{
		description: "",
		unit:        "",
	},
	"media_errors_corrected": label{
		description: "",
		unit:        "",
	},
	"channel_pos": label{
		description: "",
		unit:        "",
	},
	"channel_id": label{
		description: "",
		unit:        "",
	},
	"memory_controller_id": label{
		description: "",
		unit:        "",
	},
	"socket_id": label{
		description: "",
		unit:        "",
	},
}

func (nc *NvdimmCollector) DiscoverDevices() int {
	nc.Memory_topology_count = C.nvm_get_memory_topology_count()
	if nc.Memory_topology_count < 0 {
    	logError(int(nc.Memory_topology_count))
    	return int(nc.Memory_topology_count)
	} else {
		nc.Memory_topology = make([]C.struct_memory_topology, nc.Memory_topology_count) // Allocate
		memory_topology_ptr := (*C.struct_memory_topology)(unsafe.Pointer(&nc.Memory_topology[0]))
		C.nvm_get_memory_topology(memory_topology_ptr, C.NVM_UINT8(nc.Memory_topology_count))
		nc.Device_discovery_count = C.nvm_get_device_count()
		if nc.Device_discovery_count > 0 {
		    nc.Device_array = make([]C.struct_device_discovery, nc.Device_discovery_count)
			nc.Device_details = make([]C.struct_device_details, nc.Device_discovery_count)
			device_array_ptr := (*C.struct_device_discovery)(unsafe.Pointer(&nc.Device_array[0]))
			C.nvm_get_devices(device_array_ptr, C.NVM_UINT8(nc.Device_discovery_count))
			
			for i := 0; i < int(nc.Device_discovery_count); i++ {
				C.nvm_get_device_details(&nc.Device_array[i].uid[0], &nc.Device_details[i])
			}
			log.Debug("Discover process finished, discovered %d devices!", int(nc.Device_discovery_count))
		} else {
            logError(int(nc.Device_discovery_count))
		    return int(nc.Device_discovery_count)
		}
	}
	return 0;
}

func (nc *NvdimmCollector) GetNvdimmMetrics(namespaces []plugin.Namespace) []plugin.Metric {
	var results []plugin.Metric
	if nc.Memory_topology_count <= 0 {
	    return results;
	}
	for _, namespace := range namespaces {
		metricName := namespace[len(namespace)-1].Value // e.g. "capacity"
		var met plugin.Metric
		for i, elem := range nc.Device_array {
			switch metricName {
			case "device_capacity":
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(elem.capacity),
				}
			case "interleave_set_id":
				// TODO
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint16(0),
				}
			case "pool_id":
				// TODO
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint16(0),
				}
			case "channel_pos":
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint16(elem.channel_pos),
				}
			case "channel_id":
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint16(elem.channel_id),
				}
			case "memory_controller_id":
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint16(elem.memory_controller_id),
				}
			case "socket_id":
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint16(elem.socket_id),
				}
			case "volatile_capacity": // CAPACITIES
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(nc.Device_details[i].capacities.memory_capacity),
				}
			case "persistent_capacity": // CAPACITIES
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(nc.Device_details[i].capacities.app_direct_capacity),
				}
			case "unconfigured_capacity": // CAPACITIES
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(nc.Device_details[i].capacities.unconfigured_capacity),
				}
			case "inaccessible_capacity": // CAPACITIES
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(nc.Device_details[i].capacities.inaccessible_capacity),
				}
			case "reserved_capacity": // CAPACITIES
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(nc.Device_details[i].capacities.reserved_capacity),
				}
			case "speed":
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(nc.Device_details[i].speed),
				}
			case "power_management_enabled":
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint8(nc.Device_details[i].power_management_enabled),
				}
			case "power_limit":
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint8(nc.Device_details[i].power_limit),
				}
			case "peak_power_budget":
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint16(nc.Device_details[i].peak_power_budget),
				}
			case "avg_power_budget":
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint16(nc.Device_details[i].avg_power_budget),
				}
			case "bytes_read": // PERFORMANCE
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(nc.Device_details[i].performance.bytes_read),
				}
			case "host_reads": // PERFORMANCE
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(nc.Device_details[i].performance.host_reads),
				}
			case "bytes_written": // PERFORMANCE
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(nc.Device_details[i].performance.bytes_written),
				}
			case "host_writes": // PERFORMANCE
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(nc.Device_details[i].performance.host_writes),
				}
			case "spare_capacity": // SENSOR
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(C.getSensorReading(nc.Device_details[i], 1)),
					Unit: "%",
				}
			case "wear_level": // SENSOR
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(C.getSensorReading(nc.Device_details[i], 2)),
					Unit: "%",
				}
			case "power_cycles": // SENSOR
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(C.getSensorReading(nc.Device_details[i], 3)),
				}
			case "unsafe_shutdowns": // SENSOR
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(C.getSensorReading(nc.Device_details[i], 6)),
				}
			case "media_errors_uncorrectable": // SENSOR
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(C.getSensorReading(nc.Device_details[i], 9)),
				}
			case "media_errors_corrected": // SENSOR
				met = plugin.Metric{
					Namespace: plugin.NewNamespace(namespace[0].Value, namespace[1].Value,
						namespace[2].Value, strconv.Itoa(int(elem.physical_id)),
						metricName),
					Data: uint64(C.getSensorReading(nc.Device_details[i], 10)),
				}
			}
			if namespace[3].Value == "*" {
				results = append(results, met)
			} else {
				if namespace[3].Value == strconv.Itoa(int(elem.physical_id)) {
					results = append(results, met)
				}
			}
		}
	}
	return results
}
