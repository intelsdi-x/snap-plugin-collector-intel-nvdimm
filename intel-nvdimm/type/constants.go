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
	log "github.com/sirupsen/logrus"
)

// #cgo LDFLAGS: -L/lib64 -lixpdimm
// #include <nvm_management.h>
// #include <nvm_types.h>
import "C"

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
	HEALTH_UNKNOWN        InterleaveSetHealth = 0
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

func logError(code int) {
	switch code {
	case C.NVM_ERR_INVALIDPERMISSIONS:
		log.Error("Invalid permissions")
	case C.NVM_ERR_NOTSUPPORTED:
		log.Error("This method is not supported in the current context")
	case C.NVM_ERR_NOMEMORY:
		log.Error("Not enough memory to complete the requested operation")
	case C.NVM_ERR_UNKNOWN:
		log.Error("An unknown error occurred")
	case C.NVM_ERR_NOSIMULATOR:
		log.Error("No simulator is loaded")
	case C.NVM_ERR_BADDRIVER:
		log.Error("The underlying software is missing or incompatible")
	case C.NVM_ERR_NOTMANAGEABLE:
		log.Error("The device is not manageable by the management software")
	case C.NVM_ERR_DATATRANSFERERROR:
		log.Error("There was an error in the data transfer")
	case C.NVM_SUCCESS:
		log.Debug("Success")
	case C.NVM_ERR_BADERRORCODE:
		log.Error("The return code was not valid")
	case C.NVM_ERR_DEVICEERROR:
		log.Error("There was an internal error in the device")
	case C.NVM_ERR_DEVICEBUSY:
		log.Error("The device is currently busy processing a long operation command")
	case C.NVM_ERR_BADPASSPHRASE:
		log.Error("The passphrase is not valid")
	case C.NVM_ERR_INVALIDPASSPHRASE:
		log.Error("The new passphrase does not meet the minimum requirements")
	case C.NVM_ERR_SECURITYFROZEN:
		log.Error("No changes can be made to the security state of the device")
	case C.NVM_ERR_LIMITPASSPHRASE:
		log.Error("The maximum passphrase submission limit has been reached")
	case C.NVM_ERR_SECURITYDISABLED:
		log.Error("Data at rest security is not enabled")
	case C.NVM_ERR_BADDEVICE:
		log.Error("The device identifier is not valid")
	case C.NVM_ERR_ARRAYTOOSMALL:
		log.Error("The array is not big enough")
	case C.NVM_ERR_BADCALLBACK:
		log.Error("The callback identifier is not valid")
	case C.NVM_ERR_BADFILE:
		log.Error("The file is not valid")
	case C.NVM_ERR_BADPOOL:
		log.Error("The pool identifier is not valid")
	case C.NVM_ERR_BADNAMESPACE:
		log.Error("The namespace identifier is not valid")
	case C.NVM_ERR_BADBLOCKSIZE:
		log.Error("The specified block size is not valid")
	case C.NVM_ERR_BADSIZE:
		log.Error("The size specified is not valid")
	case C.NVM_ERR_BADFIRMWARE:
		log.Error("The firmware image is not valid for the device")
	case C.NVM_ERR_DRIVERFAILED:
		log.Error("The device driver failed the requested operation")
	case C.NVM_ERR_BADSOCKET:
		log.Error("The processor socket identifier is not valid")
	case C.NVM_ERR_BADSECURITYSTATE:
		log.Error("Device security state does not permit the request")
	case C.NVM_ERR_REQUIRESFORCE:
		log.Error("This method requires the force flag to proceed")
	case C.NVM_ERR_NAMESPACESEXIST:
		log.Error("Existing namespaces must be deleted first")
	case C.NVM_ERR_NOTFOUND:
		log.Error("The requested item was not found")
	case C.NVM_ERR_BADDEVICECONFIG:
		log.Error("The configuration data is invalid or unrecognized.")
	case C.NVM_ERR_DRIVERNOTALLOWED:
		log.Error("Driver is not allowing this command")
	case C.NVM_ERR_BADALIGNMENT:
		log.Error("The specified size does not have the required alignment")
	case C.NVM_ERR_BADTHRESHOLD:
		log.Error("The threshold value is invalid.")
	case C.NVM_ERR_EXCEEDSMAXSUBSCRIBERS:
		log.Error("Exceeded maximum number of notify subscribers")
	case C.NVM_ERR_BADNAMESPACETYPE:
		log.Error("The specified namespace type is not valid")
	case C.NVM_ERR_BADNAMESPACEENABLESTATE:
		log.Error("The specified namespace enable state is not valid")
	case C.NVM_ERR_BADNAMESPACESETTINGS:
		log.Error("Could not create ns with specified settings")
	case C.NVM_ERR_BADPCAT:
		log.Error("The PCAT table is invalid")
	case C.NVM_ERR_TOOMANYNAMESPACES:
		log.Error("The maximum number of namespaces is already present")
	case C.NVM_ERR_CONFIGNOTSUPPORTED:
		log.Error("The requested configuration is not supported")
	case C.NVM_ERR_SKUVIOLATION:
		log.Error("The method is not supported because of a license violation")
	case C.NVM_ERR_ARSINPROGRESS:
		log.Error("Address range scrub in progress")
	case C.NVM_ERR_BADSECURITYGOAL:
		log.Error("No dimm found with matching security goal to create a NS")
	case C.NVM_ERR_INVALIDPASSPHRASEFILE:
		log.Error("The passphrase file is invalid")
	case C.NVM_ERR_GOALPENDING:
		log.Error("Memory allocation goal is pending reboot")
	case C.NVM_ERR_BADPOOLHEALTH:
		log.Error("Underlying persistent memory is unavailable")
	case C.NVM_ERR_INVALIDMEMORYTYPE:
		log.Error("The address does not match the specified memory type")
	case C.NVM_ERR_INCOMPATIBLEFW:
		log.Error("The firmware image is not compatible with this version of software")
	case C.NVM_ERR_NAMESPACEBUSY:
		log.Error("The namespace cannot be changed because it is in use by a file system")
	case C.NVM_ERR_FWALREADYSTAGED:
		log.Error("A firmware image is already staged for execution. A power cycle is required before another can be staged.")
	case C.NVM_ERR_BADNFIT:
		log.Error("The NFIT table is invalid")
	default:
		log.Error("Unrecognized error code. Please notify us.")
	}
}

func convertSensorUnits(unit int) string {
	var unit_name string
	switch unit {
	case 1: // UNIT_COUNT
		unit_name = ""
	case 2: // UNIT_CELSIUS
		unit_name = "C"
	case 21: // UNIT_SECONDS
		unit_name = "s"
	case 22: // UNIT_MINUTES
		unit_name = "m"
	case 23: // UNIT_HOURS
		unit_name = "h"
	case 39: // UNIT_CYCLES
		unit_name = ""
	case 65: // UNIT_PERCENT
		unit_name = "%"
	}
	return unit_name
}
