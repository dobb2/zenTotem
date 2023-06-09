// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	entity "github.com/dobb2/zenTotem/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// Storer is an autogenerated mock type for the Storer type
type Storer struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *Storer) Create(user entity.User) (entity.User, error) {
	ret := _m.Called(user)

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.User) (entity.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(entity.User) entity.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStorer interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorer creates a new instance of Storer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorer(t mockConstructorTestingTNewStorer) *Storer {
	mock := &Storer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
