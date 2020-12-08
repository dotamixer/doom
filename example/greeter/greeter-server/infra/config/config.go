package config

import "github.com/dotamixer/doom/pkg/lion"

type Config struct {
}

func NewConfig() *Config {
	c := &Config{}
	_ = lion.Get("logic").Scan(c)

	return c
}
