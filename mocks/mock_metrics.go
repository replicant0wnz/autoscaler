// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/replicant0wnz/autoscaler/metrics (interfaces: Collector)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockCollector is a mock of Collector interface
type MockCollector struct {
	ctrl     *gomock.Controller
	recorder *MockCollectorMockRecorder
}

// MockCollectorMockRecorder is the mock recorder for MockCollector
type MockCollectorMockRecorder struct {
	mock *MockCollector
}

// NewMockCollector creates a new mock instance
func NewMockCollector(ctrl *gomock.Controller) *MockCollector {
	mock := &MockCollector{ctrl: ctrl}
	mock.recorder = &MockCollectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollector) EXPECT() *MockCollectorMockRecorder {
	return m.recorder
}

// IncrServerCreateError mocks base method
func (m *MockCollector) IncrServerCreateError() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncrServerCreateError")
}

// IncrServerCreateError indicates an expected call of IncrServerCreateError
func (mr *MockCollectorMockRecorder) IncrServerCreateError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrServerCreateError", reflect.TypeOf((*MockCollector)(nil).IncrServerCreateError))
}

// IncrServerInitError mocks base method
func (m *MockCollector) IncrServerInitError() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncrServerInitError")
}

// IncrServerInitError indicates an expected call of IncrServerInitError
func (mr *MockCollectorMockRecorder) IncrServerInitError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrServerInitError", reflect.TypeOf((*MockCollector)(nil).IncrServerInitError))
}

// IncrServerSetupError mocks base method
func (m *MockCollector) IncrServerSetupError() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncrServerSetupError")
}

// IncrServerSetupError indicates an expected call of IncrServerSetupError
func (mr *MockCollectorMockRecorder) IncrServerSetupError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrServerSetupError", reflect.TypeOf((*MockCollector)(nil).IncrServerSetupError))
}

// TrackServerCreateTime mocks base method
func (m *MockCollector) TrackServerCreateTime(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TrackServerCreateTime", arg0)
}

// TrackServerCreateTime indicates an expected call of TrackServerCreateTime
func (mr *MockCollectorMockRecorder) TrackServerCreateTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackServerCreateTime", reflect.TypeOf((*MockCollector)(nil).TrackServerCreateTime), arg0)
}

// TrackServerInitTime mocks base method
func (m *MockCollector) TrackServerInitTime(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TrackServerInitTime", arg0)
}

// TrackServerInitTime indicates an expected call of TrackServerInitTime
func (mr *MockCollectorMockRecorder) TrackServerInitTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackServerInitTime", reflect.TypeOf((*MockCollector)(nil).TrackServerInitTime), arg0)
}

// TrackServerSetupTime mocks base method
func (m *MockCollector) TrackServerSetupTime(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TrackServerSetupTime", arg0)
}

// TrackServerSetupTime indicates an expected call of TrackServerSetupTime
func (mr *MockCollectorMockRecorder) TrackServerSetupTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackServerSetupTime", reflect.TypeOf((*MockCollector)(nil).TrackServerSetupTime), arg0)
}
