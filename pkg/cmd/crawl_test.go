// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/nimbleway-cli/internal/mocktest"
	"github.com/stainless-sdks/nimbleway-cli/internal/requestflag"
)

func TestCrawlList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "list",
		"--status", "pending",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestCrawlRoot(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "root",
		"--url", "https://example.com",
		"--allow-external-links=false",
		"--allow-subdomains=false",
		"--callback", "{url: https://example.com/webhook, events: [page], headers: {X-Custom-Header: bar}, metadata: {crawlId: bar}}",
		"--crawl-entire-domain=false",
		"--exclude-path", "/exclude-this-path",
		"--exclude-path", "/and-this-path",
		"--extract-options", "{debug_options: {collect_har: true, no_retry_mode: true, record_screen: true, redact: true, show_cursor: true, solve_captcha: true, trace: true, upload_engine_logs: true, verbose: true, with_proxy_usage: true}, url: https://example.com/page, browser: chrome, city: Los Angeles, client_timeout: 25000, consent_header: true, cookies: [{creation: creation, domain: domain, expires: expires, extensions: [string], hostOnly: true, httpOnly: true, lastAccessed: lastAccessed, maxAge: Infinity, name: name, path: path, pathIsDefault: true, sameSite: strict, secure: true, value: value}], country: US, device: desktop, disable_ip_check: false, driver: vx8, dynamic_parser: {myParser: bar}, expected_status_codes: [200, 201], export_userbrowser: false, format: json, headers: {User-Agent: CustomBot/1.0, Accept-Language: en-US}, http2: true, ip6: false, is_xhr: true, locale: en-US, markdown: false, metadata: {account_name: acme-corp, definition_id: 456, definition_name: product-scraper, endpoint: /api/v2/scrape, execution_id: exec-abc123, flowit_task_id: task-xyz789, input_id: input-123, pipeline_execution_id: 12345, query_template_id: template-qry-001, source: web-app, template_id: 789, template_name: e-commerce-template}, method: GET, native_mode: requester, network_capture: [{method: GET, resource_type: document, status_code: 100, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}], no_html: false, no_userbrowser: false, os: windows, parse: true, parse_options: {merge_dynamic: true}, parser: {myParser: bar}, proxy_provider: brightdata, proxy_providers: {brightdata: 70, oxylabs: 30}, query_template: {id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, api_type: WEB, pagination: {next_page_params: {foo: bar}}, params: {foo: bar}}, raw_headers: true, referrer_type: random, render: true, render_flow: [{wait: bar}, {click: bar}], render_options: {adblock: true, blocked_domains: [ads.example.com, tracker.com], browser_engine: chrome, cache: false, connector_type: webit-cdp, disabled_resources: [image, stylesheet], enable_2captcha: true, extensions: [extension-id-1, extension-id-2], fingerprint_id: fp-abc123, hackium_configuration: {collect_logs: true, do_not_fix_math_salt: true, enable_document_element_spoof: true, enable_document_has_focus: true, enable_fake_navigation_history: true, enable_key_ordering: true, enable_sniffer: true, enable_verbose_logs: true}, headless: true, include_iframes: true, load_local_storage: true, local_storage_keys_to_load: [authToken, userId], mouse_strategy: linear, no_accept_encoding: true, override_permissions: true, random_header_order: true, render_type: load, store_local_storage: true, timeout: 30000, typing_interval: 100, typing_strategy: simple, userbrowser: true, wait_until: networkidle2, with_performance_metrics: true}, request_timeout: 30000, return_response_headers_as_header: true, save_userbrowser: false, session: {id: id, prefetch_userbrowser: true, retry: true, timeout: 1}, skill: dynamic-content, skip_ubct: false, state: CA, tag: campaign-2024-q1, template: {name: x, params: {foo: bar}}, type: generic, userbrowser_creation_template_rendered: {id: id, allowed_parameter_names: [x], render_flow_rendered: [{foo: bar}]}}",
		"--ignore-query-parameters=false",
		"--include-path", "/include-this-path",
		"--include-path", "/and-this-path",
		"--limit", "100",
		"--max-discovery-depth", "3",
		"--name", "The best crawl ever",
		"--sitemap", "include",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(crawlRoot)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "root",
		"--url", "https://example.com",
		"--allow-external-links=false",
		"--allow-subdomains=false",
		"--callback", "{url: https://example.com/webhook, events: [page], headers: {X-Custom-Header: bar}, metadata: {crawlId: bar}}",
		"--crawl-entire-domain=false",
		"--exclude-path", "/exclude-this-path",
		"--exclude-path", "/and-this-path",
		"--extract-options.debug-options", "{collect_har: true, no_retry_mode: true, record_screen: true, redact: true, show_cursor: true, solve_captcha: true, trace: true, upload_engine_logs: true, verbose: true, with_proxy_usage: true}",
		"--extract-options.url", "https://example.com/page",
		"--extract-options.browser", "chrome",
		"--extract-options.city", "Los Angeles",
		"--extract-options.client-timeout", "25000",
		"--extract-options.consent-header=true",
		"--extract-options.cookies", "[{creation: creation, domain: domain, expires: expires, extensions: [string], hostOnly: true, httpOnly: true, lastAccessed: lastAccessed, maxAge: Infinity, name: name, path: path, pathIsDefault: true, sameSite: strict, secure: true, value: value}]",
		"--extract-options.country", "US",
		"--extract-options.device", "desktop",
		"--extract-options.disable-ip-check=false",
		"--extract-options.driver", "vx8",
		"--extract-options.dynamic-parser", "{myParser: bar}",
		"--extract-options.expected-status-codes", "[200, 201]",
		"--extract-options.export-userbrowser=false",
		"--extract-options.format", "json",
		"--extract-options.headers", "{User-Agent: CustomBot/1.0, Accept-Language: en-US}",
		"--extract-options.http2=true",
		"--extract-options.ip6=false",
		"--extract-options.is-xhr=true",
		"--extract-options.locale", "en-US",
		"--extract-options.markdown=false",
		"--extract-options.metadata", "{account_name: acme-corp, definition_id: 456, definition_name: product-scraper, endpoint: /api/v2/scrape, execution_id: exec-abc123, flowit_task_id: task-xyz789, input_id: input-123, pipeline_execution_id: 12345, query_template_id: template-qry-001, source: web-app, template_id: 789, template_name: e-commerce-template}",
		"--extract-options.method", "GET",
		"--extract-options.native-mode", "requester",
		"--extract-options.network-capture", "[{method: GET, resource_type: document, status_code: 100, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}]",
		"--extract-options.no-html=false",
		"--extract-options.no-userbrowser=false",
		"--extract-options.os", "windows",
		"--extract-options.parse=true",
		"--extract-options.parse-options", "{merge_dynamic: true}",
		"--extract-options.parser", "{myParser: bar}",
		"--extract-options.proxy-provider", "brightdata",
		"--extract-options.proxy-providers", "{brightdata: 70, oxylabs: 30}",
		"--extract-options.query-template", "{id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, api_type: WEB, pagination: {next_page_params: {foo: bar}}, params: {foo: bar}}",
		"--extract-options.raw-headers=true",
		"--extract-options.referrer-type", "random",
		"--extract-options.render=true",
		"--extract-options.render-flow", "[{wait: bar}, {click: bar}]",
		"--extract-options.render-options", "{adblock: true, blocked_domains: [ads.example.com, tracker.com], browser_engine: chrome, cache: false, connector_type: webit-cdp, disabled_resources: [image, stylesheet], enable_2captcha: true, extensions: [extension-id-1, extension-id-2], fingerprint_id: fp-abc123, hackium_configuration: {collect_logs: true, do_not_fix_math_salt: true, enable_document_element_spoof: true, enable_document_has_focus: true, enable_fake_navigation_history: true, enable_key_ordering: true, enable_sniffer: true, enable_verbose_logs: true}, headless: true, include_iframes: true, load_local_storage: true, local_storage_keys_to_load: [authToken, userId], mouse_strategy: linear, no_accept_encoding: true, override_permissions: true, random_header_order: true, render_type: load, store_local_storage: true, timeout: 30000, typing_interval: 100, typing_strategy: simple, userbrowser: true, wait_until: networkidle2, with_performance_metrics: true}",
		"--extract-options.request-timeout", "30000",
		"--extract-options.return-response-headers-as-header=true",
		"--extract-options.save-userbrowser=false",
		"--extract-options.session", "{id: id, prefetch_userbrowser: true, retry: true, timeout: 1}",
		"--extract-options.skill", "dynamic-content",
		"--extract-options.skip-ubct=false",
		"--extract-options.state", "CA",
		"--extract-options.tag", "campaign-2024-q1",
		"--extract-options.template", "{name: x, params: {foo: bar}}",
		"--extract-options.type", "generic",
		"--extract-options.userbrowser-creation-template-rendered", "{id: id, allowed_parameter_names: [x], render_flow_rendered: [{foo: bar}]}",
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
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "status",
		"--id", "123e4567-e89b-12d3-a456-426614174000",
	)
}

func TestCrawlTerminate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"crawl", "terminate",
		"--id", "123e4567-e89b-12d3-a456-426614174000",
	)
}
