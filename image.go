package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/image"
)

func (dc *decortClient) Image() *image.Image {
	return image.New(dc)
}
