package bootstrap

import "google.golang.org/grpc"

type Options struct {
	registerGrpcCallback func(server *grpc.Server)
	withRedis            bool
	withMongo            bool
	serviceName          string
}

type Option func(*Options)

func NewOptions(opts ...Option) *Options {
	options := &Options{}

	for _, o := range opts {
		o(options)
	}

	return options
}

func WithRegisterGrpcCallback(fn func(server *grpc.Server)) Option {
	return func(o *Options) {
		o.registerGrpcCallback = fn
	}
}

func WithRedis(is bool) Option {
	return func(o *Options) {
		o.withRedis = is
	}
}

func WithServiceName(name string) Option {
	return func(o *Options) {
		o.serviceName = name
	}
}
