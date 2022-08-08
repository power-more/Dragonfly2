// Code generated by MockGen. DO NOT EDIT.
// Source: searcher.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "d7y.io/api/pkg/apis/manager/v1"
	model "d7y.io/dragonfly/v2/manager/model"
	gomock "github.com/golang/mock/gomock"
)

// MockSearcher is a mock of Searcher interface.
type MockSearcher struct {
	ctrl     *gomock.Controller
	recorder *MockSearcherMockRecorder
}

// MockSearcherMockRecorder is the mock recorder for MockSearcher.
type MockSearcherMockRecorder struct {
	mock *MockSearcher
}

// NewMockSearcher creates a new mock instance.
func NewMockSearcher(ctrl *gomock.Controller) *MockSearcher {
	mock := &MockSearcher{ctrl: ctrl}
	mock.recorder = &MockSearcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearcher) EXPECT() *MockSearcherMockRecorder {
	return m.recorder
}

// FindSchedulerClusters mocks base method.
func (m *MockSearcher) FindSchedulerClusters(arg0 context.Context, arg1 []model.SchedulerCluster, arg2 *v1.ListSchedulersRequest) ([]model.SchedulerCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSchedulerClusters", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.SchedulerCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSchedulerClusters indicates an expected call of FindSchedulerClusters.
func (mr *MockSearcherMockRecorder) FindSchedulerClusters(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSchedulerClusters", reflect.TypeOf((*MockSearcher)(nil).FindSchedulerClusters), arg0, arg1, arg2)
}