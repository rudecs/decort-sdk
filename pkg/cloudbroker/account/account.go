package account

import "github.com/rudecs/decort-sdk/interfaces"

type Account struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *Account {
	return &Account{
		client: client,
	}
}
