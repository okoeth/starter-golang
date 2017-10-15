package main

import mgo "gopkg.in/mgo.v2"

//go:generate mockgen -source database_wrapper.go -destination=database_wrapper_mock.go -package=main

//////////////////////////////////////////////////////////////////////////
// Database Wrapper Interfaces

// SessionWrapper is an abstraction for mgo session
type SessionWrapper interface {
	DB(name string) DatabaseWrapper
	Close()
}

// DatabaseWrapper is an abstraction for mgo database
type DatabaseWrapper interface {
	C(name string) CollectionWrapper
}

// CollectionWrapper is an abstraction for mgo collection
type CollectionWrapper interface {
	Find(query interface{}) QueryWrapper
	FindId(id interface{}) QueryWrapper
	Count() (n int, err error)
	Insert(docs ...interface{}) error
	Remove(selector interface{}) error
	RemoveId(id interface{}) error
	Update(selector interface{}, update interface{}) error
	UpdateId(id interface{}, update interface{}) error
}

// QueryWrapper is an abstraction for mgo collection
type QueryWrapper interface {
	All(result interface{}) error
	One(result interface{}) error
}

//////////////////////////////////////////////////////////////////////////
// Database Default Implementation

// MongoSessionWrapper implements session interface for MongoDB
type MongoSessionWrapper struct {
	*mgo.Session
}

// DB implements session interface for MongoDB
func (s MongoSessionWrapper) DB(name string) DatabaseWrapper {
	return &MongoDatabaseWrapper{Database: s.Session.DB(name)}
}

// MongoDatabaseWrapper implements database interface for MongoDB
type MongoDatabaseWrapper struct {
	*mgo.Database
}

// C implements database interface for MongoDB
func (d MongoDatabaseWrapper) C(name string) CollectionWrapper {
	return &MongoCollectionWrapper{Collection: d.Database.C(name)}
}

// MongoCollectionWrapper implements collection interface for MongoDB
type MongoCollectionWrapper struct {
	*mgo.Collection
}

// Find implements collection interface for MongoDB
func (c MongoCollectionWrapper) Find(query interface{}) QueryWrapper {
	return &MongoQueryWrapper{Query: c.Collection.Find(query)}
}

// FindId implements collection interface for MongoDB
func (c MongoCollectionWrapper) FindId(id interface{}) QueryWrapper {
	return &MongoQueryWrapper{Query: c.Collection.FindId(id)}
}

// MongoQueryWrapper implements query interface for MongoDB
type MongoQueryWrapper struct {
	*mgo.Query
}
