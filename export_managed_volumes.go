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
	"github.com/prometheus/client_golang/prometheus"
)

// ArchiveLocation ...
type ManagedVolume struct {
	SnapshotCount *prometheus.GaugeVec
	UsedSize      *prometheus.GaugeVec
	VolumeSize    *prometheus.GaugeVec
}

// Describe ...
func (e ManagedVolume) Describe(ch chan<- *prometheus.Desc) {
	e.SnapshotCount.Describe(ch)
	e.UsedSize.Describe(ch)
	e.VolumeSize.Describe(ch)
}

// Collect ...
func (e *ManagedVolume) Collect(ch chan<- prometheus.Metric) {

	volumes := rubrikAPI.GetManagedVolumes()
	for _, l := range volumes {

		var g prometheus.Gauge

		g = e.SnapshotCount.WithLabelValues(l.Name, l.ID, l.State)
		g.Set(float64(l.SnapshotCount))
		g.Collect(ch)
		g = e.VolumeSize.WithLabelValues(l.Name, l.ID, l.State)
		g.Set(l.VolumeSize)
		g.Collect(ch)
		g = e.UsedSize.WithLabelValues(l.Name, l.ID, l.State)
		g.Set(l.UsedSize)
		g.Collect(ch)
	}

}

// NewAManagedVolume ...
func NewManagedVolume() *ManagedVolume {
	return &ManagedVolume{
		SnapshotCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "managed_volume_snapshot_count",
			Help: "Snapshot Count on given Volume",
		}, []string{"name", "id", "state"}),
		UsedSize: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "managed_volume_used_size_bytes",
			Help: "Used size on Volume in bytes",
		}, []string{"name", "id", "state"}),
		VolumeSize: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "managed_volume_size_bytes",
			Help: "Available size on volume in bytes",
		}, []string{"name", "id", "state"}),
	}

}
