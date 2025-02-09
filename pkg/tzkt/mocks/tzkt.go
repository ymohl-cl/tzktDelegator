// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	tzkt "github.com/ymohl-cl/tzktDelegator/pkg/tzkt"
)

// TzKT is an autogenerated mock type for the TzKT type
type TzKT struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *TzKT) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Listen provides a mock function with given fields:
func (_m *TzKT) Listen() (tzkt.Message, error) {
	ret := _m.Called()

	var r0 tzkt.Message
	var r1 error
	if rf, ok := ret.Get(0).(func() (tzkt.Message, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() tzkt.Message); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(tzkt.Message)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTzKT creates a new instance of TzKT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTzKT(t interface {
	mock.TestingT
	Cleanup(func())
}) *TzKT {
	mock := &TzKT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
