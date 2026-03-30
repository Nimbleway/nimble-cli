// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
)

func TestAgentsGenerationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:generations", "create",
			"--agent-name", "agent_name",
			"--prompt", "prompt",
			"--url", "url",
			"--input-schema", "{}",
			"--metadata", "{}",
			"--output-schema", "{}",
			"--from-agent", "from_agent",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent_name: agent_name\n" +
			"prompt: prompt\n" +
			"url: url\n" +
			"input_schema: {}\n" +
			"metadata: {}\n" +
			"output_schema: {}\n" +
			"from_agent: from_agent\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents:generations", "create",
		)
	})
}

func TestAgentsGenerationsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:generations", "get",
			"--generation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
