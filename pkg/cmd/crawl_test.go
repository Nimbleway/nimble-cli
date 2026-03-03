// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestCrawlList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "list",
		"--api-key", "string",
		"--cursor", "cursor",
		"--limit", "10",
		"--status", "queued",
	)
}

func TestCrawlRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "run",
		"--api-key", "string",
		"--url", "url",
		"--allow-external-links=false",
		"--allow-subdomains=false",
		"--callback", "{url: https://example.com, events: [started], headers: {foo: string}, metadata: {foo: bar}}",
		"--crawl-entire-domain=false",
		"--exclude-path", "/exclude-this-path",
		"--exclude-path", "/and-this-path",
		"--extract-options", "{browser: chrome, browser_actions: [{goto: https://example.com/login}, {wait_for_element: '#login-form'}, {fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {click: '#submit'}, {screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}], city: Los Angeles, consent_header: true, cookies: [{creation: creation, domain: domain, expires: expires, extensions: [string], hostOnly: true, httpOnly: true, lastAccessed: lastAccessed, maxAge: Infinity, name: name, path: path, pathIsDefault: true, sameSite: strict, secure: true, value: value}], country: US, device: desktop, driver: vx8, expected_status_codes: [200, 201], formats: [html], headers: {User-Agent: CustomBot/1.0, Accept-Language: en-US}, http2: true, is_xhr: true, locale: en-US, method: GET, network_capture: [{method: GET, resource_type: document, status_code: 100, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}], os: windows, parse: true, parser: {myParser: bar}, referrer_type: random, render: true, request_timeout: 30000, session: {id: id, prefetch_userbrowser: true, retry: true, timeout: 1}, skill: dynamic-content, state: CA, tag: campaign-2024-q1, url: url}",
		"--ignore-query-parameters=false",
		"--include-path", "/include-this-path",
		"--include-path", "/and-this-path",
		"--limit", "100",
		"--max-discovery-depth", "3",
		"--name", "The best crawl ever",
		"--sitemap", "include",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(crawlRun)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "run",
		"--url", "url",
		"--allow-external-links=false",
		"--allow-subdomains=false",
		"--callback", "{url: https://example.com, events: [started], headers: {foo: string}, metadata: {foo: bar}}",
		"--crawl-entire-domain=false",
		"--exclude-path", "/exclude-this-path",
		"--exclude-path", "/and-this-path",
		"--extract-options.browser", "chrome",
		"--extract-options.browser-actions", "[{goto: https://example.com/login}, {wait_for_element: '#login-form'}, {fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {click: '#submit'}, {screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}]",
		"--extract-options.city", "Los Angeles",
		"--extract-options.consent-header=true",
		"--extract-options.cookies", "[{creation: creation, domain: domain, expires: expires, extensions: [string], hostOnly: true, httpOnly: true, lastAccessed: lastAccessed, maxAge: Infinity, name: name, path: path, pathIsDefault: true, sameSite: strict, secure: true, value: value}]",
		"--extract-options.country", "US",
		"--extract-options.device", "desktop",
		"--extract-options.driver", "vx8",
		"--extract-options.expected-status-codes", "[200, 201]",
		"--extract-options.formats", "[html]",
		"--extract-options.headers", "{User-Agent: CustomBot/1.0, Accept-Language: en-US}",
		"--extract-options.http2=true",
		"--extract-options.is-xhr=true",
		"--extract-options.locale", "en-US",
		"--extract-options.method", "GET",
		"--extract-options.network-capture", "[{method: GET, resource_type: document, status_code: 100, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}]",
		"--extract-options.os", "windows",
		"--extract-options.parse=true",
		"--extract-options.parser", "{myParser: bar}",
		"--extract-options.referrer-type", "random",
		"--extract-options.render=true",
		"--extract-options.request-timeout", "30000",
		"--extract-options.session", "{id: id, prefetch_userbrowser: true, retry: true, timeout: 1}",
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
}

func TestCrawlStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "status",
		"--api-key", "string",
		"--id", "123e4567-e89b-12d3-a456-426614174000",
	)
}

func TestCrawlTerminate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "terminate",
		"--api-key", "string",
		"--id", "123e4567-e89b-12d3-a456-426614174000",
	)
}
