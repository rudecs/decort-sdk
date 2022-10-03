package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/account"
)

func (ca *CloudApi) Account() *account.Account {
	return account.New(ca.client)
}
