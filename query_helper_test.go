package main

import (
	"net/http"
	"net/url"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

//////////////////////////////////////////////////////////////////////////
// Test Query Builder

// TODO: Refactor to array based test cases
func TestBuildSimpleQuery(t *testing.T) {
	q := []QueryElement{
		QueryElement{Key: "key0", Op: "op0", Val: "val0"},
		QueryElement{Key: "key1", Op: "op1", Val: "val1"},
	}
	s := BuildQuery("http://localhost/v1/test", q)
	u, err := url.Parse(s)
	if err != nil {
		t.Error(err)
	}
	if u.Query().Get("k0") != "key0" {
		t.Errorf("Issue with k0: %v\n", u.Query().Get("k0"))
	}
	if u.Query().Get("o0") != "op0" {
		t.Errorf("Issue with o0: %v\n", u.Query().Get("o0"))
	}
	if u.Query().Get("v0") != "val0" {
		t.Errorf("Issue with v0: %v\n", u.Query().Get("v0"))
	}
	if u.Query().Get("k1") != "key1" {
		t.Errorf("Issue with k1: %v\n", u.Query().Get("k1"))
	}
	if u.Query().Get("o1") != "op1" {
		t.Errorf("Issue with o1: %v\n", u.Query().Get("o1"))
	}
	if u.Query().Get("v1") != "val1" {
		t.Errorf("Issue with v1: %v\n", u.Query().Get("v1"))
	}
}

func TestBuildNilQuery(t *testing.T) {
	s := BuildQuery("http://localhost/v1/test", nil)
	if s != "http://localhost/v1/test" {
		t.Errorf("This query failed: %s\n", s)
	}
}

func TestBuildEmptyQuery(t *testing.T) {
	q := []QueryElement{}
	s := BuildQuery("http://localhost/v1/test", q)
	if s != "http://localhost/v1/test" {
		t.Errorf("This query failed: %s\n", s)
	}
}

func TestBuildIllegalQuery(t *testing.T) {
	s := BuildQuery("df%=df0()", nil)
	if s != "" {
		t.Errorf("This query failed: %s\n", s)
	}
}

func TestBuildEncodedQuery(t *testing.T) {
	q := []QueryElement{
		QueryElement{Key: "key/0", Op: "op?0", Val: "val:0"},
	}
	s := BuildQuery("http://localhost/v1/test", q)
	u, err := url.Parse(s)
	if err != nil {
		t.Error(err)
	}
	if u.Query().Get("k0") != "key/0" {
		t.Errorf("Issue with k0: %v\n", u.Query().Get("k0"))
	}
	if u.Query().Get("o0") != "op?0" {
		t.Errorf("Issue with o0: %v\n", u.Query().Get("o0"))
	}
	if u.Query().Get("v0") != "val:0" {
		t.Errorf("Issue with v0: %v\n", u.Query().Get("v0"))
	}
}

//////////////////////////////////////////////////////////////////////////
// Test Query Extractor

func HandleSimpleExtract(w http.ResponseWriter, r *http.Request) {
	MainLogger.Printf("Request: %s %s", r.Method, r.URL.Path)
	q := ExtractQuery(r)
	MainLogger.Printf("Query %d, %v\n", len(q), q)
	v0, ok := q["key0"]
	if !ok || v0 != "val0" {
		MainLogger.Printf("ERROR: Key 0 status: %v, value: %v\n", ok, v0)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	v1, ok := q["key1"]
	if !ok || v1 != "val1" {
		MainLogger.Printf("ERROR: Key 1 status: %v, value: %v\n", ok, v1)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func TestExtractSimpleQuery(t *testing.T) {
	MainLogger.Printf("Start TestBuildSimpleQuery\n")
	q := []QueryElement{
		QueryElement{Key: "key0", Op: "eq", Val: "val0"},
		QueryElement{Key: "key1", Op: "eq", Val: "val1"},
	}
	url := BuildQuery(GreetingServerURL+"/test/extract/simple", q)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Failing with status: %d\n", res.StatusCode)
	}
	t.Logf("Response %d", res.StatusCode)
}

func HandleRegExExtract(w http.ResponseWriter, r *http.Request) {
	q := ExtractQuery(r)
	_, ok := q["key0"].(*bson.RegEx)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func TestExtractRegExQuery(t *testing.T) {
	q := []QueryElement{
		QueryElement{Key: "key0", Op: "regex", Val: "val*"},
	}
	url := BuildQuery(GreetingServerURL+"/test/extract/regex", q)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Failing with status: %d\n", res.StatusCode)
	}
	t.Logf("Response %d", res.StatusCode)
}
