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

type LocationList struct {
	*ResultList
	Data []Location `json:"data"`
}

type Location struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	LocationType string `json:"locationType"`
	IsActive     bool   `json:"isActive"`
	IPAddress    string `json:"ipAddress"`
	Bucket       string `json:"bucket"`
}

// GetArchiveLocations ...
func (r Rubrik) GetArchiveLocations() []Location {
	resp, _ := r.makeRequest("GET", "/api/internal/archive/location", RequestParams{})
	var data LocationList
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&data)
	return data.Data
}
