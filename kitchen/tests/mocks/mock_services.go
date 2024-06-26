// Code generated by MockGen. DO NOT EDIT.
// Source: internal/ports/services.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	context "context"
	models "kitchenService/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKitchenService is a mock of KitchenService interface.
type MockKitchenService struct {
	ctrl     *gomock.Controller
	recorder *MockKitchenServiceMockRecorder
}

// MockKitchenServiceMockRecorder is the mock recorder for MockKitchenService.
type MockKitchenServiceMockRecorder struct {
	mock *MockKitchenService
}

// NewMockKitchenService creates a new mock instance.
func NewMockKitchenService(ctrl *gomock.Controller) *MockKitchenService {
	mock := &MockKitchenService{ctrl: ctrl}
	mock.recorder = &MockKitchenServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKitchenService) EXPECT() *MockKitchenServiceMockRecorder {
	return m.recorder
}

// ChangeOrderStatus mocks base method.
func (m *MockKitchenService) ChangeOrderStatus(ctx context.Context, orderId string, status *models.OrderStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOrderStatus", ctx, orderId, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeOrderStatus indicates an expected call of ChangeOrderStatus.
func (mr *MockKitchenServiceMockRecorder) ChangeOrderStatus(ctx, orderId, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOrderStatus", reflect.TypeOf((*MockKitchenService)(nil).ChangeOrderStatus), ctx, orderId, status)
}

// ProcessOrder mocks base method.
func (m *MockKitchenService) ProcessOrder(ctx context.Context, order *models.Order) (string, *models.OrderStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessOrder", ctx, order)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*models.OrderStatus)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ProcessOrder indicates an expected call of ProcessOrder.
func (mr *MockKitchenServiceMockRecorder) ProcessOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessOrder", reflect.TypeOf((*MockKitchenService)(nil).ProcessOrder), ctx, order)
}

// MockOrderProxy is a mock of OrderProxy interface.
type MockOrderProxy struct {
	ctrl     *gomock.Controller
	recorder *MockOrderProxyMockRecorder
}

// MockOrderProxyMockRecorder is the mock recorder for MockOrderProxy.
type MockOrderProxyMockRecorder struct {
	mock *MockOrderProxy
}

// NewMockOrderProxy creates a new mock instance.
func NewMockOrderProxy(ctrl *gomock.Controller) *MockOrderProxy {
	mock := &MockOrderProxy{ctrl: ctrl}
	mock.recorder = &MockOrderProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderProxy) EXPECT() *MockOrderProxyMockRecorder {
	return m.recorder
}

// ChangeOrderStatus mocks base method.
func (m *MockOrderProxy) ChangeOrderStatus(ctx context.Context, orderId, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOrderStatus", ctx, orderId, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeOrderStatus indicates an expected call of ChangeOrderStatus.
func (mr *MockOrderProxyMockRecorder) ChangeOrderStatus(ctx, orderId, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOrderStatus", reflect.TypeOf((*MockOrderProxy)(nil).ChangeOrderStatus), ctx, orderId, status)
}
