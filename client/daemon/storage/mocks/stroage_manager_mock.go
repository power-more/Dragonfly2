// Code generated by MockGen. DO NOT EDIT.
// Source: storage_manager.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	v1 "d7y.io/api/pkg/apis/common/v1"
	storage "d7y.io/dragonfly/v2/client/daemon/storage"
	util "d7y.io/dragonfly/v2/client/util"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskStorageDriver is a mock of TaskStorageDriver interface.
type MockTaskStorageDriver struct {
	ctrl     *gomock.Controller
	recorder *MockTaskStorageDriverMockRecorder
}

// MockTaskStorageDriverMockRecorder is the mock recorder for MockTaskStorageDriver.
type MockTaskStorageDriverMockRecorder struct {
	mock *MockTaskStorageDriver
}

// NewMockTaskStorageDriver creates a new mock instance.
func NewMockTaskStorageDriver(ctrl *gomock.Controller) *MockTaskStorageDriver {
	mock := &MockTaskStorageDriver{ctrl: ctrl}
	mock.recorder = &MockTaskStorageDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskStorageDriver) EXPECT() *MockTaskStorageDriverMockRecorder {
	return m.recorder
}

// GetExtendAttribute mocks base method.
func (m *MockTaskStorageDriver) GetExtendAttribute(ctx context.Context, req *storage.PeerTaskMetadata) (*v1.ExtendAttribute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtendAttribute", ctx, req)
	ret0, _ := ret[0].(*v1.ExtendAttribute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtendAttribute indicates an expected call of GetExtendAttribute.
func (mr *MockTaskStorageDriverMockRecorder) GetExtendAttribute(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtendAttribute", reflect.TypeOf((*MockTaskStorageDriver)(nil).GetExtendAttribute), ctx, req)
}

// GetPieces mocks base method.
func (m *MockTaskStorageDriver) GetPieces(ctx context.Context, req *v1.PieceTaskRequest) (*v1.PiecePacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPieces", ctx, req)
	ret0, _ := ret[0].(*v1.PiecePacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPieces indicates an expected call of GetPieces.
func (mr *MockTaskStorageDriverMockRecorder) GetPieces(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPieces", reflect.TypeOf((*MockTaskStorageDriver)(nil).GetPieces), ctx, req)
}

// GetTotalPieces mocks base method.
func (m *MockTaskStorageDriver) GetTotalPieces(ctx context.Context, req *storage.PeerTaskMetadata) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalPieces", ctx, req)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalPieces indicates an expected call of GetTotalPieces.
func (mr *MockTaskStorageDriverMockRecorder) GetTotalPieces(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalPieces", reflect.TypeOf((*MockTaskStorageDriver)(nil).GetTotalPieces), ctx, req)
}

// IsInvalid mocks base method.
func (m *MockTaskStorageDriver) IsInvalid(req *storage.PeerTaskMetadata) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInvalid", req)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInvalid indicates an expected call of IsInvalid.
func (mr *MockTaskStorageDriverMockRecorder) IsInvalid(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInvalid", reflect.TypeOf((*MockTaskStorageDriver)(nil).IsInvalid), req)
}

// ReadAllPieces mocks base method.
func (m *MockTaskStorageDriver) ReadAllPieces(ctx context.Context, req *storage.ReadAllPiecesRequest) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllPieces", ctx, req)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllPieces indicates an expected call of ReadAllPieces.
func (mr *MockTaskStorageDriverMockRecorder) ReadAllPieces(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllPieces", reflect.TypeOf((*MockTaskStorageDriver)(nil).ReadAllPieces), ctx, req)
}

// ReadPiece mocks base method.
func (m *MockTaskStorageDriver) ReadPiece(ctx context.Context, req *storage.ReadPieceRequest) (io.Reader, io.Closer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPiece", ctx, req)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(io.Closer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadPiece indicates an expected call of ReadPiece.
func (mr *MockTaskStorageDriverMockRecorder) ReadPiece(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPiece", reflect.TypeOf((*MockTaskStorageDriver)(nil).ReadPiece), ctx, req)
}

// Store mocks base method.
func (m *MockTaskStorageDriver) Store(ctx context.Context, req *storage.StoreRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockTaskStorageDriverMockRecorder) Store(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockTaskStorageDriver)(nil).Store), ctx, req)
}

// UpdateTask mocks base method.
func (m *MockTaskStorageDriver) UpdateTask(ctx context.Context, req *storage.UpdateTaskRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockTaskStorageDriverMockRecorder) UpdateTask(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockTaskStorageDriver)(nil).UpdateTask), ctx, req)
}

// ValidateDigest mocks base method.
func (m *MockTaskStorageDriver) ValidateDigest(req *storage.PeerTaskMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateDigest", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateDigest indicates an expected call of ValidateDigest.
func (mr *MockTaskStorageDriverMockRecorder) ValidateDigest(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateDigest", reflect.TypeOf((*MockTaskStorageDriver)(nil).ValidateDigest), req)
}

// WritePiece mocks base method.
func (m *MockTaskStorageDriver) WritePiece(ctx context.Context, req *storage.WritePieceRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WritePiece", ctx, req)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WritePiece indicates an expected call of WritePiece.
func (mr *MockTaskStorageDriverMockRecorder) WritePiece(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePiece", reflect.TypeOf((*MockTaskStorageDriver)(nil).WritePiece), ctx, req)
}

// MockReclaimer is a mock of Reclaimer interface.
type MockReclaimer struct {
	ctrl     *gomock.Controller
	recorder *MockReclaimerMockRecorder
}

// MockReclaimerMockRecorder is the mock recorder for MockReclaimer.
type MockReclaimerMockRecorder struct {
	mock *MockReclaimer
}

// NewMockReclaimer creates a new mock instance.
func NewMockReclaimer(ctrl *gomock.Controller) *MockReclaimer {
	mock := &MockReclaimer{ctrl: ctrl}
	mock.recorder = &MockReclaimerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReclaimer) EXPECT() *MockReclaimerMockRecorder {
	return m.recorder
}

// CanReclaim mocks base method.
func (m *MockReclaimer) CanReclaim() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanReclaim")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanReclaim indicates an expected call of CanReclaim.
func (mr *MockReclaimerMockRecorder) CanReclaim() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanReclaim", reflect.TypeOf((*MockReclaimer)(nil).CanReclaim))
}

// MarkReclaim mocks base method.
func (m *MockReclaimer) MarkReclaim() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkReclaim")
}

// MarkReclaim indicates an expected call of MarkReclaim.
func (mr *MockReclaimerMockRecorder) MarkReclaim() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkReclaim", reflect.TypeOf((*MockReclaimer)(nil).MarkReclaim))
}

// Reclaim mocks base method.
func (m *MockReclaimer) Reclaim() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reclaim")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reclaim indicates an expected call of Reclaim.
func (mr *MockReclaimerMockRecorder) Reclaim() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reclaim", reflect.TypeOf((*MockReclaimer)(nil).Reclaim))
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Alive mocks base method.
func (m *MockManager) Alive(alive time.Duration) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alive", alive)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Alive indicates an expected call of Alive.
func (mr *MockManagerMockRecorder) Alive(alive interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alive", reflect.TypeOf((*MockManager)(nil).Alive), alive)
}

// CleanUp mocks base method.
func (m *MockManager) CleanUp() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CleanUp")
}

// CleanUp indicates an expected call of CleanUp.
func (mr *MockManagerMockRecorder) CleanUp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanUp", reflect.TypeOf((*MockManager)(nil).CleanUp))
}

// FindCompletedSubTask mocks base method.
func (m *MockManager) FindCompletedSubTask(taskID string) *storage.ReusePeerTask {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCompletedSubTask", taskID)
	ret0, _ := ret[0].(*storage.ReusePeerTask)
	return ret0
}

// FindCompletedSubTask indicates an expected call of FindCompletedSubTask.
func (mr *MockManagerMockRecorder) FindCompletedSubTask(taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCompletedSubTask", reflect.TypeOf((*MockManager)(nil).FindCompletedSubTask), taskID)
}

// FindCompletedTask mocks base method.
func (m *MockManager) FindCompletedTask(taskID string) *storage.ReusePeerTask {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCompletedTask", taskID)
	ret0, _ := ret[0].(*storage.ReusePeerTask)
	return ret0
}

// FindCompletedTask indicates an expected call of FindCompletedTask.
func (mr *MockManagerMockRecorder) FindCompletedTask(taskID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCompletedTask", reflect.TypeOf((*MockManager)(nil).FindCompletedTask), taskID)
}

// FindPartialCompletedTask mocks base method.
func (m *MockManager) FindPartialCompletedTask(taskID string, rg *util.Range) *storage.ReusePeerTask {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPartialCompletedTask", taskID, rg)
	ret0, _ := ret[0].(*storage.ReusePeerTask)
	return ret0
}

// FindPartialCompletedTask indicates an expected call of FindPartialCompletedTask.
func (mr *MockManagerMockRecorder) FindPartialCompletedTask(taskID, rg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPartialCompletedTask", reflect.TypeOf((*MockManager)(nil).FindPartialCompletedTask), taskID, rg)
}

// GetExtendAttribute mocks base method.
func (m *MockManager) GetExtendAttribute(ctx context.Context, req *storage.PeerTaskMetadata) (*v1.ExtendAttribute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtendAttribute", ctx, req)
	ret0, _ := ret[0].(*v1.ExtendAttribute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtendAttribute indicates an expected call of GetExtendAttribute.
func (mr *MockManagerMockRecorder) GetExtendAttribute(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtendAttribute", reflect.TypeOf((*MockManager)(nil).GetExtendAttribute), ctx, req)
}

// GetPieces mocks base method.
func (m *MockManager) GetPieces(ctx context.Context, req *v1.PieceTaskRequest) (*v1.PiecePacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPieces", ctx, req)
	ret0, _ := ret[0].(*v1.PiecePacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPieces indicates an expected call of GetPieces.
func (mr *MockManagerMockRecorder) GetPieces(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPieces", reflect.TypeOf((*MockManager)(nil).GetPieces), ctx, req)
}

// GetTotalPieces mocks base method.
func (m *MockManager) GetTotalPieces(ctx context.Context, req *storage.PeerTaskMetadata) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalPieces", ctx, req)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalPieces indicates an expected call of GetTotalPieces.
func (mr *MockManagerMockRecorder) GetTotalPieces(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalPieces", reflect.TypeOf((*MockManager)(nil).GetTotalPieces), ctx, req)
}

// IsInvalid mocks base method.
func (m *MockManager) IsInvalid(req *storage.PeerTaskMetadata) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInvalid", req)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInvalid indicates an expected call of IsInvalid.
func (mr *MockManagerMockRecorder) IsInvalid(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInvalid", reflect.TypeOf((*MockManager)(nil).IsInvalid), req)
}

// Keep mocks base method.
func (m *MockManager) Keep() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Keep")
}

// Keep indicates an expected call of Keep.
func (mr *MockManagerMockRecorder) Keep() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keep", reflect.TypeOf((*MockManager)(nil).Keep))
}

// ReadAllPieces mocks base method.
func (m *MockManager) ReadAllPieces(ctx context.Context, req *storage.ReadAllPiecesRequest) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllPieces", ctx, req)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllPieces indicates an expected call of ReadAllPieces.
func (mr *MockManagerMockRecorder) ReadAllPieces(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllPieces", reflect.TypeOf((*MockManager)(nil).ReadAllPieces), ctx, req)
}

// ReadPiece mocks base method.
func (m *MockManager) ReadPiece(ctx context.Context, req *storage.ReadPieceRequest) (io.Reader, io.Closer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPiece", ctx, req)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(io.Closer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadPiece indicates an expected call of ReadPiece.
func (mr *MockManagerMockRecorder) ReadPiece(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPiece", reflect.TypeOf((*MockManager)(nil).ReadPiece), ctx, req)
}

// RegisterSubTask mocks base method.
func (m *MockManager) RegisterSubTask(ctx context.Context, req *storage.RegisterSubTaskRequest) (storage.TaskStorageDriver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterSubTask", ctx, req)
	ret0, _ := ret[0].(storage.TaskStorageDriver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterSubTask indicates an expected call of RegisterSubTask.
func (mr *MockManagerMockRecorder) RegisterSubTask(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterSubTask", reflect.TypeOf((*MockManager)(nil).RegisterSubTask), ctx, req)
}

// RegisterTask mocks base method.
func (m *MockManager) RegisterTask(ctx context.Context, req *storage.RegisterTaskRequest) (storage.TaskStorageDriver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterTask", ctx, req)
	ret0, _ := ret[0].(storage.TaskStorageDriver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterTask indicates an expected call of RegisterTask.
func (mr *MockManagerMockRecorder) RegisterTask(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTask", reflect.TypeOf((*MockManager)(nil).RegisterTask), ctx, req)
}

// Store mocks base method.
func (m *MockManager) Store(ctx context.Context, req *storage.StoreRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockManagerMockRecorder) Store(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockManager)(nil).Store), ctx, req)
}

// UnregisterTask mocks base method.
func (m *MockManager) UnregisterTask(ctx context.Context, req storage.CommonTaskRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterTask", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterTask indicates an expected call of UnregisterTask.
func (mr *MockManagerMockRecorder) UnregisterTask(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterTask", reflect.TypeOf((*MockManager)(nil).UnregisterTask), ctx, req)
}

// UpdateTask mocks base method.
func (m *MockManager) UpdateTask(ctx context.Context, req *storage.UpdateTaskRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockManagerMockRecorder) UpdateTask(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockManager)(nil).UpdateTask), ctx, req)
}

// ValidateDigest mocks base method.
func (m *MockManager) ValidateDigest(req *storage.PeerTaskMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateDigest", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateDigest indicates an expected call of ValidateDigest.
func (mr *MockManagerMockRecorder) ValidateDigest(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateDigest", reflect.TypeOf((*MockManager)(nil).ValidateDigest), req)
}

// WritePiece mocks base method.
func (m *MockManager) WritePiece(ctx context.Context, req *storage.WritePieceRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WritePiece", ctx, req)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WritePiece indicates an expected call of WritePiece.
func (mr *MockManagerMockRecorder) WritePiece(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePiece", reflect.TypeOf((*MockManager)(nil).WritePiece), ctx, req)
}