package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gopkg.in/mgo.v2/bson"
)

//////////////////////////////////////////////////////////////////////////
// Helper functions

var theID string
var session *MockSessionWrapper
var database *MockDatabaseWrapper
var collection *MockCollectionWrapper
var query *MockQueryWrapper

func mock(t *testing.T, mockFunction func(t *testing.T)) *gomock.Controller {
	if testGreetingController.Session != nil {
		// No mocking required
		t.Logf("No mcking required, skipping mock set-up")
		return nil
	}
	mc := gomock.NewController(t)
	session = NewMockSessionWrapper(mc)
	database = NewMockDatabaseWrapper(mc)
	collection = NewMockCollectionWrapper(mc)
	query = NewMockQueryWrapper(mc)
	session.EXPECT().DB("starterdb").Return(database).AnyTimes()
	database.EXPECT().C("greetings").Return(collection).AnyTimes()
	testGreetingController.Session = session
	mockFunction(t)
	return mc
}

func unmock(mc *gomock.Controller) {
	if mc != nil {
		mc.Finish()
		// Remove hooks to mocks
		testGreetingController.Session = nil
	}
}

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
// Test and mock functions

func mockCreateGreeting(t *testing.T) {
	theID = "52f6aef226f149b7048b4567"
	gomock.InOrder(
		// Step 1
		collection.EXPECT().Insert(gomock.Any()).Return(nil).Times(1),
	)
}

func TestCreateGreeting(t *testing.T) {
	mc := mock(t, mockCreateGreeting)
	defer unmock(mc)
	t.Logf("Start TestCreateGreeting with ID: %s", theID)
	// Step 1: Create greeting
	gm, err := createTestGreeting("test")
	if err != nil {
		t.Fatal(err)
	}
	theID = gm.ID.Hex()
	t.Logf("Created greeting %s", theID)
}

func mockGetGreeting(t *testing.T) {
	theID = "52f6aef226f149b7048b4567"
	model1 := GreetingModel{
		ID:      bson.ObjectIdHex(theID),
		Title:   "test",
		Message: "test",
	}
	gomock.InOrder(
		// Step 1
		collection.EXPECT().FindId(bson.ObjectIdHex(theID)).Return(query).Times(1),
		query.EXPECT().One(gomock.Any()).SetArg(0, model1).Return(nil).Times(1),
	)
}

func TestGetGreeting(t *testing.T) {
	mc := mock(t, mockGetGreeting)
	defer unmock(mc)
	t.Logf("Start TestGetGreeting with ID: %s", theID)
	// Step 1: Get greeting
	gm, err := ClientGetGreeting(theID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got greeting %s", gm.ID.Hex())
}

func mockUpdateGreeting(t *testing.T) {
	theID = "52f6aef226f149b7048b4567"
	model1 := GreetingModel{
		ID:      bson.ObjectIdHex(theID),
		Title:   "test",
		Message: "test",
	}
	model2 := GreetingModel{
		ID:      bson.ObjectIdHex(theID),
		Title:   "Updated title",
		Message: "test",
	}
	gomock.InOrder(
		// Step 1
		collection.EXPECT().FindId(bson.ObjectIdHex(theID)).Return(query).Times(1),
		query.EXPECT().One(gomock.Any()).SetArg(0, model1).Return(nil).Times(1),
		// Step 2
		collection.EXPECT().FindId(bson.ObjectIdHex(theID)).Return(query).Times(1),
		query.EXPECT().One(gomock.Any()).SetArg(0, model1).Return(nil).Times(1),
		collection.EXPECT().UpdateId(gomock.Any(), gomock.Any()).Return(nil).Times(1),
		// Step 3
		collection.EXPECT().FindId(bson.ObjectIdHex(theID)).Return(query).Times(1),
		query.EXPECT().One(gomock.Any()).SetArg(0, model2).Return(nil).Times(1),
	)
}

func TestUpdateGreeting(t *testing.T) {
	mc := mock(t, mockUpdateGreeting)
	defer unmock(mc)
	t.Logf("Start TestUpdateGreeting with ID: %s", theID)
	// Step 1: First get greeting
	gm, err := ClientGetGreeting(theID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got greeting with title: %s", gm.Title)

	// Step 2: Then update post with new title
	gm.Title = "Updated title"
	err = ClientUpdateGreeting(gm)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Updated greeting with ID: %s", gm.ID.Hex())

	// Step 3: Finally, read post once more and check title
	gm2, err := ClientGetGreeting(theID)
	if err != nil {
		t.Fatal(err)
	}
	if gm.Title != gm2.Title {
		t.Errorf("Expected title %s but got %s", gm.Title, gm2.Title)
	}
	t.Logf("Got greeting with title: %s", gm2.Title)
}

func mockDeleteGreeting(t *testing.T) {
	theID = "52f6aef226f149b7048b4567"
	gomock.InOrder(
		// Step 1
		collection.EXPECT().RemoveId(gomock.Any()).Return(nil).Times(1),
	)
}

func TestDeleteGreeting(t *testing.T) {
	mc := mock(t, mockDeleteGreeting)
	defer unmock(mc)
	t.Logf("Start TestDeleteGreeting with ID: %s", theID)
	gm := &GreetingModel{
		ID: bson.ObjectIdHex(theID),
	}
	err := ClientDeleteGreeting(gm)
	if err != nil {
		t.Fatal(err)
	}
}

func mockGetGreetingByTitle(t *testing.T) {
	models1 := []GreetingModel{
		GreetingModel{
			ID:      bson.ObjectIdHex("52f6aef226f149b7048b4567"),
			Title:   "alpha-1",
			Message: "test",
		},
	}
	models2 := []GreetingModel{
		GreetingModel{
			ID:      bson.ObjectIdHex("52f6aef226f149b7048b4567"),
			Title:   "alpha-1",
			Message: "test",
		},
		GreetingModel{
			ID:      bson.ObjectIdHex("52f6aef226f149b7048b4567"),
			Title:   "alpha-2",
			Message: "test",
		},
		GreetingModel{
			ID:      bson.ObjectIdHex("52f6aef226f149b7048b4567"),
			Title:   "alpha-3",
			Message: "test",
		},
	}
	gomock.InOrder(
		// Step 1
		collection.EXPECT().Insert(gomock.Any()).Return(nil).Times(3),
		// Step 2
		collection.EXPECT().Find(gomock.Any()).Return(query).Times(1),
		query.EXPECT().All(gomock.Any()).SetArg(0, models1).Return(nil).Times(1),
		// Step 3
		collection.EXPECT().Find(gomock.Any()).Return(query).Times(1),
		query.EXPECT().All(gomock.Any()).SetArg(0, models2).Return(nil).Times(1),
		// Step 4
		collection.EXPECT().RemoveId(gomock.Any()).Return(nil).Times(3),
	)
}

func TestGetGreetingByTitle(t *testing.T) {
	mc := mock(t, mockGetGreetingByTitle)
	defer unmock(mc)
	t.Logf("Start TestDeleteGreeting")
	// Step 1: Create three records
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
	// Step 2: Simple query
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
	// Step 3: Wildcard query
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
	// Step 4: Delete three records
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
