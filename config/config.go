package config

type Config struct {
	AppID         string
	AppSecret     string
	SSOURL        string
	DecortURL     string
	Retries       uint64
	SSLSkipVerify bool
}
