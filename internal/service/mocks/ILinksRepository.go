// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ILinksRepository is an autogenerated mock type for the ILinksRepository type
type ILinksRepository struct {
	mock.Mock
}

// GetOrigin provides a mock function with given fields: ctx, shortUrl
func (_m *ILinksRepository) GetOrigin(ctx context.Context, shortUrl string) (string, error) {
	ret := _m.Called(ctx, shortUrl)

	if len(ret) == 0 {
		panic("no return value specified for GetOrigin")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, shortUrl)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, shortUrl)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetShortUrl provides a mock function with given fields: ctx, url
func (_m *ILinksRepository) GetShortUrl(ctx context.Context, url string) (string, error) {
	ret := _m.Called(ctx, url)

	if len(ret) == 0 {
		panic("no return value specified for GetShortUrl")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, url)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, url)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUrl provides a mock function with given fields: ctx, url, shortUrl
func (_m *ILinksRepository) SaveUrl(ctx context.Context, url string, shortUrl string) (string, error) {
	ret := _m.Called(ctx, url, shortUrl)

	if len(ret) == 0 {
		panic("no return value specified for SaveUrl")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, url, shortUrl)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, url, shortUrl)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, url, shortUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewILinksRepository creates a new instance of ILinksRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewILinksRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ILinksRepository {
	mock := &ILinksRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
