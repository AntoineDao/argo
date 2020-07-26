// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// Gatekeeper is an autogenerated mock type for the Gatekeeper type
type Gatekeeper struct {
	mock.Mock
}

// Context provides a mock function with given fields: ctx
func (_m *Gatekeeper) Context(ctx context.Context) (context.Context, error) {
	ret := _m.Called(ctx)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context) context.Context); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StreamServerInterceptor provides a mock function with given fields:
func (_m *Gatekeeper) StreamServerInterceptor() grpc.StreamServerInterceptor {
	ret := _m.Called()

	var r0 grpc.StreamServerInterceptor
	if rf, ok := ret.Get(0).(func() grpc.StreamServerInterceptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grpc.StreamServerInterceptor)
		}
	}

	return r0
}

// UnaryServerInterceptor provides a mock function with given fields:
func (_m *Gatekeeper) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	ret := _m.Called()

	var r0 grpc.UnaryServerInterceptor
	if rf, ok := ret.Get(0).(func() grpc.UnaryServerInterceptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grpc.UnaryServerInterceptor)
		}
	}

	return r0
}
