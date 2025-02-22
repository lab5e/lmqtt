// Code generated by MockGen. DO NOT EDIT.
// Source: persistence/session/session.go

// Package session is a generated GoMock package.
package session

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/lab5e/lmqtt/pkg/entities"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockStore) Set(session *entities.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", session)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockStoreMockRecorder) Set(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockStore)(nil).Set), session)
}

// Remove mocks base method
func (m *MockStore) Remove(clientID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockStoreMockRecorder) Remove(clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockStore)(nil).Remove), clientID)
}

// Get mocks base method
func (m *MockStore) Get(clientID string) (*entities.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", clientID)
	ret0, _ := ret[0].(*entities.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockStoreMockRecorder) Get(clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), clientID)
}

// Iterate mocks base method
func (m *MockStore) Iterate(fn IterateFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterate", fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Iterate indicates an expected call of Iterate
func (mr *MockStoreMockRecorder) Iterate(fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterate", reflect.TypeOf((*MockStore)(nil).Iterate), fn)
}

// SetSessionExpiry mocks base method
func (m *MockStore) SetSessionExpiry(clientID string, expiry uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSessionExpiry", clientID, expiry)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSessionExpiry indicates an expected call of SetSessionExpiry
func (mr *MockStoreMockRecorder) SetSessionExpiry(clientID, expiry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSessionExpiry", reflect.TypeOf((*MockStore)(nil).SetSessionExpiry), clientID, expiry)
}
