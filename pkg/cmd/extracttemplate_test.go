// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestExtractTemplatesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract:templates", "update",
			"--extract-template-name", "extract_template_name",
			"--body", "{op: add, path: path, from: from, value: {}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(extractTemplatesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract:templates", "update",
			"--extract-template-name", "extract_template_name",
			"--body.op", "add",
			"--body.path", "path",
			"--body.from", "from",
			"--body.value", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- op: add\n" +
			"  path: path\n" +
			"  from: from\n" +
			"  value: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"extract:templates", "update",
			"--extract-template-name", "extract_template_name",
		)
	})
}

func TestExtractTemplatesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract:templates", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestExtractTemplatesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract:templates", "delete",
			"--extract-template-name", "extract_template_name",
		)
	})
}

func TestExtractTemplatesAsync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract:templates", "async",
			"--params", "{foo: bar}",
			"--template", "template",
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
			"params:\n" +
			"  foo: bar\n" +
			"template: template\n" +
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
			"extract:templates", "async",
		)
	})
}

func TestExtractTemplatesBatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract:templates", "batch",
			"--input", "{formats: [html, markdown], localization: true, params: {foo: bar}}",
			"--shared-inputs", "{template: template, formats: [html, markdown], localization: true, params: {foo: bar}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(extractTemplatesBatch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract:templates", "batch",
			"--input.formats", "[html, markdown]",
			"--input.localization=true",
			"--input.params", "{foo: bar}",
			"--shared-inputs.template", "template",
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
			"  template: template\n" +
			"  formats:\n" +
			"    - html\n" +
			"    - markdown\n" +
			"  localization: true\n" +
			"  params:\n" +
			"    foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"extract:templates", "batch",
		)
	})
}

func TestExtractTemplatesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract:templates", "get",
			"--extract-template-name", "extract_template_name",
		)
	})
}

func TestExtractTemplatesRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract:templates", "run",
			"--params", "{foo: bar}",
			"--template", "template",
			"--format", "html",
			"--format", "markdown",
			"--localization=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"params:\n" +
			"  foo: bar\n" +
			"template: template\n" +
			"formats:\n" +
			"  - html\n" +
			"  - markdown\n" +
			"localization: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"extract:templates", "run",
		)
	})
}
