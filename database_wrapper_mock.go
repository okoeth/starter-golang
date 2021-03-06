// Code generated by MockGen. DO NOT EDIT.
// Source: database_wrapper.go

// Package main is a generated GoMock package.
package main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSessionWrapper is a mock of SessionWrapper interface
type MockSessionWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockSessionWrapperMockRecorder
}

// MockSessionWrapperMockRecorder is the mock recorder for MockSessionWrapper
type MockSessionWrapperMockRecorder struct {
	mock *MockSessionWrapper
}

// NewMockSessionWrapper creates a new mock instance
func NewMockSessionWrapper(ctrl *gomock.Controller) *MockSessionWrapper {
	mock := &MockSessionWrapper{ctrl: ctrl}
	mock.recorder = &MockSessionWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionWrapper) EXPECT() *MockSessionWrapperMockRecorder {
	return m.recorder
}

// DB mocks base method
func (m *MockSessionWrapper) DB(name string) DatabaseWrapper {
	ret := m.ctrl.Call(m, "DB", name)
	ret0, _ := ret[0].(DatabaseWrapper)
	return ret0
}

// DB indicates an expected call of DB
func (mr *MockSessionWrapperMockRecorder) DB(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockSessionWrapper)(nil).DB), name)
}

// Close mocks base method
func (m *MockSessionWrapper) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockSessionWrapperMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSessionWrapper)(nil).Close))
}

// MockDatabaseWrapper is a mock of DatabaseWrapper interface
type MockDatabaseWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseWrapperMockRecorder
}

// MockDatabaseWrapperMockRecorder is the mock recorder for MockDatabaseWrapper
type MockDatabaseWrapperMockRecorder struct {
	mock *MockDatabaseWrapper
}

// NewMockDatabaseWrapper creates a new mock instance
func NewMockDatabaseWrapper(ctrl *gomock.Controller) *MockDatabaseWrapper {
	mock := &MockDatabaseWrapper{ctrl: ctrl}
	mock.recorder = &MockDatabaseWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseWrapper) EXPECT() *MockDatabaseWrapperMockRecorder {
	return m.recorder
}

// C mocks base method
func (m *MockDatabaseWrapper) C(name string) CollectionWrapper {
	ret := m.ctrl.Call(m, "C", name)
	ret0, _ := ret[0].(CollectionWrapper)
	return ret0
}

// C indicates an expected call of C
func (mr *MockDatabaseWrapperMockRecorder) C(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "C", reflect.TypeOf((*MockDatabaseWrapper)(nil).C), name)
}

// MockCollectionWrapper is a mock of CollectionWrapper interface
type MockCollectionWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionWrapperMockRecorder
}

// MockCollectionWrapperMockRecorder is the mock recorder for MockCollectionWrapper
type MockCollectionWrapperMockRecorder struct {
	mock *MockCollectionWrapper
}

// NewMockCollectionWrapper creates a new mock instance
func NewMockCollectionWrapper(ctrl *gomock.Controller) *MockCollectionWrapper {
	mock := &MockCollectionWrapper{ctrl: ctrl}
	mock.recorder = &MockCollectionWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollectionWrapper) EXPECT() *MockCollectionWrapperMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockCollectionWrapper) Find(query interface{}) QueryWrapper {
	ret := m.ctrl.Call(m, "Find", query)
	ret0, _ := ret[0].(QueryWrapper)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockCollectionWrapperMockRecorder) Find(query interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCollectionWrapper)(nil).Find), query)
}

// FindId mocks base method
func (m *MockCollectionWrapper) FindId(id interface{}) QueryWrapper {
	ret := m.ctrl.Call(m, "FindId", id)
	ret0, _ := ret[0].(QueryWrapper)
	return ret0
}

// FindId indicates an expected call of FindId
func (mr *MockCollectionWrapperMockRecorder) FindId(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindId", reflect.TypeOf((*MockCollectionWrapper)(nil).FindId), id)
}

// Count mocks base method
func (m *MockCollectionWrapper) Count() (int, error) {
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockCollectionWrapperMockRecorder) Count() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockCollectionWrapper)(nil).Count))
}

// Insert mocks base method
func (m *MockCollectionWrapper) Insert(docs ...interface{}) error {
	varargs := []interface{}{}
	for _, a := range docs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockCollectionWrapperMockRecorder) Insert(docs ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCollectionWrapper)(nil).Insert), docs...)
}

// Remove mocks base method
func (m *MockCollectionWrapper) Remove(selector interface{}) error {
	ret := m.ctrl.Call(m, "Remove", selector)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockCollectionWrapperMockRecorder) Remove(selector interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCollectionWrapper)(nil).Remove), selector)
}

// RemoveId mocks base method
func (m *MockCollectionWrapper) RemoveId(id interface{}) error {
	ret := m.ctrl.Call(m, "RemoveId", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveId indicates an expected call of RemoveId
func (mr *MockCollectionWrapperMockRecorder) RemoveId(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveId", reflect.TypeOf((*MockCollectionWrapper)(nil).RemoveId), id)
}

// Update mocks base method
func (m *MockCollectionWrapper) Update(selector, update interface{}) error {
	ret := m.ctrl.Call(m, "Update", selector, update)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockCollectionWrapperMockRecorder) Update(selector, update interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCollectionWrapper)(nil).Update), selector, update)
}

// UpdateId mocks base method
func (m *MockCollectionWrapper) UpdateId(id, update interface{}) error {
	ret := m.ctrl.Call(m, "UpdateId", id, update)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateId indicates an expected call of UpdateId
func (mr *MockCollectionWrapperMockRecorder) UpdateId(id, update interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateId", reflect.TypeOf((*MockCollectionWrapper)(nil).UpdateId), id, update)
}

// MockQueryWrapper is a mock of QueryWrapper interface
type MockQueryWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockQueryWrapperMockRecorder
}

// MockQueryWrapperMockRecorder is the mock recorder for MockQueryWrapper
type MockQueryWrapperMockRecorder struct {
	mock *MockQueryWrapper
}

// NewMockQueryWrapper creates a new mock instance
func NewMockQueryWrapper(ctrl *gomock.Controller) *MockQueryWrapper {
	mock := &MockQueryWrapper{ctrl: ctrl}
	mock.recorder = &MockQueryWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQueryWrapper) EXPECT() *MockQueryWrapperMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockQueryWrapper) All(result interface{}) error {
	ret := m.ctrl.Call(m, "All", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All
func (mr *MockQueryWrapperMockRecorder) All(result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockQueryWrapper)(nil).All), result)
}

// One mocks base method
func (m *MockQueryWrapper) One(result interface{}) error {
	ret := m.ctrl.Call(m, "One", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// One indicates an expected call of One
func (mr *MockQueryWrapperMockRecorder) One(result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockQueryWrapper)(nil).One), result)
}
