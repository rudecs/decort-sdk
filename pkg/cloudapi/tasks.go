package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/tasks"
)

func (ca *CloudApi) Tasks() *tasks.Tasks {
	return tasks.New(ca.client)
}
