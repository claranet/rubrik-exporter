//
// rubrik-exporter
//
// Exports metrics from rubrik backup for prometheus
//
// License: Apache License Version 2.0,
// Organization: Claranet GmbH
// Author: Martin Weber <martin.weber@de.clara.net>
//

package rubrik

import (
	"encoding/json"
)

type VirtualMachine struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	EffectiveSLADomainID string `json:"effectiveSlaDomainId"`
}

type VirtualMachineList struct {
	*ResultList
	Data []VirtualMachine `json:"data"`
}

// ListAllVM retrieves a list of all Virtual Machine ID and Name
// for All kinds of hypervisors (vmware, nutanix, hyperv)
func (r Rubrik) ListAllVM() []VirtualMachine {
	var list []VirtualMachine
	list = append(list, r.ListVmwareVM()...)
	list = append(list, r.ListNutanixVM()...)
	list = append(list, r.ListHypervVM()...)

	return list
}

// ListVmwareVM retrieve a List of all known VMware VM's
func (r Rubrik) ListVmwareVM() []VirtualMachine {
	resp, _ := r.makeRequest("GET", "/api/v1/vmware/vm", RequestParams{})

	data := json.NewDecoder(resp.Body)
	var s VirtualMachineList
	data.Decode(&s)
	return s.Data
}

// ListNutanixVM retrieve a List of all known VMware VM's
func (r Rubrik) ListNutanixVM() []VirtualMachine {
	resp, _ := r.makeRequest("GET", "/api/internal/nutanix/vm", RequestParams{})

	data := json.NewDecoder(resp.Body)
	var s VirtualMachineList
	data.Decode(&s)
	return s.Data
}

// ListHypervVM retrieve a List of all known VMware VM's
func (r Rubrik) ListHypervVM() []VirtualMachine {
	resp, _ := r.makeRequest("GET", "/api/internal/hyperv/vm", RequestParams{})

	data := json.NewDecoder(resp.Body)
	var s VirtualMachineList
	data.Decode(&s)
	return s.Data
}
