// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/krobus00/auth-service/internal/model (interfaces: PermissionRepository)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	redis "github.com/go-redis/redis/v8"
	gomock "github.com/golang/mock/gomock"
	model "github.com/krobus00/auth-service/internal/model"
	gorm "gorm.io/gorm"
)

// MockPermissionRepository is a mock of PermissionRepository interface.
type MockPermissionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPermissionRepositoryMockRecorder
}

// MockPermissionRepositoryMockRecorder is the mock recorder for MockPermissionRepository.
type MockPermissionRepositoryMockRecorder struct {
	mock *MockPermissionRepository
}

// NewMockPermissionRepository creates a new mock instance.
func NewMockPermissionRepository(ctrl *gomock.Controller) *MockPermissionRepository {
	mock := &MockPermissionRepository{ctrl: ctrl}
	mock.recorder = &MockPermissionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPermissionRepository) EXPECT() *MockPermissionRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPermissionRepository) Create(arg0 context.Context, arg1 *model.Permission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPermissionRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPermissionRepository)(nil).Create), arg0, arg1)
}

// DeleteByID mocks base method.
func (m *MockPermissionRepository) DeleteByID(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockPermissionRepositoryMockRecorder) DeleteByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockPermissionRepository)(nil).DeleteByID), arg0, arg1)
}

// FindByID mocks base method.
func (m *MockPermissionRepository) FindByID(arg0 context.Context, arg1 string) (*model.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockPermissionRepositoryMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockPermissionRepository)(nil).FindByID), arg0, arg1)
}

// FindByName mocks base method.
func (m *MockPermissionRepository) FindByName(arg0 context.Context, arg1 string) (*model.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0, arg1)
	ret0, _ := ret[0].(*model.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockPermissionRepositoryMockRecorder) FindByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockPermissionRepository)(nil).FindByName), arg0, arg1)
}

// InjectDB mocks base method.
func (m *MockPermissionRepository) InjectDB(arg0 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectDB", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectDB indicates an expected call of InjectDB.
func (mr *MockPermissionRepositoryMockRecorder) InjectDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectDB", reflect.TypeOf((*MockPermissionRepository)(nil).InjectDB), arg0)
}

// InjectRedisClient mocks base method.
func (m *MockPermissionRepository) InjectRedisClient(arg0 *redis.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectRedisClient", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectRedisClient indicates an expected call of InjectRedisClient.
func (mr *MockPermissionRepositoryMockRecorder) InjectRedisClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectRedisClient", reflect.TypeOf((*MockPermissionRepository)(nil).InjectRedisClient), arg0)
}

// Update mocks base method.
func (m *MockPermissionRepository) Update(arg0 context.Context, arg1 *model.Permission) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPermissionRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPermissionRepository)(nil).Update), arg0, arg1)
}
