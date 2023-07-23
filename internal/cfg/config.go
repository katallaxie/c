package cfg

import "os"

// Config is the configuration for the application.
type Config struct {
	// Verbose enables verbose logging.
	Verbose bool
	// Template is the template to use.
	Template string
	// Force forces the creation of the file.
	Force bool
	// Prefix is the prefix to use.
	Prefix string
}

// Cwd is the current working directory.
func (c *Config) Cwd() (string, error) {
	p, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return p, nil
}

// New returns a new Config.
func New() *Config {
	return &Config{}
}

// Default returns the default configuration.
func Default() *Config {
	return &Config{
		Verbose: false,
	}
}
