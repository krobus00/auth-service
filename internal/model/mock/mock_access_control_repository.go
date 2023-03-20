// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/krobus00/auth-service/internal/model (interfaces: AccessControlRepository)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/krobus00/auth-service/internal/model"
	gorm "gorm.io/gorm"
)

// MockAccessControlRepository is a mock of AccessControlRepository interface.
type MockAccessControlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccessControlRepositoryMockRecorder
}

// MockAccessControlRepositoryMockRecorder is the mock recorder for MockAccessControlRepository.
type MockAccessControlRepositoryMockRecorder struct {
	mock *MockAccessControlRepository
}

// NewMockAccessControlRepository creates a new mock instance.
func NewMockAccessControlRepository(ctrl *gomock.Controller) *MockAccessControlRepository {
	mock := &MockAccessControlRepository{ctrl: ctrl}
	mock.recorder = &MockAccessControlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessControlRepository) EXPECT() *MockAccessControlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccessControlRepository) Create(arg0 context.Context, arg1 *model.AccessControl) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAccessControlRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccessControlRepository)(nil).Create), arg0, arg1)
}

// InjectDB mocks base method.
func (m *MockAccessControlRepository) InjectDB(arg0 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectDB", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectDB indicates an expected call of InjectDB.
func (mr *MockAccessControlRepositoryMockRecorder) InjectDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectDB", reflect.TypeOf((*MockAccessControlRepository)(nil).InjectDB), arg0)
}

// Remove mocks base method.
func (m *MockAccessControlRepository) Remove(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockAccessControlRepositoryMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockAccessControlRepository)(nil).Remove), arg0, arg1)
}

// UpdateByID mocks base method.
func (m *MockAccessControlRepository) UpdateByID(arg0 context.Context, arg1 *model.AccessControl) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockAccessControlRepositoryMockRecorder) UpdateByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockAccessControlRepository)(nil).UpdateByID), arg0, arg1)
}
