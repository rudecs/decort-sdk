package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/image"
)

func (ca *CloudApi) Image() *image.Image {
	return image.New(ca.client)
}
