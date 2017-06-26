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
	config   map[string]map[string]string
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
			return errors.New("namespace is nil")
		}
		if configKey == "" {
			return errors.New("config key is nil")
		}
		if configValue == "" {
			return errors.New("config value is nil")
		}

		config[namespace][configKey] = configValue
		return nil
	}
}

func Get(namespace string, options map[string]*proto.Option) func(string) (string, error) {
	return func(configKey string) (string, error) {
		if configKey == "" {
			return "", errors.Errorf("config key is nil")
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
		return "", errors.Errorf("config `%s` for namespace `%s` does not exist", configKey, namespace)
	}
}

func MustGet(namespace string, options map[string]*proto.Option) func(string) string {
	return func(configKey string) string {
		v, _ := Get(namespace, options)(configKey)
		return v
	}
}
