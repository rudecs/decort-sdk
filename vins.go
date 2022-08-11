package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/vins"
)

func (dc *decortClient) Vins() *vins.Vins {
	return vins.New(dc)
}
