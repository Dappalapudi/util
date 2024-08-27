package env

import "github.com/kelseyhightower/envconfig"

// Load loads into struct fields os envs vars using this library https://github.com/kelseyhightower/envconfig.
func Load(v interface{}) error {
	return envconfig.Process(string(Get()), v)
}
