package main

import "strings"

// BuildURL creates URL from base URL (configured here), path with
// placeholders (provided by caller) and values for placeholders
// (provided by caller)
func BuildURL(path string, params map[string]string) string {
	if params != nil {
		for key, val := range params {
			path = strings.Replace(path, key, val, -1)
		}
	}
	return path
}
