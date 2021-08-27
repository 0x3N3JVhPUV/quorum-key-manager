// Code generated by MockGen. DO NOT EDIT.
// Source: authorizator.go

// Package mock is a generated GoMock package.
package mock

import (
	types "github.com/consensys/quorum-key-manager/src/auth/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuthorizator is a mock of Authorizator interface
type MockAuthorizator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizatorMockRecorder
}

// MockAuthorizatorMockRecorder is the mock recorder for MockAuthorizator
type MockAuthorizatorMockRecorder struct {
	mock *MockAuthorizator
}

// NewMockAuthorizator creates a new mock instance
func NewMockAuthorizator(ctrl *gomock.Controller) *MockAuthorizator {
	mock := &MockAuthorizator{ctrl: ctrl}
	mock.recorder = &MockAuthorizatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorizator) EXPECT() *MockAuthorizatorMockRecorder {
	return m.recorder
}

// CheckPermission mocks base method
func (m *MockAuthorizator) CheckPermission(ops ...*types.Operation) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range ops {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckPermission", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckPermission indicates an expected call of CheckPermission
func (mr *MockAuthorizatorMockRecorder) CheckPermission(ops ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPermission", reflect.TypeOf((*MockAuthorizator)(nil).CheckPermission), ops...)
}

// CheckAccess mocks base method
func (m *MockAuthorizator) CheckAccess(allowedTenants []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAccess", allowedTenants)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckAccess indicates an expected call of CheckAccess
func (mr *MockAuthorizatorMockRecorder) CheckAccess(allowedTenants interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccess", reflect.TypeOf((*MockAuthorizator)(nil).CheckAccess), allowedTenants)
}