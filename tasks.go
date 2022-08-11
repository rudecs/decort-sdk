package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/tasks"
)

func (dc *decortClient) Tasks() *tasks.Tasks {
	return tasks.New(dc)
}
