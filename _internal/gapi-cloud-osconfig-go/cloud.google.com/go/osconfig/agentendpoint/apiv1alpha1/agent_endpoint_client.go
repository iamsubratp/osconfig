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
	"context"
	"math"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	agentendpointpb "github.com/GoogleCloudPlatform/osconfig/_internal/gapi-cloud-osconfig-go/google.golang.org/genproto/googleapis/cloud/osconfig/agentendpoint/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ReceiveTaskNotification      []gax.CallOption
	ReportTaskStart              []gax.CallOption
	ReportTaskProgress           []gax.CallOption
	ReportTaskComplete           []gax.CallOption
	LookupEffectiveGuestPolicies []gax.CallOption
}

func defaultClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("osconfig.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	retry := map[[2]string][]gax.CallOption{}
	return &CallOptions{
		ReceiveTaskNotification:      retry[[2]string{"default", "non_idempotent"}],
		ReportTaskStart:              retry[[2]string{"default", "non_idempotent"}],
		ReportTaskProgress:           retry[[2]string{"default", "non_idempotent"}],
		ReportTaskComplete:           retry[[2]string{"default", "non_idempotent"}],
		LookupEffectiveGuestPolicies: retry[[2]string{"default", "non_idempotent"}],
	}
}

// Client is a client for interacting with Cloud OS Config API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type Client struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	client agentendpointpb.AgentEndpointServiceClient

	// The call options for this service.
	CallOptions *CallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new agent endpoint service client.
//
// OS Config agent endpoint API.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:        conn,
		CallOptions: defaultCallOptions(),

		client: agentendpointpb.NewAgentEndpointServiceClient(conn),
	}
	c.setGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *Client) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ReceiveTaskNotification stream established by client to receive Task notifications.
// This method is called by an agent and not an active developer method.
func (c *Client) ReceiveTaskNotification(ctx context.Context, req *agentendpointpb.ReceiveTaskNotificationRequest, opts ...gax.CallOption) (agentendpointpb.AgentEndpointService_ReceiveTaskNotificationClient, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ReceiveTaskNotification[0:len(c.CallOptions.ReceiveTaskNotification):len(c.CallOptions.ReceiveTaskNotification)], opts...)
	var resp agentendpointpb.AgentEndpointService_ReceiveTaskNotificationClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReceiveTaskNotification(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ReportTaskStart signals the start of a task execution and returns the task info.
// This method is called by an agent and not an active developer method.
func (c *Client) ReportTaskStart(ctx context.Context, req *agentendpointpb.ReportTaskStartRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskStartResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ReportTaskStart[0:len(c.CallOptions.ReportTaskStart):len(c.CallOptions.ReportTaskStart)], opts...)
	var resp *agentendpointpb.ReportTaskStartResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReportTaskStart(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ReportTaskProgress signals an intermediary progress checkpoint in task execution.
// This method is called by an agent and not an active developer method.
func (c *Client) ReportTaskProgress(ctx context.Context, req *agentendpointpb.ReportTaskProgressRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskProgressResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ReportTaskProgress[0:len(c.CallOptions.ReportTaskProgress):len(c.CallOptions.ReportTaskProgress)], opts...)
	var resp *agentendpointpb.ReportTaskProgressResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReportTaskProgress(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ReportTaskComplete signals that the task execution is complete and optionally returns the next
// task.
// This method is called by an agent and not an active developer method.
func (c *Client) ReportTaskComplete(ctx context.Context, req *agentendpointpb.ReportTaskCompleteRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskCompleteResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ReportTaskComplete[0:len(c.CallOptions.ReportTaskComplete):len(c.CallOptions.ReportTaskComplete)], opts...)
	var resp *agentendpointpb.ReportTaskCompleteResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReportTaskComplete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LookupEffectiveGuestPolicies looks up the effective guest policies for an instance.
func (c *Client) LookupEffectiveGuestPolicies(ctx context.Context, req *agentendpointpb.LookupEffectiveGuestPoliciesRequest, opts ...gax.CallOption) (*agentendpointpb.LookupEffectiveGuestPoliciesResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.LookupEffectiveGuestPolicies[0:len(c.CallOptions.LookupEffectiveGuestPolicies):len(c.CallOptions.LookupEffectiveGuestPolicies)], opts...)
	var resp *agentendpointpb.LookupEffectiveGuestPoliciesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.LookupEffectiveGuestPolicies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
