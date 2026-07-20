// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestAgentsRunsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:runs", "create",
			"--agent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--input", "input",
			"--effort", "low",
			"--enable-events=true",
			"--input-data", "[{foo: bar}]",
			"--output-schema", "{foo: bar}",
			"--previous-interaction-id", "previous_interaction_id",
			"--sources", "{allow: [{domains: [string], title: title, order: 0}], avoid: avoid, block: [{domains: [string], title: title, order: 0}], prioritize: prioritize}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsRunsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:runs", "create",
			"--agent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--input", "input",
			"--effort", "low",
			"--enable-events=true",
			"--input-data", "[{foo: bar}]",
			"--output-schema", "{foo: bar}",
			"--previous-interaction-id", "previous_interaction_id",
			"--sources.allow", "[{domains: [string], title: title, order: 0}]",
			"--sources.avoid", "avoid",
			"--sources.block", "[{domains: [string], title: title, order: 0}]",
			"--sources.prioritize", "prioritize",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"input: input\n" +
			"effort: low\n" +
			"enable_events: true\n" +
			"input_data:\n" +
			"  - foo: bar\n" +
			"output_schema:\n" +
			"  foo: bar\n" +
			"previous_interaction_id: previous_interaction_id\n" +
			"sources:\n" +
			"  allow:\n" +
			"    - domains:\n" +
			"        - string\n" +
			"      title: title\n" +
			"      order: 0\n" +
			"  avoid: avoid\n" +
			"  block:\n" +
			"    - domains:\n" +
			"        - string\n" +
			"      title: title\n" +
			"      order: 0\n" +
			"  prioritize: prioritize\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents:runs", "create",
			"--agent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAgentsRunsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:runs", "list",
			"--agent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestAgentsRunsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:runs", "get",
			"--agent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "run_id",
		)
	})
}

func TestAgentsRunsResult(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:runs", "result",
			"--agent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "run_id",
		)
	})
}

func TestAgentsRunsStreamEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:runs", "stream-events",
			"--agent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--run-id", "run_id",
		)
	})
}
