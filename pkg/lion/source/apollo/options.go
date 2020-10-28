package apollo

import (
	"context"
	"github.com/dotamixer/doom/pkg/lion/source"
)

type appIDKey struct{}
type clusterKey struct{}
type namespaceKey struct{}
type ipKey struct{}

// WithAddress sets the consul address
func WithAppID(a string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, appIDKey{}, a)
	}
}

// WithPrefix sets the key prefix to use
func WithCluster(p string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, clusterKey{}, p)
	}
}

// StripPrefix indicates whether to remove the prefix from config entries, or leave it in place.
func StripNamespaces(ns []string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}

		o.Context = context.WithValue(o.Context, namespaceKey{}, ns)
	}
}

// WithPrefix sets the key prefix to use
func WithIP(ip string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, ipKey{}, ip)
	}
}
