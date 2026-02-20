// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
)

func TestCrawlList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "list",
		"--cursor", "cursor",
		"--limit", "10",
		"--status", "queued",
	)
}

func TestCrawlStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "status",
		"--id", "123e4567-e89b-12d3-a456-426614174000",
	)
}

func TestCrawlTerminate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "terminate",
		"--id", "123e4567-e89b-12d3-a456-426614174000",
	)
}
