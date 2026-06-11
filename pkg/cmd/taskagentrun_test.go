// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
)

func TestTaskAgentRunsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent:runs", "list",
			"--agent-id", "agent_id",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestTaskAgentRunsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent:runs", "cancel",
			"--agent-id", "agent_id",
			"--run-id", "run_id",
		)
	})
}

func TestTaskAgentRunsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent:runs", "get",
			"--agent-id", "agent_id",
			"--run-id", "run_id",
		)
	})
}

func TestTaskAgentRunsGetResult(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent:runs", "get-result",
			"--agent-id", "agent_id",
			"--run-id", "run_id",
		)
	})
}

func TestTaskAgentRunsStreamEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent:runs", "stream-events",
			"--agent-id", "agent_id",
			"--run-id", "run_id",
		)
	})
}
