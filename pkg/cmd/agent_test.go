// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestAgentList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent", "list",
			"--limit", "1",
			"--managed-by", "nimble",
			"--offset", "0",
			"--privacy", "public",
			"--search", "search",
		)
	})
}

func TestAgentGenerate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent", "generate",
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
			"agent", "generate",
		)
	})
}

func TestAgentGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent", "get",
			"--template-name", "template_name",
		)
	})
}

func TestAgentGetGeneration(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent", "get-generation",
			"--generation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAgentPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent", "publish",
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
			"agent", "publish",
			"--agent-name", "agent_name",
		)
	})
}

func TestAgentRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent", "run",
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
			"agent", "run",
		)
	})
}

func TestAgentRunAsync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent", "run-async",
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
			"agent", "run-async",
		)
	})
}

func TestAgentRunBatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent", "run-batch",
			"--input", "{formats: [html, markdown], localization: true, params: {foo: bar}}",
			"--shared-inputs", "{agent: agent, formats: [html, markdown], localization: true, params: {foo: bar}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentRunBatch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent", "run-batch",
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
			"agent", "run-batch",
		)
	})
}
