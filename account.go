package client

import (
	"github.com/rudecs/decort-sdk/account"
)

func (dc *decortClient) Account() *account.Account {
	return account.New(dc)
}
