package client

import "github.com/rudecs/decort-sdk/lb"

func (dc *decortClient) LB() *lb.LB {
	return lb.New(dc)
}
