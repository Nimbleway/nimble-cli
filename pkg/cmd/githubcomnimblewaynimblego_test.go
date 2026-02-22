// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestExtract(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"extract",
		"--url", "url",
		"--browser", "chrome",
		"--browser-action", "{goto: https://example.com/login}",
		"--browser-action", "{wait_for_element: '#login-form'}",
		"--browser-action", "{fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}",
		"--browser-action", "{fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}",
		"--browser-action", "{click: '#submit'}",
		"--browser-action", "{screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}",
		"--city", "Los Angeles",
		"--consent-header=true",
		"--cookies", "[{creation: creation, domain: domain, expires: expires, extensions: [string], hostOnly: true, httpOnly: true, lastAccessed: lastAccessed, maxAge: Infinity, name: name, path: path, pathIsDefault: true, sameSite: strict, secure: true, value: value}]",
		"--country", "US",
		"--device", "desktop",
		"--driver", "vx8",
		"--expected-status-code", "200",
		"--expected-status-code", "201",
		"--format", "html",
		"--headers", "{User-Agent: CustomBot/1.0, Accept-Language: en-US}",
		"--http2=true",
		"--is-xhr=true",
		"--locale", "en-US",
		"--method", "GET",
		"--network-capture", "{method: GET, resource_type: document, status_code: 100, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}",
		"--os", "windows",
		"--parse=true",
		"--parser", "{myParser: bar}",
		"--referrer-type", "random",
		"--render=true",
		"--request-timeout", "30000",
		"--session", "{id: id, prefetch_userbrowser: true, retry: true, timeout: 1}",
		"--skill", "dynamic-content",
		"--state", "CA",
		"--tag", "campaign-2024-q1",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(extract)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"extract",
		"--url", "url",
		"--browser", "chrome",
		"--browser-action", "{goto: https://example.com/login}",
		"--browser-action", "{wait_for_element: '#login-form'}",
		"--browser-action", "{fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}",
		"--browser-action", "{fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}",
		"--browser-action", "{click: '#submit'}",
		"--browser-action", "{screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}",
		"--city", "Los Angeles",
		"--consent-header=true",
		"--cookies", "[{creation: creation, domain: domain, expires: expires, extensions: [string], hostOnly: true, httpOnly: true, lastAccessed: lastAccessed, maxAge: Infinity, name: name, path: path, pathIsDefault: true, sameSite: strict, secure: true, value: value}]",
		"--country", "US",
		"--device", "desktop",
		"--driver", "vx8",
		"--expected-status-code", "200",
		"--expected-status-code", "201",
		"--format", "html",
		"--headers", "{User-Agent: CustomBot/1.0, Accept-Language: en-US}",
		"--http2=true",
		"--is-xhr=true",
		"--locale", "en-US",
		"--method", "GET",
		"--network-capture.method", "GET",
		"--network-capture.resource-type", "document",
		"--network-capture.status-code", "100",
		"--network-capture.url", "{value: value, type: exact}",
		"--network-capture.validation=true",
		"--network-capture.wait-for-requests-count", "0",
		"--network-capture.wait-for-requests-count-timeout", "1",
		"--os", "windows",
		"--parse=true",
		"--parser", "{myParser: bar}",
		"--referrer-type", "random",
		"--render=true",
		"--request-timeout", "30000",
		"--session.id", "id",
		"--session.prefetch-userbrowser=true",
		"--session.retry=true",
		"--session.timeout", "1",
		"--skill", "dynamic-content",
		"--state", "CA",
		"--tag", "campaign-2024-q1",
	)
}

func TestExtractAsync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"extract-async",
		"--url", "url",
		"--browser", "chrome",
		"--browser-action", "{goto: https://example.com/login}",
		"--browser-action", "{wait_for_element: '#login-form'}",
		"--browser-action", "{fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}",
		"--browser-action", "{fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}",
		"--browser-action", "{click: '#submit'}",
		"--browser-action", "{screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}",
		"--callback-url", "https://example.com/webhook/callback",
		"--city", "Los Angeles",
		"--consent-header=true",
		"--cookies", "[{creation: creation, domain: domain, expires: expires, extensions: [string], hostOnly: true, httpOnly: true, lastAccessed: lastAccessed, maxAge: Infinity, name: name, path: path, pathIsDefault: true, sameSite: strict, secure: true, value: value}]",
		"--country", "US",
		"--device", "desktop",
		"--driver", "vx8",
		"--expected-status-code", "200",
		"--expected-status-code", "201",
		"--format", "html",
		"--headers", "{User-Agent: CustomBot/1.0, Accept-Language: en-US}",
		"--http2=true",
		"--is-xhr=true",
		"--locale", "en-US",
		"--method", "GET",
		"--network-capture", "{method: GET, resource_type: document, status_code: 100, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}",
		"--os", "windows",
		"--parse=true",
		"--parser", "{myParser: bar}",
		"--referrer-type", "random",
		"--render=true",
		"--request-timeout", "30000",
		"--session", "{id: id, prefetch_userbrowser: true, retry: true, timeout: 1}",
		"--skill", "dynamic-content",
		"--state", "CA",
		"--storage-compress=true",
		"--storage-object-name", "result-2024-01-15.json",
		"--storage-type", "s3",
		"--storage-url", "s3://bucket-name/path/to/object",
		"--tag", "campaign-2024-q1",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(extractAsync)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"extract-async",
		"--url", "url",
		"--browser", "chrome",
		"--browser-action", "{goto: https://example.com/login}",
		"--browser-action", "{wait_for_element: '#login-form'}",
		"--browser-action", "{fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}",
		"--browser-action", "{fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}",
		"--browser-action", "{click: '#submit'}",
		"--browser-action", "{screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}",
		"--callback-url", "https://example.com/webhook/callback",
		"--city", "Los Angeles",
		"--consent-header=true",
		"--cookies", "[{creation: creation, domain: domain, expires: expires, extensions: [string], hostOnly: true, httpOnly: true, lastAccessed: lastAccessed, maxAge: Infinity, name: name, path: path, pathIsDefault: true, sameSite: strict, secure: true, value: value}]",
		"--country", "US",
		"--device", "desktop",
		"--driver", "vx8",
		"--expected-status-code", "200",
		"--expected-status-code", "201",
		"--format", "html",
		"--headers", "{User-Agent: CustomBot/1.0, Accept-Language: en-US}",
		"--http2=true",
		"--is-xhr=true",
		"--locale", "en-US",
		"--method", "GET",
		"--network-capture.method", "GET",
		"--network-capture.resource-type", "document",
		"--network-capture.status-code", "100",
		"--network-capture.url", "{value: value, type: exact}",
		"--network-capture.validation=true",
		"--network-capture.wait-for-requests-count", "0",
		"--network-capture.wait-for-requests-count-timeout", "1",
		"--os", "windows",
		"--parse=true",
		"--parser", "{myParser: bar}",
		"--referrer-type", "random",
		"--render=true",
		"--request-timeout", "30000",
		"--session.id", "id",
		"--session.prefetch-userbrowser=true",
		"--session.retry=true",
		"--session.timeout", "1",
		"--skill", "dynamic-content",
		"--state", "CA",
		"--storage-compress=true",
		"--storage-object-name", "result-2024-01-15.json",
		"--storage-type", "s3",
		"--storage-url", "s3://bucket-name/path/to/object",
		"--tag", "campaign-2024-q1",
	)
}

func TestMap(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"map",
		"--url", "url",
		"--country", "US",
		"--domain-filter", "all",
		"--limit", "1000",
		"--locale", "en-US",
		"--sitemap", "include",
	)
}

func TestSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"search",
		"--query", "x",
		"--content-type", "[string]",
		"--country", "country",
		"--deep-search=true",
		"--end-date", "end_date",
		"--exclude-domain", "[string]",
		"--include-answer=true",
		"--include-domain", "[string]",
		"--locale", "locale",
		"--max-subagents", "1",
		"--num-results", "1",
		"--parsing-type", "plain_text",
		"--search-engine", "google_search",
		"--start-date", "start_date",
		"--time-range", "hour",
		"--topic", "string",
	)
}
