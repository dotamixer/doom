package auth

type Options struct {
	Secret            string
	HeaderAuthorize   string
	EffectiveDuration int
}

func defaultOptions() *Options {
	return &Options{
		Secret:            "secret-key",
		HeaderAuthorize:   "authorization",
		EffectiveDuration: 15,
	}
}

func Config(opts *Options) {
	defaultOpt := defaultOptions()
	if opts.Secret == "" {
		opts.Secret = defaultOpt.Secret
	}

	if opts.HeaderAuthorize == "" {
		opts.HeaderAuthorize = defaultOpt.HeaderAuthorize
	}

	defaultAuth.opts = opts
}
