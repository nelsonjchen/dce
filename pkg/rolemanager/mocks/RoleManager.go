// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	awsiface "github.com/Optum/dce/pkg/awsiface"
	errors "github.com/Optum/dce/pkg/errors"

	mock "github.com/stretchr/testify/mock"

	rolemanager "github.com/Optum/dce/pkg/rolemanager"
)

// RoleManager is an autogenerated mock type for the RoleManager type
type RoleManager struct {
	mock.Mock
}

// CreateRoleWithPolicy provides a mock function with given fields: input
func (_m *RoleManager) CreateRoleWithPolicy(input *rolemanager.CreateRoleWithPolicyInput) (*rolemanager.CreateRoleWithPolicyOutput, error) {
	ret := _m.Called(input)

	var r0 *rolemanager.CreateRoleWithPolicyOutput
	if rf, ok := ret.Get(0).(func(*rolemanager.CreateRoleWithPolicyInput) *rolemanager.CreateRoleWithPolicyOutput); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rolemanager.CreateRoleWithPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*rolemanager.CreateRoleWithPolicyInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DestroyRoleWithPolicy provides a mock function with given fields: input
func (_m *RoleManager) DestroyRoleWithPolicy(input *rolemanager.DestroyRoleWithPolicyInput) (*rolemanager.DestroyRoleWithPolicyOutput, *errors.MultiError) {
	ret := _m.Called(input)

	var r0 *rolemanager.DestroyRoleWithPolicyOutput
	if rf, ok := ret.Get(0).(func(*rolemanager.DestroyRoleWithPolicyInput) *rolemanager.DestroyRoleWithPolicyOutput); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rolemanager.DestroyRoleWithPolicyOutput)
		}
	}

	var r1 *errors.MultiError
	if rf, ok := ret.Get(1).(func(*rolemanager.DestroyRoleWithPolicyInput) *errors.MultiError); ok {
		r1 = rf(input)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.MultiError)
		}
	}

	return r0, r1
}

// SetIAMClient provides a mock function with given fields: iamClient
func (_m *RoleManager) SetIAMClient(iamClient awsiface.IAM) {
	_m.Called(iamClient)
}
