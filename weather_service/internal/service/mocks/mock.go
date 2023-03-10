// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockWeather is a mock of Weather interface.
type MockWeather struct {
	ctrl     *gomock.Controller
	recorder *MockWeatherMockRecorder
}

// MockWeatherMockRecorder is the mock recorder for MockWeather.
type MockWeatherMockRecorder struct {
	mock *MockWeather
}

// NewMockWeather creates a new mock instance.
func NewMockWeather(ctrl *gomock.Controller) *MockWeather {
	mock := &MockWeather{ctrl: ctrl}
	mock.recorder = &MockWeatherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeather) EXPECT() *MockWeatherMockRecorder {
	return m.recorder
}

// GetOpenWeatherFree mocks base method.
func (m *MockWeather) GetOpenWeatherFree(apiKey, city string, date time.Time) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenWeatherFree", apiKey, city, date)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenWeatherFree indicates an expected call of GetOpenWeatherFree.
func (mr *MockWeatherMockRecorder) GetOpenWeatherFree(apiKey, city, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenWeatherFree", reflect.TypeOf((*MockWeather)(nil).GetOpenWeatherFree), apiKey, city, date)
}

// GetOpenWeatherPaid mocks base method.
func (m *MockWeather) GetOpenWeatherPaid(apiKey, city string, date time.Time) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenWeatherPaid", apiKey, city, date)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenWeatherPaid indicates an expected call of GetOpenWeatherPaid.
func (mr *MockWeatherMockRecorder) GetOpenWeatherPaid(apiKey, city, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenWeatherPaid", reflect.TypeOf((*MockWeather)(nil).GetOpenWeatherPaid), apiKey, city, date)
}

// GetWeatherFree mocks base method.
func (m *MockWeather) GetWeatherFree(apiKey, city string, date time.Time) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeatherFree", apiKey, city, date)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWeatherFree indicates an expected call of GetWeatherFree.
func (mr *MockWeatherMockRecorder) GetWeatherFree(apiKey, city, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeatherFree", reflect.TypeOf((*MockWeather)(nil).GetWeatherFree), apiKey, city, date)
}

// MockLocation is a mock of Location interface.
type MockLocation struct {
	ctrl     *gomock.Controller
	recorder *MockLocationMockRecorder
}

// MockLocationMockRecorder is the mock recorder for MockLocation.
type MockLocationMockRecorder struct {
	mock *MockLocation
}

// NewMockLocation creates a new mock instance.
func NewMockLocation(ctrl *gomock.Controller) *MockLocation {
	mock := &MockLocation{ctrl: ctrl}
	mock.recorder = &MockLocationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocation) EXPECT() *MockLocationMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockLocation) Get(ip, apiKey string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ip, apiKey)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLocationMockRecorder) Get(ip, apiKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLocation)(nil).Get), ip, apiKey)
}

// MockIp is a mock of Ip interface.
type MockIp struct {
	ctrl     *gomock.Controller
	recorder *MockIpMockRecorder
}

// MockIpMockRecorder is the mock recorder for MockIp.
type MockIpMockRecorder struct {
	mock *MockIp
}

// NewMockIp creates a new mock instance.
func NewMockIp(ctrl *gomock.Controller) *MockIp {
	mock := &MockIp{ctrl: ctrl}
	mock.recorder = &MockIpMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIp) EXPECT() *MockIpMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIp) Get() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIpMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIp)(nil).Get))
}
