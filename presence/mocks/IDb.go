// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/nymtech/nym-directory/models"

// IDb is an autogenerated mock type for the IDb type
type IDb struct {
	mock.Mock
}

// AddCoco provides a mock function with given fields: _a0
func (_m *IDb) AddCoco(_a0 models.CocoPresence) {
	_m.Called(_a0)
}

// AddGateway provides a mock function with given fields: _a0
func (_m *IDb) AddGateway(_a0 models.GatewayPresence) {
	_m.Called(_a0)
}

// AddMix provides a mock function with given fields: _a0
func (_m *IDb) AddMix(_a0 models.MixNodePresence) {
	_m.Called(_a0)
}

// AddMixProvider provides a mock function with given fields: _a0
func (_m *IDb) AddMixProvider(_a0 models.MixProviderPresence) {
	_m.Called(_a0)
}

// Allow provides a mock function with given fields: _a0
func (_m *IDb) Allow(_a0 string) {
	_m.Called(_a0)
}

// Disallow provides a mock function with given fields: _a0
func (_m *IDb) Disallow(_a0 string) {
	_m.Called(_a0)
}

// ListDisallowed provides a mock function with given fields:
func (_m *IDb) ListDisallowed() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Topology provides a mock function with given fields:
func (_m *IDb) Topology() models.Topology {
	ret := _m.Called()

	var r0 models.Topology
	if rf, ok := ret.Get(0).(func() models.Topology); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Topology)
	}

	return r0
}
