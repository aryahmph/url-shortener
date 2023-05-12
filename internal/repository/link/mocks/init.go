// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	link "github.com/aryahmph/url-shortener/internal/model/link"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// CreateLink provides a mock function with given fields: ctx, arg
func (_m *Repository) CreateLink(ctx context.Context, arg link.Link) error {
	ret := _m.Called(ctx, arg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, link.Link) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_CreateLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateLink'
type Repository_CreateLink_Call struct {
	*mock.Call
}

// CreateLink is a helper method to define mock.On call
//   - ctx context.Context
//   - arg link.Link
func (_e *Repository_Expecter) CreateLink(ctx interface{}, arg interface{}) *Repository_CreateLink_Call {
	return &Repository_CreateLink_Call{Call: _e.mock.On("CreateLink", ctx, arg)}
}

func (_c *Repository_CreateLink_Call) Run(run func(ctx context.Context, arg link.Link)) *Repository_CreateLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(link.Link))
	})
	return _c
}

func (_c *Repository_CreateLink_Call) Return(_a0 error) *Repository_CreateLink_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_CreateLink_Call) RunAndReturn(run func(context.Context, link.Link) error) *Repository_CreateLink_Call {
	_c.Call.Return(run)
	return _c
}

// GetLink provides a mock function with given fields: ctx, id
func (_m *Repository) GetLink(ctx context.Context, id string) (link.Link, error) {
	ret := _m.Called(ctx, id)

	var r0 link.Link
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (link.Link, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) link.Link); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(link.Link)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLink'
type Repository_GetLink_Call struct {
	*mock.Call
}

// GetLink is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Repository_Expecter) GetLink(ctx interface{}, id interface{}) *Repository_GetLink_Call {
	return &Repository_GetLink_Call{Call: _e.mock.On("GetLink", ctx, id)}
}

func (_c *Repository_GetLink_Call) Run(run func(ctx context.Context, id string)) *Repository_GetLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Repository_GetLink_Call) Return(_a0 link.Link, _a1 error) *Repository_GetLink_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetLink_Call) RunAndReturn(run func(context.Context, string) (link.Link, error)) *Repository_GetLink_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}