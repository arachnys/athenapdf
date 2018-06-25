package config

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
	"sync"

	"github.com/arachnys/athenapdf/pkg/proto"
)

var (
	config   = make(map[string]map[string]string)
	configMu sync.RWMutex
)

func toOptionName(namespace, configKey string) string {
	return strings.ToLower(fmt.Sprintf("%s_%s", namespace, configKey))
}

func toEnvVar(namespace, configKey string) string {
	return strings.ToUpper(fmt.Sprintf("ATHENAPDF_CONFIG_%s_%s", namespace, configKey))
}

func Set(namespace string) func(string, string) error {
	return func(configKey, configValue string) error {
		configMu.Lock()
		defer configMu.Unlock()

		if namespace == "" {
			return ConfigError{err: errors.New("namespace is nil")}
		}
		if configKey == "" {
			return ConfigError{errors.New("config key is nil"), namespace}
		}
		if configValue == "" {
			return ConfigError{errors.New("config value is nil"), namespace}
		}

		if _, ok := config[namespace]; !ok {
			config[namespace] = make(map[string]string)
		}

		config[namespace][configKey] = configValue
		return nil
	}
}

func Get(namespace string, options map[string]*proto.Option) func(string) (string, error) {
	return func(configKey string) (string, error) {
		if namespace == "" {
			return "", ConfigError{err: errors.New("namespace is nil")}
		}
		if configKey == "" {
			return "", ConfigError{errors.New("config key is nil"), namespace}
		}

		// Attempt to get config from options (context)
		if v, ok := options[toOptionName(namespace, configKey)]; ok && v.GetStringValue() != "" {
			return v.GetStringValue(), nil
		}
		// Attempt to get config from global config registry
		if v, ok := config[namespace][configKey]; ok && v != "" {
			return v, nil
		}
		// Attempt to get config from environment (default)
		if v, ok := os.LookupEnv(toEnvVar(namespace, configKey)); ok && v != "" {
			return v, nil
		}

		return "", ConfigError{errors.Errorf("config `%s` does not exist", configKey), namespace}
	}
}

func MustGet(namespace string, options map[string]*proto.Option) func(string) string {
	return func(configKey string) string {
		v, _ := Get(namespace, options)(configKey)
		return v
	}
}