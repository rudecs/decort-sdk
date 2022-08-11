package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/sizes"
)

func (dc *decortClient) Sizes() *sizes.Sizes {
	return sizes.New(dc)
}
