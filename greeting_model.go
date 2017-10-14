package main

import "gopkg.in/mgo.v2/bson"
import "time"

type (
	// GreetingModel represents the response of the hello service
	GreetingModel struct {
		ID        bson.ObjectId `json:"id" bson:"_id"`
		Title     string        `json:"title" bson:"title"`
		Message   string        `json:"message" bson:"message"`
		CreatedAt time.Time     `json:"created_at" bson:"created_at"`
		CreatedBy string        `json:"created_by" bson:"created_by"`
		UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
		UpdatedBy string        `json:"updated_by" bson:"updated_by"`
	}
)

// Clone copies business data from other object
func (gm *GreetingModel) Clone(gmreq GreetingModel) {
	gm.Title = gmreq.Title
	gm.Message = gmreq.Message
}
