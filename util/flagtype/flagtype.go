package flagtype

import (
	"time"

	"github.com/spf13/pflag"
)

type durationValue struct {
	v *time.Duration
}

func WrapDuration(d *time.Duration) pflag.Value {
	return durationValue{
		v: d,
	}
}

func (d durationValue) String() string {
	if *d.v == 0 {
		return ""
	}
	return d.v.String()
}

func (d durationValue) Set(s string) (err error) {
	*d.v, err = time.ParseDuration(s)
	return
}

func (d durationValue) Type() string {
	return "string"
}
