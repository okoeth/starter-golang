package main

import (
	"net/http"
	"net/url"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

// QueryElement holds key / operand / value of a query
type QueryElement struct {
	Key string
	Op  string
	Val string
}

// BuildQuery creates a URL query string from an array of QueryElements
func BuildQuery(b string, q []QueryElement) string {
	var u *url.URL
	u, err := url.Parse(b)
	if err != nil {
		return ""
	}
	if q == nil || len(q) == 0 {
		return u.String()
	}
	p := url.Values{}
	for i, v := range q {
		MainLogger.Printf("Adding key %v with op %v and val %v", v.Key, v.Op, v.Val)
		p.Add("k"+strconv.Itoa(i), v.Key)
		p.Add("o"+strconv.Itoa(i), v.Op)
		p.Add("v"+strconv.Itoa(i), v.Val)
	}
	u.RawQuery = p.Encode()
	return u.String()
}

// ExtractQuery creates a mongo query from URL query params
func ExtractQuery(r *http.Request) bson.M {
	q := make(bson.M)
	for i := 0; i < 10; i++ {
		key := r.URL.Query().Get("k" + strconv.Itoa(i))
		op := r.URL.Query().Get("o" + strconv.Itoa(i))
		val := r.URL.Query().Get("v" + strconv.Itoa(i))
		if key == "" || op == "" || val == "" {
			break
		}
		MainLogger.Printf("Extracting key %v with op %v and val %v", key, op, val)
		if op == "eq" {
			q[key] = val
		} else if op == "regex" {
			sq := new(bson.RegEx)
			sq.Pattern = val
			sq.Options = ""
			q[key] = sq
		} else if op == "lt" || op == "gt" {
		}
	}
	if len(q) == 0 {
		return nil
	}
	return q
}
