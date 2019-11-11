// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gapic-generator. DO NOT EDIT.

package agentendpoint

import (
	agentendpointpb "google.golang.org/genproto/googleapis/cloud/osconfig/agentendpoint/v1alpha1"
)

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/api/option"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	gstatus "google.golang.org/grpc/status"
)

var _ = io.EOF
var _ = ptypes.MarshalAny
var _ status.Status

type mockAgentEndpointServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	agentendpointpb.AgentEndpointServiceServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockAgentEndpointServer) ReceiveTaskNotification(req *agentendpointpb.ReceiveTaskNotificationRequest, stream agentendpointpb.AgentEndpointService_ReceiveTaskNotificationServer) error {
	md, _ := metadata.FromIncomingContext(stream.Context())
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return s.err
	}
	for _, v := range s.resps {
		if err := stream.Send(v.(*agentendpointpb.ReceiveTaskNotificationResponse)); err != nil {
			return err
		}
	}
	return nil
}

func (s *mockAgentEndpointServer) ReportTaskStart(ctx context.Context, req *agentendpointpb.ReportTaskStartRequest) (*agentendpointpb.ReportTaskStartResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*agentendpointpb.ReportTaskStartResponse), nil
}

func (s *mockAgentEndpointServer) ReportTaskProgress(ctx context.Context, req *agentendpointpb.ReportTaskProgressRequest) (*agentendpointpb.ReportTaskProgressResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*agentendpointpb.ReportTaskProgressResponse), nil
}

func (s *mockAgentEndpointServer) ReportTaskComplete(ctx context.Context, req *agentendpointpb.ReportTaskCompleteRequest) (*agentendpointpb.ReportTaskCompleteResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*agentendpointpb.ReportTaskCompleteResponse), nil
}

func (s *mockAgentEndpointServer) LookupEffectiveGuestPolicies(ctx context.Context, req *agentendpointpb.LookupEffectiveGuestPoliciesRequest) (*agentendpointpb.LookupEffectiveGuestPoliciesResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*agentendpointpb.LookupEffectiveGuestPoliciesResponse), nil
}

// clientOpt is the option tests should use to connect to the test server.
// It is initialized by TestMain.
var clientOpt option.ClientOption

var (
	mockAgentEndpoint mockAgentEndpointServer
)

func TestMain(m *testing.M) {
	flag.Parse()

	serv := grpc.NewServer()
	agentendpointpb.RegisterAgentEndpointServiceServer(serv, &mockAgentEndpoint)

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	go serv.Serve(lis)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clientOpt = option.WithGRPCConn(conn)

	os.Exit(m.Run())
}

func TestAgentEndpointServiceReceiveTaskNotification(t *testing.T) {
	var expectedResponse *agentendpointpb.ReceiveTaskNotificationResponse = &agentendpointpb.ReceiveTaskNotificationResponse{}

	mockAgentEndpoint.err = nil
	mockAgentEndpoint.reqs = nil

	mockAgentEndpoint.resps = append(mockAgentEndpoint.resps[:0], expectedResponse)

	var instanceIdToken string = "instanceIdToken1328651839"
	var agentVersion string = "agentVersion-26267234"
	var request = &agentendpointpb.ReceiveTaskNotificationRequest{
		InstanceIdToken: instanceIdToken,
		AgentVersion:    agentVersion,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	stream, err := c.ReceiveTaskNotification(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := stream.Recv()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockAgentEndpoint.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestAgentEndpointServiceReceiveTaskNotificationError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockAgentEndpoint.err = gstatus.Error(errCode, "test error")

	var instanceIdToken string = "instanceIdToken1328651839"
	var agentVersion string = "agentVersion-26267234"
	var request = &agentendpointpb.ReceiveTaskNotificationRequest{
		InstanceIdToken: instanceIdToken,
		AgentVersion:    agentVersion,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	stream, err := c.ReceiveTaskNotification(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := stream.Recv()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestAgentEndpointServiceReportTaskStart(t *testing.T) {
	var expectedResponse *agentendpointpb.ReportTaskStartResponse = &agentendpointpb.ReportTaskStartResponse{}

	mockAgentEndpoint.err = nil
	mockAgentEndpoint.reqs = nil

	mockAgentEndpoint.resps = append(mockAgentEndpoint.resps[:0], expectedResponse)

	var instanceIdToken string = "instanceIdToken1328651839"
	var request = &agentendpointpb.ReportTaskStartRequest{
		InstanceIdToken: instanceIdToken,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ReportTaskStart(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockAgentEndpoint.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestAgentEndpointServiceReportTaskStartError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockAgentEndpoint.err = gstatus.Error(errCode, "test error")

	var instanceIdToken string = "instanceIdToken1328651839"
	var request = &agentendpointpb.ReportTaskStartRequest{
		InstanceIdToken: instanceIdToken,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ReportTaskStart(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestAgentEndpointServiceReportTaskProgress(t *testing.T) {
	var expectedResponse *agentendpointpb.ReportTaskProgressResponse = &agentendpointpb.ReportTaskProgressResponse{}

	mockAgentEndpoint.err = nil
	mockAgentEndpoint.reqs = nil

	mockAgentEndpoint.resps = append(mockAgentEndpoint.resps[:0], expectedResponse)

	var instanceIdToken string = "instanceIdToken1328651839"
	var taskId string = "taskId-1537240555"
	var taskType agentendpointpb.TaskType = agentendpointpb.TaskType_TASK_TYPE_UNSPECIFIED
	var request = &agentendpointpb.ReportTaskProgressRequest{
		InstanceIdToken: instanceIdToken,
		TaskId:          taskId,
		TaskType:        taskType,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ReportTaskProgress(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockAgentEndpoint.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestAgentEndpointServiceReportTaskProgressError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockAgentEndpoint.err = gstatus.Error(errCode, "test error")

	var instanceIdToken string = "instanceIdToken1328651839"
	var taskId string = "taskId-1537240555"
	var taskType agentendpointpb.TaskType = agentendpointpb.TaskType_TASK_TYPE_UNSPECIFIED
	var request = &agentendpointpb.ReportTaskProgressRequest{
		InstanceIdToken: instanceIdToken,
		TaskId:          taskId,
		TaskType:        taskType,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ReportTaskProgress(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestAgentEndpointServiceReportTaskComplete(t *testing.T) {
	var expectedResponse *agentendpointpb.ReportTaskCompleteResponse = &agentendpointpb.ReportTaskCompleteResponse{}

	mockAgentEndpoint.err = nil
	mockAgentEndpoint.reqs = nil

	mockAgentEndpoint.resps = append(mockAgentEndpoint.resps[:0], expectedResponse)

	var instanceIdToken string = "instanceIdToken1328651839"
	var taskId string = "taskId-1537240555"
	var taskType agentendpointpb.TaskType = agentendpointpb.TaskType_TASK_TYPE_UNSPECIFIED
	var errorMessage string = "errorMessage-1938755376"
	var request = &agentendpointpb.ReportTaskCompleteRequest{
		InstanceIdToken: instanceIdToken,
		TaskId:          taskId,
		TaskType:        taskType,
		ErrorMessage:    errorMessage,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ReportTaskComplete(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockAgentEndpoint.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestAgentEndpointServiceReportTaskCompleteError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockAgentEndpoint.err = gstatus.Error(errCode, "test error")

	var instanceIdToken string = "instanceIdToken1328651839"
	var taskId string = "taskId-1537240555"
	var taskType agentendpointpb.TaskType = agentendpointpb.TaskType_TASK_TYPE_UNSPECIFIED
	var errorMessage string = "errorMessage-1938755376"
	var request = &agentendpointpb.ReportTaskCompleteRequest{
		InstanceIdToken: instanceIdToken,
		TaskId:          taskId,
		TaskType:        taskType,
		ErrorMessage:    errorMessage,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ReportTaskComplete(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestAgentEndpointServiceLookupEffectiveGuestPolicies(t *testing.T) {
	var expectedResponse *agentendpointpb.LookupEffectiveGuestPoliciesResponse = &agentendpointpb.LookupEffectiveGuestPoliciesResponse{}

	mockAgentEndpoint.err = nil
	mockAgentEndpoint.reqs = nil

	mockAgentEndpoint.resps = append(mockAgentEndpoint.resps[:0], expectedResponse)

	var instanceIdToken string = "instanceIdToken1328651839"
	var osShortName string = "osShortName2110754217"
	var osVersion string = "osVersion672836989"
	var osArchitecture string = "osArchitecture561018830"
	var request = &agentendpointpb.LookupEffectiveGuestPoliciesRequest{
		InstanceIdToken: instanceIdToken,
		OsShortName:     osShortName,
		OsVersion:       osVersion,
		OsArchitecture:  osArchitecture,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.LookupEffectiveGuestPolicies(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockAgentEndpoint.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestAgentEndpointServiceLookupEffectiveGuestPoliciesError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockAgentEndpoint.err = gstatus.Error(errCode, "test error")

	var instanceIdToken string = "instanceIdToken1328651839"
	var osShortName string = "osShortName2110754217"
	var osVersion string = "osVersion672836989"
	var osArchitecture string = "osArchitecture561018830"
	var request = &agentendpointpb.LookupEffectiveGuestPoliciesRequest{
		InstanceIdToken: instanceIdToken,
		OsShortName:     osShortName,
		OsVersion:       osVersion,
		OsArchitecture:  osArchitecture,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.LookupEffectiveGuestPolicies(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
