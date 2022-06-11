// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	models "github.com/duvansh91/xmen/internal/human/models"
	mock "github.com/stretchr/testify/mock"
)

// ValidateIsMutant is an autogenerated mock type for the ValidateIsMutant type
type ValidateIsMutant struct {
	mock.Mock
}

// Validate provides a mock function with given fields: h
func (_m *ValidateIsMutant) Validate(h *models.Human) (bool, error) {
	ret := _m.Called(h)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*models.Human) bool); ok {
		r0 = rf(h)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Human) error); ok {
		r1 = rf(h)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}