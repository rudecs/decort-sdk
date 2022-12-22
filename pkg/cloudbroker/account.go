package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudbroker/account"
)

// Accessing the Account method group
func (cb *CloudBroker) Account() *account.Account {
	return account.New(cb.client)
}
