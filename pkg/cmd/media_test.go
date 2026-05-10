// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestMediaRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"media", "run",
			"--url", "https://example.com",
			"--country", "country",
			"--expected-mime-type", "string",
			"--locale", "locale",
			"--storage", "{url: url, object_name: object_name, type: s3}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mediaRun)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"media", "run",
			"--url", "https://example.com",
			"--country", "country",
			"--expected-mime-type", "string",
			"--locale", "locale",
			"--storage.url", "url",
			"--storage.object-name", "object_name",
			"--storage.type", "s3",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: https://example.com\n" +
			"country: country\n" +
			"expected_mime_types:\n" +
			"  - string\n" +
			"locale: locale\n" +
			"storage:\n" +
			"  url: url\n" +
			"  object_name: object_name\n" +
			"  type: s3\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"media", "run",
		)
	})
}

func TestMediaRunAsync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"media", "run-async",
			"--url", "https://example.com",
			"--callback-url", "https://example.com/webhook/callback",
			"--country", "country",
			"--expected-mime-type", "string",
			"--locale", "locale",
			"--storage", "{url: url, object_name: object_name, type: s3}",
			"--storage-compress=true",
			"--storage-object-name", "result-2024-01-15.json",
			"--storage-type", "s3",
			"--storage-url", "s3://bucket-name/path/to/object",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mediaRunAsync)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"media", "run-async",
			"--url", "https://example.com",
			"--callback-url", "https://example.com/webhook/callback",
			"--country", "country",
			"--expected-mime-type", "string",
			"--locale", "locale",
			"--storage.url", "url",
			"--storage.object-name", "object_name",
			"--storage.type", "s3",
			"--storage-compress=true",
			"--storage-object-name", "result-2024-01-15.json",
			"--storage-type", "s3",
			"--storage-url", "s3://bucket-name/path/to/object",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: https://example.com\n" +
			"callback_url: https://example.com/webhook/callback\n" +
			"country: country\n" +
			"expected_mime_types:\n" +
			"  - string\n" +
			"locale: locale\n" +
			"storage:\n" +
			"  url: url\n" +
			"  object_name: object_name\n" +
			"  type: s3\n" +
			"storage_compress: true\n" +
			"storage_object_name: result-2024-01-15.json\n" +
			"storage_type: s3\n" +
			"storage_url: s3://bucket-name/path/to/object\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"media", "run-async",
		)
	})
}
