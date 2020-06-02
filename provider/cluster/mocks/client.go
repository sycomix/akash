// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cluster "github.com/ovrclk/akash/provider/cluster"

	manifest "github.com/ovrclk/akash/manifest"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ovrclk/akash/x/market/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Deploy provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) Deploy(_a0 context.Context, _a1 types.LeaseID, _a2 *manifest.Group) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.LeaseID, *manifest.Group) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deployments provides a mock function with given fields: _a0
func (_m *Client) Deployments(_a0 context.Context) ([]cluster.Deployment, error) {
	ret := _m.Called(_a0)

	var r0 []cluster.Deployment
	if rf, ok := ret.Get(0).(func(context.Context) []cluster.Deployment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cluster.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Inventory provides a mock function with given fields: _a0
func (_m *Client) Inventory(_a0 context.Context) ([]cluster.Node, error) {
	ret := _m.Called(_a0)

	var r0 []cluster.Node
	if rf, ok := ret.Get(0).(func(context.Context) []cluster.Node); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cluster.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LeaseStatus provides a mock function with given fields: _a0, _a1
func (_m *Client) LeaseStatus(_a0 context.Context, _a1 types.LeaseID) (*cluster.LeaseStatus, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *cluster.LeaseStatus
	if rf, ok := ret.Get(0).(func(context.Context, types.LeaseID) *cluster.LeaseStatus); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.LeaseStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.LeaseID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceLogs provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Client) ServiceLogs(_a0 context.Context, _a1 types.LeaseID, _a2 int64, _a3 bool) ([]*cluster.ServiceLog, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []*cluster.ServiceLog
	if rf, ok := ret.Get(0).(func(context.Context, types.LeaseID, int64, bool) []*cluster.ServiceLog); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*cluster.ServiceLog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.LeaseID, int64, bool) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceStatus provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) ServiceStatus(_a0 context.Context, _a1 types.LeaseID, _a2 string) (*cluster.ServiceStatus, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *cluster.ServiceStatus
	if rf, ok := ret.Get(0).(func(context.Context, types.LeaseID, string) *cluster.ServiceStatus); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.ServiceStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.LeaseID, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeardownLease provides a mock function with given fields: _a0, _a1
func (_m *Client) TeardownLease(_a0 context.Context, _a1 types.LeaseID) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.LeaseID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}