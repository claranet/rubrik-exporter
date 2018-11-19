//
// rubrik-exporter
//
// Exports metrics from rubrik backup for prometheus
//
// License: Apache License Version 2.0,
// Organization: Claranet GmbH
// Author: Martin Weber <martin.weber@de.clara.net>
//

package main

import (
	"strings"

	"github.com/claranet/rubrik-exporter/rubrik"
	"github.com/prometheus/client_golang/prometheus"
)

// VMStats ...
type VMStats struct {
	VMIsProtected         *prometheus.GaugeVec
	VMLogicalBytes        *prometheus.GaugeVec
	VMIngestedBytes       *prometheus.GaugeVec
	VMExclusiveBytes      *prometheus.GaugeVec
	VMSharedPhysicalbytes *prometheus.GaugeVec
	VMIndexStorageBytes   *prometheus.GaugeVec
}

// Describe ...
func (e VMStats) Describe(ch chan<- *prometheus.Desc) {
	e.VMIsProtected.Describe(ch)
	e.VMExclusiveBytes.Describe(ch)
	e.VMIndexStorageBytes.Describe(ch)
	e.VMIngestedBytes.Describe(ch)
	e.VMLogicalBytes.Describe(ch)
	e.VMSharedPhysicalbytes.Describe(ch)
}

// Collect ...
func (e *VMStats) Collect(ch chan<- prometheus.Metric) {
	storages := make(map[string]rubrik.VmStorage)

	for _, s := range rubrikAPI.GetPerVMStorage() {
		storages[s.ID] = s
	}

	vms := rubrikAPI.ListAllVM()
	for _, vm := range vms {
		shortID := strings.Split(vm.ID, ":::")[1]
		strg := storages[shortID]

		var g prometheus.Gauge

		g = e.VMIsProtected.WithLabelValues(vm.Name, vm.ID)
		if vm.EffectiveSLADomainID == "UNPROTECTED" {
			g.Set(0)
		} else {
			g.Set(1)
		}
		g.Collect(ch)

		g = e.VMExclusiveBytes.WithLabelValues(vm.Name, vm.ID)
		g.Set(float64(strg.ExclusivePhysicalBytes))
		g.Collect(ch)
		g = e.VMIndexStorageBytes.WithLabelValues(vm.Name, vm.ID)
		g.Set(float64(strg.IndexStorageBytes))
		g.Collect(ch)
		g = e.VMIngestedBytes.WithLabelValues(vm.Name, vm.ID)
		g.Set(float64(strg.IndexStorageBytes))
		g.Collect(ch)
		g = e.VMLogicalBytes.WithLabelValues(vm.Name, vm.ID)
		g.Set(float64(strg.Logicalbytes))
		g.Collect(ch)
		g = e.VMSharedPhysicalbytes.WithLabelValues(vm.Name, vm.ID)
		g.Set(float64(strg.SharedPhysicalBytes))
		g.Collect(ch)
	}

}

// NewVMStatsExport ...
func NewVMStatsExport() *VMStats {
	return &VMStats{
		VMIsProtected: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "vm_protected",
			Help: "...",
		}, []string{"vmname", "vmid"}),
		VMExclusiveBytes: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "vm_consumed_exclusive_bytes",
			Help: "...",
		}, []string{"vmname", "vmid"}),
		VMIndexStorageBytes: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "vm_consumed_index_storage_bytes",
			Help: "...",
		}, []string{"vmname", "vmid"}),
		VMIngestedBytes: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "vm_consumed_ingested_bytes",
			Help: "...",
		}, []string{"vmname", "vmid"}),
		VMLogicalBytes: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "vm_consumed_logical_bytes",
			Help: "...",
		}, []string{"vmname", "vmid"}),
		VMSharedPhysicalbytes: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "vm_consumed_shared_physical_bytes",
			Help: "...",
		}, []string{"vmname", "vmid"}),
	}
}
