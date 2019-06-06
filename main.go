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
	"net/http"

	"github.com/claranet/rubrik-exporter/rubrik"
	"github.com/namsral/flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/log"
)

var rubrikAPI *rubrik.Rubrik
var vmIDNameMap map[string]string

var (
	namespace      = "rubrik"
	rubrikURL      = flag.String("rubrik.url", "", "Rubrik URL to connect https://rubrik.local.host")
	rubrikUser     = flag.String("rubrik.username", "", "Rubrik API User")
	rubrikPassword = flag.String("rubrik.password", "", "Rubrik API User Password")
	listenAddress  = flag.String("listen-address", ":9477", "The address to lisiten on for HTTP requests.")
)

func main() {
	flag.Parse()

	log.Debug("Create Rubrik Exporter instance")
	rubrikAPI = rubrik.NewRubrik(*rubrikURL, *rubrikUser, *rubrikPassword)

	prometheus.MustRegister(NewRubrikStatsExport())
	prometheus.MustRegister(NewVMStatsExport())
	prometheus.MustRegister(NewArchiveLocation())
	prometheus.MustRegister(NewManagedVolume())

	http.Handle("/metrics", prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html><head><title>Rubrik Exporter</title></head><body><h1>Rubrik Exporter</h1><p><a href="/metrics">Metrics</a></p></body></html>`))
	})

	log.Debugf("Starting Server: %s", *listenAddress)
	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
