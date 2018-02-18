// Code generated by MockGen. DO NOT EDIT.
// Source: httpwrapper.go

// Package http_wrapper is a generated GoMock package.
package http_wrapper

import (
	reflect "reflect"

	log "github.com/FenixAra/grocery-app/utils/log"
	gomock "github.com/golang/mock/gomock"
)

// MockIHTTPWrapper is a mock of IHTTPWrapper interface
type MockIHTTPWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockIHTTPWrapperMockRecorder
}

// MockIHTTPWrapperMockRecorder is the mock recorder for MockIHTTPWrapper
type MockIHTTPWrapperMockRecorder struct {
	mock *MockIHTTPWrapper
}

// NewMockIHTTPWrapper creates a new mock instance
func NewMockIHTTPWrapper(ctrl *gomock.Controller) *MockIHTTPWrapper {
	mock := &MockIHTTPWrapper{ctrl: ctrl}
	mock.recorder = &MockIHTTPWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIHTTPWrapper) EXPECT() *MockIHTTPWrapperMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockIHTTPWrapper) Init(l *log.Logger) {
	m.ctrl.Call(m, "Init", l)
}

// Init indicates an expected call of Init
func (mr *MockIHTTPWrapperMockRecorder) Init(l interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockIHTTPWrapper)(nil).Init), l)
}

// MakeRequest mocks base method
func (m *MockIHTTPWrapper) MakeRequest(method string, pType int, u string, payload interface{}, auth, pass string, v interface{}) (int, error) {
	ret := m.ctrl.Call(m, "MakeRequest", method, pType, u, payload, auth, pass, v)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeRequest indicates an expected call of MakeRequest
func (mr *MockIHTTPWrapperMockRecorder) MakeRequest(method, pType, u, payload, auth, pass, v interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRequest", reflect.TypeOf((*MockIHTTPWrapper)(nil).MakeRequest), method, pType, u, payload, auth, pass, v)
}
