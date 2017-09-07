# snap plugin collector - Intel NVDIMM

## Collected Metrics
This plugin has the ability to gather the following metrics:

Metric Name | Unit | Description
----------- | ---- | -------------------
/intel/nvdimm/pool/*/capacity | UINT64 | The total user-visible capacity of the pool in bytes. If the pool is mirrored, the user-visible capacity is less than the capacity occupied on the NVDIMMs
/intel/nvdimm/pool/*/free_capacity | UINT64 | The capacity in bytes that is not currently being used in a namespace
/intel/nvdimm/pool/*/socket_id | INT16 | The processor socket identifer where the pool resides. -1 for a system level pool
/intel/nvdimm/pool/*/nvdimm_count | UINT16 | The number of underlying NVDIMMs on which the pool capacity reside
/intel/nvdimm/pool/*/interleave_set_count | UINT16 | The number of interleave sets in the pool
/intel/nvdimm/pool/*/pool_health | enum | UNKNOWN Health cannot be determined. NORMAL All underlying NVDIMMs are available. DEGRADED One or more underlying NVDIMMs are missing or failed but the capacity is mirrored. FAILED One or more underlying NVDIMMs are missing or failed.
| |
/intel/nvdimm/interleave set/*/pool_id | GUID | Unique identifer of pool consisting this interleave set.
/intel/nvdimm/interleave set/*/size | UINT64 | The total size of the interleave set in bytes.
/intel/nvdimm/interleave set/*/available_size | UINT64 | The capacity in bytes not yet allocated to a namespace.
/intel/nvdimm/interleave set/*/dimm_count | UINT8 | The number of NVDIMMs in this interleave set
/intel/nvdimm/interleave set/*/mirrored | UINT8 | Whether the interleave set is mirrored.
/intel/nvdimm/interleave set/*/interleave_set_health | enum | UNKNOWN Health cannot be determined. NORMAL All underlying NVDIMMs are available. DEGRADED One or more underlying NVDIMMs are missing or failed but the capacity is mirrored. FAILED One or more underlying NVDIMMs are missing or failed.
/intel/nvdimm/interleave set/*/encryption_status | enum | Lockstates of all NVDIMMs
| |
/intel/nvdimm/namespace/*/pool_id | GUID | Unique identifer of pool contain-ing this namespace
/intel/nvdimm/namespace/*/type | enum | UNKNOWN Cannot be determined STORAGE Storage APP DIRECT App-Direct
/intel/nvdimm/namespace/*/block size | UINT32 | The logical size in bytes for read/write operations
/intel/nvdimm/namespace/*/block_count | UINT64 | The total number of blocks of memory that make up in the namespace
/intel/nvdimm/namespace/*/health | enum | UNKNOWN Health cannot be determined. NORMAL All underlying NVDIMMs are available. NONCRITICAL Non-critical health issue reported CRITICAL Critical health issue reported
/intel/nvdimm/namespace/*/enable_state | enum | UNKNOWN The enable state cannot be determined ENABLED The namespace is exposed to the operating system DISABLED The namespace is hidden from the operating system
/intel/nvdimm/namespace/*/btt | UINT8 | The namespace is optimized for speed vs. resources. NOTE: If the namespace is disabled, the btt state will always be 0
| |
/intel/nvdimm/nvdimm/*/interleave_set_id | UINT16 | Unique identifer of interleave set containing this device.
/intel/nvdimm/nvdimm/*/pool id | UINT16 | Unique identifer of pool containing this device.
/intel/nvdimm/nvdimm/*/channel pos | UINT8 | NVDIMM number in the memory channel
/intel/nvdimm/nvdimm/*/channel id | UINT8 | The memory channel number
/intel/nvdimm/nvdimm/*/memory_controller_id | UINT16 | The associated memory controller identifier
/intel/nvdimm/nvdimm/*/socket_id | UINT16 | The processor socket identifier where the NVDIMM is installed
/intel/nvdimm/nvdimm/*/device_capacity | UINT64 | Raw device capacity in bytes
/intel/nvdimm/nvdimm/*/volatile_capacity | UINT64 | Capacity configured to be used as memory mode
/intel/nvdimm/nvdimm/*/persistent_capacity | UINT64 | Capacity configured to be used as persisted mode
/intel/nvdimm/nvdimm/*/unconfigured_capacity | UINT64 | Capacity that requires further configuration
/intel/nvdimm/nvdimm/*/inaccessible_capacity | UINT64 | Capacity not supported or licensed by SKU
/intel/nvdimm/nvdimm/*/reserved_capacity | UINT64 | Reserved capacity for metadata and alignment
/intel/nvdimm/nvdimm/*/speed | UINT64 |The spped in nanoseconds
/intel/nvdimm/nvdimm/*/power_management_enabled | UINT8 | Enablement state of firmware power management policy
/intel/nvdimm/nvdimm/*/power_limit | UINT8 | Power limit in Watts
/intel/nvdimm/nvdimm/*/peak_power_budget | UINT16 | Power budget in mW used for instantaneous power
/intel/nvdimm/nvdimm/*/avg_power_budget | UINT16 | Power budget in mW used for averaged power
/intel/nvdimm/nvdimm/*/bytes_read | UINT64 | Total bytes read from NVDIMM
/intel/nvdimm/nvdimm/*/host_reads | UINT64 | Total number of read requests
/intel/nvdimm/nvdimm/*/bytes_written | UINT64 | Total number of bytes written
/intel/nvdimm/nvdimm/*/host_writes | UINT64 | Total number of write requests
/intel/nvdimm/nvdimm/*/spare_capacity | UINT64 | The percentage of spare capacity remaining
/intel/nvdimm/nvdimm/*/wear_level | UINT64 | An estimate of the NVDIMM life as a percentage
/intel/nvdimm/nvdimm/*/power_cycles | UINT64 | The number of power cycles over the lifetime of the NVDIMM
/intel/nvdimm/nvdimm/*/unsafe_shutdowns | UINT64 | The number of shutdowns without notification over the lifetime of the NVDIMM
/intel/nvdimm/nvdimm/*/media_errors_uncorrectable | UINT64 | The number of ECC uncorrectable errors
/intel/nvdimm/nvdimm/*/media_errors_corrected | UINT64 | The number of ECC corrected errors
