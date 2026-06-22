// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestTaskAgentCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent", "create",
			"--agent-name", "agent_name",
			"--description", "description",
			"--display-name", "display_name",
			"--domain-expertise", "domain_expertise",
			"--effort", "effort",
			"--goal", "string",
			"--icon", "icon",
			"--is-active=true",
			"--output-schema", "{foo: bar}",
			"--sources", "{allow: [{domains: [string], title: title, order: 0}], avoid: avoid, block: [{domains: [string], title: title, order: 0}], prioritize: prioritize}",
			"--suggested-question", "string",
			"--template", "template",
			"--use-case", "research",
			"--workspace-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(taskAgentCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent", "create",
			"--agent-name", "agent_name",
			"--description", "description",
			"--display-name", "display_name",
			"--domain-expertise", "domain_expertise",
			"--effort", "effort",
			"--goal", "string",
			"--icon", "icon",
			"--is-active=true",
			"--output-schema", "{foo: bar}",
			"--sources.allow", "[{domains: [string], title: title, order: 0}]",
			"--sources.avoid", "avoid",
			"--sources.block", "[{domains: [string], title: title, order: 0}]",
			"--sources.prioritize", "prioritize",
			"--suggested-question", "string",
			"--template", "template",
			"--use-case", "research",
			"--workspace-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent_name: agent_name\n" +
			"description: description\n" +
			"display_name: display_name\n" +
			"domain_expertise: domain_expertise\n" +
			"effort: effort\n" +
			"goals:\n" +
			"  - string\n" +
			"icon: icon\n" +
			"is_active: true\n" +
			"output_schema:\n" +
			"  foo: bar\n" +
			"sources:\n" +
			"  allow:\n" +
			"    - domains:\n" +
			"        - string\n" +
			"      title: title\n" +
			"      order: 0\n" +
			"  avoid: avoid\n" +
			"  block:\n" +
			"    - domains:\n" +
			"        - string\n" +
			"      title: title\n" +
			"      order: 0\n" +
			"  prioritize: prioritize\n" +
			"suggested_questions:\n" +
			"  - string\n" +
			"template: template\n" +
			"use_case: research\n" +
			"workspace_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"task-agent", "create",
		)
	})
}

func TestTaskAgentUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent", "update",
			"--agent-id", "agent_id",
			"--body", "{op: replace, path: path, value: {}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(taskAgentUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent", "update",
			"--agent-id", "agent_id",
			"--body.op", "replace",
			"--body.path", "path",
			"--body.value", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- op: replace\n" +
			"  path: path\n" +
			"  value: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"task-agent", "update",
			"--agent-id", "agent_id",
		)
	})
}

func TestTaskAgentList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent", "list",
			"--effort", "effort",
			"--limit", "1",
			"--offset", "0",
			"--use-case", "use_case",
		)
	})
}

func TestTaskAgentDeactivate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent", "deactivate",
			"--agent-id", "agent_id",
		)
	})
}

func TestTaskAgentGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent", "get",
			"--agent-id", "agent_id",
		)
	})
}

func TestTaskAgentRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent", "run",
			"--agent-id", "agent_id",
			"--input", "input",
			"--enable-events=true",
			"--output-schema", "{foo: bar}",
			"--sources", "{allow: [{domains: [string], title: title, order: 0}], avoid: avoid, block: [{domains: [string], title: title, order: 0}], prioritize: prioritize}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(taskAgentRun)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task-agent", "run",
			"--agent-id", "agent_id",
			"--input", "input",
			"--enable-events=true",
			"--output-schema", "{foo: bar}",
			"--sources.allow", "[{domains: [string], title: title, order: 0}]",
			"--sources.avoid", "avoid",
			"--sources.block", "[{domains: [string], title: title, order: 0}]",
			"--sources.prioritize", "prioritize",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"input: input\n" +
			"enable_events: true\n" +
			"output_schema:\n" +
			"  foo: bar\n" +
			"sources:\n" +
			"  allow:\n" +
			"    - domains:\n" +
			"        - string\n" +
			"      title: title\n" +
			"      order: 0\n" +
			"  avoid: avoid\n" +
			"  block:\n" +
			"    - domains:\n" +
			"        - string\n" +
			"      title: title\n" +
			"      order: 0\n" +
			"  prioritize: prioritize\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"task-agent", "run",
			"--agent-id", "agent_id",
		)
	})
}
