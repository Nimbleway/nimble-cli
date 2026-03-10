// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
)

func TestAgentList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent", "list",
			"--api-key", "string",
			"--limit", "1",
			"--managed-by", "nimble",
			"--offset", "0",
			"--privacy", "public",
			"--search", "search",
		)
	})
}

func TestAgentGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent", "get",
			"--api-key", "string",
			"--template-name", "template_name",
		)
	})
}

func TestAgentRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent", "run",
			"--api-key", "string",
			"--agent", "agent",
			"--params", "{foo: bar}",
			"--localization=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent: agent\n" +
			"params:\n" +
			"  foo: bar\n" +
			"localization: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "agent", "run",
			"--api-key", "string",
		)
	})
}

func TestAgentRunAsync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "agent", "run-async",
			"--api-key", "string",
			"--agent", "agent",
			"--params", "{foo: bar}",
			"--callback-url", "https://example.com/webhook/callback",
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
			"localization: true\n" +
			"storage_compress: true\n" +
			"storage_object_name: result-2024-01-15.json\n" +
			"storage_type: s3\n" +
			"storage_url: s3://bucket-name/path/to/object\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "agent", "run-async",
			"--api-key", "string",
		)
	})
}
