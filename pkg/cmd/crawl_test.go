// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestCrawlList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"crawl", "list",
			"--cursor", "cursor",
			"--limit", "10",
			"--status", "queued",
		)
	})
}

func TestCrawlRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"crawl", "run",
			"--url", "url",
			"--allow-external-links=false",
			"--allow-subdomains=false",
			"--callback", "{url: https://example.com, events: [started], headers: {foo: string}, metadata: {foo: bar}}",
			"--crawl-entire-domain=false",
			"--exclude-path", "/exclude-this-path",
			"--exclude-path", "/and-this-path",
			"--extract-options", "{auto_driver_configuration: {vx10: 2, vx10-pro: 0, vx6-fast: 1, vx6-stealth: 1, vx8: 5, vx8-pro: 5}, body: {key: value}, browser: chrome, browser_actions: [{goto: https://example.com/login}, {wait_for_element: '#login-form'}, {fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {click: '#submit'}, {screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}], city: Los Angeles, consent_header: true, cookies: sessionId=abc123; userId=user456, country: US, device: desktop, driver: vx8, expected_status_codes: [200, 201], formats: [html], headers: {Accept-Language: en-US, User-Agent: CustomBot/1.0}, http2: true, is_xhr: true, locale: en-US, markdown_backend: full_page, method: GET, network_capture: [{method: GET, resource_type: document, status_code: 100, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}], os: windows, parse: true, parser: {myParser: bar}, realtime_total_timeout: 15000, referrer_type: random, render: true, request_timeout: 30000, session: {id: id, prefetch_userbrowser: true, renew_on_blocked: true, retry: true, timeout: 1}, skill: dynamic-content, state: CA, tag: campaign-2024-q1, url: url}",
			"--ignore-query-parameters=false",
			"--include-path", "/include-this-path",
			"--include-path", "/and-this-path",
			"--limit", "100",
			"--max-discovery-depth", "3",
			"--name", "The best crawl ever",
			"--sitemap", "include",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(crawlRun)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"crawl", "run",
			"--url", "url",
			"--allow-external-links=false",
			"--allow-subdomains=false",
			"--callback", "{url: https://example.com, events: [started], headers: {foo: string}, metadata: {foo: bar}}",
			"--crawl-entire-domain=false",
			"--exclude-path", "/exclude-this-path",
			"--exclude-path", "/and-this-path",
			"--extract-options.auto-driver-configuration", "{vx10: 2, vx10-pro: 0, vx6-fast: 1, vx6-stealth: 1, vx8: 5, vx8-pro: 5}",
			"--extract-options.body", "{key: value}",
			"--extract-options.browser", "chrome",
			"--extract-options.browser-actions", "[{goto: https://example.com/login}, {wait_for_element: '#login-form'}, {fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {click: '#submit'}, {screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}]",
			"--extract-options.city", "Los Angeles",
			"--extract-options.consent-header=true",
			"--extract-options.cookies", "sessionId=abc123; userId=user456",
			"--extract-options.country", "US",
			"--extract-options.device", "desktop",
			"--extract-options.driver", "vx8",
			"--extract-options.expected-status-codes", "[200, 201]",
			"--extract-options.formats", "[html]",
			"--extract-options.headers", "{Accept-Language: en-US, User-Agent: CustomBot/1.0}",
			"--extract-options.http2=true",
			"--extract-options.is-xhr=true",
			"--extract-options.locale", "en-US",
			"--extract-options.markdown-backend", "full_page",
			"--extract-options.method", "GET",
			"--extract-options.network-capture", "[{method: GET, resource_type: document, status_code: 100, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}]",
			"--extract-options.os", "windows",
			"--extract-options.parse=true",
			"--extract-options.parser", "{myParser: bar}",
			"--extract-options.realtime-total-timeout", "15000",
			"--extract-options.referrer-type", "random",
			"--extract-options.render=true",
			"--extract-options.request-timeout", "30000",
			"--extract-options.session", "{id: id, prefetch_userbrowser: true, renew_on_blocked: true, retry: true, timeout: 1}",
			"--extract-options.skill", "dynamic-content",
			"--extract-options.state", "CA",
			"--extract-options.tag", "campaign-2024-q1",
			"--extract-options.url", "url",
			"--ignore-query-parameters=false",
			"--include-path", "/include-this-path",
			"--include-path", "/and-this-path",
			"--limit", "100",
			"--max-discovery-depth", "3",
			"--name", "The best crawl ever",
			"--sitemap", "include",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: url\n" +
			"allow_external_links: false\n" +
			"allow_subdomains: false\n" +
			"callback:\n" +
			"  url: https://example.com\n" +
			"  events:\n" +
			"    - started\n" +
			"  headers:\n" +
			"    foo: string\n" +
			"  metadata:\n" +
			"    foo: bar\n" +
			"crawl_entire_domain: false\n" +
			"exclude_paths:\n" +
			"  - /exclude-this-path\n" +
			"  - /and-this-path\n" +
			"extract_options:\n" +
			"  auto_driver_configuration:\n" +
			"    vx10: 2\n" +
			"    vx10-pro: 0\n" +
			"    vx6-fast: 1\n" +
			"    vx6-stealth: 1\n" +
			"    vx8: 5\n" +
			"    vx8-pro: 5\n" +
			"  body:\n" +
			"    key: value\n" +
			"  browser: chrome\n" +
			"  browser_actions:\n" +
			"    - goto: https://example.com/login\n" +
			"    - wait_for_element: '#login-form'\n" +
			"    - fill:\n" +
			"        selector: '#username'\n" +
			"        value: user@example.com\n" +
			"        click_on_element: true\n" +
			"        delay: 1000\n" +
			"        mode: type\n" +
			"        mouse_movement_strategy: linear\n" +
			"        required: 'true'\n" +
			"        scroll: true\n" +
			"        skip: 'true'\n" +
			"        timeout: 0\n" +
			"        typing_interval: 1000\n" +
			"        typing_strategy: simple\n" +
			"        visible: true\n" +
			"    - fill:\n" +
			"        selector: '#password'\n" +
			"        value: password123\n" +
			"        click_on_element: true\n" +
			"        delay: 1000\n" +
			"        mode: type\n" +
			"        mouse_movement_strategy: linear\n" +
			"        required: 'true'\n" +
			"        scroll: true\n" +
			"        skip: 'true'\n" +
			"        timeout: 0\n" +
			"        typing_interval: 1000\n" +
			"        typing_strategy: simple\n" +
			"        visible: true\n" +
			"    - click: '#submit'\n" +
			"    - screenshot:\n" +
			"        format: png\n" +
			"        full_page: true\n" +
			"        quality: 0\n" +
			"        required: 'true'\n" +
			"        skip: 'true'\n" +
			"  city: Los Angeles\n" +
			"  consent_header: true\n" +
			"  cookies: sessionId=abc123; userId=user456\n" +
			"  country: US\n" +
			"  device: desktop\n" +
			"  driver: vx8\n" +
			"  expected_status_codes:\n" +
			"    - 200\n" +
			"    - 201\n" +
			"  formats:\n" +
			"    - html\n" +
			"  headers:\n" +
			"    Accept-Language: en-US\n" +
			"    User-Agent: CustomBot/1.0\n" +
			"  http2: true\n" +
			"  is_xhr: true\n" +
			"  locale: en-US\n" +
			"  markdown_backend: full_page\n" +
			"  method: GET\n" +
			"  network_capture:\n" +
			"    - method: GET\n" +
			"      resource_type: document\n" +
			"      status_code: 100\n" +
			"      url:\n" +
			"        value: value\n" +
			"        type: exact\n" +
			"      validation: true\n" +
			"      wait_for_requests_count: 0\n" +
			"      wait_for_requests_count_timeout: 1\n" +
			"  os: windows\n" +
			"  parse: true\n" +
			"  parser:\n" +
			"    myParser: bar\n" +
			"  realtime_total_timeout: 15000\n" +
			"  referrer_type: random\n" +
			"  render: true\n" +
			"  request_timeout: 30000\n" +
			"  session:\n" +
			"    id: id\n" +
			"    prefetch_userbrowser: true\n" +
			"    renew_on_blocked: true\n" +
			"    retry: true\n" +
			"    timeout: 1\n" +
			"  skill: dynamic-content\n" +
			"  state: CA\n" +
			"  tag: campaign-2024-q1\n" +
			"  url: url\n" +
			"ignore_query_parameters: false\n" +
			"include_paths:\n" +
			"  - /include-this-path\n" +
			"  - /and-this-path\n" +
			"limit: 100\n" +
			"max_discovery_depth: 3\n" +
			"name: The best crawl ever\n" +
			"sitemap: include\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"crawl", "run",
		)
	})
}

func TestCrawlStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"crawl", "status",
			"--id", "123e4567-e89b-12d3-a456-426614174000",
		)
	})
}

func TestCrawlTerminate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"crawl", "terminate",
			"--id", "123e4567-e89b-12d3-a456-426614174000",
		)
	})
}
