package main

import (
	"net/http/httptest"
	"os"
	"testing"

	"goji.io"
	"goji.io/pat"
	mgo "gopkg.in/mgo.v2"
)

var theSession *mgo.Session

func getLocalSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://user:password@localhost:27017/starterdb")
	if err != nil {
		return nil
	}
	return s
}

func TestMain(m *testing.M) {
	// Setup database connection
	theSession = getLocalSession()
	if theSession == nil {
		MainLogger.Printf("Skipping test, no local database.\n")
		os.Exit(0)
	}
	MainLogger.Printf("Connected to local database.\n")
	// Setup web server
	mux := goji.NewMux()
	// Routes for greeting_controller_test
	AddGreetingController(theSession, mux)
	// Routes for query_helper_test
	mux.HandleFunc(pat.Get("/test/extract/simple"), HandleSimpleExtract)
	mux.HandleFunc(pat.Get("/test/extract/regex"), HandleRegExExtract)
	MainLogger.Printf("Created routes: %v", mux)
	server := httptest.NewServer(mux)
	GreetingServerURL = server.URL
	MainLogger.Printf("Creating server: %s", GreetingServerURL)
	// Run tests
	i := m.Run()
	defer server.Close()
	os.Exit(i)
}
