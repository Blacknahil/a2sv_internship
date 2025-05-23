// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	domain "clean_task_manager_api_tested/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserRepositoryInterface is an autogenerated mock type for the UserRepositoryInterface type
type UserRepositoryInterface struct {
	mock.Mock
}

// Login provides a mock function with given fields: c, loginRequest
func (_m *UserRepositoryInterface) Login(c context.Context, loginRequest domain.LoginRequest) (domain.LoginResponse, error) {
	ret := _m.Called(c, loginRequest)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 domain.LoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.LoginRequest) (domain.LoginResponse, error)); ok {
		return rf(c, loginRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.LoginRequest) domain.LoginResponse); ok {
		r0 = rf(c, loginRequest)
	} else {
		r0 = ret.Get(0).(domain.LoginResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.LoginRequest) error); ok {
		r1 = rf(c, loginRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Promote provides a mock function with given fields: c, userID
func (_m *UserRepositoryInterface) Promote(c context.Context, userID string) error {
	ret := _m.Called(c, userID)

	if len(ret) == 0 {
		panic("no return value specified for Promote")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: c, user
func (_m *UserRepositoryInterface) Register(c context.Context, user domain.User) error {
	ret := _m.Called(c, user)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) error); ok {
		r0 = rf(c, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepositoryInterface creates a new instance of UserRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepositoryInterface {
	mock := &UserRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
