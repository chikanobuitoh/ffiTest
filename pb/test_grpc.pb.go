// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: pb/test.proto

package pb

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

// SampleSerciveClient is the client API for SampleSercive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleSerciveClient interface {
	// giftUpload
	Check(ctx context.Context, opts ...grpc.CallOption) (SampleSercive_CheckClient, error)
}

type sampleSerciveClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleSerciveClient(cc grpc.ClientConnInterface) SampleSerciveClient {
	return &sampleSerciveClient{cc}
}

func (c *sampleSerciveClient) Check(ctx context.Context, opts ...grpc.CallOption) (SampleSercive_CheckClient, error) {
	stream, err := c.cc.NewStream(ctx, &SampleSercive_ServiceDesc.Streams[0], "/grpcsample.SampleSercive/Check", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleSerciveCheckClient{stream}
	return x, nil
}

type SampleSercive_CheckClient interface {
	Send(*CheckRequest) error
	CloseAndRecv() (*CheckResponce, error)
	grpc.ClientStream
}

type sampleSerciveCheckClient struct {
	grpc.ClientStream
}

func (x *sampleSerciveCheckClient) Send(m *CheckRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sampleSerciveCheckClient) CloseAndRecv() (*CheckResponce, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CheckResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SampleSerciveServer is the server API for SampleSercive service.
// All implementations must embed UnimplementedSampleSerciveServer
// for forward compatibility
type SampleSerciveServer interface {
	// giftUpload
	Check(SampleSercive_CheckServer) error
	mustEmbedUnimplementedSampleSerciveServer()
}

// UnimplementedSampleSerciveServer must be embedded to have forward compatible implementations.
type UnimplementedSampleSerciveServer struct {
}

func (UnimplementedSampleSerciveServer) Check(SampleSercive_CheckServer) error {
	return status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedSampleSerciveServer) mustEmbedUnimplementedSampleSerciveServer() {}

// UnsafeSampleSerciveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleSerciveServer will
// result in compilation errors.
type UnsafeSampleSerciveServer interface {
	mustEmbedUnimplementedSampleSerciveServer()
}

func RegisterSampleSerciveServer(s grpc.ServiceRegistrar, srv SampleSerciveServer) {
	s.RegisterService(&SampleSercive_ServiceDesc, srv)
}

func _SampleSercive_Check_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SampleSerciveServer).Check(&sampleSerciveCheckServer{stream})
}

type SampleSercive_CheckServer interface {
	SendAndClose(*CheckResponce) error
	Recv() (*CheckRequest, error)
	grpc.ServerStream
}

type sampleSerciveCheckServer struct {
	grpc.ServerStream
}

func (x *sampleSerciveCheckServer) SendAndClose(m *CheckResponce) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sampleSerciveCheckServer) Recv() (*CheckRequest, error) {
	m := new(CheckRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SampleSercive_ServiceDesc is the grpc.ServiceDesc for SampleSercive service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SampleSercive_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcsample.SampleSercive",
	HandlerType: (*SampleSerciveServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Check",
			Handler:       _SampleSercive_Check_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pb/test.proto",
}
