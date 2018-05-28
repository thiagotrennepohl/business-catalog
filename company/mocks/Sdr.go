package mocks

import (
	"encoding/csv"

	mock "github.com/stretchr/testify/mock"
)

type Sdr struct {
	mock.Mock
}

// ReadCsvFile provides a mock function with given fields: _a0
func (_m *Sdr) ReadCSV(_a0 string) (*csv.Reader, error) {
	ret := _m.Called(_a0)

	var r0 *csv.Reader
	if rf, ok := ret.Get(0).(func(string) *csv.Reader); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*csv.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Sdr) ParseHeaders(file *csv.Reader) ([]string, error) {
	ret := _m.Called(file)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*csv.Reader) []string); ok {
		r0 = rf(file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*csv.Reader) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Sdr) Extract(file *csv.Reader, headers []string) ([]map[string]interface{}, error) {
	ret := _m.Called(file, headers)

	var r0 []map[string]interface{}
	if rf, ok := ret.Get(0).(func(*csv.Reader, []string) []map[string]interface{}); ok {
		r0 = rf(file, headers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*csv.Reader) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadCSV(string) (*csv.Reader, error)
// ParseHeaders(*csv.Reader) ([]string, error)
// Extract(*csv.Reader, []string) ([]map[string]interface{}, error)
