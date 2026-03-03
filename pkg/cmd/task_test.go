// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
)

func TestTasksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tasks", "list",
		"--api-key", "string",
		"--cursor", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--limit", "10",
	)
}

func TestTasksGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tasks", "get",
		"--api-key", "string",
		"--task-id", "123e4567-e89b-12d3-a456-426614174000",
	)
}

func TestTasksResults(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tasks", "results",
		"--api-key", "string",
		"--task-id", "123e4567-e89b-12d3-a456-426614174000",
	)
}
