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
	log.Debugf("Is logged in: %t", r.isLoggedIn)

	_url := r.url + action

	log.Infof("Requested action: %s", action)

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	var netClient = http.Client{Transport: tr}

	body := p.body

	_url += "?" + p.params.Encode()
	log.Debugf("Request full URL: %s", _url)

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

// NewRubrik - Creates a new Rubrik API instance and login to it
func NewRubrik(url string, username string, password string) *Rubrik {

	log.Debug("Create new API Instance")
	session := &Rubrik{
		url:          url,
		username:     username,
		password:     password,
		sessionToken: "",
		isLoggedIn:   false,
	}
	session.Login()
	log.Info("Session-Token: %s", session.sessionToken)

	return session
}
