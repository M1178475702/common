// Code generated by MockGen. DO NOT EDIT.
// Source: mock/mock.go

// Package mock_mock is a generated GoMock package.
package mock_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMock is a mock of Mock interface.
type MockMock struct {
	ctrl     *gomock.Controller
	recorder *MockMockMockRecorder
}

// MockMockMockRecorder is the mock recorder for MockMock.
type MockMockMockRecorder struct {
	mock *MockMock
}

// NewMockMock creates a new mock instance.
func NewMockMock(ctrl *gomock.Controller) *MockMock {
	mock := &MockMock{ctrl: ctrl}
	mock.recorder = &MockMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMock) EXPECT() *MockMockMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockMock) Get(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockMockMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMock)(nil).Get), key)
}
