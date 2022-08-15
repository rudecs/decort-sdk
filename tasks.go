package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/tasks"
)

func (dc *Client) Tasks() *tasks.Tasks {
	return tasks.New(dc)
}
