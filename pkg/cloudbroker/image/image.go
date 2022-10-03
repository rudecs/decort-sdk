package image

import "github.com/rudecs/decort-sdk/interfaces"

type Image struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *Image {
	return &Image{
		client: client,
	}
}
