package client

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/rudecs/decort-sdk/config"
)

// NewLegacyHttpClient creates legacy HTTP Client
func NewLegacyHttpClient(cfg config.LegacyConfig) *http.Client {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{
			//nolint:gosec
			InsecureSkipVerify: cfg.SSLSkipVerify,
		},
	}

	return &http.Client{
		Transport: &transportLegacy{
			base:      transCfg,
			username:  cfg.Username,
			password:  cfg.Password,
			retries:   cfg.Retries,
			token:     cfg.Token,
			decortURL: cfg.DecortURL,
		},

		Timeout: 5 * time.Minute,
	}
}
