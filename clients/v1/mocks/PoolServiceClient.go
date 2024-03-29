// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	gax "github.com/googleapis/gax-go/v2"
	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	poolsv1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"

	v1 "github.com/coinbase/waas-client-library-go/clients/v1"
)

// PoolServiceClient is an autogenerated mock type for the PoolServiceClient type
type PoolServiceClient struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *PoolServiceClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Connection provides a mock function with given fields:
func (_m *PoolServiceClient) Connection() *grpc.ClientConn {
	ret := _m.Called()

	var r0 *grpc.ClientConn
	if rf, ok := ret.Get(0).(func() *grpc.ClientConn); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpc.ClientConn)
		}
	}

	return r0
}

// CreatePool provides a mock function with given fields: ctx, req, opts
func (_m *PoolServiceClient) CreatePool(ctx context.Context, req *poolsv1.CreatePoolRequest, opts ...gax.CallOption) (*poolsv1.Pool, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *poolsv1.Pool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *poolsv1.CreatePoolRequest, ...gax.CallOption) (*poolsv1.Pool, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *poolsv1.CreatePoolRequest, ...gax.CallOption) *poolsv1.Pool); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*poolsv1.Pool)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *poolsv1.CreatePoolRequest, ...gax.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPool provides a mock function with given fields: ctx, req, opts
func (_m *PoolServiceClient) GetPool(ctx context.Context, req *poolsv1.GetPoolRequest, opts ...gax.CallOption) (*poolsv1.Pool, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *poolsv1.Pool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *poolsv1.GetPoolRequest, ...gax.CallOption) (*poolsv1.Pool, error)); ok {
		return rf(ctx, req, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *poolsv1.GetPoolRequest, ...gax.CallOption) *poolsv1.Pool); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*poolsv1.Pool)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *poolsv1.GetPoolRequest, ...gax.CallOption) error); ok {
		r1 = rf(ctx, req, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPools provides a mock function with given fields: ctx, req, opts
func (_m *PoolServiceClient) ListPools(ctx context.Context, req *poolsv1.ListPoolsRequest, opts ...gax.CallOption) v1.PoolIterator {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, req)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 v1.PoolIterator
	if rf, ok := ret.Get(0).(func(context.Context, *poolsv1.ListPoolsRequest, ...gax.CallOption) v1.PoolIterator); ok {
		r0 = rf(ctx, req, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.PoolIterator)
		}
	}

	return r0
}

// NewPoolServiceClient creates a new instance of PoolServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPoolServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *PoolServiceClient {
	mock := &PoolServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
