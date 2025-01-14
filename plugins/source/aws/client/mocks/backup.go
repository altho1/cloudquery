// Code generated by MockGen. DO NOT EDIT.
// Source: backup.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	backup "github.com/aws/aws-sdk-go-v2/service/backup"
	gomock "github.com/golang/mock/gomock"
)

// MockBackupClient is a mock of BackupClient interface.
type MockBackupClient struct {
	ctrl     *gomock.Controller
	recorder *MockBackupClientMockRecorder
}

// MockBackupClientMockRecorder is the mock recorder for MockBackupClient.
type MockBackupClientMockRecorder struct {
	mock *MockBackupClient
}

// NewMockBackupClient creates a new mock instance.
func NewMockBackupClient(ctrl *gomock.Controller) *MockBackupClient {
	mock := &MockBackupClient{ctrl: ctrl}
	mock.recorder = &MockBackupClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackupClient) EXPECT() *MockBackupClientMockRecorder {
	return m.recorder
}

// DescribeBackupJob mocks base method.
func (m *MockBackupClient) DescribeBackupJob(arg0 context.Context, arg1 *backup.DescribeBackupJobInput, arg2 ...func(*backup.Options)) (*backup.DescribeBackupJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBackupJob", varargs...)
	ret0, _ := ret[0].(*backup.DescribeBackupJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeBackupJob indicates an expected call of DescribeBackupJob.
func (mr *MockBackupClientMockRecorder) DescribeBackupJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBackupJob", reflect.TypeOf((*MockBackupClient)(nil).DescribeBackupJob), varargs...)
}

// DescribeBackupVault mocks base method.
func (m *MockBackupClient) DescribeBackupVault(arg0 context.Context, arg1 *backup.DescribeBackupVaultInput, arg2 ...func(*backup.Options)) (*backup.DescribeBackupVaultOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBackupVault", varargs...)
	ret0, _ := ret[0].(*backup.DescribeBackupVaultOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeBackupVault indicates an expected call of DescribeBackupVault.
func (mr *MockBackupClientMockRecorder) DescribeBackupVault(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBackupVault", reflect.TypeOf((*MockBackupClient)(nil).DescribeBackupVault), varargs...)
}

// DescribeCopyJob mocks base method.
func (m *MockBackupClient) DescribeCopyJob(arg0 context.Context, arg1 *backup.DescribeCopyJobInput, arg2 ...func(*backup.Options)) (*backup.DescribeCopyJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCopyJob", varargs...)
	ret0, _ := ret[0].(*backup.DescribeCopyJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCopyJob indicates an expected call of DescribeCopyJob.
func (mr *MockBackupClientMockRecorder) DescribeCopyJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCopyJob", reflect.TypeOf((*MockBackupClient)(nil).DescribeCopyJob), varargs...)
}

// DescribeFramework mocks base method.
func (m *MockBackupClient) DescribeFramework(arg0 context.Context, arg1 *backup.DescribeFrameworkInput, arg2 ...func(*backup.Options)) (*backup.DescribeFrameworkOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFramework", varargs...)
	ret0, _ := ret[0].(*backup.DescribeFrameworkOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFramework indicates an expected call of DescribeFramework.
func (mr *MockBackupClientMockRecorder) DescribeFramework(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFramework", reflect.TypeOf((*MockBackupClient)(nil).DescribeFramework), varargs...)
}

// DescribeGlobalSettings mocks base method.
func (m *MockBackupClient) DescribeGlobalSettings(arg0 context.Context, arg1 *backup.DescribeGlobalSettingsInput, arg2 ...func(*backup.Options)) (*backup.DescribeGlobalSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeGlobalSettings", varargs...)
	ret0, _ := ret[0].(*backup.DescribeGlobalSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeGlobalSettings indicates an expected call of DescribeGlobalSettings.
func (mr *MockBackupClientMockRecorder) DescribeGlobalSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeGlobalSettings", reflect.TypeOf((*MockBackupClient)(nil).DescribeGlobalSettings), varargs...)
}

// DescribeProtectedResource mocks base method.
func (m *MockBackupClient) DescribeProtectedResource(arg0 context.Context, arg1 *backup.DescribeProtectedResourceInput, arg2 ...func(*backup.Options)) (*backup.DescribeProtectedResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeProtectedResource", varargs...)
	ret0, _ := ret[0].(*backup.DescribeProtectedResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeProtectedResource indicates an expected call of DescribeProtectedResource.
func (mr *MockBackupClientMockRecorder) DescribeProtectedResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeProtectedResource", reflect.TypeOf((*MockBackupClient)(nil).DescribeProtectedResource), varargs...)
}

// DescribeRecoveryPoint mocks base method.
func (m *MockBackupClient) DescribeRecoveryPoint(arg0 context.Context, arg1 *backup.DescribeRecoveryPointInput, arg2 ...func(*backup.Options)) (*backup.DescribeRecoveryPointOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRecoveryPoint", varargs...)
	ret0, _ := ret[0].(*backup.DescribeRecoveryPointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRecoveryPoint indicates an expected call of DescribeRecoveryPoint.
func (mr *MockBackupClientMockRecorder) DescribeRecoveryPoint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRecoveryPoint", reflect.TypeOf((*MockBackupClient)(nil).DescribeRecoveryPoint), varargs...)
}

// DescribeRegionSettings mocks base method.
func (m *MockBackupClient) DescribeRegionSettings(arg0 context.Context, arg1 *backup.DescribeRegionSettingsInput, arg2 ...func(*backup.Options)) (*backup.DescribeRegionSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRegionSettings", varargs...)
	ret0, _ := ret[0].(*backup.DescribeRegionSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRegionSettings indicates an expected call of DescribeRegionSettings.
func (mr *MockBackupClientMockRecorder) DescribeRegionSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRegionSettings", reflect.TypeOf((*MockBackupClient)(nil).DescribeRegionSettings), varargs...)
}

// DescribeReportJob mocks base method.
func (m *MockBackupClient) DescribeReportJob(arg0 context.Context, arg1 *backup.DescribeReportJobInput, arg2 ...func(*backup.Options)) (*backup.DescribeReportJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReportJob", varargs...)
	ret0, _ := ret[0].(*backup.DescribeReportJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReportJob indicates an expected call of DescribeReportJob.
func (mr *MockBackupClientMockRecorder) DescribeReportJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReportJob", reflect.TypeOf((*MockBackupClient)(nil).DescribeReportJob), varargs...)
}

// DescribeReportPlan mocks base method.
func (m *MockBackupClient) DescribeReportPlan(arg0 context.Context, arg1 *backup.DescribeReportPlanInput, arg2 ...func(*backup.Options)) (*backup.DescribeReportPlanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReportPlan", varargs...)
	ret0, _ := ret[0].(*backup.DescribeReportPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReportPlan indicates an expected call of DescribeReportPlan.
func (mr *MockBackupClientMockRecorder) DescribeReportPlan(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReportPlan", reflect.TypeOf((*MockBackupClient)(nil).DescribeReportPlan), varargs...)
}

// DescribeRestoreJob mocks base method.
func (m *MockBackupClient) DescribeRestoreJob(arg0 context.Context, arg1 *backup.DescribeRestoreJobInput, arg2 ...func(*backup.Options)) (*backup.DescribeRestoreJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRestoreJob", varargs...)
	ret0, _ := ret[0].(*backup.DescribeRestoreJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRestoreJob indicates an expected call of DescribeRestoreJob.
func (mr *MockBackupClientMockRecorder) DescribeRestoreJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRestoreJob", reflect.TypeOf((*MockBackupClient)(nil).DescribeRestoreJob), varargs...)
}

// GetBackupPlan mocks base method.
func (m *MockBackupClient) GetBackupPlan(arg0 context.Context, arg1 *backup.GetBackupPlanInput, arg2 ...func(*backup.Options)) (*backup.GetBackupPlanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackupPlan", varargs...)
	ret0, _ := ret[0].(*backup.GetBackupPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackupPlan indicates an expected call of GetBackupPlan.
func (mr *MockBackupClientMockRecorder) GetBackupPlan(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupPlan", reflect.TypeOf((*MockBackupClient)(nil).GetBackupPlan), varargs...)
}

// GetBackupPlanFromJSON mocks base method.
func (m *MockBackupClient) GetBackupPlanFromJSON(arg0 context.Context, arg1 *backup.GetBackupPlanFromJSONInput, arg2 ...func(*backup.Options)) (*backup.GetBackupPlanFromJSONOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackupPlanFromJSON", varargs...)
	ret0, _ := ret[0].(*backup.GetBackupPlanFromJSONOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackupPlanFromJSON indicates an expected call of GetBackupPlanFromJSON.
func (mr *MockBackupClientMockRecorder) GetBackupPlanFromJSON(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupPlanFromJSON", reflect.TypeOf((*MockBackupClient)(nil).GetBackupPlanFromJSON), varargs...)
}

// GetBackupPlanFromTemplate mocks base method.
func (m *MockBackupClient) GetBackupPlanFromTemplate(arg0 context.Context, arg1 *backup.GetBackupPlanFromTemplateInput, arg2 ...func(*backup.Options)) (*backup.GetBackupPlanFromTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackupPlanFromTemplate", varargs...)
	ret0, _ := ret[0].(*backup.GetBackupPlanFromTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackupPlanFromTemplate indicates an expected call of GetBackupPlanFromTemplate.
func (mr *MockBackupClientMockRecorder) GetBackupPlanFromTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupPlanFromTemplate", reflect.TypeOf((*MockBackupClient)(nil).GetBackupPlanFromTemplate), varargs...)
}

// GetBackupSelection mocks base method.
func (m *MockBackupClient) GetBackupSelection(arg0 context.Context, arg1 *backup.GetBackupSelectionInput, arg2 ...func(*backup.Options)) (*backup.GetBackupSelectionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackupSelection", varargs...)
	ret0, _ := ret[0].(*backup.GetBackupSelectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackupSelection indicates an expected call of GetBackupSelection.
func (mr *MockBackupClientMockRecorder) GetBackupSelection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupSelection", reflect.TypeOf((*MockBackupClient)(nil).GetBackupSelection), varargs...)
}

// GetBackupVaultAccessPolicy mocks base method.
func (m *MockBackupClient) GetBackupVaultAccessPolicy(arg0 context.Context, arg1 *backup.GetBackupVaultAccessPolicyInput, arg2 ...func(*backup.Options)) (*backup.GetBackupVaultAccessPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackupVaultAccessPolicy", varargs...)
	ret0, _ := ret[0].(*backup.GetBackupVaultAccessPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackupVaultAccessPolicy indicates an expected call of GetBackupVaultAccessPolicy.
func (mr *MockBackupClientMockRecorder) GetBackupVaultAccessPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupVaultAccessPolicy", reflect.TypeOf((*MockBackupClient)(nil).GetBackupVaultAccessPolicy), varargs...)
}

// GetBackupVaultNotifications mocks base method.
func (m *MockBackupClient) GetBackupVaultNotifications(arg0 context.Context, arg1 *backup.GetBackupVaultNotificationsInput, arg2 ...func(*backup.Options)) (*backup.GetBackupVaultNotificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackupVaultNotifications", varargs...)
	ret0, _ := ret[0].(*backup.GetBackupVaultNotificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackupVaultNotifications indicates an expected call of GetBackupVaultNotifications.
func (mr *MockBackupClientMockRecorder) GetBackupVaultNotifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupVaultNotifications", reflect.TypeOf((*MockBackupClient)(nil).GetBackupVaultNotifications), varargs...)
}

// GetRecoveryPointRestoreMetadata mocks base method.
func (m *MockBackupClient) GetRecoveryPointRestoreMetadata(arg0 context.Context, arg1 *backup.GetRecoveryPointRestoreMetadataInput, arg2 ...func(*backup.Options)) (*backup.GetRecoveryPointRestoreMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRecoveryPointRestoreMetadata", varargs...)
	ret0, _ := ret[0].(*backup.GetRecoveryPointRestoreMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecoveryPointRestoreMetadata indicates an expected call of GetRecoveryPointRestoreMetadata.
func (mr *MockBackupClientMockRecorder) GetRecoveryPointRestoreMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecoveryPointRestoreMetadata", reflect.TypeOf((*MockBackupClient)(nil).GetRecoveryPointRestoreMetadata), varargs...)
}

// GetSupportedResourceTypes mocks base method.
func (m *MockBackupClient) GetSupportedResourceTypes(arg0 context.Context, arg1 *backup.GetSupportedResourceTypesInput, arg2 ...func(*backup.Options)) (*backup.GetSupportedResourceTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSupportedResourceTypes", varargs...)
	ret0, _ := ret[0].(*backup.GetSupportedResourceTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSupportedResourceTypes indicates an expected call of GetSupportedResourceTypes.
func (mr *MockBackupClientMockRecorder) GetSupportedResourceTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportedResourceTypes", reflect.TypeOf((*MockBackupClient)(nil).GetSupportedResourceTypes), varargs...)
}

// ListBackupJobs mocks base method.
func (m *MockBackupClient) ListBackupJobs(arg0 context.Context, arg1 *backup.ListBackupJobsInput, arg2 ...func(*backup.Options)) (*backup.ListBackupJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBackupJobs", varargs...)
	ret0, _ := ret[0].(*backup.ListBackupJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBackupJobs indicates an expected call of ListBackupJobs.
func (mr *MockBackupClientMockRecorder) ListBackupJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackupJobs", reflect.TypeOf((*MockBackupClient)(nil).ListBackupJobs), varargs...)
}

// ListBackupPlanTemplates mocks base method.
func (m *MockBackupClient) ListBackupPlanTemplates(arg0 context.Context, arg1 *backup.ListBackupPlanTemplatesInput, arg2 ...func(*backup.Options)) (*backup.ListBackupPlanTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBackupPlanTemplates", varargs...)
	ret0, _ := ret[0].(*backup.ListBackupPlanTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBackupPlanTemplates indicates an expected call of ListBackupPlanTemplates.
func (mr *MockBackupClientMockRecorder) ListBackupPlanTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackupPlanTemplates", reflect.TypeOf((*MockBackupClient)(nil).ListBackupPlanTemplates), varargs...)
}

// ListBackupPlanVersions mocks base method.
func (m *MockBackupClient) ListBackupPlanVersions(arg0 context.Context, arg1 *backup.ListBackupPlanVersionsInput, arg2 ...func(*backup.Options)) (*backup.ListBackupPlanVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBackupPlanVersions", varargs...)
	ret0, _ := ret[0].(*backup.ListBackupPlanVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBackupPlanVersions indicates an expected call of ListBackupPlanVersions.
func (mr *MockBackupClientMockRecorder) ListBackupPlanVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackupPlanVersions", reflect.TypeOf((*MockBackupClient)(nil).ListBackupPlanVersions), varargs...)
}

// ListBackupPlans mocks base method.
func (m *MockBackupClient) ListBackupPlans(arg0 context.Context, arg1 *backup.ListBackupPlansInput, arg2 ...func(*backup.Options)) (*backup.ListBackupPlansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBackupPlans", varargs...)
	ret0, _ := ret[0].(*backup.ListBackupPlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBackupPlans indicates an expected call of ListBackupPlans.
func (mr *MockBackupClientMockRecorder) ListBackupPlans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackupPlans", reflect.TypeOf((*MockBackupClient)(nil).ListBackupPlans), varargs...)
}

// ListBackupSelections mocks base method.
func (m *MockBackupClient) ListBackupSelections(arg0 context.Context, arg1 *backup.ListBackupSelectionsInput, arg2 ...func(*backup.Options)) (*backup.ListBackupSelectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBackupSelections", varargs...)
	ret0, _ := ret[0].(*backup.ListBackupSelectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBackupSelections indicates an expected call of ListBackupSelections.
func (mr *MockBackupClientMockRecorder) ListBackupSelections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackupSelections", reflect.TypeOf((*MockBackupClient)(nil).ListBackupSelections), varargs...)
}

// ListBackupVaults mocks base method.
func (m *MockBackupClient) ListBackupVaults(arg0 context.Context, arg1 *backup.ListBackupVaultsInput, arg2 ...func(*backup.Options)) (*backup.ListBackupVaultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBackupVaults", varargs...)
	ret0, _ := ret[0].(*backup.ListBackupVaultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBackupVaults indicates an expected call of ListBackupVaults.
func (mr *MockBackupClientMockRecorder) ListBackupVaults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackupVaults", reflect.TypeOf((*MockBackupClient)(nil).ListBackupVaults), varargs...)
}

// ListCopyJobs mocks base method.
func (m *MockBackupClient) ListCopyJobs(arg0 context.Context, arg1 *backup.ListCopyJobsInput, arg2 ...func(*backup.Options)) (*backup.ListCopyJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCopyJobs", varargs...)
	ret0, _ := ret[0].(*backup.ListCopyJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCopyJobs indicates an expected call of ListCopyJobs.
func (mr *MockBackupClientMockRecorder) ListCopyJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCopyJobs", reflect.TypeOf((*MockBackupClient)(nil).ListCopyJobs), varargs...)
}

// ListFrameworks mocks base method.
func (m *MockBackupClient) ListFrameworks(arg0 context.Context, arg1 *backup.ListFrameworksInput, arg2 ...func(*backup.Options)) (*backup.ListFrameworksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFrameworks", varargs...)
	ret0, _ := ret[0].(*backup.ListFrameworksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFrameworks indicates an expected call of ListFrameworks.
func (mr *MockBackupClientMockRecorder) ListFrameworks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFrameworks", reflect.TypeOf((*MockBackupClient)(nil).ListFrameworks), varargs...)
}

// ListProtectedResources mocks base method.
func (m *MockBackupClient) ListProtectedResources(arg0 context.Context, arg1 *backup.ListProtectedResourcesInput, arg2 ...func(*backup.Options)) (*backup.ListProtectedResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProtectedResources", varargs...)
	ret0, _ := ret[0].(*backup.ListProtectedResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProtectedResources indicates an expected call of ListProtectedResources.
func (mr *MockBackupClientMockRecorder) ListProtectedResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProtectedResources", reflect.TypeOf((*MockBackupClient)(nil).ListProtectedResources), varargs...)
}

// ListRecoveryPointsByBackupVault mocks base method.
func (m *MockBackupClient) ListRecoveryPointsByBackupVault(arg0 context.Context, arg1 *backup.ListRecoveryPointsByBackupVaultInput, arg2 ...func(*backup.Options)) (*backup.ListRecoveryPointsByBackupVaultOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRecoveryPointsByBackupVault", varargs...)
	ret0, _ := ret[0].(*backup.ListRecoveryPointsByBackupVaultOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRecoveryPointsByBackupVault indicates an expected call of ListRecoveryPointsByBackupVault.
func (mr *MockBackupClientMockRecorder) ListRecoveryPointsByBackupVault(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecoveryPointsByBackupVault", reflect.TypeOf((*MockBackupClient)(nil).ListRecoveryPointsByBackupVault), varargs...)
}

// ListRecoveryPointsByResource mocks base method.
func (m *MockBackupClient) ListRecoveryPointsByResource(arg0 context.Context, arg1 *backup.ListRecoveryPointsByResourceInput, arg2 ...func(*backup.Options)) (*backup.ListRecoveryPointsByResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRecoveryPointsByResource", varargs...)
	ret0, _ := ret[0].(*backup.ListRecoveryPointsByResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRecoveryPointsByResource indicates an expected call of ListRecoveryPointsByResource.
func (mr *MockBackupClientMockRecorder) ListRecoveryPointsByResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecoveryPointsByResource", reflect.TypeOf((*MockBackupClient)(nil).ListRecoveryPointsByResource), varargs...)
}

// ListReportJobs mocks base method.
func (m *MockBackupClient) ListReportJobs(arg0 context.Context, arg1 *backup.ListReportJobsInput, arg2 ...func(*backup.Options)) (*backup.ListReportJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListReportJobs", varargs...)
	ret0, _ := ret[0].(*backup.ListReportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReportJobs indicates an expected call of ListReportJobs.
func (mr *MockBackupClientMockRecorder) ListReportJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReportJobs", reflect.TypeOf((*MockBackupClient)(nil).ListReportJobs), varargs...)
}

// ListReportPlans mocks base method.
func (m *MockBackupClient) ListReportPlans(arg0 context.Context, arg1 *backup.ListReportPlansInput, arg2 ...func(*backup.Options)) (*backup.ListReportPlansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListReportPlans", varargs...)
	ret0, _ := ret[0].(*backup.ListReportPlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReportPlans indicates an expected call of ListReportPlans.
func (mr *MockBackupClientMockRecorder) ListReportPlans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReportPlans", reflect.TypeOf((*MockBackupClient)(nil).ListReportPlans), varargs...)
}

// ListRestoreJobs mocks base method.
func (m *MockBackupClient) ListRestoreJobs(arg0 context.Context, arg1 *backup.ListRestoreJobsInput, arg2 ...func(*backup.Options)) (*backup.ListRestoreJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRestoreJobs", varargs...)
	ret0, _ := ret[0].(*backup.ListRestoreJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRestoreJobs indicates an expected call of ListRestoreJobs.
func (mr *MockBackupClientMockRecorder) ListRestoreJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRestoreJobs", reflect.TypeOf((*MockBackupClient)(nil).ListRestoreJobs), varargs...)
}

// ListTags mocks base method.
func (m *MockBackupClient) ListTags(arg0 context.Context, arg1 *backup.ListTagsInput, arg2 ...func(*backup.Options)) (*backup.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTags", varargs...)
	ret0, _ := ret[0].(*backup.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags.
func (mr *MockBackupClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockBackupClient)(nil).ListTags), varargs...)
}
