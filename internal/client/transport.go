package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type transport struct {
	base         http.RoundTripper
	retries      uint64
	clientId     string
	clientSecret string
	token        string
	ssoUrl       string
	expiryTime   time.Time
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.token == "" || time.Now().After(t.expiryTime) {
		body := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&response_type=id_token", t.clientId, t.clientSecret)
		bodyReader := strings.NewReader(body)

		req, _ := http.NewRequest("POST", t.ssoUrl+"/v1/oauth/access_token", bodyReader)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, err := t.base.RoundTrip(req)
		if err != nil {
			return nil, fmt.Errorf("cannot get token: %v", err)
		}

		tokenBytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("cannot get token: %s", tokenBytes)
		}

		token := string(tokenBytes)

		t.token = token
		t.expiryTime = time.Now().AddDate(0, 0, 1)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "bearer "+t.token)
	req.Header.Set("Accept", "application/json")

	var resp *http.Response
	var err error
	for i := uint64(0); i < t.retries; i++ {
		resp, err = t.base.RoundTrip(req)
		if err == nil {
			if resp.StatusCode == 200 {
				return resp, nil
			}
			respBytes, _ := ioutil.ReadAll(resp.Body)
			err = fmt.Errorf("%s", respBytes)
			resp.Body.Close()
		}
		//logrus.Errorf("Could not execute request: %v. Retrying %d/%d", err, i+1, t.retries)
		time.Sleep(time.Second * 5)
	}
	return nil, fmt.Errorf("could not execute request: %v", err)
}
