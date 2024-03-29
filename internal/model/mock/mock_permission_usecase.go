// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/krobus00/auth-service/internal/model (interfaces: PermissionUsecase)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/krobus00/auth-service/internal/model"
)

// MockPermissionUsecase is a mock of PermissionUsecase interface.
type MockPermissionUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockPermissionUsecaseMockRecorder
}

// MockPermissionUsecaseMockRecorder is the mock recorder for MockPermissionUsecase.
type MockPermissionUsecaseMockRecorder struct {
	mock *MockPermissionUsecase
}

// NewMockPermissionUsecase creates a new mock instance.
func NewMockPermissionUsecase(ctrl *gomock.Controller) *MockPermissionUsecase {
	mock := &MockPermissionUsecase{ctrl: ctrl}
	mock.recorder = &MockPermissionUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPermissionUsecase) EXPECT() *MockPermissionUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPermissionUsecase) Create(arg0 context.Context, arg1 *model.CreatePermissionPayload) (*model.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*model.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPermissionUsecaseMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPermissionUsecase)(nil).Create), arg0, arg1)
}

// DeleteByID mocks base method.
func (m *MockPermissionUsecase) DeleteByID(arg0 context.Context, arg1 *model.DeletePermissionByIDPayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockPermissionUsecaseMockRecorder) DeleteByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockPermissionUsecase)(nil).DeleteByID), arg0, arg1)
}

// FindByID mocks base method.
func (m *MockPermissionUsecase) FindByID(arg0 context.Context, arg1 *model.FindPermissionByIDPayload) (*model.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockPermissionUsecaseMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockPermissionUsecase)(nil).FindByID), arg0, arg1)
}

// FindByName mocks base method.
func (m *MockPermissionUsecase) FindByName(arg0 context.Context, arg1 *model.FindPermissionByNamePayload) (*model.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0, arg1)
	ret0, _ := ret[0].(*model.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockPermissionUsecaseMockRecorder) FindByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockPermissionUsecase)(nil).FindByName), arg0, arg1)
}

// InjectAuthUsecase mocks base method.
func (m *MockPermissionUsecase) InjectAuthUsecase(arg0 model.AuthUsecase) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectAuthUsecase", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectAuthUsecase indicates an expected call of InjectAuthUsecase.
func (mr *MockPermissionUsecaseMockRecorder) InjectAuthUsecase(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectAuthUsecase", reflect.TypeOf((*MockPermissionUsecase)(nil).InjectAuthUsecase), arg0)
}

// InjectPermissionRepo mocks base method.
func (m *MockPermissionUsecase) InjectPermissionRepo(arg0 model.PermissionRepository) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectPermissionRepo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectPermissionRepo indicates an expected call of InjectPermissionRepo.
func (mr *MockPermissionUsecaseMockRecorder) InjectPermissionRepo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectPermissionRepo", reflect.TypeOf((*MockPermissionUsecase)(nil).InjectPermissionRepo), arg0)
}

// Update mocks base method.
func (m *MockPermissionUsecase) Update(arg0 context.Context, arg1 *model.UpdatePermissionPayload) (*model.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*model.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPermissionUsecaseMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPermissionUsecase)(nil).Update), arg0, arg1)
}
