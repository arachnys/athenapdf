package config

import (
	"os"
	"testing"

	"github.com/arachnys/athenapdf/pkg/proto"
)

func TestSet(t *testing.T) {
	testCases := []struct {
		name      string
		namespace string
		key       string
		value     string
		hasError  bool
	}{
		{name: "no namespace", hasError: true},
		{name: "no config key", namespace: "namespace", hasError: true},
		{"no config value", "namespace", "configKey", "", true},
		{"valid config", "namespace", "configKey", "configValue", false},
		{"another valid config", "namespace2", "configKey2", "configValue2", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Set(tc.namespace)(tc.key, tc.value)
			if err != nil && !tc.hasError {
				t.Fatalf("failed to register config, unexpected error: %+v", err)
			}
			if err == nil && tc.hasError {
				t.Fatalf("expected error setting config without correct arguments")
			}

			if !tc.hasError {
				got, err := Get(tc.namespace, nil)(tc.key)
				if err != nil {
					t.Errorf("failed to retrieve config value from key, unexpected error: %+v", err)
				}

				if got != tc.value {
					t.Errorf("got %+v; want %+v", got, tc.value)
				}
			}
		})
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		name      string
		init      func()
		namespace string
		key       string
		options   map[string]*proto.Option
		want      string
		hasError  bool
	}{
		{name: "no namespace", init: func() {}, hasError: true},
		{name: "no config key", init: func() {}, namespace: "namespace", hasError: true},
		{
			"config from options",
			func() {},
			"options",
			"key",
			map[string]*proto.Option{
				"options_key": &proto.Option{
					&proto.Option_StringValue{
						"value in options",
					},
				},
			},
			"value in options",
			false,
		},
		{
			"config from registry",
			func() {
				if err := Set("registry")("key", "value in registry"); err != nil {
					t.Fatalf("unexpected error: %+v", err)
				}
			},
			"registry",
			"key",
			nil,
			"value in registry",
			false,
		},
		{
			"config from environment",
			func() {
				os.Setenv("ATHENAPDF_CONFIG_ENVIRONMENT_KEY", "value in environment")
			},
			"environment",
			"key",
			nil,
			"value in environment",
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.init()

			got, err := Get(tc.namespace, tc.options)(tc.key)
			if err != nil && !tc.hasError {
				t.Fatalf("failed to get config, unexpected error: %+v", err)
			}
			if err == nil && tc.hasError {
				t.Fatalf("expected error getting config without correct arguments")
			}

			if !tc.hasError && got != tc.want {
				t.Errorf("got %+v; want %+v", got, tc.want)
			}
		})
	}
}

func TestMustGet(t *testing.T) {
	if err := Set("namespace")("key", "valid"); err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	testCases := []struct {
		name      string
		namespace string
		key       string
		want      string
	}{
		{"no namespace", "", "", ""},
		{"no config key", "namespace", "", ""},
		{"non-existent config", "namespace", "non_existent", ""},
		{"existing config", "namespace", "key", "valid"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MustGet(tc.namespace, nil)(tc.key); got != tc.want {
				t.Errorf("got %+v; want %+v", got, tc.want)
			}
		})
	}
}
