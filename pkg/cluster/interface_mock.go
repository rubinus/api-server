// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cluster/interface.go

// Package cluster is a generated GoMock package.
package cluster

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClusterRegistrationOperator is a mock of ClusterRegistrationOperator interface.
type MockClusterRegistrationOperator struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRegistrationOperatorMockRecorder
}

// MockClusterRegistrationOperatorMockRecorder is the mock recorder for MockClusterRegistrationOperator.
type MockClusterRegistrationOperatorMockRecorder struct {
	mock *MockClusterRegistrationOperator
}

// NewMockClusterRegistrationOperator creates a new mock instance.
func NewMockClusterRegistrationOperator(ctrl *gomock.Controller) *MockClusterRegistrationOperator {
	mock := &MockClusterRegistrationOperator{ctrl: ctrl}
	mock.recorder = &MockClusterRegistrationOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRegistrationOperator) EXPECT() *MockClusterRegistrationOperatorMockRecorder {
	return m.recorder
}

// GetArgocdURL mocks base method.
func (m *MockClusterRegistrationOperator) GetArgocdURL() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArgocdURL")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArgocdURL indicates an expected call of GetArgocdURL.
func (mr *MockClusterRegistrationOperatorMockRecorder) GetArgocdURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArgocdURL", reflect.TypeOf((*MockClusterRegistrationOperator)(nil).GetArgocdURL))
}

// GetTraefikNodePortToHostCluster mocks base method.
func (m *MockClusterRegistrationOperator) GetTraefikNodePortToHostCluster(tenantLocalPath, hostCluster string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTraefikNodePortToHostCluster", tenantLocalPath, hostCluster)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTraefikNodePortToHostCluster indicates an expected call of GetTraefikNodePortToHostCluster.
func (mr *MockClusterRegistrationOperatorMockRecorder) GetTraefikNodePortToHostCluster(tenantLocalPath, hostCluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTraefikNodePortToHostCluster", reflect.TypeOf((*MockClusterRegistrationOperator)(nil).GetTraefikNodePortToHostCluster), tenantLocalPath, hostCluster)
}

// InitializeDependencies mocks base method.
func (m *MockClusterRegistrationOperator) InitializeDependencies(param *ClusterRegistrationParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeDependencies", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeDependencies indicates an expected call of InitializeDependencies.
func (mr *MockClusterRegistrationOperatorMockRecorder) InitializeDependencies(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeDependencies", reflect.TypeOf((*MockClusterRegistrationOperator)(nil).InitializeDependencies), param)
}

// Remove mocks base method.
func (m *MockClusterRegistrationOperator) Remove() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove")
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockClusterRegistrationOperatorMockRecorder) Remove() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockClusterRegistrationOperator)(nil).Remove))
}

// Save mocks base method.
func (m *MockClusterRegistrationOperator) Save() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockClusterRegistrationOperatorMockRecorder) Save() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockClusterRegistrationOperator)(nil).Save))
}
