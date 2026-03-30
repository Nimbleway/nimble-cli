// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestAgentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "list",
			"--limit", "1",
			"--managed-by", "nimble",
			"--offset", "0",
			"--privacy", "public",
			"--search", "search",
		)
	})
}

func TestAgentsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "get",
			"--template-name", "template_name",
		)
	})
}

func TestAgentsPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "publish",
			"--agent-name", "agent_name",
			"--version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("version_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "publish",
			"--agent-name", "agent_name",
		)
	})
}

func TestAgentsRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "run",
			"--agent", "agent",
			"--params", "{foo: bar}",
			"--format", "html",
			"--format", "markdown",
			"--localization=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent: agent\n" +
			"params:\n" +
			"  foo: bar\n" +
			"formats:\n" +
			"  - html\n" +
			"  - markdown\n" +
			"localization: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "run",
		)
	})
}

func TestAgentsRunAsync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "run-async",
			"--agent", "agent",
			"--params", "{foo: bar}",
			"--callback-url", "https://example.com/webhook/callback",
			"--format", "html",
			"--format", "markdown",
			"--localization=true",
			"--storage-compress=true",
			"--storage-object-name", "result-2024-01-15.json",
			"--storage-type", "s3",
			"--storage-url", "s3://bucket-name/path/to/object",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent: agent\n" +
			"params:\n" +
			"  foo: bar\n" +
			"callback_url: https://example.com/webhook/callback\n" +
			"formats:\n" +
			"  - html\n" +
			"  - markdown\n" +
			"localization: true\n" +
			"storage_compress: true\n" +
			"storage_object_name: result-2024-01-15.json\n" +
			"storage_type: s3\n" +
			"storage_url: s3://bucket-name/path/to/object\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "run-async",
		)
	})
}

func TestAgentsRunBatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "run-batch",
			"--input", "{formats: [html, markdown], localization: true, params: {foo: bar}}",
			"--shared-inputs", "{agent: agent, formats: [html, markdown], localization: true, params: {foo: bar}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsRunBatch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "run-batch",
			"--input.formats", "[html, markdown]",
			"--input.localization=true",
			"--input.params", "{foo: bar}",
			"--shared-inputs.agent", "agent",
			"--shared-inputs.formats", "[html, markdown]",
			"--shared-inputs.localization=true",
			"--shared-inputs.params", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"inputs:\n" +
			"  - formats:\n" +
			"      - html\n" +
			"      - markdown\n" +
			"    localization: true\n" +
			"    params:\n" +
			"      foo: bar\n" +
			"shared_inputs:\n" +
			"  agent: agent\n" +
			"  formats:\n" +
			"    - html\n" +
			"    - markdown\n" +
			"  localization: true\n" +
			"  params:\n" +
			"    foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "run-batch",
		)
	})
}
