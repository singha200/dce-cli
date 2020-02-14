// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import io "io"
import mock "github.com/stretchr/testify/mock"

// TFTemplater is an autogenerated mock type for the TFTemplater type
type TFTemplater struct {
	mock.Mock
}

// AddVariable provides a mock function with given fields: name, vartype, vardefault
func (_m *TFTemplater) AddVariable(name string, vartype string, vardefault string) error {
	ret := _m.Called(name, vartype, vardefault)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(name, vartype, vardefault)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: w
func (_m *TFTemplater) Write(w io.Writer) error {
	ret := _m.Called(w)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer) error); ok {
		r0 = rf(w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
