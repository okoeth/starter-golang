package main

import (
	"encoding/json"
	"net/http"
	"time"

	"goji.io"

	"goji.io/pat"
	"gopkg.in/mgo.v2/bson"
)

type (
	// GreetingController represents the controller for working with this app
	GreetingController struct {
		Session SessionWrapper
	}
)

// NewGreetingController provides a reference to an EventController with
func NewGreetingController(session SessionWrapper) *GreetingController {
	return &GreetingController{session}
}

// AddHandlers inserts new greeting
func (gc *GreetingController) AddHandlers(mux *goji.Mux) {
	mux.HandleFunc(pat.Post("/v1/greetings"), gc.CreateGreeting)
	mux.HandleFunc(pat.Get("/v1/greetings"), gc.GetGreetings)
	mux.HandleFunc(pat.Get("/v1/greetings/:id"), gc.GetGreeting)
	mux.HandleFunc(pat.Delete("/v1/greetings/:id"), gc.DeleteGreeting)
	mux.HandleFunc(pat.Put("/v1/greetings/:id"), gc.UpdateGreeting)
}

// CreateGreeting inserts new greeting
func (gc *GreetingController) CreateGreeting(w http.ResponseWriter, r *http.Request) {
	MainLogger.Printf("Request: %s %s", r.Method, r.URL.Path)
	var gm GreetingModel
	err := json.NewDecoder(r.Body).Decode(&gm)
	if err != nil {
		MainLogger.Printf("Error decoding body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gm.ID = bson.NewObjectId()
	gm.CreatedAt = time.Now()
	gm.CreatedBy = "gopher"
	gm.UpdatedAt = gm.CreatedAt
	gm.UpdatedBy = gm.CreatedBy
	gc.Session.DB("starterdb").C("greetings").Insert(gm)
	gmj, err := json.Marshal(gm)
	if err != nil {
		MainLogger.Println("Error marshaling into JSON")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(gmj)
}

// GetGreetings retrieves all greetings
func (gc *GreetingController) GetGreetings(w http.ResponseWriter, r *http.Request) {
	MainLogger.Printf("Request: %s %s", r.Method, r.URL.Path)
	var gms []GreetingModel
	err := gc.Session.DB("starterdb").C("greetings").Find(ExtractQuery(r)).All(&gms)
	if err != nil {
		MainLogger.Println("Error reading from db")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	gmsj, err := json.Marshal(gms)
	if err != nil {
		MainLogger.Println("Error marshaling into JSON")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(gmsj)
}

// GetGreeting retrieves specific greeting
func (gc *GreetingController) GetGreeting(w http.ResponseWriter, r *http.Request) {
	MainLogger.Printf("Request: %s %s", r.Method, r.URL.Path)
	id := pat.Param(r, "id")
	if !bson.IsObjectIdHex(id) {
		MainLogger.Println("Invalid id")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	idh := bson.ObjectIdHex(id)
	gm := GreetingModel{}
	err := gc.Session.DB("starterdb").C("greetings").FindId(idh).One(&gm)
	if err != nil {
		MainLogger.Println("Unknown id")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	gmj, err := json.Marshal(gm)
	if err != nil {
		MainLogger.Println("Error marshaling into JSON")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(gmj)
}

// UpdateGreeting upwrite new data to existing greeting
func (gc *GreetingController) UpdateGreeting(w http.ResponseWriter, r *http.Request) {
	MainLogger.Printf("Request: %s %s", r.Method, r.URL.Path)
	id := pat.Param(r, "id")
	if !bson.IsObjectIdHex(id) {
		MainLogger.Println("Invalid id")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	idh := bson.ObjectIdHex(id)
	gmdb := GreetingModel{}
	err := gc.Session.DB("starterdb").C("greetings").FindId(idh).One(&gmdb)
	if err != nil {
		MainLogger.Println("Unknown id")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	gmreq := GreetingModel{}
	err = json.NewDecoder(r.Body).Decode(&gmreq)
	if err != nil {
		MainLogger.Println("Error decoding body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gmdb.UpdatedAt = time.Now()
	gmdb.UpdatedBy = "gopher"
	gmdb.Clone(gmreq)
	err = gc.Session.DB("starterdb").C("greetings").UpdateId(idh, gmdb)
	if err != nil {
		MainLogger.Printf("Error updating database: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteGreeting removes existing greeting
func (gc *GreetingController) DeleteGreeting(w http.ResponseWriter, r *http.Request) {
	MainLogger.Printf("Request: %s %s", r.Method, r.URL.Path)
	id := pat.Param(r, "id")
	if !bson.IsObjectIdHex(id) {
		MainLogger.Println("Invalid id")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	idh := bson.ObjectIdHex(id)
	err := gc.Session.DB("starterdb").C("greetings").RemoveId(idh)
	if err != nil {
		MainLogger.Println("Unknown id")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
