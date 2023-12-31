// Package container provides the dependency injection.
package container

import (
	"github.com/Drynok/fhome_ha/env"

	"go.uber.org/dig"
)

// New create container with dependency injection and default dependencies.
func New() (*dig.Container, error) {
	cnt := dig.New()

	for _, err := range []error{
		cnt.Provide(env.New),
	} {
		if err != nil {
			return nil, err
		}
	}

	return cnt, nil
}
