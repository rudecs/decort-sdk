package config

// Legacy client configuration
type LegacyConfig struct {
	// ServiceAccount username
	// Required: true
	// Example : "osh_mikoev"
	Username string

	// ServiceAccount password
	// Required: true
	// Example: "[1o>hYkjnJr)HI78q7t&#%8Lm"
	Password string

	// Platform token
	// Required: false
	// Example: "158e76424b0d4810b6086hgbhj928fc4a6bc06e"
	Token string

	// Address of the platform on which the actions are planned
	// Required: true
	// Example: "https://mr4.digitalenergy.online"
	DecortURL string

	// Amount platform request attempts
	// Default value: 5
	// Required: false
	Retries uint64

	// Skip verify, true by default
	// Required: false
	SSLSkipVerify bool
}
