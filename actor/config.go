package actor

type Config struct {
}

func Configure(opts ...ConfigOption) *Config {
	config := &Config{}
	for _, opt := range opts {
		opt(config)
	}
	return config
}
