package redis

// Config represents the mongo configuration.
type Config struct {
	Network  string
	Address  string
	Password string
	DB       int
}

// NewConfig builds a new configuration with default values.
func NewConfig() *Config {
	return &Config{
		Network:  "tcp",
		Address:  "127.0.0.1:6379",
		Password: "",
		DB:       1,
	}
}
