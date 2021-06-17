// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mock is a generated GoMock package.
package mock

import (
	manager "github.com/consensysquorum/quorum-key-manager/src/manifests/manager"
	manifest "github.com/consensysquorum/quorum-key-manager/src/manifests/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Subscribe mock base method
func (m *MockManager) Subscribe(kinds []manifest.Kind, messages chan<- []manager.Message) manager.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", kinds, messages)
	ret0, _ := ret[0].(manager.Subscription)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockManagerMockRecorder) Subscribe(kinds, messages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockManager)(nil).Subscribe), kinds, messages)
}

// MockSubscription is a mock of Subscription interface
type MockSubscription struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionMockRecorder
}

// MockSubscriptionMockRecorder is the mock recorder for MockSubscription
type MockSubscriptionMockRecorder struct {
	mock *MockSubscription
}

// NewMockSubscription creates a new mock instance
func NewMockSubscription(ctrl *gomock.Controller) *MockSubscription {
	mock := &MockSubscription{ctrl: ctrl}
	mock.recorder = &MockSubscriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubscription) EXPECT() *MockSubscriptionMockRecorder {
	return m.recorder
}

// Unsubscribe mock base method
func (m *MockSubscription) Unsubscribe() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe")
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockSubscriptionMockRecorder) Unsubscribe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockSubscription)(nil).Unsubscribe))
}

// Error mock base method
func (m *MockSubscription) Error() <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// Error indicates an expected call of Error
func (mr *MockSubscriptionMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockSubscription)(nil).Error))
}
