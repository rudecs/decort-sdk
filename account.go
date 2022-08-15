package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/account"
)

func (dc *Client) Account() *account.Account {
	return account.New(dc)
}
