package main

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

var theID string

//////////////////////////////////////////////////////////////////////////
// Helper functions

func createTestGreeting(data string) (*GreetingModel, error) {
	gm := &GreetingModel{
		Title:   data,
		Message: data,
	}
	return ClientCreateGreeting(gm)
}

//////////////////////////////////////////////////////////////////////////
// Test function

func TestCreateGreeting(t *testing.T) {
	t.Logf("Start TestCreateGreeting with ID: %s\n", theID)
	gm, err := createTestGreeting("test")
	if err != nil {
		t.Fatal(err)
	}
	theID = gm.ID.Hex()
	t.Logf("Created greeting %s", theID)
}

func TestGetGreeting(t *testing.T) {
	t.Logf("Start TestGetGreeting with ID: %s\n", theID)
	gm, err := ClientGetGreeting(theID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got greeting %s", gm.ID.Hex())
}

func TestUpdateGreeting(t *testing.T) {
	t.Logf("Start TestUpdateGreeting with ID: %s\n", theID)
	// First read post
	gm, err := ClientGetGreeting(theID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got greeting with title: %s", gm.Title)

	// Then update post with new title
	gm.Title = "Updated title"
	err = ClientUpdateGreeting(gm)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Updated greeting with ID: %s", gm.ID.Hex())

	// Finally, read post once more and check title
	gm2, err := ClientGetGreeting(theID)
	if err != nil {
		t.Fatal(err)
	}
	if gm.Title != gm2.Title {
		t.Errorf("Expected ttile %s but got %s", gm.Title, gm2.Title)
	}
	t.Logf("Got greeting with title: %s", gm2.Title)
}

func TestDeleteGreeting(t *testing.T) {
	t.Logf("Start TestDeleteGreeting with ID: %s\n", theID)
	gm := &GreetingModel{
		ID: bson.ObjectIdHex(theID),
	}
	err := ClientDeleteGreeting(gm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetGreetingByTitle(t *testing.T) {
	gm1, err := createTestGreeting("alpha-1")
	if err != nil {
		t.Fatal(err)
	}
	gm2, err := createTestGreeting("alpha-2")
	if err != nil {
		t.Fatal(err)
	}
	gm3, err := createTestGreeting("alpha-3")
	if err != nil {
		t.Fatal(err)
	}
	// Simple query
	q1 := []QueryElement{
		QueryElement{Key: "title", Op: "eq", Val: "alpha-1"},
	}
	gms1, err := ClientGetGreetings(q1)
	if err != nil {
		t.Fatal(err)
	}
	if len(gms1) != 1 {
		t.Errorf("Expected to find one record, found: %d", len(gms1))
	}
	// Wildcard query
	q2 := []QueryElement{
		QueryElement{Key: "title", Op: "regex", Val: "alpha-*"},
	}
	gms2, err := ClientGetGreetings(q2)
	if err != nil {
		t.Fatal(err)
	}
	if len(gms2) != 3 {
		t.Errorf("Expected to find three records, found: %d", len(gms2))
	}
	err = ClientDeleteGreeting(gm1)
	if err != nil {
		t.Fatal(err)
	}
	err = ClientDeleteGreeting(gm2)
	if err != nil {
		t.Fatal(err)
	}
	err = ClientDeleteGreeting(gm3)
	if err != nil {
		t.Fatal(err)
	}
}
