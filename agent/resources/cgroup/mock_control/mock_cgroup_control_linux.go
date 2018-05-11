// Copyright 2015-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/resources/cgroup (interfaces: Control)

// Package mock_cgroup is a generated GoMock package.
package mock_cgroup

import (
	reflect "reflect"

	cgroup "github.com/aws/amazon-ecs-agent/agent/resources/cgroup"
	cgroups "github.com/containerd/cgroups"
	gomock "github.com/golang/mock/gomock"
)

// MockControl is a mock of Control interface
type MockControl struct {
	ctrl     *gomock.Controller
	recorder *MockControlMockRecorder
}

// MockControlMockRecorder is the mock recorder for MockControl
type MockControlMockRecorder struct {
	mock *MockControl
}

// NewMockControl creates a new mock instance
func NewMockControl(ctrl *gomock.Controller) *MockControl {
	mock := &MockControl{ctrl: ctrl}
	mock.recorder = &MockControlMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockControl) EXPECT() *MockControlMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockControl) Create(arg0 *cgroup.Spec) (cgroups.Cgroup, error) {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(cgroups.Cgroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockControlMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockControl)(nil).Create), arg0)
}

// Exists mocks base method
func (m *MockControl) Exists(arg0 string) bool {
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists
func (mr *MockControlMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockControl)(nil).Exists), arg0)
}

// Init mocks base method
func (m *MockControl) Init() error {
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockControlMockRecorder) Init() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockControl)(nil).Init))
}

// Remove mocks base method
func (m *MockControl) Remove(arg0 string) error {
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockControlMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockControl)(nil).Remove), arg0)
}
