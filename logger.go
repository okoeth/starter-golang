package main

import "net/http"
import "fmt"

// Logger is the actual implementation of the middleware
func Logger(handler http.HandlerFunc) http.HandlerFunc {
	logger := func(w http.ResponseWriter, r *http.Request) {
		MainLogger.Printf("Request: %s/%d.%d %s %s", r.Proto, r.ProtoMajor, r.ProtoMinor, r.Method, r.URL.Path)
		MainLogger.Printf("  Header: %s", logHeader(r.Header))
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(logger)
}

// logHeader prints the header
func logHeader(header http.Header) string {
	hs := fmt.Sprintf("{ ")
	for name, values := range header {
		for _, value := range values {
			hs = fmt.Sprintf("%s %s: %s, ", hs, name, value)
		}
	}
	return fmt.Sprintf("%s }", hs)
}
