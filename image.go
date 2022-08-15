package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/image"
)

func (dc *Client) Image() *image.Image {
	return image.New(dc)
}
