// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
)

func TestTaskAgentTemplatesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent:templates", "list",
			"--effort", "effort",
			"--limit", "1",
			"--offset", "0",
			"--use-case", "use_case",
		)
	})
}

func TestTaskAgentTemplatesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent:templates", "get",
			"--template-name", "template_name",
		)
	})
}
