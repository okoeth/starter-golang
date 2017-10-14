package main

import (
	"testing"
)

type (
	// testCase represents input and output data of a single test case
	testCase struct {
		InputURL    string
		OutputURL   string
		InputParams map[string]string
	}
)

// New test cases can be added by simply creating new entries in this array
var tcs = []testCase{
	// Build nil URL
	testCase{
		"/v1/test", "/v1/test", nil},
	// Build empty URL
	testCase{
		"/v1/test", "/v1/test", map[string]string{}},
	// Build URL with unknown param
	testCase{
		"/v1/test/:id", "/v1/test/:id", map[string]string{":xx": "XX"}},
	// Build URL with too many params
	testCase{
		"/v1/test/:id1", "/v1/test/ID1", map[string]string{":id1": "ID1", ":id2": "ID2"}},
	// Build URL with too many params
	testCase{
		"/v1/test/:id1", "/v1/test/ID1", map[string]string{":id1": "ID1", ":id2": "ID2"}},
	// Build URL with too few params
	testCase{
		"/v1/test/:id1/test/:id2", "/v1/test/ID1/test/:id2", map[string]string{":id1": "ID1"}},
	// Build URL with two params
	testCase{
		"/v1/test/:id1/test/:id2", "/v1/test/ID1/test/ID2", map[string]string{":id1": "ID1", ":id2": "ID2"}},
	// Build URL with param in front
	testCase{
		":id1/v1/test", "ID1/v1/test", map[string]string{":id1": "ID1", ":id2": "ID2"}},
	// Build URL with param in middle
	testCase{
		"/v1/:id1/test", "/v1/ID1/test", map[string]string{":id1": "ID1", ":id2": "ID2"}},
	// Build URL with param in end
	testCase{
		"/v1/test/:id1", "/v1/test/ID1", map[string]string{":id1": "ID1", ":id2": "ID2"}},
}

// TestBuildURL iterates over test cases
func TestBuildURL(t *testing.T) {
	for i := range tcs {
		url := BuildURL(tcs[i].InputURL, tcs[i].InputParams)
		if url != tcs[i].OutputURL {
			t.Errorf("Build URL failed. Expected %s but returned %s\n", tcs[i].OutputURL, url)
		}
	}
}
