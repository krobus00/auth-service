// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/krobus00/auth-service/internal/model (interfaces: GroupRepository)

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

// MockGroupRepository is a mock of GroupRepository interface.
type MockGroupRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGroupRepositoryMockRecorder
}

// MockGroupRepositoryMockRecorder is the mock recorder for MockGroupRepository.
type MockGroupRepositoryMockRecorder struct {
	mock *MockGroupRepository
}

// NewMockGroupRepository creates a new mock instance.
func NewMockGroupRepository(ctrl *gomock.Controller) *MockGroupRepository {
	mock := &MockGroupRepository{ctrl: ctrl}
	mock.recorder = &MockGroupRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupRepository) EXPECT() *MockGroupRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockGroupRepository) Create(arg0 context.Context, arg1 *model.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockGroupRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGroupRepository)(nil).Create), arg0, arg1)
}

// DeleteByID mocks base method.
func (m *MockGroupRepository) DeleteByID(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockGroupRepositoryMockRecorder) DeleteByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockGroupRepository)(nil).DeleteByID), arg0, arg1)
}

// FindByID mocks base method.
func (m *MockGroupRepository) FindByID(arg0 context.Context, arg1 string) (*model.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockGroupRepositoryMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockGroupRepository)(nil).FindByID), arg0, arg1)
}

// FindByName mocks base method.
func (m *MockGroupRepository) FindByName(arg0 context.Context, arg1 string) (*model.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0, arg1)
	ret0, _ := ret[0].(*model.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockGroupRepositoryMockRecorder) FindByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockGroupRepository)(nil).FindByName), arg0, arg1)
}

// InjectDB mocks base method.
func (m *MockGroupRepository) InjectDB(arg0 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectDB", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectDB indicates an expected call of InjectDB.
func (mr *MockGroupRepositoryMockRecorder) InjectDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectDB", reflect.TypeOf((*MockGroupRepository)(nil).InjectDB), arg0)
}

// InjectRedisClient mocks base method.
func (m *MockGroupRepository) InjectRedisClient(arg0 *redis.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectRedisClient", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectRedisClient indicates an expected call of InjectRedisClient.
func (mr *MockGroupRepositoryMockRecorder) InjectRedisClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectRedisClient", reflect.TypeOf((*MockGroupRepository)(nil).InjectRedisClient), arg0)
}

// Update mocks base method.
func (m *MockGroupRepository) Update(arg0 context.Context, arg1 *model.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockGroupRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGroupRepository)(nil).Update), arg0, arg1)
}
