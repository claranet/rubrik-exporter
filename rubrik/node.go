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
	"fmt"
	"net/url"
)

type NodeList struct {
	*ResultList
	Data []Node `json:"data"`
}

// Node - Descripe a Rubrik node
type Node struct {
	ID              string `json:"id"`
	BrikID          string `json:"brikId"`
	Status          string `json:"status"`
	IPAddress       string `json:"ipAddress"`
	NeedsInspection bool   `json:"needsInspection"`
}

type NodeStat struct {
	ID              string       `json:"id"`
	BrikID          string       `json:"brikId"`
	Status          string       `json:"status"`
	IPAddress       string       `json:"ipAddress"`
	NeedsInspection bool         `json:"needsInspection"`
	NetworkStat     NetworkStat  `json:"networkStat"`
	Iops            Iops         `json:"iops"`
	IOThroughput    IoThroughput `json:"ioThroughput"`
	CPUStat         []TimeStat   `json:"cpuStat"`
}

// GetNodes - Returns the List of all Rubrik Nodes
func (r Rubrik) GetNodes() []Node {
	resp, _ := r.makeRequest("GET", "/api/internal/node", RequestParams{})

	var l NodeList
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&l)

	return l.Data
}

// GetNodeStats ...
func (r Rubrik) GetNodeStats(id string) NodeStat {
	resp, _ := r.makeRequest(
		"GET",
		fmt.Sprintf("/api/internal/node/%s/stats", id),
		RequestParams{params: url.Values{"range": []string{"-10min"}}})

	var result NodeStat
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&result)

	return result
}
