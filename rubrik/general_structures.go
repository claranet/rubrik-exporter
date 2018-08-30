package rubrik

type ResultList struct {
	HasMore bool `json:"hasMore"`
	Total   int  `json:"total"`
	// Data    []interface{} `json:"data"`
}

type Iops struct {
	ReadsPerSecond  []TimeStat `json:"readsPerSecond"`
	WritesPerSecond []TimeStat `json:"writesPerSecond"`
}

type IoThroughput struct {
	ReadBytePerSecond  []TimeStat `json:"readBytePerSecond"`
	WriteBytePerSecond []TimeStat `json:"writeBytePerSecond"`
}

type NetworkStat struct {
	BytesReceived    []TimeStat `json:"bytesReceived"`
	BytesTransmitted []TimeStat `json:"bytesTransmitted"`
}

type TimeStat struct {
	Time string `json:"time"`
	Stat int    `json:"stat"`
}
