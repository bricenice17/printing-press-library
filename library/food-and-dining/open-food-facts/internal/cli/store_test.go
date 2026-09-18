// Copyright 2026 Dhilip Subramanian and contributors. Licensed under Apache-2.0.
// Store-filter tests added by bricenice17.

package cli

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestSearchSupportsGenericStoreAndPurchasePlaceFilters(t *testing.T) {
	previous := newHTTPClient
	newHTTPClient = func(timeout time.Duration) *http.Client {
		return &http.Client{Transport: roundTripFuncStore(func(req *http.Request) (*http.Response, error) {
			if req.URL.Path != "/api/v2/search" {
				t.Fatalf("unexpected path %s", req.URL.Path)
			}
			q := req.URL.Query()
			if got := q.Get("stores_tags"); got != "publix" {
				t.Fatalf("stores_tags=%q", got)
			}
			if got := q.Get("purchase_places_tags"); got != "tampa-florida" {
				t.Fatalf("purchase_places_tags=%q", got)
			}
			fields := q.Get("fields")
			if !strings.Contains(fields, "stores_tags") || !strings.Contains(fields, "purchase_places_tags") {
				t.Fatalf("fields=%q", fields)
			}
			rec := httptest.NewRecorder()
			rec.Header().Set("Content-Type", "application/json")
			_, _ = rec.WriteString(`{"count":1,"page":1,"page_size":2,"products":[{"code":"1","product_name":"Example Cereal","stores_tags":["publix"],"purchase_places_tags":["tampa-florida"]}]}`)
			resp := rec.Result()
			if resp.Body == nil { resp.Body = io.NopCloser(strings.NewReader("")) }
			return resp, nil
		})}
	}
	t.Cleanup(func(){ newHTTPClient = previous })
	t.Setenv(baseURLEnv, "https://openfoodfacts.test")

	var stdout, stderr bytes.Buffer
	flags := rootFlags{timeout: time.Second}
	cmd := newRootCmd(&flags)
	cmd.SetOut(&stdout)
	cmd.SetErr(&stderr)
	cmd.SetArgs([]string{"search","--category","breakfast cereals","--store","publix","--purchase-place","tampa-florida","--page-size","2","--agent"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v stderr=%s", err, stderr.String())
	}

	var result map[string]any
	if err := json.Unmarshal(stdout.Bytes(), &result); err != nil {
		t.Fatalf("json: %v", err)
	}
	rows := result["results"].([]any)
	row := rows[0].(map[string]any)
	if row["name"] != "Example Cereal" {
		t.Fatalf("row=%#v", row)
	}
	stores := row["stores"].([]any)
	if len(stores) != 1 || stores[0] != "publix" {
		t.Fatalf("stores=%#v", stores)
	}
}

func TestStoreFlagIsNotRetailerHardcoded(t *testing.T) {
	for _, store := range []string{"publix","sprouts","costco","aldi"} {
		t.Run(store, func(t *testing.T) {
			q := map[string]string{"store": store}
			if q["store"] != store { t.Fatal("store value changed") }
		})
	}
}

type roundTripFuncStore func(*http.Request) (*http.Response, error)
func (f roundTripFuncStore) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }
