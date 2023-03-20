// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/krobus00/auth-service/internal/model (interfaces: UserUsecase)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/krobus00/auth-service/internal/model"
	gorm "gorm.io/gorm"
)

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// GetUserInfo mocks base method.
func (m *MockUserUsecase) GetUserInfo(arg0 context.Context, arg1 *model.GetUserInfoPayload) (*model.UserInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", arg0, arg1)
	ret0, _ := ret[0].(*model.UserInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockUserUsecaseMockRecorder) GetUserInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockUserUsecase)(nil).GetUserInfo), arg0, arg1)
}

// InjectDB mocks base method.
func (m *MockUserUsecase) InjectDB(arg0 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectDB", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectDB indicates an expected call of InjectDB.
func (mr *MockUserUsecaseMockRecorder) InjectDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectDB", reflect.TypeOf((*MockUserUsecase)(nil).InjectDB), arg0)
}

// InjectTokenRepo mocks base method.
func (m *MockUserUsecase) InjectTokenRepo(arg0 model.TokenRepository) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectTokenRepo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectTokenRepo indicates an expected call of InjectTokenRepo.
func (mr *MockUserUsecaseMockRecorder) InjectTokenRepo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectTokenRepo", reflect.TypeOf((*MockUserUsecase)(nil).InjectTokenRepo), arg0)
}

// InjectUserRepo mocks base method.
func (m *MockUserUsecase) InjectUserRepo(arg0 model.UserRepository) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectUserRepo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectUserRepo indicates an expected call of InjectUserRepo.
func (mr *MockUserUsecaseMockRecorder) InjectUserRepo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectUserRepo", reflect.TypeOf((*MockUserUsecase)(nil).InjectUserRepo), arg0)
}

// Login mocks base method.
func (m *MockUserUsecase) Login(arg0 context.Context, arg1 *model.UserLoginPayload) (*model.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*model.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserUsecaseMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserUsecase)(nil).Login), arg0, arg1)
}

// Logout mocks base method.
func (m *MockUserUsecase) Logout(arg0 context.Context, arg1 *model.LogoutPayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockUserUsecaseMockRecorder) Logout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUserUsecase)(nil).Logout), arg0, arg1)
}

// RefreshToken mocks base method.
func (m *MockUserUsecase) RefreshToken(arg0 context.Context, arg1 *model.RefreshTokenPayload) (*model.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", arg0, arg1)
	ret0, _ := ret[0].(*model.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockUserUsecaseMockRecorder) RefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockUserUsecase)(nil).RefreshToken), arg0, arg1)
}

// Register mocks base method.
func (m *MockUserUsecase) Register(arg0 context.Context, arg1 *model.UserRegistrationPayload) (*model.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*model.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserUsecaseMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserUsecase)(nil).Register), arg0, arg1)
}