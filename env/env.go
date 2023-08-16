// Package env provides the ability to validate and parse environment variables.
package env

import (
	"fmt"
	"path/filepath"
	"runtime"

	env "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

// Environment environment variables configuration.
type Environment struct {
	ServerMode string `env:SERVER_MODE,default=`
	ServerPort string `env:"SERVER_PORT,default=5050"`
}

// New initialize the environment
func New() (*Environment, error) {
	var (
		_, b, _, _ = runtime.Caller(0)
		base       = filepath.Dir(b)
		_          = godotenv.Load(fmt.Sprintf("%s/../../.env", base))
		config     = new(Environment)
	)

	_, err := env.UnmarshalFromEnviron(config)
	return config, err
}
