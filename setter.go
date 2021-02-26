package email

import "github.com/go-redis/redis/v7"

// Setter ...
type Setter func(config *config) error

// WithHost ...
func WithHost(host string) Setter {
	return func(config *config) error {
		config.host = host
		return nil
	}
}

// WithPort ...
func WithPort(port string) Setter {
	return func(config *config) error {
		config.port = port
		return nil
	}
}

// WithAddress ...
func WithAddress(address string) Setter {
	return func(config *config) error {
		config.address = address
		return nil
	}
}

// WithUsername ...
func WithUsername(username string) Setter {
	return func(config *config) error {
		config.username = username
		return nil
	}
}

// WithPassword ...
func WithPassword(password string) Setter {
	return func(config *config) error {
		config.password = password
		return nil
	}
}

// WithCache ...
func WithCache(client *redis.Client) Setter {
	return func(config *config) error {
		config.client = client
		return nil
	}
}

// WithIgnoreError ...
func WithIgnoreError(ignoreError bool) Setter {
	return func(config *config) error {
		config.ignoreError = ignoreError
		return nil
	}
}
