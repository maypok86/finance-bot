// Code generated by MockGen. DO NOT EDIT.
// Source: category.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/maypok86/finance-bot/internal/model"
)

// MockCategory is a mock of Category interface.
type MockCategory struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryMockRecorder
}

// MockCategoryMockRecorder is the mock recorder for MockCategory.
type MockCategoryMockRecorder struct {
	mock *MockCategory
}

// NewMockCategory creates a new mock instance.
func NewMockCategory(ctrl *gomock.Controller) *MockCategory {
	mock := &MockCategory{ctrl: ctrl}
	mock.recorder = &MockCategoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategory) EXPECT() *MockCategoryMockRecorder {
	return m.recorder
}

// GetAllCategories mocks base method.
func (m *MockCategory) GetAllCategories(ctx context.Context) ([]*model.DBCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCategories", ctx)
	ret0, _ := ret[0].([]*model.DBCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCategories indicates an expected call of GetAllCategories.
func (mr *MockCategoryMockRecorder) GetAllCategories(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCategories", reflect.TypeOf((*MockCategory)(nil).GetAllCategories), ctx)
}
