// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
)

func TestJobsRunsArtifactsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs:runs:artifacts", "list",
			"--run-id", "run_id",
		)
	})
}

func TestJobsRunsArtifactsDownloadURL(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs:runs:artifacts", "download-url",
			"--run-id", "run_id",
			"--artifact-id", "artifact_id",
		)
	})
}

func TestJobsRunsArtifactsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs:runs:artifacts", "get",
			"--run-id", "run_id",
			"--artifact-id", "artifact_id",
		)
	})
}

func TestJobsRunsArtifactsPreview(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs:runs:artifacts", "preview",
			"--run-id", "run_id",
			"--artifact-id", "artifact_id",
		)
	})
}
