// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/telia-oss/concourse-github-lambda (interfaces: RepoClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/github"
	reflect "reflect"
)

// MockRepoClient is a mock of RepoClient interface
type MockRepoClient struct {
	ctrl     *gomock.Controller
	recorder *MockRepoClientMockRecorder
}

// MockRepoClientMockRecorder is the mock recorder for MockRepoClient
type MockRepoClientMockRecorder struct {
	mock *MockRepoClient
}

// NewMockRepoClient creates a new mock instance
func NewMockRepoClient(ctrl *gomock.Controller) *MockRepoClient {
	mock := &MockRepoClient{ctrl: ctrl}
	mock.recorder = &MockRepoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepoClient) EXPECT() *MockRepoClientMockRecorder {
	return m.recorder
}

// CreateKey mocks base method
func (m *MockRepoClient) CreateKey(arg0 context.Context, arg1, arg2 string, arg3 *github.Key) (*github.Key, *github.Response, error) {
	ret := m.ctrl.Call(m, "CreateKey", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.Key)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateKey indicates an expected call of CreateKey
func (mr *MockRepoClientMockRecorder) CreateKey(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockRepoClient)(nil).CreateKey), arg0, arg1, arg2, arg3)
}

// DeleteKey mocks base method
func (m *MockRepoClient) DeleteKey(arg0 context.Context, arg1, arg2 string, arg3 int64) (*github.Response, error) {
	ret := m.ctrl.Call(m, "DeleteKey", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteKey indicates an expected call of DeleteKey
func (mr *MockRepoClientMockRecorder) DeleteKey(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKey", reflect.TypeOf((*MockRepoClient)(nil).DeleteKey), arg0, arg1, arg2, arg3)
}

// ListKeys mocks base method
func (m *MockRepoClient) ListKeys(arg0 context.Context, arg1, arg2 string, arg3 *github.ListOptions) ([]*github.Key, *github.Response, error) {
	ret := m.ctrl.Call(m, "ListKeys", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*github.Key)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListKeys indicates an expected call of ListKeys
func (mr *MockRepoClientMockRecorder) ListKeys(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockRepoClient)(nil).ListKeys), arg0, arg1, arg2, arg3)
}
