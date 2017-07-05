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

//go:generate stringer -type=EncryptionString
//go:generate stringer -type=PoolHealthString
//go:generate stringer -type=InterleaveSetHealth
//go:generate stringer -type=TypeNamespace
//go:generate stringer -type=HealthNamespace
//go:generate stringer -type=StateNamespace

type EncryptionString int
type PoolHealthString int
type InterleaveSetHealth int
type TypeNamespace int
type HealthNamespace int
type StateNamespace int

const (
	NVM_ENCRYPTION_OFF    EncryptionString    = 0
	NVM_ENCRYPTION_ON     EncryptionString    = 1
	NVM_ENCRYPTION_IGNORE EncryptionString    = 2
	UNKNOWN               PoolHealthString    = 0 // The pool health cannot be determined.
	NORMAL                PoolHealthString    = 1 // All underlying AEP DIMM Persistent memory capacity is available.
	PENDING               PoolHealthString    = 2 // A new memory allocation goal has been created but not applied.
	ERROR                 PoolHealthString    = 3 // There is an issue with some or all of the underlying
	LOCKED                PoolHealthString    = 4 // One or more of the underlying AEP DIMMs are locked.
	HEAHLTH_UNKNOWN       InterleaveSetHealth = 0
	HEALTH_NORMAL         InterleaveSetHealth = 1 // Available and underlying AEP DIMMs have good health.
	DEGRADED              InterleaveSetHealth = 2 // In danger of failure, may have degraded performance.
	FAILED                InterleaveSetHealth = 3 // Interleave set has failed and is unavailable.

	TYPE_UNKNOWN    TypeNamespace = 0 // Type cannot be determined
	TYPE_STORAGE    TypeNamespace = 1 // Storage namespace
	TYPE_APP_DIRECT TypeNamespace = 2 // App Direct namespace

	NAMESPACE_HEALTH_UNKNOWN      HealthNamespace = 0     // Namespace health cannot be determined
	NAMESPACE_HEALTH_NORMAL       HealthNamespace = 5     // Namespace is OK
	NAMESPACE_HEALTH_NONCRITICAL  HealthNamespace = 10    // Non-critical health issue
	NAMESPACE_HEALTH_CRITICAL     HealthNamespace = 25    // Critical health issue
	NAMESPACE_HEALTH_BROKENMIRROR HealthNamespace = 65535 // Broken mirror

	STATE_UNKNOWN  StateNamespace = 0 // Cannot be determined
	STATE_ENABLED  StateNamespace = 2 // Exposed to OS
	STATE_DISABLED StateNamespace = 3 // Hidden from OS
)
