// Code generated by protoc-gen-go.
// source: binlogservice.proto
// DO NOT EDIT!

/*
Package binlogservice is a generated protocol buffer package.

It is generated from these files:
	binlogservice.proto

It has these top-level messages:
*/
package binlogservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import binlogdata "github.com/youtube/vitess/go/vt/proto/binlogdata"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for UpdateStream service

type UpdateStreamClient interface {
	// StreamUpdate streams the binlog events, to know which objects have changed.
	StreamUpdate(ctx context.Context, in *binlogdata.StreamUpdateRequest, opts ...grpc.CallOption) (UpdateStream_StreamUpdateClient, error)
	// StreamKeyRange returns the binlog transactions related to
	// the specified Keyrange.
	StreamKeyRange(ctx context.Context, in *binlogdata.StreamKeyRangeRequest, opts ...grpc.CallOption) (UpdateStream_StreamKeyRangeClient, error)
	// StreamTables returns the binlog transactions related to
	// the specified Tables.
	StreamTables(ctx context.Context, in *binlogdata.StreamTablesRequest, opts ...grpc.CallOption) (UpdateStream_StreamTablesClient, error)
}

type updateStreamClient struct {
	cc *grpc.ClientConn
}

func NewUpdateStreamClient(cc *grpc.ClientConn) UpdateStreamClient {
	return &updateStreamClient{cc}
}

func (c *updateStreamClient) StreamUpdate(ctx context.Context, in *binlogdata.StreamUpdateRequest, opts ...grpc.CallOption) (UpdateStream_StreamUpdateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UpdateStream_serviceDesc.Streams[0], c.cc, "/binlogservice.UpdateStream/StreamUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &updateStreamStreamUpdateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UpdateStream_StreamUpdateClient interface {
	Recv() (*binlogdata.StreamUpdateResponse, error)
	grpc.ClientStream
}

type updateStreamStreamUpdateClient struct {
	grpc.ClientStream
}

func (x *updateStreamStreamUpdateClient) Recv() (*binlogdata.StreamUpdateResponse, error) {
	m := new(binlogdata.StreamUpdateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *updateStreamClient) StreamKeyRange(ctx context.Context, in *binlogdata.StreamKeyRangeRequest, opts ...grpc.CallOption) (UpdateStream_StreamKeyRangeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UpdateStream_serviceDesc.Streams[1], c.cc, "/binlogservice.UpdateStream/StreamKeyRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &updateStreamStreamKeyRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UpdateStream_StreamKeyRangeClient interface {
	Recv() (*binlogdata.StreamKeyRangeResponse, error)
	grpc.ClientStream
}

type updateStreamStreamKeyRangeClient struct {
	grpc.ClientStream
}

func (x *updateStreamStreamKeyRangeClient) Recv() (*binlogdata.StreamKeyRangeResponse, error) {
	m := new(binlogdata.StreamKeyRangeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *updateStreamClient) StreamTables(ctx context.Context, in *binlogdata.StreamTablesRequest, opts ...grpc.CallOption) (UpdateStream_StreamTablesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UpdateStream_serviceDesc.Streams[2], c.cc, "/binlogservice.UpdateStream/StreamTables", opts...)
	if err != nil {
		return nil, err
	}
	x := &updateStreamStreamTablesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UpdateStream_StreamTablesClient interface {
	Recv() (*binlogdata.StreamTablesResponse, error)
	grpc.ClientStream
}

type updateStreamStreamTablesClient struct {
	grpc.ClientStream
}

func (x *updateStreamStreamTablesClient) Recv() (*binlogdata.StreamTablesResponse, error) {
	m := new(binlogdata.StreamTablesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for UpdateStream service

type UpdateStreamServer interface {
	// StreamUpdate streams the binlog events, to know which objects have changed.
	StreamUpdate(*binlogdata.StreamUpdateRequest, UpdateStream_StreamUpdateServer) error
	// StreamKeyRange returns the binlog transactions related to
	// the specified Keyrange.
	StreamKeyRange(*binlogdata.StreamKeyRangeRequest, UpdateStream_StreamKeyRangeServer) error
	// StreamTables returns the binlog transactions related to
	// the specified Tables.
	StreamTables(*binlogdata.StreamTablesRequest, UpdateStream_StreamTablesServer) error
}

func RegisterUpdateStreamServer(s *grpc.Server, srv UpdateStreamServer) {
	s.RegisterService(&_UpdateStream_serviceDesc, srv)
}

func _UpdateStream_StreamUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(binlogdata.StreamUpdateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpdateStreamServer).StreamUpdate(m, &updateStreamStreamUpdateServer{stream})
}

type UpdateStream_StreamUpdateServer interface {
	Send(*binlogdata.StreamUpdateResponse) error
	grpc.ServerStream
}

type updateStreamStreamUpdateServer struct {
	grpc.ServerStream
}

func (x *updateStreamStreamUpdateServer) Send(m *binlogdata.StreamUpdateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UpdateStream_StreamKeyRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(binlogdata.StreamKeyRangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpdateStreamServer).StreamKeyRange(m, &updateStreamStreamKeyRangeServer{stream})
}

type UpdateStream_StreamKeyRangeServer interface {
	Send(*binlogdata.StreamKeyRangeResponse) error
	grpc.ServerStream
}

type updateStreamStreamKeyRangeServer struct {
	grpc.ServerStream
}

func (x *updateStreamStreamKeyRangeServer) Send(m *binlogdata.StreamKeyRangeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UpdateStream_StreamTables_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(binlogdata.StreamTablesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpdateStreamServer).StreamTables(m, &updateStreamStreamTablesServer{stream})
}

type UpdateStream_StreamTablesServer interface {
	Send(*binlogdata.StreamTablesResponse) error
	grpc.ServerStream
}

type updateStreamStreamTablesServer struct {
	grpc.ServerStream
}

func (x *updateStreamStreamTablesServer) Send(m *binlogdata.StreamTablesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _UpdateStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binlogservice.UpdateStream",
	HandlerType: (*UpdateStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamUpdate",
			Handler:       _UpdateStream_StreamUpdate_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamKeyRange",
			Handler:       _UpdateStream_StreamKeyRange_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamTables",
			Handler:       _UpdateStream_StreamTables_Handler,
			ServerStreams: true,
		},
	},
}