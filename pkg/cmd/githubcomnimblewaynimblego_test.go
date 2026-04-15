// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Nimbleway/nimble-cli/internal/mocktest"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
)

func TestExtract(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
			"--cookies", "sessionId=abc123; userId=user456",
			"--country", "US",
			"--device", "desktop",
			"--driver", "vx8",
			"--expected-status-code", "200",
			"--expected-status-code", "201",
			"--format", "html",
			"--headers", "{Accept-Language: en-US, User-Agent: CustomBot/1.0}",
			"--http2=true",
			"--is-xhr=true",
			"--locale", "en-US",
			"--markdown-backend", "full_page",
			"--method", "GET",
			"--network-capture", "{method: GET, resource_type: document, status_code: 100, stop_on_render_flow_failure: true, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(extract)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
			"--cookies", "sessionId=abc123; userId=user456",
			"--country", "US",
			"--device", "desktop",
			"--driver", "vx8",
			"--expected-status-code", "200",
			"--expected-status-code", "201",
			"--format", "html",
			"--headers", "{Accept-Language: en-US, User-Agent: CustomBot/1.0}",
			"--http2=true",
			"--is-xhr=true",
			"--locale", "en-US",
			"--markdown-backend", "full_page",
			"--method", "GET",
			"--network-capture.method", "GET",
			"--network-capture.resource-type", "document",
			"--network-capture.status-code", "100",
			"--network-capture.stop-on-render-flow-failure=true",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: url\n" +
			"browser: chrome\n" +
			"browser_actions:\n" +
			"  - goto: https://example.com/login\n" +
			"  - wait_for_element: '#login-form'\n" +
			"  - fill:\n" +
			"      selector: '#username'\n" +
			"      value: user@example.com\n" +
			"      click_on_element: true\n" +
			"      delay: 1000\n" +
			"      mode: type\n" +
			"      mouse_movement_strategy: linear\n" +
			"      required: 'true'\n" +
			"      scroll: true\n" +
			"      skip: 'true'\n" +
			"      timeout: 0\n" +
			"      typing_interval: 1000\n" +
			"      typing_strategy: simple\n" +
			"      visible: true\n" +
			"  - fill:\n" +
			"      selector: '#password'\n" +
			"      value: password123\n" +
			"      click_on_element: true\n" +
			"      delay: 1000\n" +
			"      mode: type\n" +
			"      mouse_movement_strategy: linear\n" +
			"      required: 'true'\n" +
			"      scroll: true\n" +
			"      skip: 'true'\n" +
			"      timeout: 0\n" +
			"      typing_interval: 1000\n" +
			"      typing_strategy: simple\n" +
			"      visible: true\n" +
			"  - click: '#submit'\n" +
			"  - screenshot:\n" +
			"      format: png\n" +
			"      full_page: true\n" +
			"      quality: 0\n" +
			"      required: 'true'\n" +
			"      skip: 'true'\n" +
			"city: Los Angeles\n" +
			"consent_header: true\n" +
			"cookies: sessionId=abc123; userId=user456\n" +
			"country: US\n" +
			"device: desktop\n" +
			"driver: vx8\n" +
			"expected_status_codes:\n" +
			"  - 200\n" +
			"  - 201\n" +
			"formats:\n" +
			"  - html\n" +
			"headers:\n" +
			"  Accept-Language: en-US\n" +
			"  User-Agent: CustomBot/1.0\n" +
			"http2: true\n" +
			"is_xhr: true\n" +
			"locale: en-US\n" +
			"markdown_backend: full_page\n" +
			"method: GET\n" +
			"network_capture:\n" +
			"  - method: GET\n" +
			"    resource_type: document\n" +
			"    status_code: 100\n" +
			"    stop_on_render_flow_failure: true\n" +
			"    url:\n" +
			"      value: value\n" +
			"      type: exact\n" +
			"    validation: true\n" +
			"    wait_for_requests_count: 0\n" +
			"    wait_for_requests_count_timeout: 1\n" +
			"os: windows\n" +
			"parse: true\n" +
			"parser:\n" +
			"  myParser: bar\n" +
			"referrer_type: random\n" +
			"render: true\n" +
			"request_timeout: 30000\n" +
			"session:\n" +
			"  id: id\n" +
			"  prefetch_userbrowser: true\n" +
			"  retry: true\n" +
			"  timeout: 1\n" +
			"skill: dynamic-content\n" +
			"state: CA\n" +
			"tag: campaign-2024-q1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"extract",
		)
	})
}

func TestExtractAsync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
			"--cookies", "sessionId=abc123; userId=user456",
			"--country", "US",
			"--device", "desktop",
			"--driver", "vx8",
			"--expected-status-code", "200",
			"--expected-status-code", "201",
			"--format", "html",
			"--headers", "{Accept-Language: en-US, User-Agent: CustomBot/1.0}",
			"--http2=true",
			"--is-xhr=true",
			"--locale", "en-US",
			"--markdown-backend", "full_page",
			"--method", "GET",
			"--network-capture", "{method: GET, resource_type: document, status_code: 100, stop_on_render_flow_failure: true, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}",
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
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(extractAsync)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
			"--cookies", "sessionId=abc123; userId=user456",
			"--country", "US",
			"--device", "desktop",
			"--driver", "vx8",
			"--expected-status-code", "200",
			"--expected-status-code", "201",
			"--format", "html",
			"--headers", "{Accept-Language: en-US, User-Agent: CustomBot/1.0}",
			"--http2=true",
			"--is-xhr=true",
			"--locale", "en-US",
			"--markdown-backend", "full_page",
			"--method", "GET",
			"--network-capture.method", "GET",
			"--network-capture.resource-type", "document",
			"--network-capture.status-code", "100",
			"--network-capture.stop-on-render-flow-failure=true",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: url\n" +
			"browser: chrome\n" +
			"browser_actions:\n" +
			"  - goto: https://example.com/login\n" +
			"  - wait_for_element: '#login-form'\n" +
			"  - fill:\n" +
			"      selector: '#username'\n" +
			"      value: user@example.com\n" +
			"      click_on_element: true\n" +
			"      delay: 1000\n" +
			"      mode: type\n" +
			"      mouse_movement_strategy: linear\n" +
			"      required: 'true'\n" +
			"      scroll: true\n" +
			"      skip: 'true'\n" +
			"      timeout: 0\n" +
			"      typing_interval: 1000\n" +
			"      typing_strategy: simple\n" +
			"      visible: true\n" +
			"  - fill:\n" +
			"      selector: '#password'\n" +
			"      value: password123\n" +
			"      click_on_element: true\n" +
			"      delay: 1000\n" +
			"      mode: type\n" +
			"      mouse_movement_strategy: linear\n" +
			"      required: 'true'\n" +
			"      scroll: true\n" +
			"      skip: 'true'\n" +
			"      timeout: 0\n" +
			"      typing_interval: 1000\n" +
			"      typing_strategy: simple\n" +
			"      visible: true\n" +
			"  - click: '#submit'\n" +
			"  - screenshot:\n" +
			"      format: png\n" +
			"      full_page: true\n" +
			"      quality: 0\n" +
			"      required: 'true'\n" +
			"      skip: 'true'\n" +
			"callback_url: https://example.com/webhook/callback\n" +
			"city: Los Angeles\n" +
			"consent_header: true\n" +
			"cookies: sessionId=abc123; userId=user456\n" +
			"country: US\n" +
			"device: desktop\n" +
			"driver: vx8\n" +
			"expected_status_codes:\n" +
			"  - 200\n" +
			"  - 201\n" +
			"formats:\n" +
			"  - html\n" +
			"headers:\n" +
			"  Accept-Language: en-US\n" +
			"  User-Agent: CustomBot/1.0\n" +
			"http2: true\n" +
			"is_xhr: true\n" +
			"locale: en-US\n" +
			"markdown_backend: full_page\n" +
			"method: GET\n" +
			"network_capture:\n" +
			"  - method: GET\n" +
			"    resource_type: document\n" +
			"    status_code: 100\n" +
			"    stop_on_render_flow_failure: true\n" +
			"    url:\n" +
			"      value: value\n" +
			"      type: exact\n" +
			"    validation: true\n" +
			"    wait_for_requests_count: 0\n" +
			"    wait_for_requests_count_timeout: 1\n" +
			"os: windows\n" +
			"parse: true\n" +
			"parser:\n" +
			"  myParser: bar\n" +
			"referrer_type: random\n" +
			"render: true\n" +
			"request_timeout: 30000\n" +
			"session:\n" +
			"  id: id\n" +
			"  prefetch_userbrowser: true\n" +
			"  retry: true\n" +
			"  timeout: 1\n" +
			"skill: dynamic-content\n" +
			"state: CA\n" +
			"storage_compress: true\n" +
			"storage_object_name: result-2024-01-15.json\n" +
			"storage_type: s3\n" +
			"storage_url: s3://bucket-name/path/to/object\n" +
			"tag: campaign-2024-q1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"extract-async",
		)
	})
}

func TestExtractBatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract-batch",
			"--input", "{browser: chrome, browser_actions: [{goto: https://example.com/login}, {wait_for_element: '#login-form'}, {fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {click: '#submit'}, {screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}], callback_url: https://example.com/webhook/callback, city: Los Angeles, consent_header: true, cookies: sessionId=abc123; userId=user456, country: US, device: desktop, driver: vx8, expected_status_codes: [200, 201], formats: [html], headers: {Accept-Language: en-US, User-Agent: CustomBot/1.0}, http2: true, is_xhr: true, locale: en-US, markdown_backend: full_page, method: GET, network_capture: [{method: GET, resource_type: document, status_code: 100, stop_on_render_flow_failure: true, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}], os: windows, parse: true, parser: {myParser: bar}, referrer_type: random, render: true, request_timeout: 30000, session: {id: id, prefetch_userbrowser: true, retry: true, timeout: 1}, skill: dynamic-content, state: CA, storage_compress: true, storage_object_name: result-2024-01-15.json, storage_type: s3, storage_url: s3://bucket-name/path/to/object, tag: campaign-2024-q1, url: url}",
			"--shared-inputs", "{browser: chrome, browser_actions: [{goto: https://example.com/login}, {wait_for_element: '#login-form'}, {fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {click: '#submit'}, {screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}], callback_url: https://example.com/webhook/callback, city: Los Angeles, consent_header: true, cookies: sessionId=abc123; userId=user456, country: US, device: desktop, driver: vx8, expected_status_codes: [200, 201], formats: [html], headers: {Accept-Language: en-US, User-Agent: CustomBot/1.0}, http2: true, is_xhr: true, locale: en-US, markdown_backend: full_page, method: GET, network_capture: [{method: GET, resource_type: document, status_code: 100, stop_on_render_flow_failure: true, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}], os: windows, parse: true, parser: {myParser: bar}, referrer_type: random, render: true, request_timeout: 30000, session: {id: id, prefetch_userbrowser: true, retry: true, timeout: 1}, skill: dynamic-content, state: CA, storage_compress: true, storage_object_name: result-2024-01-15.json, storage_type: s3, storage_url: s3://bucket-name/path/to/object, tag: campaign-2024-q1, url: url}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(extractBatch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract-batch",
			"--input.browser", "chrome",
			"--input.browser-actions", "[{goto: https://example.com/login}, {wait_for_element: '#login-form'}, {fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {click: '#submit'}, {screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}]",
			"--input.callback-url", "https://example.com/webhook/callback",
			"--input.city", "Los Angeles",
			"--input.consent-header=true",
			"--input.cookies", "sessionId=abc123; userId=user456",
			"--input.country", "US",
			"--input.device", "desktop",
			"--input.driver", "vx8",
			"--input.expected-status-codes", "[200, 201]",
			"--input.formats", "[html]",
			"--input.headers", "{Accept-Language: en-US, User-Agent: CustomBot/1.0}",
			"--input.http2=true",
			"--input.is-xhr=true",
			"--input.locale", "en-US",
			"--input.markdown-backend", "full_page",
			"--input.method", "GET",
			"--input.network-capture", "[{method: GET, resource_type: document, status_code: 100, stop_on_render_flow_failure: true, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}]",
			"--input.os", "windows",
			"--input.parse=true",
			"--input.parser", "{myParser: bar}",
			"--input.referrer-type", "random",
			"--input.render=true",
			"--input.request-timeout", "30000",
			"--input.session", "{id: id, prefetch_userbrowser: true, retry: true, timeout: 1}",
			"--input.skill", "dynamic-content",
			"--input.state", "CA",
			"--input.storage-compress=true",
			"--input.storage-object-name", "result-2024-01-15.json",
			"--input.storage-type", "s3",
			"--input.storage-url", "s3://bucket-name/path/to/object",
			"--input.tag", "campaign-2024-q1",
			"--input.url", "url",
			"--shared-inputs.browser", "chrome",
			"--shared-inputs.browser-actions", "[{goto: https://example.com/login}, {wait_for_element: '#login-form'}, {fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {click: '#submit'}, {screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}]",
			"--shared-inputs.callback-url", "https://example.com/webhook/callback",
			"--shared-inputs.city", "Los Angeles",
			"--shared-inputs.consent-header=true",
			"--shared-inputs.cookies", "sessionId=abc123; userId=user456",
			"--shared-inputs.country", "US",
			"--shared-inputs.device", "desktop",
			"--shared-inputs.driver", "vx8",
			"--shared-inputs.expected-status-codes", "[200, 201]",
			"--shared-inputs.formats", "[html]",
			"--shared-inputs.headers", "{Accept-Language: en-US, User-Agent: CustomBot/1.0}",
			"--shared-inputs.http2=true",
			"--shared-inputs.is-xhr=true",
			"--shared-inputs.locale", "en-US",
			"--shared-inputs.markdown-backend", "full_page",
			"--shared-inputs.method", "GET",
			"--shared-inputs.network-capture", "[{method: GET, resource_type: document, status_code: 100, stop_on_render_flow_failure: true, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}]",
			"--shared-inputs.os", "windows",
			"--shared-inputs.parse=true",
			"--shared-inputs.parser", "{myParser: bar}",
			"--shared-inputs.referrer-type", "random",
			"--shared-inputs.render=true",
			"--shared-inputs.request-timeout", "30000",
			"--shared-inputs.session", "{id: id, prefetch_userbrowser: true, retry: true, timeout: 1}",
			"--shared-inputs.skill", "dynamic-content",
			"--shared-inputs.state", "CA",
			"--shared-inputs.storage-compress=true",
			"--shared-inputs.storage-object-name", "result-2024-01-15.json",
			"--shared-inputs.storage-type", "s3",
			"--shared-inputs.storage-url", "s3://bucket-name/path/to/object",
			"--shared-inputs.tag", "campaign-2024-q1",
			"--shared-inputs.url", "url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"inputs:\n" +
			"  - browser: chrome\n" +
			"    browser_actions:\n" +
			"      - goto: https://example.com/login\n" +
			"      - wait_for_element: '#login-form'\n" +
			"      - fill:\n" +
			"          selector: '#username'\n" +
			"          value: user@example.com\n" +
			"          click_on_element: true\n" +
			"          delay: 1000\n" +
			"          mode: type\n" +
			"          mouse_movement_strategy: linear\n" +
			"          required: 'true'\n" +
			"          scroll: true\n" +
			"          skip: 'true'\n" +
			"          timeout: 0\n" +
			"          typing_interval: 1000\n" +
			"          typing_strategy: simple\n" +
			"          visible: true\n" +
			"      - fill:\n" +
			"          selector: '#password'\n" +
			"          value: password123\n" +
			"          click_on_element: true\n" +
			"          delay: 1000\n" +
			"          mode: type\n" +
			"          mouse_movement_strategy: linear\n" +
			"          required: 'true'\n" +
			"          scroll: true\n" +
			"          skip: 'true'\n" +
			"          timeout: 0\n" +
			"          typing_interval: 1000\n" +
			"          typing_strategy: simple\n" +
			"          visible: true\n" +
			"      - click: '#submit'\n" +
			"      - screenshot:\n" +
			"          format: png\n" +
			"          full_page: true\n" +
			"          quality: 0\n" +
			"          required: 'true'\n" +
			"          skip: 'true'\n" +
			"    callback_url: https://example.com/webhook/callback\n" +
			"    city: Los Angeles\n" +
			"    consent_header: true\n" +
			"    cookies: sessionId=abc123; userId=user456\n" +
			"    country: US\n" +
			"    device: desktop\n" +
			"    driver: vx8\n" +
			"    expected_status_codes:\n" +
			"      - 200\n" +
			"      - 201\n" +
			"    formats:\n" +
			"      - html\n" +
			"    headers:\n" +
			"      Accept-Language: en-US\n" +
			"      User-Agent: CustomBot/1.0\n" +
			"    http2: true\n" +
			"    is_xhr: true\n" +
			"    locale: en-US\n" +
			"    markdown_backend: full_page\n" +
			"    method: GET\n" +
			"    network_capture:\n" +
			"      - method: GET\n" +
			"        resource_type: document\n" +
			"        status_code: 100\n" +
			"        stop_on_render_flow_failure: true\n" +
			"        url:\n" +
			"          value: value\n" +
			"          type: exact\n" +
			"        validation: true\n" +
			"        wait_for_requests_count: 0\n" +
			"        wait_for_requests_count_timeout: 1\n" +
			"    os: windows\n" +
			"    parse: true\n" +
			"    parser:\n" +
			"      myParser: bar\n" +
			"    referrer_type: random\n" +
			"    render: true\n" +
			"    request_timeout: 30000\n" +
			"    session:\n" +
			"      id: id\n" +
			"      prefetch_userbrowser: true\n" +
			"      retry: true\n" +
			"      timeout: 1\n" +
			"    skill: dynamic-content\n" +
			"    state: CA\n" +
			"    storage_compress: true\n" +
			"    storage_object_name: result-2024-01-15.json\n" +
			"    storage_type: s3\n" +
			"    storage_url: s3://bucket-name/path/to/object\n" +
			"    tag: campaign-2024-q1\n" +
			"    url: url\n" +
			"shared_inputs:\n" +
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
			"  callback_url: https://example.com/webhook/callback\n" +
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
			"      stop_on_render_flow_failure: true\n" +
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
			"  referrer_type: random\n" +
			"  render: true\n" +
			"  request_timeout: 30000\n" +
			"  session:\n" +
			"    id: id\n" +
			"    prefetch_userbrowser: true\n" +
			"    retry: true\n" +
			"    timeout: 1\n" +
			"  skill: dynamic-content\n" +
			"  state: CA\n" +
			"  storage_compress: true\n" +
			"  storage_object_name: result-2024-01-15.json\n" +
			"  storage_type: s3\n" +
			"  storage_url: s3://bucket-name/path/to/object\n" +
			"  tag: campaign-2024-q1\n" +
			"  url: url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"extract-batch",
		)
	})
}

func TestMap(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"map",
			"--url", "url",
			"--country", "US",
			"--domain-filter", "all",
			"--limit", "1000",
			"--locale", "en-US",
			"--sitemap", "include",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: url\n" +
			"country: US\n" +
			"domain_filter: all\n" +
			"limit: 1000\n" +
			"locale: en-US\n" +
			"sitemap: include\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"map",
		)
	})
}

func TestSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"search",
			"--query", "x",
			"--content-type", "[string]",
			"--country", "country",
			"--deep-search=true",
			"--end-date", "end_date",
			"--exclude-domain", "[string]",
			"--focus", "string",
			"--include-answer=true",
			"--include-domain", "[string]",
			"--locale", "locale",
			"--max-results", "1",
			"--max-subagents", "1",
			"--output-format", "plain_text",
			"--search-depth", "lite",
			"--start-date", "start_date",
			"--time-range", "hour",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: x\n" +
			"content_type:\n" +
			"  - string\n" +
			"country: country\n" +
			"deep_search: true\n" +
			"end_date: end_date\n" +
			"exclude_domains:\n" +
			"  - string\n" +
			"focus: string\n" +
			"include_answer: true\n" +
			"include_domains:\n" +
			"  - string\n" +
			"locale: locale\n" +
			"max_results: 1\n" +
			"max_subagents: 1\n" +
			"output_format: plain_text\n" +
			"search_depth: lite\n" +
			"start_date: start_date\n" +
			"time_range: hour\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"search",
		)
	})
}
