// API Actor API for managing account
package account

import "github.com/rudecs/decort-sdk/interfaces"

// Structure for creating request to account
type Account struct {
	client interfaces.Caller
}

// Builder for account endpoints
func New(client interfaces.Caller) *Account {
	return &Account{
		client: client,
	}
}
