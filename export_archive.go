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
	"github.com/claranet/rubrik-exporter/rubrik"
	"github.com/prometheus/client_golang/prometheus"
)

// ArchiveLocation ...
type ArchiveLocation struct {
	ArchiveLocationStatus *prometheus.GaugeVec
}

// Describe ...
func (e ArchiveLocation) Describe(ch chan<- *prometheus.Desc) {
	e.ArchiveLocationStatus.Describe(ch)
}

// Collect ...
func (e *ArchiveLocation) Collect(ch chan<- prometheus.Metric) {
	storages := make(map[string]rubrik.VmStorage)

	for _, s := range rubrikAPI.GetPerVMStorage() {
		storages[s.ID] = s
	}

	locations := rubrikAPI.GetArchiveLocations()
	for _, l := range locations {

		var g prometheus.Gauge

		g = e.ArchiveLocationStatus.WithLabelValues(l.Name, l.Bucket, l.IPAddress)
		if l.IsActive {
			g.Set(1)
		} else {
			g.Set(0)
		}

		g.Collect(ch)
	}

}

// NewArchiveLocation ...
func NewArchiveLocation() *ArchiveLocation {
	return &ArchiveLocation{
		ArchiveLocationStatus: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "archive_location_status",
			Help: "Archive Loction Status - 1: Active, 0: Inactive",
		}, []string{"name", "bucket", "target"}),
	}

}
