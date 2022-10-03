package cloudbroker

import "github.com/rudecs/decort-sdk/interfaces"

type CloudBroker struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *CloudBroker {
	return &CloudBroker{
		client: client,
	}
}
