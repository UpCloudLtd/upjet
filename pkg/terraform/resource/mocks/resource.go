// Copyright 2021 Upbound Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/crossplane-contrib/terrajet/pkg/terraform/resource (interfaces: SecretClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockSecretClient is a mock of SecretClient interface.
type MockSecretClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretClientMockRecorder
}

// MockSecretClientMockRecorder is the mock recorder for MockSecretClient.
type MockSecretClientMockRecorder struct {
	mock *MockSecretClient
}

// NewMockSecretClient creates a new mock instance.
func NewMockSecretClient(ctrl *gomock.Controller) *MockSecretClient {
	mock := &MockSecretClient{ctrl: ctrl}
	mock.recorder = &MockSecretClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretClient) EXPECT() *MockSecretClientMockRecorder {
	return m.recorder
}

// GetSecretData mocks base method.
func (m *MockSecretClient) GetSecretData(arg0 context.Context, arg1 v1.SecretReference) (map[string][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretData", arg0, arg1)
	ret0, _ := ret[0].(map[string][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretData indicates an expected call of GetSecretData.
func (mr *MockSecretClientMockRecorder) GetSecretData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretData", reflect.TypeOf((*MockSecretClient)(nil).GetSecretData), arg0, arg1)
}

// GetSecretValue mocks base method.
func (m *MockSecretClient) GetSecretValue(arg0 context.Context, arg1 v1.SecretKeySelector) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretValue", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretValue indicates an expected call of GetSecretValue.
func (mr *MockSecretClientMockRecorder) GetSecretValue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockSecretClient)(nil).GetSecretValue), arg0, arg1)
}
