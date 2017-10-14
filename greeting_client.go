package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	createGreetingURL = "/v1/greetings"
	getGreetingURL    = "/v1/greetings/:id"
	getGreetingsURL   = "/v1/greetings"
	updateGreetingURL = "/v1/greetings/:id"
	deleteGreetingURL = "/v1/greetings/:id"
)

// GreetingServerURL holds the URL to the server
var GreetingServerURL string

// ClientCreateGreeting calls server to create a Greeting
func ClientCreateGreeting(gm *GreetingModel) (*GreetingModel, error) {
	gmj, err := json.Marshal(gm)
	if err != nil {
		MainLogger.Printf("Error marshaling into JSON: %v", err)
		return nil, err
	}
	url := GreetingServerURL + createGreetingURL
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(gmj))
	if err != nil {
		MainLogger.Printf("Error creating HTTP request: %v", err)
		return nil, err
	}
	MainLogger.Printf("Client Greeting with URL %s", url)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		MainLogger.Printf("Error sending HTTP request: %v", err)
		return nil, err
	}
	if res.StatusCode != 201 {
		MainLogger.Printf("Unexpected HTTP response code %d", res.StatusCode)
		return nil, errors.New("Response code not 201")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		MainLogger.Printf("Error reading response body: %v", err)
		return nil, err
	}
	var gmr GreetingModel
	err = json.Unmarshal(body, &gmr)
	if err != nil {
		MainLogger.Printf("Error de-marshaling model from body: %v", err)
		return nil, err
	}
	MainLogger.Printf("Created Greeting with id %v\n", gmr.ID)
	return &gmr, nil
}

// ClientGetGreetings returns Greetings from query
func ClientGetGreetings(q []QueryElement) ([]GreetingModel, error) {
	url := BuildQuery(GreetingServerURL+getGreetingsURL, q)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		MainLogger.Printf("Error creating HTTP request: %v", err)
		return nil, err
	}
	MainLogger.Printf("Client GET with URL %s", url)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		MainLogger.Printf("Error sending HTTP request: %v", err)
		return nil, err
	}
	if res.StatusCode != 200 {
		MainLogger.Printf("Unexpected HTTP response code %d", res.StatusCode)
		return nil, errors.New("Response code not 200")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		MainLogger.Printf("Error reading response body: %v", err)
		return nil, err
	}
	var gms []GreetingModel
	err = json.Unmarshal(body, &gms)
	if err != nil {
		MainLogger.Printf("Error de-marshaling model from body: %v", err)
		return nil, err
	}
	MainLogger.Printf("Retrieved %d number of greetings\n", len(gms))
	return gms, nil
}

// ClientGetGreeting returns Greeting from ID
func ClientGetGreeting(ID string) (*GreetingModel, error) {
	url := GreetingServerURL + BuildURL(getGreetingURL,
		map[string]string{":id": ID})
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		MainLogger.Printf("Error creating HTTP request: %v", err)
		return nil, err
	}
	MainLogger.Printf("Client GET with URL %s", url)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		MainLogger.Printf("Error sending HTTP request: %v", err)
		return nil, err
	}
	if res.StatusCode != 200 {
		MainLogger.Printf("Unexpected HTTP response code %d", res.StatusCode)
		return nil, errors.New("Response code not 200")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		MainLogger.Printf("Error reading response body: %v", err)
		return nil, err
	}
	var gm GreetingModel
	err = json.Unmarshal(body, &gm)
	if err != nil {
		MainLogger.Printf("Error de-marshaling model from body: %v", err)
		return nil, err
	}
	MainLogger.Printf("Retrieved Greeting with id %v\n", gm.ID)
	return &gm, nil
}

// ClientUpdateGreeting calls server to update a Greeting
func ClientUpdateGreeting(gm *GreetingModel) error {
	gmj, err := json.Marshal(gm)
	if err != nil {
		MainLogger.Printf("Error de-marshaling model from JSON: %v", err)
		return err
	}
	url := GreetingServerURL + BuildURL(updateGreetingURL,
		map[string]string{":id": gm.ID.Hex()})
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(gmj))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		MainLogger.Printf("Error creating HTTP request: %v", err)
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		MainLogger.Printf("Error sending HTTP request: %v", err)
		return err
	}
	if res.StatusCode != 200 {
		MainLogger.Printf("Unexpected HTTP response code %d", res.StatusCode)
		return errors.New("Response code not 200")
	}
	MainLogger.Printf("Updated Greeting with id %v\n", gm.ID)
	return nil
}

// ClientDeleteGreeting calls server to delete a Greeting
func ClientDeleteGreeting(gm *GreetingModel) error {
	url := GreetingServerURL + BuildURL(deleteGreetingURL,
		map[string]string{":id": gm.ID.Hex()})
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		MainLogger.Printf("Error creating HTTP request: %v", err)
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		MainLogger.Printf("Error sending HTTP request: %v", err)
		return err
	}
	if res.StatusCode != 200 {
		MainLogger.Printf("Unexpected HTTP response code %d", res.StatusCode)
		return errors.New("Response code not 200")
	}
	MainLogger.Printf("Deleted Greeting with id %v\n", gm.ID)
	return nil
}
