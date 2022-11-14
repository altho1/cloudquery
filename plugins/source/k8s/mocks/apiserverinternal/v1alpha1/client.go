// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1 (interfaces: InternalV1alpha1Interface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
	rest "k8s.io/client-go/rest"
)

// MockInternalV1alpha1Interface is a mock of InternalV1alpha1Interface interface.
type MockInternalV1alpha1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockInternalV1alpha1InterfaceMockRecorder
}

// MockInternalV1alpha1InterfaceMockRecorder is the mock recorder for MockInternalV1alpha1Interface.
type MockInternalV1alpha1InterfaceMockRecorder struct {
	mock *MockInternalV1alpha1Interface
}

// NewMockInternalV1alpha1Interface creates a new mock instance.
func NewMockInternalV1alpha1Interface(ctrl *gomock.Controller) *MockInternalV1alpha1Interface {
	mock := &MockInternalV1alpha1Interface{ctrl: ctrl}
	mock.recorder = &MockInternalV1alpha1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInternalV1alpha1Interface) EXPECT() *MockInternalV1alpha1InterfaceMockRecorder {
	return m.recorder
}

// RESTClient mocks base method.
func (m *MockInternalV1alpha1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockInternalV1alpha1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockInternalV1alpha1Interface)(nil).RESTClient))
}

// StorageVersions mocks base method.
func (m *MockInternalV1alpha1Interface) StorageVersions() v1alpha1.StorageVersionInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageVersions")
	ret0, _ := ret[0].(v1alpha1.StorageVersionInterface)
	return ret0
}

// StorageVersions indicates an expected call of StorageVersions.
func (mr *MockInternalV1alpha1InterfaceMockRecorder) StorageVersions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageVersions", reflect.TypeOf((*MockInternalV1alpha1Interface)(nil).StorageVersions))
}