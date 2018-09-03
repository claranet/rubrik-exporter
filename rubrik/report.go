package rubrik

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type ReportList struct {
	*ResultList
	Data []Report `json:"data"`
}

type Report struct {
	Name           string `json:"name"`
	ReportType     string `json:"reportType"`
	UpdateTime     string `json:"updateTime"`
	ID             string `json:"id"`
	ReportTemplate string `json:"reportTemplate"`
	UpdateStatus   string `json:"updateStatus"`
}

type ReportData struct {
	ID          string             `json:"id"`
	Attribute   string             `json:"attribute"`
	ChartType   string             `json:"chartType"`
	Name        string             `json:"name"`
	Measure     string             `json:"measure"`
	DataColumns []ReportDataColumn `json:"dataColumns"`
}

type ReportDataColumn struct {
	Label      string            `json:"label"`
	DataPoints []ReportDataPoint `json:"dataPoints"`
}

type ReportDataPoint struct {
	Measure string  `json:"measure"`
	Value   float64 `json:"value"`
}

func (r Rubrik) GetReports(params map[string]string) []Report {
	_params := &RequestParams{params: url.Values{}}
	for k, v := range params {
		_params.params[k] = []string{v}
	}

	resp, _ := r.makeRequest("GET", "/api/internal/report", *_params)

	var l ReportList
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&l)

	return l.Data
}

// GetTaskDetails - Returned the reported TaskStatus in last 24h
// returns  map[succeeded:3 failed:1 canceled:2]
func (r Rubrik) GetTaskDetails() map[string]float64 {
	report := r.GetReports(map[string]string{
		"type": "Canned", "report_template": "ProtectionTasksDetails",
	})[0]

	_params := &RequestParams{params: url.Values{"chart_id": []string{"chart0"}}}
	_url := fmt.Sprintf("/api/internal/report/%s/chart", report.ID)

	resp, _ := r.makeRequest("GET", _url, *_params)

	var result = make(map[string]float64)

	var data []ReportData
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&data)

	for _, c := range data[0].DataColumns {
		_key := strings.ToLower(c.Label)
		result[_key] = c.DataPoints[0].Value
	}

	return result
}
