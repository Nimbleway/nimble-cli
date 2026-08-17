// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestAgentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "create",
			"--agent-name", "agent_name",
			"--description", "description",
			"--display-name", "display_name",
			"--effort", "low",
			"--goal", "string",
			"--icon", "icon",
			"--is-active=true",
			"--output-schema", "{foo: bar}",
			"--skill", "skill",
			"--sources", "{allow: [{domains: [string], title: title, order: 0}], avoid: avoid, block: [{domains: [string], title: title, order: 0}], prioritize: prioritize}",
			"--suggested-question", "string",
			"--template", "template",
			"--use-case", "research",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "create",
			"--agent-name", "agent_name",
			"--description", "description",
			"--display-name", "display_name",
			"--effort", "low",
			"--goal", "string",
			"--icon", "icon",
			"--is-active=true",
			"--output-schema", "{foo: bar}",
			"--skill", "skill",
			"--sources.allow", "[{domains: [string], title: title, order: 0}]",
			"--sources.avoid", "avoid",
			"--sources.block", "[{domains: [string], title: title, order: 0}]",
			"--sources.prioritize", "prioritize",
			"--suggested-question", "string",
			"--template", "template",
			"--use-case", "research",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agent_name: agent_name\n" +
			"description: description\n" +
			"display_name: display_name\n" +
			"effort: low\n" +
			"goals:\n" +
			"  - string\n" +
			"icon: icon\n" +
			"is_active: true\n" +
			"output_schema:\n" +
			"  foo: bar\n" +
			"skill: skill\n" +
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
			"use_case: research\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "create",
		)
	})
}

func TestAgentsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "update",
			"--agent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body", "{op: add, path: path, from: from, value: {}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "update",
			"--agent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body.op", "add",
			"--body.path", "path",
			"--body.from", "from",
			"--body.value", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- op: add\n" +
			"  path: path\n" +
			"  from: from\n" +
			"  value: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "update",
			"--agent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAgentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestAgentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "delete",
			"--agent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAgentsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "get",
			"--agent-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAgentsRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "run",
			"--input", "input",
			"--agent-name", "agent_name",
			"--effort", "low",
			"--enable-events=true",
			"--input-data", "[{foo: bar}]",
			"--origin", "api",
			"--output-schema", "{foo: bar}",
			"--previous-interaction-id", "previous_interaction_id",
			"--skill", "skill",
			"--sources", "{allow: [{domains: [string], title: title, order: 0}], avoid: avoid, block: [{domains: [string], title: title, order: 0}], prioritize: prioritize}",
			"--use-case", "research",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsRun)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "run",
			"--input", "input",
			"--agent-name", "agent_name",
			"--effort", "low",
			"--enable-events=true",
			"--input-data", "[{foo: bar}]",
			"--origin", "api",
			"--output-schema", "{foo: bar}",
			"--previous-interaction-id", "previous_interaction_id",
			"--skill", "skill",
			"--sources.allow", "[{domains: [string], title: title, order: 0}]",
			"--sources.avoid", "avoid",
			"--sources.block", "[{domains: [string], title: title, order: 0}]",
			"--sources.prioritize", "prioritize",
			"--use-case", "research",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"input: input\n" +
			"agent_name: agent_name\n" +
			"effort: low\n" +
			"enable_events: true\n" +
			"input_data:\n" +
			"  - foo: bar\n" +
			"origin: api\n" +
			"output_schema:\n" +
			"  foo: bar\n" +
			"previous_interaction_id: previous_interaction_id\n" +
			"skill: skill\n" +
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
			"use_case: research\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "run",
		)
	})
}
