// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/camsas/poseidon/pkg/firmament (interfaces: FirmamentSchedulerClient)

package mock_firmament

import (
	gomock "github.com/golang/mock/gomock"
	firmament "github.com/kubernetes-sigs/poseidon/pkg/firmament"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Mock of FirmamentSchedulerClient interface
type MockFirmamentSchedulerClient struct {
	ctrl     *gomock.Controller
	recorder *_MockFirmamentSchedulerClientRecorder
}

// Recorder for MockFirmamentSchedulerClient (not exported)
type _MockFirmamentSchedulerClientRecorder struct {
	mock *MockFirmamentSchedulerClient
}

func NewMockFirmamentSchedulerClient(ctrl *gomock.Controller) *MockFirmamentSchedulerClient {
	mock := &MockFirmamentSchedulerClient{ctrl: ctrl}
	mock.recorder = &_MockFirmamentSchedulerClientRecorder{mock}
	return mock
}

func (_m *MockFirmamentSchedulerClient) EXPECT() *_MockFirmamentSchedulerClientRecorder {
	return _m.recorder
}

func (_m *MockFirmamentSchedulerClient) AddNodeStats(_param0 context.Context, _param1 *firmament.ResourceStats, _param2 ...grpc.CallOption) (*firmament.ResourceStatsResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "AddNodeStats", _s...)
	ret0, _ := ret[0].(*firmament.ResourceStatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFirmamentSchedulerClientRecorder) AddNodeStats(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddNodeStats", _s...)
}

func (_m *MockFirmamentSchedulerClient) AddTaskStats(_param0 context.Context, _param1 *firmament.TaskStats, _param2 ...grpc.CallOption) (*firmament.TaskStatsResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "AddTaskStats", _s...)
	ret0, _ := ret[0].(*firmament.TaskStatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFirmamentSchedulerClientRecorder) AddTaskStats(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddTaskStats", _s...)
}

func (_m *MockFirmamentSchedulerClient) NodeAdded(_param0 context.Context, _param1 *firmament.ResourceTopologyNodeDescriptor, _param2 ...grpc.CallOption) (*firmament.NodeAddedResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "NodeAdded", _s...)
	ret0, _ := ret[0].(*firmament.NodeAddedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFirmamentSchedulerClientRecorder) NodeAdded(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeAdded", _s...)
}

func (_m *MockFirmamentSchedulerClient) NodeFailed(_param0 context.Context, _param1 *firmament.ResourceUID, _param2 ...grpc.CallOption) (*firmament.NodeFailedResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "NodeFailed", _s...)
	ret0, _ := ret[0].(*firmament.NodeFailedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFirmamentSchedulerClientRecorder) NodeFailed(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeFailed", _s...)
}

func (_m *MockFirmamentSchedulerClient) NodeRemoved(_param0 context.Context, _param1 *firmament.ResourceUID, _param2 ...grpc.CallOption) (*firmament.NodeRemovedResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "NodeRemoved", _s...)
	ret0, _ := ret[0].(*firmament.NodeRemovedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFirmamentSchedulerClientRecorder) NodeRemoved(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeRemoved", _s...)
}

func (_m *MockFirmamentSchedulerClient) NodeUpdated(_param0 context.Context, _param1 *firmament.ResourceTopologyNodeDescriptor, _param2 ...grpc.CallOption) (*firmament.NodeUpdatedResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "NodeUpdated", _s...)
	ret0, _ := ret[0].(*firmament.NodeUpdatedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFirmamentSchedulerClientRecorder) NodeUpdated(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeUpdated", _s...)
}

func (_m *MockFirmamentSchedulerClient) Schedule(_param0 context.Context, _param1 *firmament.ScheduleRequest, _param2 ...grpc.CallOption) (*firmament.SchedulingDeltas, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Schedule", _s...)
	ret0, _ := ret[0].(*firmament.SchedulingDeltas)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFirmamentSchedulerClientRecorder) Schedule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Schedule", _s...)
}

func (_m *MockFirmamentSchedulerClient) TaskCompleted(_param0 context.Context, _param1 *firmament.TaskUID, _param2 ...grpc.CallOption) (*firmament.TaskCompletedResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "TaskCompleted", _s...)
	ret0, _ := ret[0].(*firmament.TaskCompletedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFirmamentSchedulerClientRecorder) TaskCompleted(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskCompleted", _s...)
}

func (_m *MockFirmamentSchedulerClient) TaskFailed(_param0 context.Context, _param1 *firmament.TaskUID, _param2 ...grpc.CallOption) (*firmament.TaskFailedResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "TaskFailed", _s...)
	ret0, _ := ret[0].(*firmament.TaskFailedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFirmamentSchedulerClientRecorder) TaskFailed(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskFailed", _s...)
}

func (_m *MockFirmamentSchedulerClient) TaskRemoved(_param0 context.Context, _param1 *firmament.TaskUID, _param2 ...grpc.CallOption) (*firmament.TaskRemovedResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "TaskRemoved", _s...)
	ret0, _ := ret[0].(*firmament.TaskRemovedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFirmamentSchedulerClientRecorder) TaskRemoved(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskRemoved", _s...)
}

func (_m *MockFirmamentSchedulerClient) TaskSubmitted(_param0 context.Context, _param1 *firmament.TaskDescription, _param2 ...grpc.CallOption) (*firmament.TaskSubmittedResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "TaskSubmitted", _s...)
	ret0, _ := ret[0].(*firmament.TaskSubmittedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFirmamentSchedulerClientRecorder) TaskSubmitted(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskSubmitted", _s...)
}

func (_m *MockFirmamentSchedulerClient) TaskUpdated(_param0 context.Context, _param1 *firmament.TaskDescription, _param2 ...grpc.CallOption) (*firmament.TaskUpdatedResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "TaskUpdated", _s...)
	ret0, _ := ret[0].(*firmament.TaskUpdatedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFirmamentSchedulerClientRecorder) TaskUpdated(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskUpdated", _s...)
}
