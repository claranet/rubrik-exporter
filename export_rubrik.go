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

// RubrikStats ...
type RubrikStats struct {
	StreamCount          *prometheus.GaugeVec
	AverageStorageGrowth *prometheus.GaugeVec
	RunawayRemaining     *prometheus.GaugeVec

	SucceededTask *prometheus.GaugeVec
	FailedTask    *prometheus.GaugeVec
	CancledTask   *prometheus.GaugeVec

	NodeCount              *prometheus.GaugeVec
	NodeNetworkReceived    *prometheus.GaugeVec
	NodeNetworkTransmitted *prometheus.GaugeVec
	NodeIOPRead            *prometheus.GaugeVec
	NodeIOPWrite           *prometheus.GaugeVec
	NodeThroughputRead     *prometheus.GaugeVec
	NodeThroughputWrite    *prometheus.GaugeVec

	SystemPhysicalIngest *prometheus.GaugeVec

	SystemStorageTotal         *prometheus.GaugeVec
	SystemStorageUsed          *prometheus.GaugeVec
	SystemStorageAvailable     *prometheus.GaugeVec
	SystemStorageSnapshot      *prometheus.GaugeVec
	SystemStorageLiveMount     *prometheus.GaugeVec
	SystemStorageMiscellaneous *prometheus.GaugeVec

	ArchiveStorageBandwith        *prometheus.GaugeVec
	ArchiveStorageArchivedVM      *prometheus.GaugeVec
	ArchiveStorageArchivedFileSet *prometheus.GaugeVec
	ArchiveStorageDataDownloaded  *prometheus.GaugeVec
	ArchiveStorageDataArchived    *prometheus.GaugeVec
}

// Describe ...
func (e *RubrikStats) Describe(ch chan<- *prometheus.Desc) {
	e.StreamCount.Describe(ch)
	e.AverageStorageGrowth.Describe(ch)
	e.RunawayRemaining.Describe(ch)

	e.SucceededTask.Describe(ch)
	e.FailedTask.Describe(ch)
	e.CancledTask.Describe(ch)

	e.NodeCount.Describe(ch)
	e.NodeNetworkReceived.Describe(ch)
	e.NodeNetworkTransmitted.Describe(ch)
	e.NodeIOPRead.Describe(ch)
	e.NodeIOPWrite.Describe(ch)
	e.NodeThroughputRead.Describe(ch)
	e.NodeThroughputWrite.Describe(ch)

	e.SystemPhysicalIngest.Describe(ch)

	e.SystemStorageTotal.Describe(ch)
	e.SystemStorageUsed.Describe(ch)
	e.SystemStorageAvailable.Describe(ch)
	e.SystemStorageSnapshot.Describe(ch)
	e.SystemStorageLiveMount.Describe(ch)
	e.SystemStorageMiscellaneous.Describe(ch)

	e.ArchiveStorageBandwith.Describe(ch)
	e.ArchiveStorageArchivedFileSet.Describe(ch)
	e.ArchiveStorageArchivedVM.Describe(ch)
	e.ArchiveStorageDataArchived.Describe(ch)
	e.ArchiveStorageDataDownloaded.Describe(ch)
}

// Collect ...
func (e *RubrikStats) Collect(ch chan<- prometheus.Metric) {
	var g prometheus.Gauge

	g = e.StreamCount.WithLabelValues()
	g.Set(float64(rubrikAPI.GetStreamCount()))
	g.Collect(ch)
	g = e.RunawayRemaining.WithLabelValues()
	g.Set(float64(rubrikAPI.GetRunawayRemaining()))
	g.Collect(ch)
	g = e.AverageStorageGrowth.WithLabelValues()
	g.Set(float64(rubrikAPI.GetAverageStorageGrowthPerDay()))
	g.Collect(ch)

	taskStat := rubrikAPI.GetTaskDetails()

	g = e.SucceededTask.WithLabelValues()
	g.Set(taskStat["succeeded"])
	g.Collect(ch)
	g = e.FailedTask.WithLabelValues()
	g.Set(taskStat["failed"])
	g.Collect(ch)
	g = e.CancledTask.WithLabelValues()
	g.Set(taskStat["cancled"])
	g.Collect(ch)

	nodes := rubrikAPI.GetNodes()
	{
		_nodes := make(map[string]int)
		for _, n := range nodes {
			if _, ok := _nodes[n.BrikID]; !ok {
				_nodes[n.BrikID] = 0
			}
			_nodes[n.BrikID]++
		}
		for bID, c := range _nodes {
			g := e.NodeCount.WithLabelValues(bID)
			g.Set(float64(c))
			g.Collect(ch)
		}
	}

	for _, v := range nodes {
		nodeStat := rubrikAPI.GetNodeStats(v.ID)

		g = e.NodeNetworkReceived.WithLabelValues(v.ID)
		g.Set(float64(nodeStat.NetworkStat.BytesReceived[0].Stat))
		g.Collect(ch)
		g = e.NodeNetworkTransmitted.WithLabelValues(v.ID)
		g.Set(float64(nodeStat.NetworkStat.BytesTransmitted[0].Stat))
		g.Collect(ch)

		g = e.NodeIOPRead.WithLabelValues(v.ID)
		g.Set(float64(nodeStat.Iops.ReadsPerSecond[0].Stat))
		g.Collect(ch)
		g = e.NodeIOPWrite.WithLabelValues(v.ID)
		g.Set(float64(nodeStat.Iops.WritesPerSecond[0].Stat))
		g.Collect(ch)

		g = e.NodeThroughputRead.WithLabelValues(v.ID)
		g.Set(float64(nodeStat.IOThroughput.ReadBytePerSecond[0].Stat))
		g.Collect(ch)
		g = e.NodeThroughputWrite.WithLabelValues(v.ID)
		g.Set(float64(nodeStat.IOThroughput.WriteBytePerSecond[0].Stat))
		g.Collect(ch)

	}

	systemStorage := rubrikAPI.GetSystemStorage()

	g = e.SystemStorageAvailable.WithLabelValues()
	g.Set(float64(systemStorage.Available))
	g.Collect(ch)
	g = e.SystemStorageLiveMount.WithLabelValues()
	g.Set(float64(systemStorage.LiveMount))
	g.Collect(ch)
	g = e.SystemStorageMiscellaneous.WithLabelValues()
	g.Set(float64(systemStorage.Miscellaneous))
	g.Collect(ch)
	g = e.SystemStorageSnapshot.WithLabelValues()
	g.Set(float64(systemStorage.Snapshot))
	g.Collect(ch)
	g = e.SystemStorageTotal.WithLabelValues()
	g.Set(float64(systemStorage.Total))
	g.Collect(ch)
	g = e.SystemStorageUsed.WithLabelValues()
	g.Set(float64(systemStorage.Used))
	g.Collect(ch)

	locations := rubrikAPI.GetArchiveLocations()
	usages := rubrikAPI.GetDataLocationUsage()
	for _, l := range locations {
		var usage rubrik.DataLocationUsage
		for _, u := range usages {
			if u.LocationID == l.ID {
				usage = u
				break
			}
		}

		g = e.ArchiveStorageBandwith.WithLabelValues(l.Name, l.IPAddress)
		val := rubrikAPI.GetArchivalBandwith(l.ID, "-10min")[0].Stat
		g.Set(float64(val))
		g.Collect(ch)

		g = e.ArchiveStorageDataArchived.WithLabelValues(l.Name, l.IPAddress)
		g.Set(float64(usage.DataArchived))
		g.Collect(ch)
		g = e.ArchiveStorageDataDownloaded.WithLabelValues(l.Name, l.IPAddress)
		g.Set(float64(usage.DataDownloaded))
		g.Collect(ch)

		g = e.ArchiveStorageArchivedVM.WithLabelValues(l.Name, l.IPAddress, "vmware")
		g.Set(float64(usage.NumVMsArchived))
		g.Collect(ch)
		g = e.ArchiveStorageArchivedVM.WithLabelValues(l.Name, l.IPAddress, "nutanix")
		g.Set(float64(usage.NumNutanixVmsArchived))
		g.Collect(ch)
		g = e.ArchiveStorageArchivedVM.WithLabelValues(l.Name, l.IPAddress, "hyperv")
		g.Set(float64(usage.NumHypervVmsArchived))
		g.Collect(ch)

		g = e.ArchiveStorageArchivedFileSet.WithLabelValues(l.Name, l.IPAddress, "linux")
		g.Set(float64(usage.NumLinuxFilesetsArchived))
		g.Collect(ch)
		g = e.ArchiveStorageArchivedFileSet.WithLabelValues(l.Name, l.IPAddress, "windows")
		g.Set(float64(usage.NumWindowsFilesetsArchived))
		g.Collect(ch)
		g = e.ArchiveStorageArchivedFileSet.WithLabelValues(l.Name, l.IPAddress, "share")
		g.Set(float64(usage.NumShareFilesetsArchived))
		g.Collect(ch)
		g = e.ArchiveStorageArchivedFileSet.WithLabelValues(l.Name, l.IPAddress, "fileset")
		g.Set(float64(usage.NumFilesetsArchived))
		g.Collect(ch)
	}

	ingest := rubrikAPI.GetPhysicalIngest()
	g = e.SystemPhysicalIngest.WithLabelValues()
	g.Set(float64(ingest[0].Stat))
	g.Collect(ch)
}

// NewRubrikStatsExport ...
func NewRubrikStatsExport() *RubrikStats {
	return &RubrikStats{
		StreamCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "count_streams",
			Help: "Count Rubrik Backup Streams",
		}, []string{}),
		AverageStorageGrowth: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "stat_average_storage_growth_per_day",
			Help: "Get average storage growth per day in bytes",
		}, []string{}),
		RunawayRemaining: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "stat_runaway_remaining",
			Help: "Get the number of days remaining before the system fills up",
		}, []string{}),

		SucceededTask: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "report_task_succeded",
			Help: "...",
		}, []string{}),
		FailedTask: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "report_task_failed",
			Help: "...",
		}, []string{}),
		CancledTask: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "report_task_cancled",
			Help: "...",
		}, []string{}),

		NodeCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "count_nodes",
			Help: "Count Rubrik Nodes in a Brick",
		}, []string{"brik"}),
		NodeNetworkReceived: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "node_network_received",
			Help: "Node Network Byte received",
		}, []string{"node"}),
		NodeNetworkTransmitted: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "node_network_transmitted",
			Help: "Node Network Byte transmitted",
		}, []string{"node"}),

		NodeIOPRead: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "node_io_read",
			Help: "Node Read IO per second",
		}, []string{"node"}),
		NodeIOPWrite: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "node_io_write",
			Help: "Node Write IO per second",
		}, []string{"node"}),
		NodeThroughputRead: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "node_throughput_read",
			Help: "Node Read Throughput per second",
		}, []string{"node"}),
		NodeThroughputWrite: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "node_throughput_write",
			Help: "Node Write Throughput per second",
		}, []string{"node"}),

		SystemPhysicalIngest: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "system_physical_ingest_bytes",
			Help: "...",
		}, []string{}),

		SystemStorageAvailable: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "system_storage_available",
			Help: "Available Storage Bytes",
		}, []string{}),
		SystemStorageLiveMount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "system_storage_live_mount",
			Help: "...",
		}, []string{}),
		SystemStorageMiscellaneous: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "system_storage_miscellaneous",
			Help: "...",
		}, []string{}),
		SystemStorageSnapshot: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "system_storage_snapshot",
			Help: "storage bytes used by snapshots",
		}, []string{}),
		SystemStorageTotal: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "system_storage_total",
			Help: "total available bytes",
		}, []string{}),
		SystemStorageUsed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "system_storage_used",
			Help: "used bytes on storage",
		}, []string{}),

		ArchiveStorageBandwith: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "archive_storage_bandwidth",
			Help: "...",
		}, []string{"name", "target"}),
		ArchiveStorageArchivedFileSet: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "archive_storage_archived_fileset",
			Help: "...",
		}, []string{"name", "target", "type"}),
		ArchiveStorageArchivedVM: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "archive_storage_archived_vm",
			Help: "...",
		}, []string{"name", "target", "type"}),
		ArchiveStorageDataArchived: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "archive_storage_data_archived",
			Help: "...",
		}, []string{"name", "target"}),
		ArchiveStorageDataDownloaded: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace, Name: "archive_storage_data_downloaded",
			Help: "...",
		}, []string{"name", "target"}),
	}
}
