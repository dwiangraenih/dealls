// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/dwiangraenihantik/go/src/github.com/dwiangraeni/dealls/utils/util_generate_password.go

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	reflect "reflect"

	middleware "github.com/dwiangraeni/dealls/middleware"
	model "github.com/dwiangraeni/dealls/model"
	gomock "github.com/golang/mock/gomock"
)

// MockPasswordHasher is a mock of PasswordHasher interface.
type MockPasswordHasher struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordHasherMockRecorder
}

// MockPasswordHasherMockRecorder is the mock recorder for MockPasswordHasher.
type MockPasswordHasherMockRecorder struct {
	mock *MockPasswordHasher
}

// NewMockPasswordHasher creates a new mock instance.
func NewMockPasswordHasher(ctrl *gomock.Controller) *MockPasswordHasher {
	mock := &MockPasswordHasher{ctrl: ctrl}
	mock.recorder = &MockPasswordHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordHasher) EXPECT() *MockPasswordHasherMockRecorder {
	return m.recorder
}

// CheckPasswordHash mocks base method.
func (m *MockPasswordHasher) CheckPasswordHash(password, hash string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPasswordHash", password, hash)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckPasswordHash indicates an expected call of CheckPasswordHash.
func (mr *MockPasswordHasherMockRecorder) CheckPasswordHash(password, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPasswordHash", reflect.TypeOf((*MockPasswordHasher)(nil).CheckPasswordHash), password, hash)
}

// GeneratePassword mocks base method.
func (m *MockPasswordHasher) GeneratePassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeneratePassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeneratePassword indicates an expected call of GeneratePassword.
func (mr *MockPasswordHasherMockRecorder) GeneratePassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeneratePassword", reflect.TypeOf((*MockPasswordHasher)(nil).GeneratePassword), password)
}

// GenerateToken mocks base method.
func (m *MockPasswordHasher) GenerateToken(account model.AccountBaseModel, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", account, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockPasswordHasherMockRecorder) GenerateToken(account, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockPasswordHasher)(nil).GenerateToken), account, key)
}

// VerifyToken mocks base method.
func (m *MockPasswordHasher) VerifyToken(token, key string) (*middleware.AccessTokenClaim, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", token, key)
	ret0, _ := ret[0].(*middleware.AccessTokenClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockPasswordHasherMockRecorder) VerifyToken(token, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockPasswordHasher)(nil).VerifyToken), token, key)
}
