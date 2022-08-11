package client

import (
	"github.com/rudecs/decort-sdk/image"
)

func (dc *decortClient) Image() *image.Image {
	return image.New(dc)
}
