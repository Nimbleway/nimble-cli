// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
)

func TestBatchesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"batches", "list",
		)
	})
}

func TestBatchesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"batches", "get",
			"--batch-id", "123e4567-e89b-12d3-a456-426614174000",
		)
	})
}

func TestBatchesProgress(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"batches", "progress",
			"--batch-id", "123e4567-e89b-12d3-a456-426614174000",
		)
	})
}
