// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
	mock "github.com/stretchr/testify/mock"
)

// PoolServiceServer is an autogenerated mock type for the PoolServiceServer type
type PoolServiceServer struct {
	mock.Mock
}

// CreatePool provides a mock function with given fields: _a0, _a1
func (_m *PoolServiceServer) CreatePool(_a0 context.Context, _a1 *v1.CreatePoolRequest) (*v1.Pool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.Pool
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreatePoolRequest) *v1.Pool); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Pool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CreatePoolRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPool provides a mock function with given fields: _a0, _a1
func (_m *PoolServiceServer) GetPool(_a0 context.Context, _a1 *v1.GetPoolRequest) (*v1.Pool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.Pool
	if rf, ok := ret.Get(0).(func(context.Context, *v1.GetPoolRequest) *v1.Pool); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Pool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.GetPoolRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPools provides a mock function with given fields: _a0, _a1
func (_m *PoolServiceServer) ListPools(_a0 context.Context, _a1 *v1.ListPoolsRequest) (*v1.ListPoolsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.ListPoolsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ListPoolsRequest) *v1.ListPoolsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListPoolsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.ListPoolsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedPoolServiceServer provides a mock function with given fields:
func (_m *PoolServiceServer) mustEmbedUnimplementedPoolServiceServer() {
	_m.Called()
}

type NewPoolServiceServerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewPoolServiceServer creates a new instance of PoolServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPoolServiceServer(t NewPoolServiceServerT) *PoolServiceServer {
	mock := &PoolServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
