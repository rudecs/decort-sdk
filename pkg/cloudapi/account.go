package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/account"
)

// Accessing the Account method group
func (ca *CloudAPI) Account() *account.Account {
	return account.New(ca.client)
}
