package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type transportLegacy struct {
	base      http.RoundTripper
	username  string
	password  string
	retries   uint64
	token     string
	decortURL string
}

func (t *transportLegacy) RoundTrip(request *http.Request) (*http.Response, error) {
	if t.token == "" {
		body := fmt.Sprintf("username=%s&password=%s", t.username, t.password)
		bodyReader := strings.NewReader(body)

		req, _ := http.NewRequestWithContext(request.Context(), "POST", t.decortURL+"/restmachine/cloudapi/user/authenticate", bodyReader)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, err := t.base.RoundTrip(req)
		if err != nil {
			return nil, fmt.Errorf("unable to get token: %w", err)
		}

		tokenBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("unable to get token: %s", tokenBytes)
		}

		token := string(tokenBytes)
		t.token = token
	}

	tokenValue := fmt.Sprintf("authkey=%s", t.token)
	tokenReader := strings.NewReader(tokenValue)

	req, _ := http.NewRequestWithContext(request.Context(), request.Method, request.URL.String(), tokenReader)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	var resp *http.Response
	var err error
	for i := uint64(0); i < t.retries; i++ {
		resp, err = t.base.RoundTrip(req)
		if err == nil {
			if resp.StatusCode == 200 {
				return resp, nil
			}
			respBytes, _ := io.ReadAll(resp.Body)
			err = fmt.Errorf("%s", respBytes)
			resp.Body.Close()
		}

		time.Sleep(time.Second * 5)
	}
	return nil, fmt.Errorf("could not execute request: %w", err)
}
