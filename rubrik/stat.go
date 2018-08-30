package rubrik

import (
	// "fmt"
	"encoding/json"
	"io/ioutil"
)

type VmStorageList struct {
	*ResultList
	Data []VmStorage `json:"data"`
}

type VmStorage struct {
	ID                     string
	Logicalbytes           float64 `json:"logicalBytes"`
	IngestedBytes          float64 `json:"ingestedBytes"`
	ExclusivePhysicalBytes float64 `json:"exclusivePhysicalBytes"`
	SharedPhysicalBytes    float64 `json:"sharedPhysicalBytes"`
	IndexStorageBytes      float64 `json:"indexStorageBytes"`
}

type SystemStorage struct {
	Total         int `json:"total"`
	Used          int `json:"used"`
	Available     int `json:"available"`
	Snapshot      int `json:"snapshot"`
	LiveMount     int `json:"liveMount"`
	Miscellaneous int `json:"miscellaneous"`
}

type DataLocationUsageList struct {
	*ResultList
	Data []DataLocationUsage `json:"data"`
}

type DataLocationUsage struct {
	LocationID                 string `json:"locationId"`
	DataDownloaded             int    `json:"dataDownloaded"`
	DataArchived               int    `json:"dataArchived"`
	NumVMsArchived             int    `json:"numVMsArchived"`
	NumFilesetsArchived        int    `json:"numFilesetsArchived"`
	NumLinuxFilesetsArchived   int    `json:"numLinuxFilesetsArchived"`
	NumWindowsFilesetsArchived int    `json:"numWindowsFilesetsArchived"`
	NumShareFilesetsArchived   int    `json:"numShareFilesetsArchived"`
	NumMssqlDbsArchived        int    `json:"numMssqlDbsArchived"`
	NumHypervVmsArchived       int    `json:"numHypervVmsArchived"`
	NumNutanixVmsArchived      int    `json:"numNutanixVmsArchived"`
	NumManagedVolumesArchived  int    `json:"numManagedVolumesArchived"`
}

// GetSystemStorage ...
func (r Rubrik) GetSystemStorage() SystemStorage {
	resp, _ := r.makeRequest("GET", "/api/internal/stats/system_storage", RequestParams{})

	data := json.NewDecoder(resp.Body)
	var d SystemStorage
	data.Decode(&d)

	return d
}

// GetPerVMStorage ...
func (r Rubrik) GetPerVMStorage() []VmStorage {
	resp, _ := r.makeRequest("GET", "/api/internal/stats/per_vm_storage", RequestParams{})

	data := json.NewDecoder(resp.Body)
	var d VmStorageList
	data.Decode(&d)

	return d.Data
}

// GetStreamCount ...
func (r Rubrik) GetStreamCount() int {
	resp, _ := r.makeRequest("GET", "/api/internal/stats/streams/count", RequestParams{})
	body, _ := ioutil.ReadAll(resp.Body)

	var data map[string]int
	json.Unmarshal(body, &data)

	return data["count"]
}

// GetDataLocationUsage ...
func (r Rubrik) GetDataLocationUsage() []DataLocationUsage {
	resp, _ := r.makeRequest("GET", "/api/internal/stats/data_location/usage", RequestParams{})

	var data DataLocationUsageList
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&data)

	return data.Data
}
