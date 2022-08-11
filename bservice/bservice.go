package bservice

import "github.com/rudecs/decort-sdk/interfaces"

type BService struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *BService {
	return &BService{
		client,
	}
}
