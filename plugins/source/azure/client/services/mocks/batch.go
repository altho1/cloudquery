// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: BatchAccountsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	batch "github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"
	gomock "github.com/golang/mock/gomock"
)

// MockBatchAccountsClient is a mock of BatchAccountsClient interface.
type MockBatchAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockBatchAccountsClientMockRecorder
}

// MockBatchAccountsClientMockRecorder is the mock recorder for MockBatchAccountsClient.
type MockBatchAccountsClientMockRecorder struct {
	mock *MockBatchAccountsClient
}

// NewMockBatchAccountsClient creates a new mock instance.
func NewMockBatchAccountsClient(ctrl *gomock.Controller) *MockBatchAccountsClient {
	mock := &MockBatchAccountsClient{ctrl: ctrl}
	mock.recorder = &MockBatchAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBatchAccountsClient) EXPECT() *MockBatchAccountsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockBatchAccountsClient) List(arg0 context.Context) (batch.AccountListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(batch.AccountListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockBatchAccountsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockBatchAccountsClient)(nil).List), arg0)
}
