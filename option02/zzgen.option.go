// Code generated by gozz:option github.com/go-zing/gozz. DO NOT EDIT.

package option02

import ()

// functional options type for Config
type ConfigOption func(*Config)

// apply functional options for Config
func (o *Config) applyOptions(opts ...ConfigOption) {
	for _, opt := range opts {
		opt(o)
	}
}

// connect host
func WithHost(v string) ConfigOption { return func(o *Config) { o.Host = v } }

// connect port
func WithPort(v string) ConfigOption { return func(o *Config) { o.Port = v } }

// database username
func WithUsername(v string) ConfigOption { return func(o *Config) { o.Username = v } }

// database password
func WithPassword(v string) ConfigOption { return func(o *Config) { o.Password = v } }

// functional options type for Config2
type Config2Option func(*Config2)

// apply functional options for Config2
func (o *Config2) applyOptions(opts ...Config2Option) {
	for _, opt := range opts {
		opt(o)
	}
}

// config url
func WithUrl(v string) Config2Option { return func(o *Config2) { o.Url = v } }