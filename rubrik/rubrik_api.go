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
	//	"os"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/prometheus/log"
)

type RequestParams struct {
	body, header string
	params       url.Values
}

type Rubrik struct {
	url      string
	username string
	password string

	sessionToken string
	isLoggedIn   bool
}

func (r *Rubrik) makeRequest(reqType string, action string, p RequestParams) (*http.Response, error) {
	if !r.isLoggedIn {
		r.Login()
	}

	_url := r.url + action

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	var netClient = http.Client{Transport: tr}

	body := p.body

	_url += "?" + p.params.Encode()

	req, err := http.NewRequest(reqType, _url, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "text/JSON")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.sessionToken))

	resp, err := netClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp, nil
}

func NewRubrik(url string, username string, password string) *Rubrik {

	return &Rubrik{
		url:          url,
		username:     username,
		password:     password,
		sessionToken: "",
		isLoggedIn:   false,
	}

}
