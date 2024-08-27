// Package gives access to global environment variables used by all services.
package env

import (
	"strings"
)

type Environment string

const (
	Test        Environment = "test"
	Development             = "dev"
	QA                      = "qa"
	Stage                   = "stg"
	Production              = "prod"
)

const DefaultEnvironment = Test

var current Environment = DefaultEnvironment

// Set default environment.
func Set(newEnv string) {
	current = Environment(strings.ToLower(newEnv))
}

// Get returns the standard BR environment derived from the BR_ENV environment variable.
// If no environment is set DefaultEnvironment returned.
func Get() Environment {
	if current == "" {
		return DefaultEnvironment
	}
	return current
}

// Is returns whether the current environments is one of any of the given environments.
func Is(env ...Environment) bool {
	currentEnv := Get()

	for _, e := range env {
		if e == currentEnv {
			return true
		}
	}
	return false
}

// IsDev returns true if it is dev environment.
func IsDev() bool {
	return Get() == Development
}

// IsTest returns true if it is test environment.
func IsTest() bool {
	return Get() == Test
}

// IsStg returns true if it is staging environment.
func IsStg() bool {
	return Get() == Stage
}

// IsProd returns true if it is prod environment.
func IsProd() bool {
	return Get() == Production
}

func IsQA() bool {
	return Get() == QA
}
