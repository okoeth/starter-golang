package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"goji.io"
	"goji.io/pat"
	mgo "gopkg.in/mgo.v2"
)

// MainLogger is the logger for the app
var MainLogger = log.New(os.Stdout, "STARTER: ", log.Lshortfile|log.LstdFlags)

func init() {
	flag.Parse()
}

func main() {
	// Set-up routes
	mux := goji.NewMux()
	AddGreetingController(getSession(), mux)
	mux.Handle(pat.Get("/html/*"), http.FileServer(http.Dir("./")))
	MainLogger.Printf("System is ready.\n")
	http.ListenAndServe("0.0.0.0:8000", mux)
}

func getSession() *mgo.Session {
	host := os.Getenv("MONGODB_HOST")
	if host == "" {
		MainLogger.Panicln("Missing mongo hostname in env. Try: export MONGODB_HOST=<hostname>")
		return nil
	}
	MainLogger.Printf("Connecting to mongo on %s.\n", host)
	s, err := mgo.Dial("mongodb://user:password@" + host + ":27017/starterdb")
	if err != nil {
		MainLogger.Panicf("Cannot connect to MongoDB on host %s: %s\n", host, err)
		return nil
	}
	MainLogger.Printf("Connection to mongo established.\n")
	return s
}
