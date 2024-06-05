// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	proxyrpc "github.com/google/syzkaller/vm/proxyapp/proxyrpc"
	mock "github.com/stretchr/testify/mock"
)

// ProxyAppInterface is an autogenerated mock type for the ProxyAppInterface type
type ProxyAppInterface struct {
	mock.Mock
}

// Close provides a mock function with given fields: in, out
func (_m *ProxyAppInterface) Close(in proxyrpc.CloseParams, out *proxyrpc.CloseReply) error {
	ret := _m.Called(in, out)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(proxyrpc.CloseParams, *proxyrpc.CloseReply) error); ok {
		r0 = rf(in, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Copy provides a mock function with given fields: in, out
func (_m *ProxyAppInterface) Copy(in proxyrpc.CopyParams, out *proxyrpc.CopyResult) error {
	ret := _m.Called(in, out)

	if len(ret) == 0 {
		panic("no return value specified for Copy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(proxyrpc.CopyParams, *proxyrpc.CopyResult) error); ok {
		r0 = rf(in, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateInstance provides a mock function with given fields: in, out
func (_m *ProxyAppInterface) CreateInstance(in proxyrpc.CreateInstanceParams, out *proxyrpc.CreateInstanceResult) error {
	ret := _m.Called(in, out)

	if len(ret) == 0 {
		panic("no return value specified for CreateInstance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(proxyrpc.CreateInstanceParams, *proxyrpc.CreateInstanceResult) error); ok {
		r0 = rf(in, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePool provides a mock function with given fields: in, out
func (_m *ProxyAppInterface) CreatePool(in proxyrpc.CreatePoolParams, out *proxyrpc.CreatePoolResult) error {
	ret := _m.Called(in, out)

	if len(ret) == 0 {
		panic("no return value specified for CreatePool")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(proxyrpc.CreatePoolParams, *proxyrpc.CreatePoolResult) error); ok {
		r0 = rf(in, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Diagnose provides a mock function with given fields: in, out
func (_m *ProxyAppInterface) Diagnose(in proxyrpc.DiagnoseParams, out *proxyrpc.DiagnoseReply) error {
	ret := _m.Called(in, out)

	if len(ret) == 0 {
		panic("no return value specified for Diagnose")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(proxyrpc.DiagnoseParams, *proxyrpc.DiagnoseReply) error); ok {
		r0 = rf(in, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Forward provides a mock function with given fields: in, out
func (_m *ProxyAppInterface) Forward(in proxyrpc.ForwardParams, out *proxyrpc.ForwardResult) error {
	ret := _m.Called(in, out)

	if len(ret) == 0 {
		panic("no return value specified for Forward")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(proxyrpc.ForwardParams, *proxyrpc.ForwardResult) error); ok {
		r0 = rf(in, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PoolLogs provides a mock function with given fields: in, out
func (_m *ProxyAppInterface) PoolLogs(in proxyrpc.PoolLogsParam, out *proxyrpc.PoolLogsReply) error {
	ret := _m.Called(in, out)

	if len(ret) == 0 {
		panic("no return value specified for PoolLogs")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(proxyrpc.PoolLogsParam, *proxyrpc.PoolLogsReply) error); ok {
		r0 = rf(in, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunReadProgress provides a mock function with given fields: in, out
func (_m *ProxyAppInterface) RunReadProgress(in proxyrpc.RunReadProgressParams, out *proxyrpc.RunReadProgressReply) error {
	ret := _m.Called(in, out)

	if len(ret) == 0 {
		panic("no return value specified for RunReadProgress")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(proxyrpc.RunReadProgressParams, *proxyrpc.RunReadProgressReply) error); ok {
		r0 = rf(in, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunStart provides a mock function with given fields: in, out
func (_m *ProxyAppInterface) RunStart(in proxyrpc.RunStartParams, out *proxyrpc.RunStartReply) error {
	ret := _m.Called(in, out)

	if len(ret) == 0 {
		panic("no return value specified for RunStart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(proxyrpc.RunStartParams, *proxyrpc.RunStartReply) error); ok {
		r0 = rf(in, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunStop provides a mock function with given fields: in, out
func (_m *ProxyAppInterface) RunStop(in proxyrpc.RunStopParams, out *proxyrpc.RunStopReply) error {
	ret := _m.Called(in, out)

	if len(ret) == 0 {
		panic("no return value specified for RunStop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(proxyrpc.RunStopParams, *proxyrpc.RunStopReply) error); ok {
		r0 = rf(in, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProxyAppInterface creates a new instance of ProxyAppInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProxyAppInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProxyAppInterface {
	mock := &ProxyAppInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
