package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewEnvConfig(t *testing.T) {
	t.Run("Datadog Agent Address", func(t *testing.T) {
		c := NewEnvConfig()
		assert.Equal(t, "", c.DatadogAgentAddress)

		expectValue := "123qweasdzxc:12312"
		err := os.Setenv("DATADOG_AGENT_ADDRESS", expectValue)
		assert.Nil(t, err)

		c = NewEnvConfig()
		assert.Equal(t, expectValue, c.DatadogAgentAddress)
	})

	t.Run("Datadog APM Service Name", func(t *testing.T) {
		c := NewEnvConfig()
		assert.Equal(t, "weaver", c.DatadogAPMServiceName)

		expectValue := "12o38y27345987234695thekgjr"
		err := os.Setenv("DATADOG_APM_SERVICE_NAME", expectValue)
		assert.Nil(t, err)

		c = NewEnvConfig()
		assert.Equal(t, expectValue, c.DatadogAPMServiceName)
	})
}
