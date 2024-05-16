package logger

// Option can be used to customize a new logger config.
type Option func(Config) Config

// WithLevel sets the log level.
func WithLevel(l Level) Option {
	return func(c Config) Config {
		c.Level = l

		return c
	}
}

// WithEncoding sets the logger encoding
func WithEncoding(e Encoding) Option {
	return func(c Config) Config {
		c.Encoding = e

		return c
	}
}
