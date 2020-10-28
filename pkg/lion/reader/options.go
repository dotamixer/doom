package reader

import (
	"github.com/dotamixer/doom/pkg/lion/encoder"
	"github.com/dotamixer/doom/pkg/lion/encoder/hcl"
	"github.com/dotamixer/doom/pkg/lion/encoder/json"
	"github.com/dotamixer/doom/pkg/lion/encoder/toml"
	"github.com/dotamixer/doom/pkg/lion/encoder/xml"
	"github.com/dotamixer/doom/pkg/lion/encoder/yaml"
)

type Options struct {
	Encoding map[string]encoder.Encoder
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Encoding: map[string]encoder.Encoder{
			"json": json.NewEncoder(),
			"yaml": yaml.NewEncoder(),
			"toml": toml.NewEncoder(),
			"xml":  xml.NewEncoder(),
			"hcl":  hcl.NewEncoder(),
			"yml":  yaml.NewEncoder(),
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		if o.Encoding == nil {
			o.Encoding = make(map[string]encoder.Encoder)
		}
		o.Encoding[e.String()] = e
	}
}
