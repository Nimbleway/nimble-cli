// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
)

func TestFastSerpRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fast-serp", "run",
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
			"--show-hidden-results=false",
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
			"render: false\n" +
			"show_hidden_results: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"fast-serp", "run",
		)
	})
}
