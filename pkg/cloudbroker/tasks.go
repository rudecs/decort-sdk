package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/tasks"
)

func (cb *CloudBroker) Tasks() *tasks.Tasks {
	return tasks.New(cb.client)
}
