package opts

import "github.com/rudecs/decort-sdk/typed"

type Opts struct {
	IsAdmin    bool
	AdminValue string
}

func New(options []DecortOpts) *Opts {
	if len(options) == 0 {
		return nil
	}

	option := &Opts{}

	for _, v := range options {
		if v.Type == typed.TypeAdmin {
			option.IsAdmin = true
			option.AdminValue = v.Value
		}

	}

	return option
}
