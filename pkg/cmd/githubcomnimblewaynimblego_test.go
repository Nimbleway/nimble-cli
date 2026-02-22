// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
)

func TestMap(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"map",
		"--url", "url",
		"--country", "US",
		"--domain-filter", "all",
		"--limit", "1000",
		"--locale", "en-US",
		"--sitemap", "include",
	)
}

func TestSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"search",
		"--query", "x",
		"--content-type", "[string]",
		"--country", "country",
		"--deep-search=true",
		"--end-date", "end_date",
		"--exclude-domain", "[string]",
		"--include-answer=true",
		"--include-domain", "[string]",
		"--locale", "locale",
		"--max-subagents", "1",
		"--num-results", "1",
		"--parsing-type", "plain_text",
		"--search-engine", "google_search",
		"--start-date", "start_date",
		"--time-range", "hour",
		"--topic", "string",
	)
}
