// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestJobsRunsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs:runs", "create",
			"--job-id", "job_id",
			"--inputs", "{type: s3, data: [{foo: bar}], file_path: file_path, node_data: {foo: [{foo: bar}]}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(jobsRunsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs:runs", "create",
			"--job-id", "job_id",
			"--inputs.type", "s3",
			"--inputs.data", "[{foo: bar}]",
			"--inputs.file-path", "file_path",
			"--inputs.node-data", "{foo: [{foo: bar}]}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"inputs:\n" +
			"  type: s3\n" +
			"  data:\n" +
			"    - foo: bar\n" +
			"  file_path: file_path\n" +
			"  node_data:\n" +
			"    foo:\n" +
			"      - foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"jobs:runs", "create",
			"--job-id", "job_id",
		)
	})
}

func TestJobsRunsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs:runs", "list",
			"--job-id", "job_id",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestJobsRunsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs:runs", "cancel",
			"--run-id", "run_id",
		)
	})
}

func TestJobsRunsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs:runs", "get",
			"--run-id", "run_id",
		)
	})
}
