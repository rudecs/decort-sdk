package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/tasks"
)

// Accessing the Tasks method group
func (ca *CloudAPI) Tasks() *tasks.Tasks {
	return tasks.New(ca.client)
}
