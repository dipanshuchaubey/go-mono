// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: protos/bootcamp_service/bootcamps.proto

package bootcamp_service

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

const (
	BootcampService_GetBootcampsDetails_FullMethodName = "/protos.bootcamp_service.BootcampService/GetBootcampsDetails"
)

// BootcampServiceClient is the client API for BootcampService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BootcampServiceClient interface {
	GetBootcampsDetails(ctx context.Context, in *GetBootcampsDetailsRequest, opts ...grpc.CallOption) (*GetBootcampsDetailsResponse, error)
}

type bootcampServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBootcampServiceClient(cc grpc.ClientConnInterface) BootcampServiceClient {
	return &bootcampServiceClient{cc}
}

func (c *bootcampServiceClient) GetBootcampsDetails(ctx context.Context, in *GetBootcampsDetailsRequest, opts ...grpc.CallOption) (*GetBootcampsDetailsResponse, error) {
	out := new(GetBootcampsDetailsResponse)
	err := c.cc.Invoke(ctx, BootcampService_GetBootcampsDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BootcampServiceServer is the server API for BootcampService service.
// All implementations must embed UnimplementedBootcampServiceServer
// for forward compatibility
type BootcampServiceServer interface {
	GetBootcampsDetails(context.Context, *GetBootcampsDetailsRequest) (*GetBootcampsDetailsResponse, error)
	mustEmbedUnimplementedBootcampServiceServer()
}

// UnimplementedBootcampServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBootcampServiceServer struct {
}

func (UnimplementedBootcampServiceServer) GetBootcampsDetails(context.Context, *GetBootcampsDetailsRequest) (*GetBootcampsDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBootcampsDetails not implemented")
}
func (UnimplementedBootcampServiceServer) mustEmbedUnimplementedBootcampServiceServer() {}

// UnsafeBootcampServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BootcampServiceServer will
// result in compilation errors.
type UnsafeBootcampServiceServer interface {
	mustEmbedUnimplementedBootcampServiceServer()
}

func RegisterBootcampServiceServer(s grpc.ServiceRegistrar, srv BootcampServiceServer) {
	s.RegisterService(&BootcampService_ServiceDesc, srv)
}

func _BootcampService_GetBootcampsDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBootcampsDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootcampServiceServer).GetBootcampsDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BootcampService_GetBootcampsDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootcampServiceServer).GetBootcampsDetails(ctx, req.(*GetBootcampsDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BootcampService_ServiceDesc is the grpc.ServiceDesc for BootcampService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BootcampService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.bootcamp_service.BootcampService",
	HandlerType: (*BootcampServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBootcampsDetails",
			Handler:    _BootcampService_GetBootcampsDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/bootcamp_service/bootcamps.proto",
}
