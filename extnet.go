package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/extnet"
)

func (dc *decortClient) Extnet() *extnet.Extnet {
	return extnet.New(dc)
}
