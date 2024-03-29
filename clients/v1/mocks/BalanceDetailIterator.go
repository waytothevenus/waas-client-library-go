// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	iterator "google.golang.org/api/iterator"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

// BalanceDetailIterator is an autogenerated mock type for the BalanceDetailIterator type
type BalanceDetailIterator struct {
	mock.Mock
}

// Next provides a mock function with given fields:
func (_m *BalanceDetailIterator) Next() (*v1.BalanceDetail, error) {
	ret := _m.Called()

	var r0 *v1.BalanceDetail
	var r1 error
	if rf, ok := ret.Get(0).(func() (*v1.BalanceDetail, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *v1.BalanceDetail); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.BalanceDetail)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PageInfo provides a mock function with given fields:
func (_m *BalanceDetailIterator) PageInfo() *iterator.PageInfo {
	ret := _m.Called()

	var r0 *iterator.PageInfo
	if rf, ok := ret.Get(0).(func() *iterator.PageInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iterator.PageInfo)
		}
	}

	return r0
}

// Response provides a mock function with given fields:
func (_m *BalanceDetailIterator) Response() *v1.ListBalanceDetailsResponse {
	ret := _m.Called()

	var r0 *v1.ListBalanceDetailsResponse
	if rf, ok := ret.Get(0).(func() *v1.ListBalanceDetailsResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ListBalanceDetailsResponse)
		}
	}

	return r0
}

// NewBalanceDetailIterator creates a new instance of BalanceDetailIterator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBalanceDetailIterator(t interface {
	mock.TestingT
	Cleanup(func())
}) *BalanceDetailIterator {
	mock := &BalanceDetailIterator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
