// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
)

func TestMap(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"map",
			"--url", "url",
			"--country", "US",
			"--domain-filter", "all",
			"--limit", "1000",
			"--locale", "en-US",
			"--sitemap", "include",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: url\n" +
			"country: US\n" +
			"domain_filter: all\n" +
			"limit: 1000\n" +
			"locale: en-US\n" +
			"sitemap: include\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"map",
		)
	})
}

func TestSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"search",
			"--query", "x",
			"--content-type", "[string]",
			"--country", "country",
			"--deep-search=true",
			"--end-date", "end_date",
			"--exclude-domain", "[string]",
			"--focus", "string",
			"--include-answer=true",
			"--include-domain", "[string]",
			"--locale", "locale",
			"--max-results", "1",
			"--max-subagents", "1",
			"--output-format", "plain_text",
			"--search-depth", "lite",
			"--start-date", "start_date",
			"--time-range", "hour",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: x\n" +
			"content_type:\n" +
			"  - string\n" +
			"country: country\n" +
			"deep_search: true\n" +
			"end_date: end_date\n" +
			"exclude_domains:\n" +
			"  - string\n" +
			"focus: string\n" +
			"include_answer: true\n" +
			"include_domains:\n" +
			"  - string\n" +
			"locale: locale\n" +
			"max_results: 1\n" +
			"max_subagents: 1\n" +
			"output_format: plain_text\n" +
			"search_depth: lite\n" +
			"start_date: start_date\n" +
			"time_range: hour\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"search",
		)
	})
}
