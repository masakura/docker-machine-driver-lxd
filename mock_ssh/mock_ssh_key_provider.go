// Code generated by MockGen. DO NOT EDIT.
// Source: ../ssh/ssh_key_provider.go

// Package mock_ssh is a generated GoMock package.
package mock_ssh

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSSHKeyProvider is a mock of SSHKeyProvider interface
type MockSSHKeyProvider struct {
	ctrl     *gomock.Controller
	recorder *MockSSHKeyProviderMockRecorder
}

// MockSSHKeyProviderMockRecorder is the mock recorder for MockSSHKeyProvider
type MockSSHKeyProviderMockRecorder struct {
	mock *MockSSHKeyProvider
}

// NewMockSSHKeyProvider creates a new mock instance
func NewMockSSHKeyProvider(ctrl *gomock.Controller) *MockSSHKeyProvider {
	mock := &MockSSHKeyProvider{ctrl: ctrl}
	mock.recorder = &MockSSHKeyProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSSHKeyProvider) EXPECT() *MockSSHKeyProviderMockRecorder {
	return m.recorder
}

// Generate mocks base method
func (m *MockSSHKeyProvider) Generate(privateKeyPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", privateKeyPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Generate indicates an expected call of Generate
func (mr *MockSSHKeyProviderMockRecorder) Generate(privateKeyPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockSSHKeyProvider)(nil).Generate), privateKeyPath)
}

// GetPublicKey mocks base method
func (m *MockSSHKeyProvider) GetPublicKey(privateKeyPath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey", privateKeyPath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKey indicates an expected call of GetPublicKey
func (mr *MockSSHKeyProviderMockRecorder) GetPublicKey(privateKeyPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockSSHKeyProvider)(nil).GetPublicKey), privateKeyPath)
}
