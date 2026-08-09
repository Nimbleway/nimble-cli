// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestJobsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs", "create",
			"--extract-template-name", "extract_template_name",
			"--name", "name",
			"--description", "description",
			"--destination", "{path: path, type: file, format: jsonl}",
			"--display-name", "display_name",
			"--inputs", "{type: s3, data: [{foo: bar}], file_path: file_path}",
			"--schedule", "{cron: cron, enabled: true}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(jobsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs", "create",
			"--extract-template-name", "extract_template_name",
			"--name", "name",
			"--description", "description",
			"--destination.path", "path",
			"--destination.type", "file",
			"--destination.format", "jsonl",
			"--display-name", "display_name",
			"--inputs.type", "s3",
			"--inputs.data", "[{foo: bar}]",
			"--inputs.file-path", "file_path",
			"--schedule.cron", "cron",
			"--schedule.enabled=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"extract_template_name: extract_template_name\n" +
			"name: name\n" +
			"description: description\n" +
			"destination:\n" +
			"  path: path\n" +
			"  type: file\n" +
			"  format: jsonl\n" +
			"display_name: display_name\n" +
			"inputs:\n" +
			"  type: s3\n" +
			"  data:\n" +
			"    - foo: bar\n" +
			"  file_path: file_path\n" +
			"schedule:\n" +
			"  cron: cron\n" +
			"  enabled: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"jobs", "create",
		)
	})
}

func TestJobsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs", "update",
			"--job-id", "job_id",
			"--description", "description",
			"--destination", "{path: path, type: file, format: jsonl}",
			"--display-name", "display_name",
			"--inputs", "{type: s3, data: [{foo: bar}], file_path: file_path}",
			"--schedule", "{cron: cron, enabled: true}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(jobsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs", "update",
			"--job-id", "job_id",
			"--description", "description",
			"--destination.path", "path",
			"--destination.type", "file",
			"--destination.format", "jsonl",
			"--display-name", "display_name",
			"--inputs.type", "s3",
			"--inputs.data", "[{foo: bar}]",
			"--inputs.file-path", "file_path",
			"--schedule.cron", "cron",
			"--schedule.enabled=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"destination:\n" +
			"  path: path\n" +
			"  type: file\n" +
			"  format: jsonl\n" +
			"display_name: display_name\n" +
			"inputs:\n" +
			"  type: s3\n" +
			"  data:\n" +
			"    - foo: bar\n" +
			"  file_path: file_path\n" +
			"schedule:\n" +
			"  cron: cron\n" +
			"  enabled: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"jobs", "update",
			"--job-id", "job_id",
		)
	})
}

func TestJobsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestJobsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs", "delete",
			"--job-id", "job_id",
		)
	})
}

func TestJobsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"jobs", "get",
			"--job-id", "job_id",
		)
	})
}
