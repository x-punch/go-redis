package redis

// Config represents the mongo configuration.
type Config struct {
	Network  string `toml:"network"`
	Address  string `toml:"address"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

// NewConfig builds a new configuration with default values.
func NewConfig() Config {
	return Config{
		Network:  "tcp",
		Address:  ":6379",
		Password: "",
		DB:       0,
	}
}
