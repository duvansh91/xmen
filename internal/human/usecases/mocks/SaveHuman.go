// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	models "github.com/duvansh91/xmen/internal/human/models"
	mock "github.com/stretchr/testify/mock"
)

// SaveHuman is an autogenerated mock type for the SaveHuman type
type SaveHuman struct {
	mock.Mock
}

// Save provides a mock function with given fields: human
func (_m *SaveHuman) Save(human *models.Human) error {
	ret := _m.Called(human)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Human) error); ok {
		r0 = rf(human)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
