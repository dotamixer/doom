package memory

import (
	"github.com/dotamixer/doom/pkg/lion/loader"
	"github.com/dotamixer/doom/pkg/lion/reader"
	"github.com/dotamixer/doom/pkg/lion/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}
