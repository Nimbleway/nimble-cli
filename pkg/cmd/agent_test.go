// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
)

func TestAgentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agents", "list",
		"--limit", "1",
		"--managed-by", "nimble",
		"--offset", "0",
		"--privacy", "public",
	)
}

func TestAgentsAsync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agents", "async",
		"--agent", "agent",
		"--params", "{foo: bar}",
		"--callback-url", "https://example.com/webhook/callback",
		"--localization=true",
		"--storage-compress=true",
		"--storage-object-name", "result-2024-01-15.json",
		"--storage-type", "s3",
		"--storage-url", "s3://bucket-name/path/to/object",
	)
}

func TestAgentsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agents", "get",
		"--template-name", "template_name",
	)
}

func TestAgentsRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"agents", "run",
		"--agent", "agent",
		"--params", "{foo: bar}",
		"--localization=true",
	)
}
