// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/nymtech/nym-directory/models"

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// NotifyCocoNodePresence provides a mock function with given fields: up
func (_m *IService) NotifyCocoNodePresence(up models.HostInfo) {
	_m.Called(up)
}

// NotifyMixNodePresence provides a mock function with given fields: up
func (_m *IService) NotifyMixNodePresence(up models.MixHostInfo) {
	_m.Called(up)
}

// Topology provides a mock function with given fields:
func (_m *IService) Topology() models.Topology {
	ret := _m.Called()

	var r0 models.Topology
	if rf, ok := ret.Get(0).(func() models.Topology); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Topology)
	}

	return r0
}