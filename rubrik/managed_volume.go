package rubrik

import (
	"encoding/json"
)

type ManagedVolumeList struct {
	*ResultList
	Data []ManagedVolume `json:"data"`
}

type ManagedVolume struct {
	ID                      string  `json:"id"`
	State                   string  `json:"state"`
	NumChannels             float64 `json:"numChannels"`
	ConfiguredSLADomainName string  `json:"configuredSlaDomainName"`
	EffectiveSLADomainID    string  `json:"effectiveSlaDomainId"`
	PrimaryClusterID        string  `json:"primaryClusterId"`
	UsedSize                float64 `json:"usedSize"`
	SLAAssignment           string  `json:"slaAssignment"`
	// MainExport              string  `json:"mainExport"`
	ConfiguredSLADomainID  string  `json:"configuredSlaDomainId"`
	IsWritable             string  `json:"isWritable"`
	VolumeSize             float64 `json:"volumeSize"`
	EffectiveSLADomainName string  `json:"effectiveSlaDomainName"`
	SnapshotCount          float64 `json:"snapshotCount"`
	PendingSnapshotCount   float64 `json:"pendingSnapshotCount"`
	IsRelic                string  `json:"isRelic"`
	Name                   string  `json:"name"`
	// HostPatterns            string  `json:"hostPatterns"`
	// Links string `json:"links"`

}

/* GetManagedVolumes
 *
 */
func (r Rubrik) GetManagedVolumes() []ManagedVolume {
	resp, _ := r.makeRequest("GET", "/api/internal/managed_volume", RequestParams{})

	var l ManagedVolumeList
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&l)

	return l.Data
}
