// package loader manages loading from multiple sources
package loader

import (
	"context"

	"github.com/dotamixer/doom/pkg/lion/reader"
	"github.com/dotamixer/doom/pkg/lion/source"
)

// Loader manages loading sources
type Loader interface {
	// Stop the loader
	Close() error
	// Load the sources
	Load(...source.Source) error
	// A Snapshot of loaded config
	Snapshot() (*Snapshot, error)
	// Force sync of sources
	Sync() error
	// Watch for changes
	Watch(...string) (Watcher, error)
	// Name of loader
	String() string
}

//// Watcher允许您观察源并返回合并的ChangeSet
// Watcher lets you watch sources and returns a merged ChangeSet
type Watcher interface {
	// First call to next may return the current Snapshot
	// If you are watching a path then only the data from
	// that path is returned.
	//第一次调用next可能会返回当前的快照。 如果您正在查看路径，则仅返回该路径中的数据。
	Next() (*Snapshot, error)
	// Stop watching for changes
	Stop() error
}

// Snapshot is a merged ChangeSet
type Snapshot struct {
	// The merged ChangeSet
	ChangeSet *source.ChangeSet
	//快照的确定性和可比较版本
	// Deterministic and comparable version of the snapshot
	Version string
}

type Options struct {
	Reader reader.Reader
	Source []source.Source

	// for alternative data
	Context context.Context
}

type Option func(o *Options)
