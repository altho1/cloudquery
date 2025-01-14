// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: StreamAnalyticsStreamingJobsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	streamanalytics "github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2020-03-01/streamanalytics"
	gomock "github.com/golang/mock/gomock"
)

// MockStreamAnalyticsStreamingJobsClient is a mock of StreamAnalyticsStreamingJobsClient interface.
type MockStreamAnalyticsStreamingJobsClient struct {
	ctrl     *gomock.Controller
	recorder *MockStreamAnalyticsStreamingJobsClientMockRecorder
}

// MockStreamAnalyticsStreamingJobsClientMockRecorder is the mock recorder for MockStreamAnalyticsStreamingJobsClient.
type MockStreamAnalyticsStreamingJobsClientMockRecorder struct {
	mock *MockStreamAnalyticsStreamingJobsClient
}

// NewMockStreamAnalyticsStreamingJobsClient creates a new mock instance.
func NewMockStreamAnalyticsStreamingJobsClient(ctrl *gomock.Controller) *MockStreamAnalyticsStreamingJobsClient {
	mock := &MockStreamAnalyticsStreamingJobsClient{ctrl: ctrl}
	mock.recorder = &MockStreamAnalyticsStreamingJobsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamAnalyticsStreamingJobsClient) EXPECT() *MockStreamAnalyticsStreamingJobsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockStreamAnalyticsStreamingJobsClient) List(arg0 context.Context, arg1 string) (streamanalytics.StreamingJobListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(streamanalytics.StreamingJobListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStreamAnalyticsStreamingJobsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStreamAnalyticsStreamingJobsClient)(nil).List), arg0, arg1)
}
