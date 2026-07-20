// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestExtractTemplatesGenerationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract:templates:generations", "create",
			"--prompt", "prompt",
			"--url", "url",
			"--input-schema", "{foo: bar}",
			"--metadata", "{description: description, display_name: display_name, tags: [string]}",
			"--name", "name",
			"--output-schema", "{foo: bar}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(extractTemplatesGenerationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract:templates:generations", "create",
			"--prompt", "prompt",
			"--url", "url",
			"--input-schema", "{foo: bar}",
			"--metadata.description", "description",
			"--metadata.display-name", "display_name",
			"--metadata.tags", "[string]",
			"--name", "name",
			"--output-schema", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"prompt: prompt\n" +
			"url: url\n" +
			"input_schema:\n" +
			"  foo: bar\n" +
			"metadata:\n" +
			"  description: description\n" +
			"  display_name: display_name\n" +
			"  tags:\n" +
			"    - string\n" +
			"name: name\n" +
			"output_schema:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"extract:templates:generations", "create",
		)
	})
}

func TestExtractTemplatesGenerationsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract:templates:generations", "get",
			"--generation-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
