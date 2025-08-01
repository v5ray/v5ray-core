// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/v4fly/v4ray-core/v0/proxy (interfaces: Inbound,Outbound)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	net "github.com/v4fly/v4ray-core/v0/common/net"
	routing "github.com/v4fly/v4ray-core/v0/features/routing"
	transport "github.com/v4fly/v4ray-core/v0/transport"
	internet "github.com/v4fly/v4ray-core/v0/transport/internet"
)

// ProxyInbound is a mock of Inbound interface.
type ProxyInbound struct {
	ctrl     *gomock.Controller
	recorder *ProxyInboundMockRecorder
}

// ProxyInboundMockRecorder is the mock recorder for ProxyInbound.
type ProxyInboundMockRecorder struct {
	mock *ProxyInbound
}

// NewProxyInbound creates a new mock instance.
func NewProxyInbound(ctrl *gomock.Controller) *ProxyInbound {
	mock := &ProxyInbound{ctrl: ctrl}
	mock.recorder = &ProxyInboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *ProxyInbound) EXPECT() *ProxyInboundMockRecorder {
	return m.recorder
}

// Network mocks base method.
func (m *ProxyInbound) Network() []net.Network {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Network")
	ret0, _ := ret[0].([]net.Network)
	return ret0
}

// Network indicates an expected call of Network.
func (mr *ProxyInboundMockRecorder) Network() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Network", reflect.TypeOf((*ProxyInbound)(nil).Network))
}

// Process mocks base method.
func (m *ProxyInbound) Process(arg0 context.Context, arg1 net.Network, arg2 internet.Connection, arg3 routing.Dispatcher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Process indicates an expected call of Process.
func (mr *ProxyInboundMockRecorder) Process(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*ProxyInbound)(nil).Process), arg0, arg1, arg2, arg3)
}

// ProxyOutbound is a mock of Outbound interface.
type ProxyOutbound struct {
	ctrl     *gomock.Controller
	recorder *ProxyOutboundMockRecorder
}

// ProxyOutboundMockRecorder is the mock recorder for ProxyOutbound.
type ProxyOutboundMockRecorder struct {
	mock *ProxyOutbound
}

// NewProxyOutbound creates a new mock instance.
func NewProxyOutbound(ctrl *gomock.Controller) *ProxyOutbound {
	mock := &ProxyOutbound{ctrl: ctrl}
	mock.recorder = &ProxyOutboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *ProxyOutbound) EXPECT() *ProxyOutboundMockRecorder {
	return m.recorder
}

// Process mocks base method.
func (m *ProxyOutbound) Process(arg0 context.Context, arg1 *transport.Link, arg2 internet.Dialer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Process indicates an expected call of Process.
func (mr *ProxyOutboundMockRecorder) Process(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*ProxyOutbound)(nil).Process), arg0, arg1, arg2)
}
