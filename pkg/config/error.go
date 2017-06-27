package config

import (
	"fmt"
)

type ConfigError struct {
	err       error
	namespace string
}

func (e ConfigError) Error() string {
	if e.namespace != "" {
		return fmt.Sprintf("%s: %s", e.namespace, e.err.Error())
	}
	return e.err.Error()
}

func (e ConfigError) Cause() error {
	return e.err
}
