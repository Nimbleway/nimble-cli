// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestSerpRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"serp", "run",
			"--search-engine", "google_search",
			"--country", "US",
			"--device", "desktop",
			"--domain", "com",
			"--locale", "en",
			"--location", "New York, New York, United States",
			"--num-results", "10",
			"--page", "1",
			"--parse=true",
			"--query", "nimble web data",
			"--render=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"search_engine: google_search\n" +
			"country: US\n" +
			"device: desktop\n" +
			"domain: com\n" +
			"locale: en\n" +
			"location: New York, New York, United States\n" +
			"num_results: 10\n" +
			"page: 1\n" +
			"parse: true\n" +
			"query: nimble web data\n" +
			"render: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"serp", "run",
		)
	})
}

func TestSerpRunAsync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"serp", "run-async",
			"--search-engine", "google_search",
			"--callback-url", "https://example.com/webhook/callback",
			"--country", "US",
			"--device", "desktop",
			"--domain", "com",
			"--locale", "en",
			"--location", "New York, New York, United States",
			"--num-results", "10",
			"--page", "1",
			"--parse=true",
			"--query", "nimble web data",
			"--render=false",
			"--storage-compress=true",
			"--storage-object-name", "result-2024-01-15.json",
			"--storage-type", "s3",
			"--storage-url", "s3://bucket-name/path/to/object",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"search_engine: google_search\n" +
			"callback_url: https://example.com/webhook/callback\n" +
			"country: US\n" +
			"device: desktop\n" +
			"domain: com\n" +
			"locale: en\n" +
			"location: New York, New York, United States\n" +
			"num_results: 10\n" +
			"page: 1\n" +
			"parse: true\n" +
			"query: nimble web data\n" +
			"render: false\n" +
			"storage_compress: true\n" +
			"storage_object_name: result-2024-01-15.json\n" +
			"storage_type: s3\n" +
			"storage_url: s3://bucket-name/path/to/object\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"serp", "run-async",
		)
	})
}

func TestSerpRunBatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"serp", "run-batch",
			"--input", "{callback_url: https://example.com/webhook/callback, country: US, device: desktop, domain: com, locale: en, location: 'New York, New York, United States', num_results: 10, page: 1, parse: true, query: nimble web data, render: false, search_engine: google_search, storage_compress: true, storage_object_name: result-2024-01-15.json, storage_type: s3, storage_url: s3://bucket-name/path/to/object}",
			"--shared-inputs", "{callback_url: https://example.com/webhook/callback, country: US, device: desktop, domain: com, locale: en, location: 'New York, New York, United States', num_results: 10, page: 1, parse: true, query: nimble web data, render: false, search_engine: google_search, storage_compress: true, storage_object_name: result-2024-01-15.json, storage_type: s3, storage_url: s3://bucket-name/path/to/object}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(serpRunBatch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"serp", "run-batch",
			"--input.callback-url", "https://example.com/webhook/callback",
			"--input.country", "US",
			"--input.device", "desktop",
			"--input.domain", "com",
			"--input.locale", "en",
			"--input.location", "New York, New York, United States",
			"--input.num-results", "10",
			"--input.page", "1",
			"--input.parse=true",
			"--input.query", "nimble web data",
			"--input.render=false",
			"--input.search-engine", "google_search",
			"--input.storage-compress=true",
			"--input.storage-object-name", "result-2024-01-15.json",
			"--input.storage-type", "s3",
			"--input.storage-url", "s3://bucket-name/path/to/object",
			"--shared-inputs.callback-url", "https://example.com/webhook/callback",
			"--shared-inputs.country", "US",
			"--shared-inputs.device", "desktop",
			"--shared-inputs.domain", "com",
			"--shared-inputs.locale", "en",
			"--shared-inputs.location", "New York, New York, United States",
			"--shared-inputs.num-results", "10",
			"--shared-inputs.page", "1",
			"--shared-inputs.parse=true",
			"--shared-inputs.query", "nimble web data",
			"--shared-inputs.render=false",
			"--shared-inputs.search-engine", "google_search",
			"--shared-inputs.storage-compress=true",
			"--shared-inputs.storage-object-name", "result-2024-01-15.json",
			"--shared-inputs.storage-type", "s3",
			"--shared-inputs.storage-url", "s3://bucket-name/path/to/object",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"inputs:\n" +
			"  - callback_url: https://example.com/webhook/callback\n" +
			"    country: US\n" +
			"    device: desktop\n" +
			"    domain: com\n" +
			"    locale: en\n" +
			"    location: New York, New York, United States\n" +
			"    num_results: 10\n" +
			"    page: 1\n" +
			"    parse: true\n" +
			"    query: nimble web data\n" +
			"    render: false\n" +
			"    search_engine: google_search\n" +
			"    storage_compress: true\n" +
			"    storage_object_name: result-2024-01-15.json\n" +
			"    storage_type: s3\n" +
			"    storage_url: s3://bucket-name/path/to/object\n" +
			"shared_inputs:\n" +
			"  callback_url: https://example.com/webhook/callback\n" +
			"  country: US\n" +
			"  device: desktop\n" +
			"  domain: com\n" +
			"  locale: en\n" +
			"  location: New York, New York, United States\n" +
			"  num_results: 10\n" +
			"  page: 1\n" +
			"  parse: true\n" +
			"  query: nimble web data\n" +
			"  render: false\n" +
			"  search_engine: google_search\n" +
			"  storage_compress: true\n" +
			"  storage_object_name: result-2024-01-15.json\n" +
			"  storage_type: s3\n" +
			"  storage_url: s3://bucket-name/path/to/object\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"serp", "run-batch",
		)
	})
}
