// Package config is an interface for dynamic configuration.
package lion

import (
	"context"

	"github.com/dotamixer/doom/pkg/lion/loader"
	"github.com/dotamixer/doom/pkg/lion/reader"
	"github.com/dotamixer/doom/pkg/lion/source"
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	// provide the reader.Values interface
	reader.Values
	// Stop the config loader/watcher
	Close() error
	// Load config sources
	Load(source ...source.Source) error
	// Force a source changeset sync
	Sync() error
	// Watch a value for changes
	Watch(path ...string) (Watcher, error)
}

// Watcher is the config watcher
type Watcher interface {
	Next() (reader.Value, error)
	Stop() error
}

type Options struct {
	Loader loader.Loader
	Reader reader.Reader
	Source []source.Source

	// for alternative data
	Context context.Context
}

type Option func(o *Options)

var (
	// Default Config Manager
	DefaultConfig = NewConfig()
)

// NewConfig returns new config
func NewConfig(opts ...Option) Config {
	return newConfig(opts...)
}

// Return config as raw json
func Bytes() []byte {
	return DefaultConfig.Bytes()
}

// Return config as a map
func Map() map[string]interface{} {
	return DefaultConfig.Map()
}

// Scan values to a go type
func Scan(v interface{}) error {
	return DefaultConfig.Scan(v)
}

// Force a source changeset sync
func Sync() error {
	return DefaultConfig.Sync()
}

// Get a value from the config
func Get(path ...string) reader.Value {
	return DefaultConfig.Get(path...)
}

// Load config sources
func Load(source ...source.Source) error {
	return DefaultConfig.Load(source...)
}

// Watch a value for changes
func Watch(path ...string) (Watcher, error) {
	return DefaultConfig.Watch(path...)
}
