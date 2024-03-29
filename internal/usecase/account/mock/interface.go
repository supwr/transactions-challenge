// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/supwr/pismo-transactions/internal/entity"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepositoryInterface) Create(ctx context.Context, account *entity.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryInterfaceMockRecorder) Create(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepositoryInterface)(nil).Create), ctx, account)
}

// FindByDocument mocks base method.
func (m *MockRepositoryInterface) FindByDocument(ctx context.Context, document entity.Document) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDocument", ctx, document)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDocument indicates an expected call of FindByDocument.
func (mr *MockRepositoryInterfaceMockRecorder) FindByDocument(ctx, document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDocument", reflect.TypeOf((*MockRepositoryInterface)(nil).FindByDocument), ctx, document)
}

// FindById mocks base method.
func (m *MockRepositoryInterface) FindById(ctx context.Context, id int) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, id)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockRepositoryInterfaceMockRecorder) FindById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockRepositoryInterface)(nil).FindById), ctx, id)
}
