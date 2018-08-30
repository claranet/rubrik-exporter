package rubrik

import (
	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/prometheus/log"
)

type Session struct {
	Id             string `json:"id"`
	OrganizationId string `json:"organizationId"`
	Token          string `json:"token"`
	UserId         string `json:"userId"`
}

func (r *Rubrik) Login() error {
	_url := r.url + "/api/v1/session"

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	var netClient = http.Client{Transport: tr}
	req, err := http.NewRequest("POST", _url, nil)

	if err != nil {
		log.Fatal(err)
	}
	//	req.Header.Set("Content-Type", "text/JSON")
	req.SetBasicAuth(r.username, r.password)

	resp, err := netClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}

	data := json.NewDecoder(resp.Body)
	var s Session
	err = data.Decode(&s)

	r.sessionToken = s.Token
	r.isLoggedIn = true

	return nil
}

func (r *Rubrik) Logout() {
	r.makeRequest("DELETE", "/api/v1/session", RequestParams{})
}
