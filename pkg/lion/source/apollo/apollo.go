package apollo

import (
	"github.com/dotamixer/doom/pkg/lion/source"
	"github.com/philchia/agollo"
	"time"
)

type apollo struct {
	namespaces []string
	opts       source.Options
	client     *agollo.Client
}

func (a *apollo) Read() (*source.ChangeSet, error) {
	for _, ns := range a.namespaces {
		content := a.client.GetNameSpaceContent(ns, "")
		cs := &source.ChangeSet{
			Data:      []byte(content),
			Format:    a.opts.Encoder.String(),
			Source:    a.String(),
			Timestamp: time.Now(),
		}
		cs.Checksum = cs.Sum()
		return cs, nil
	}
	return nil, nil
}

func (a *apollo) Watch() (source.Watcher, error) {

}

func (a *apollo) String() string {
	return "apollo"
}

// NewSource creates a new consul source
func NewSource(opts ...source.Option) source.Source {
	options := source.NewOptions(opts...)

	defaultConfig := &agollo.Conf{}

	appID, ok := options.Context.Value(appIDKey{}).(string)
	if ok {
		defaultConfig.AppID = appID
	}
	cluster, ok := options.Context.Value(clusterKey{}).(string)
	if ok {
		defaultConfig.Cluster = cluster
	}
	ip, ok := options.Context.Value(ipKey{}).(string)
	if ok {
		defaultConfig.IP = ip
	}
	ns, ok := options.Context.Value(namespaceKey{}).([]string)
	if ok {
		defaultConfig.NameSpaceNames = ns
	}

	client := agollo.NewClient(defaultConfig)

	err := client.Start()
	if err != nil {
		return nil
	}

	a := &apollo{
		opts:       options,
		client:     client,
		namespaces: defaultConfig.NameSpaceNames,
	}

	return a
}
