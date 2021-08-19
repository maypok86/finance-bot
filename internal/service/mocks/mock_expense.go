// Code generated by MockGen. DO NOT EDIT.
// Source: expense.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	model "github.com/LazyBearCT/finance-bot/internal/model"
	times "github.com/LazyBearCT/finance-bot/pkg/times"
	gomock "github.com/golang/mock/gomock"
)

// MockExpense is a mock of Expense interface.
type MockExpense struct {
	ctrl     *gomock.Controller
	recorder *MockExpenseMockRecorder
}

// MockExpenseMockRecorder is the mock recorder for MockExpense.
type MockExpenseMockRecorder struct {
	mock *MockExpense
}

// NewMockExpense creates a new mock instance.
func NewMockExpense(ctrl *gomock.Controller) *MockExpense {
	mock := &MockExpense{ctrl: ctrl}
	mock.recorder = &MockExpenseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExpense) EXPECT() *MockExpenseMockRecorder {
	return m.recorder
}

// AddExpense mocks base method.
func (m *MockExpense) AddExpense(rawMessage string) (*model.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddExpense", rawMessage)
	ret0, _ := ret[0].(*model.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddExpense indicates an expected call of AddExpense.
func (mr *MockExpenseMockRecorder) AddExpense(rawMessage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddExpense", reflect.TypeOf((*MockExpense)(nil).AddExpense), rawMessage)
}

// DeleteByID mocks base method.
func (m *MockExpense) DeleteByID(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockExpenseMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockExpense)(nil).DeleteByID), id)
}

// GetAllByPeriod mocks base method.
func (m *MockExpense) GetAllByPeriod(period times.Period) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByPeriod", period)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetAllByPeriod indicates an expected call of GetAllByPeriod.
func (mr *MockExpenseMockRecorder) GetAllByPeriod(period interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByPeriod", reflect.TypeOf((*MockExpense)(nil).GetAllByPeriod), period)
}

// GetBaseByPeriod mocks base method.
func (m *MockExpense) GetBaseByPeriod(period times.Period) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBaseByPeriod", period)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetBaseByPeriod indicates an expected call of GetBaseByPeriod.
func (mr *MockExpenseMockRecorder) GetBaseByPeriod(period interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaseByPeriod", reflect.TypeOf((*MockExpense)(nil).GetBaseByPeriod), period)
}

// GetLastExpenses mocks base method.
func (m *MockExpense) GetLastExpenses() ([]*model.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastExpenses")
	ret0, _ := ret[0].([]*model.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastExpenses indicates an expected call of GetLastExpenses.
func (mr *MockExpenseMockRecorder) GetLastExpenses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastExpenses", reflect.TypeOf((*MockExpense)(nil).GetLastExpenses))
}