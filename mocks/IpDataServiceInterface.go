// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	service "mta-hosting-optimizer/service"

	mock "github.com/stretchr/testify/mock"
)

// IpDataServiceInterface is an autogenerated mock type for the IpDataServiceInterface type
type IpDataServiceInterface struct {
	mock.Mock
}

// GetInefficientServers provides a mock function with given fields: ipData
func (_m *IpDataServiceInterface) GetInefficientServers(ipData *[]service.IpConfig) []string {
	ret := _m.Called(ipData)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*[]service.IpConfig) []string); ok {
		r0 = rf(ipData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetServiceIpData provides a mock function with given fields:
func (_m *IpDataServiceInterface) GetServiceIpData() (*[]service.IpConfig, error) {
	ret := _m.Called()

	var r0 *[]service.IpConfig
	if rf, ok := ret.Get(0).(func() *[]service.IpConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]service.IpConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}