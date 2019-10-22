// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	v1alpha1 "github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// Resume provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Provider) Resume(_a0 *v1alpha1.AnalysisRun, _a1 v1alpha1.Metric, _a2 []v1alpha1.Argument, _a3 v1alpha1.Measurement) v1alpha1.Measurement {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 v1alpha1.Measurement
	if rf, ok := ret.Get(0).(func(*v1alpha1.AnalysisRun, v1alpha1.Metric, []v1alpha1.Argument, v1alpha1.Measurement) v1alpha1.Measurement); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(v1alpha1.Measurement)
	}

	return r0
}

// Run provides a mock function with given fields: _a0, _a1, _a2
func (_m *Provider) Run(_a0 *v1alpha1.AnalysisRun, _a1 v1alpha1.Metric, _a2 []v1alpha1.Argument) v1alpha1.Measurement {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 v1alpha1.Measurement
	if rf, ok := ret.Get(0).(func(*v1alpha1.AnalysisRun, v1alpha1.Metric, []v1alpha1.Argument) v1alpha1.Measurement); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(v1alpha1.Measurement)
	}

	return r0
}

// Terminate provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Provider) Terminate(_a0 *v1alpha1.AnalysisRun, _a1 v1alpha1.Metric, _a2 []v1alpha1.Argument, _a3 v1alpha1.Measurement) v1alpha1.Measurement {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 v1alpha1.Measurement
	if rf, ok := ret.Get(0).(func(*v1alpha1.AnalysisRun, v1alpha1.Metric, []v1alpha1.Argument, v1alpha1.Measurement) v1alpha1.Measurement); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(v1alpha1.Measurement)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *Provider) Type() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}