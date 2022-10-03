package client

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/rudecs/decort-sdk/config"
)

func NewHttpClient(cfg config.Config) *http.Client {

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{
			//nolint:gosec
			InsecureSkipVerify: cfg.SSLSkipVerify,
		},
	}

	return &http.Client{
		Transport: &transport{
			base:         transCfg,
			retries:      cfg.Retries,
			clientID:     cfg.AppID,
			clientSecret: cfg.AppSecret,
			SSOURL:       cfg.SSOURL,
			//TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},

		Timeout: 5 * time.Minute,
	}
}
