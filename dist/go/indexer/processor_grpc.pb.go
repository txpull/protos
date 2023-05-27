// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package indexer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProcessorClient is the client API for Processor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessorClient interface {
	QueueBlock(ctx context.Context, in *QueueBlockRequest, opts ...grpc.CallOption) (*QueueBlockResponse, error)
}

type processorClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessorClient(cc grpc.ClientConnInterface) ProcessorClient {
	return &processorClient{cc}
}

func (c *processorClient) QueueBlock(ctx context.Context, in *QueueBlockRequest, opts ...grpc.CallOption) (*QueueBlockResponse, error) {
	out := new(QueueBlockResponse)
	err := c.cc.Invoke(ctx, "/txpull.v1.indexer.processor.Processor/QueueBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorServer is the server API for Processor service.
// All implementations must embed UnimplementedProcessorServer
// for forward compatibility
type ProcessorServer interface {
	QueueBlock(context.Context, *QueueBlockRequest) (*QueueBlockResponse, error)
	mustEmbedUnimplementedProcessorServer()
}

// UnimplementedProcessorServer must be embedded to have forward compatible implementations.
type UnimplementedProcessorServer struct {
}

func (UnimplementedProcessorServer) QueueBlock(context.Context, *QueueBlockRequest) (*QueueBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueueBlock not implemented")
}
func (UnimplementedProcessorServer) mustEmbedUnimplementedProcessorServer() {}

// UnsafeProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessorServer will
// result in compilation errors.
type UnsafeProcessorServer interface {
	mustEmbedUnimplementedProcessorServer()
}

func RegisterProcessorServer(s grpc.ServiceRegistrar, srv ProcessorServer) {
	s.RegisterService(&Processor_ServiceDesc, srv)
}

func _Processor_QueueBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServer).QueueBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/txpull.v1.indexer.processor.Processor/QueueBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServer).QueueBlock(ctx, req.(*QueueBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Processor_ServiceDesc is the grpc.ServiceDesc for Processor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Processor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "txpull.v1.indexer.processor.Processor",
	HandlerType: (*ProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueueBlock",
			Handler:    _Processor_QueueBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "indexer/processor.proto",
}