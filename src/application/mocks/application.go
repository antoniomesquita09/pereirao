// Code generated by MockGen. DO NOT EDIT.
// Source: src/application/person.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	application "pereirao/src/application"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockPersonInterface is a mock of PersonInterface interface.
type MockPersonInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPersonInterfaceMockRecorder
}

// MockPersonInterfaceMockRecorder is the mock recorder for MockPersonInterface.
type MockPersonInterfaceMockRecorder struct {
	mock *MockPersonInterface
}

// NewMockPersonInterface creates a new mock instance.
func NewMockPersonInterface(ctrl *gomock.Controller) *MockPersonInterface {
	mock := &MockPersonInterface{ctrl: ctrl}
	mock.recorder = &MockPersonInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonInterface) EXPECT() *MockPersonInterfaceMockRecorder {
	return m.recorder
}

// GetCpf mocks base method.
func (m *MockPersonInterface) GetCpf() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCpf")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCpf indicates an expected call of GetCpf.
func (mr *MockPersonInterfaceMockRecorder) GetCpf() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCpf", reflect.TypeOf((*MockPersonInterface)(nil).GetCpf))
}

// GetEmail mocks base method.
func (m *MockPersonInterface) GetEmail() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmail")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEmail indicates an expected call of GetEmail.
func (mr *MockPersonInterfaceMockRecorder) GetEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmail", reflect.TypeOf((*MockPersonInterface)(nil).GetEmail))
}

// GetID mocks base method.
func (m *MockPersonInterface) GetID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockPersonInterfaceMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockPersonInterface)(nil).GetID))
}

// GetName mocks base method.
func (m *MockPersonInterface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockPersonInterfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockPersonInterface)(nil).GetName))
}

// IsValid mocks base method.
func (m *MockPersonInterface) IsValid() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValid indicates an expected call of IsValid.
func (mr *MockPersonInterfaceMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockPersonInterface)(nil).IsValid))
}

// MockPersonServiceInterface is a mock of PersonServiceInterface interface.
type MockPersonServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPersonServiceInterfaceMockRecorder
}

// MockPersonServiceInterfaceMockRecorder is the mock recorder for MockPersonServiceInterface.
type MockPersonServiceInterfaceMockRecorder struct {
	mock *MockPersonServiceInterface
}

// NewMockPersonServiceInterface creates a new mock instance.
func NewMockPersonServiceInterface(ctrl *gomock.Controller) *MockPersonServiceInterface {
	mock := &MockPersonServiceInterface{ctrl: ctrl}
	mock.recorder = &MockPersonServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonServiceInterface) EXPECT() *MockPersonServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPersonServiceInterface) Create(name, email, cpf string) (application.PersonInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, email, cpf)
	ret0, _ := ret[0].(application.PersonInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPersonServiceInterfaceMockRecorder) Create(name, email, cpf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPersonServiceInterface)(nil).Create), name, email, cpf)
}

// GetByCpf mocks base method.
func (m *MockPersonServiceInterface) GetByCpf(cpf string) (application.PersonInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCpf", cpf)
	ret0, _ := ret[0].(application.PersonInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCpf indicates an expected call of GetByCpf.
func (mr *MockPersonServiceInterfaceMockRecorder) GetByCpf(cpf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCpf", reflect.TypeOf((*MockPersonServiceInterface)(nil).GetByCpf), cpf)
}

// GetByID mocks base method.
func (m *MockPersonServiceInterface) GetByID(id uuid.UUID) (application.PersonInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(application.PersonInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockPersonServiceInterfaceMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockPersonServiceInterface)(nil).GetByID), id)
}

// MockPersonReader is a mock of PersonReader interface.
type MockPersonReader struct {
	ctrl     *gomock.Controller
	recorder *MockPersonReaderMockRecorder
}

// MockPersonReaderMockRecorder is the mock recorder for MockPersonReader.
type MockPersonReaderMockRecorder struct {
	mock *MockPersonReader
}

// NewMockPersonReader creates a new mock instance.
func NewMockPersonReader(ctrl *gomock.Controller) *MockPersonReader {
	mock := &MockPersonReader{ctrl: ctrl}
	mock.recorder = &MockPersonReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonReader) EXPECT() *MockPersonReaderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockPersonReader) Get(id uuid.UUID) (application.PersonInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.PersonInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPersonReaderMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPersonReader)(nil).Get), id)
}

// MockPersonWriter is a mock of PersonWriter interface.
type MockPersonWriter struct {
	ctrl     *gomock.Controller
	recorder *MockPersonWriterMockRecorder
}

// MockPersonWriterMockRecorder is the mock recorder for MockPersonWriter.
type MockPersonWriterMockRecorder struct {
	mock *MockPersonWriter
}

// NewMockPersonWriter creates a new mock instance.
func NewMockPersonWriter(ctrl *gomock.Controller) *MockPersonWriter {
	mock := &MockPersonWriter{ctrl: ctrl}
	mock.recorder = &MockPersonWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonWriter) EXPECT() *MockPersonWriterMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockPersonWriter) Save(person application.PersonInterface) (application.PersonInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", person)
	ret0, _ := ret[0].(application.PersonInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockPersonWriterMockRecorder) Save(person interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockPersonWriter)(nil).Save), person)
}

// MockPersonPersistenceInterface is a mock of PersonPersistenceInterface interface.
type MockPersonPersistenceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPersonPersistenceInterfaceMockRecorder
}

// MockPersonPersistenceInterfaceMockRecorder is the mock recorder for MockPersonPersistenceInterface.
type MockPersonPersistenceInterfaceMockRecorder struct {
	mock *MockPersonPersistenceInterface
}

// NewMockPersonPersistenceInterface creates a new mock instance.
func NewMockPersonPersistenceInterface(ctrl *gomock.Controller) *MockPersonPersistenceInterface {
	mock := &MockPersonPersistenceInterface{ctrl: ctrl}
	mock.recorder = &MockPersonPersistenceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonPersistenceInterface) EXPECT() *MockPersonPersistenceInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockPersonPersistenceInterface) Get(id uuid.UUID) (application.PersonInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.PersonInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPersonPersistenceInterfaceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPersonPersistenceInterface)(nil).Get), id)
}

// Save mocks base method.
func (m *MockPersonPersistenceInterface) Save(person application.PersonInterface) (application.PersonInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", person)
	ret0, _ := ret[0].(application.PersonInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockPersonPersistenceInterfaceMockRecorder) Save(person interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockPersonPersistenceInterface)(nil).Save), person)
}
