package client

import (
	"crypto/tls"
	"net/http"
	"time"
)

func newHttpClient(config Config) *http.Client {

	transCfg := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: config.SSLSkipVerify}}

	return &http.Client{
		Transport: &transport{
			base:         transCfg,
			retries:      config.Retries,
			clientId:     config.AppId,
			clientSecret: config.AppSecret,
			ssoUrl:       config.SSOUrl,
			//TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},

		Timeout: 5 * time.Minute,
	}
}
