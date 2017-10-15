package main

import (
	"net/http/httptest"
	"os"
	"testing"

	"goji.io"
	"goji.io/pat"
)

var testGreetingController *GreetingController

func TestMain(m *testing.M) {
	host := os.Getenv("MONGODB_HOST")
	if host != "" {
		MainLogger.Printf("Connected to database for test.")
		testGreetingController = NewGreetingController(MongoSessionWrapper{GetSession()})
	} else {
		MainLogger.Printf("Mocking database, no local database found: missing MONGODB_HOST in env")
		testGreetingController = NewGreetingController(nil)
	}

	// Setup web server
	mux := goji.NewMux()
	// Routes for controllers
	testGreetingController.AddHandlers(mux)
	// Routes for query_helper_test
	mux.HandleFunc(pat.Get("/test/extract/simple"), HandleSimpleExtract)
	mux.HandleFunc(pat.Get("/test/extract/regex"), HandleRegExExtract)
	server := httptest.NewServer(mux)
	defer server.Close()
	GreetingServerURL = server.URL
	MainLogger.Printf("Creating server: %s", GreetingServerURL)
	// Run tests
	i := m.Run()
	os.Exit(i)
}
