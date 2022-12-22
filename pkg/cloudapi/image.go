package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/image"
)

// Accessing the Image method group
func (ca *CloudAPI) Image() *image.Image {
	return image.New(ca.client)
}
