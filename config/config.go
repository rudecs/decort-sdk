package config

// Configuration for creating request to platform
type Config struct {
	// JWT platform token
	// Required: false
	// Example: "qwqwdfwv68979we0q9bfv7e9sbvd89798qrwv97ff"
	Token string

	// Application (client) identifier for authorization
	// in the cloud platform controller in oauth2 mode.
	// Required: true
	// Example: "ewqfrvea7s890avw804389qwguf234h0otfi3w4eiu"
	AppID string

	// Application (client) secret code for authorization
	// in the cloud platform controller in oauth2 mode.
	// Example: "frvet09rvesfis0c9erv9fsov0vsdfi09ovds0f"
	AppSecret string

	// Platform authentication service address
	// Required: true
	// Example: "https://sso.digitalenergy.online"
	SSOURL string

	// The address of the platform on which the actions are planned
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
