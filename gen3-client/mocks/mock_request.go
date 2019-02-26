// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/mingfeishao/go/src/github.com/uc-cdis/gen3-client/gen3-client/jwt/functions.go

// Package mocks is a generated GoMock package.
package mocks

import (
	bytes "bytes"
	http "net/http"
	url "net/url"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	jwt "github.com/uc-cdis/gen3-client/gen3-client/jwt"
)

// MockRequestInterface is a mock of RequestInterface interface
type MockRequestInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRequestInterfaceMockRecorder
}

// MockRequestInterfaceMockRecorder is the mock recorder for MockRequestInterface
type MockRequestInterfaceMockRecorder struct {
	mock *MockRequestInterface
}

// NewMockRequestInterface creates a new mock instance
func NewMockRequestInterface(ctrl *gomock.Controller) *MockRequestInterface {
	mock := &MockRequestInterface{ctrl: ctrl}
	mock.recorder = &MockRequestInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRequestInterface) EXPECT() *MockRequestInterfaceMockRecorder {
	return m.recorder
}

// MakeARequest mocks base method
func (m *MockRequestInterface) MakeARequest(arg0 *http.Client, arg1, arg2 string, arg3 map[string]string, arg4 *bytes.Buffer) (*http.Response, error) {
	ret := m.ctrl.Call(m, "MakeARequest", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeARequest indicates an expected call of MakeARequest
func (mr *MockRequestInterfaceMockRecorder) MakeARequest(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeARequest", reflect.TypeOf((*MockRequestInterface)(nil).MakeARequest), arg0, arg1, arg2, arg3, arg4)
}

// RequestNewAccessKey mocks base method
func (m *MockRequestInterface) RequestNewAccessKey(arg0 string, arg1 *jwt.Credential) {
	m.ctrl.Call(m, "RequestNewAccessKey", arg0, arg1)
}

// RequestNewAccessKey indicates an expected call of RequestNewAccessKey
func (mr *MockRequestInterfaceMockRecorder) RequestNewAccessKey(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestNewAccessKey", reflect.TypeOf((*MockRequestInterface)(nil).RequestNewAccessKey), arg0, arg1)
}

// GetPresignedURL mocks base method
func (m *MockRequestInterface) GetPresignedURL(method string, host *url.URL, endpointPostPrefix, accessKey, contentType string, body *bytes.Buffer) *http.Response {
	ret := m.ctrl.Call(m, "GetPresignedURL", method, host, endpointPostPrefix, accessKey, contentType, body)
	ret0, _ := ret[0].(*http.Response)
	return ret0
}

// GetPresignedURL indicates an expected call of GetPresignedURL
func (mr *MockRequestInterfaceMockRecorder) GetPresignedURL(method, host, endpointPostPrefix, accessKey, contentType, body interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPresignedURL", reflect.TypeOf((*MockRequestInterface)(nil).GetPresignedURL), method, host, endpointPostPrefix, accessKey, contentType, body)
}