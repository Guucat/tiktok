// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: snowflake.proto

package snowflake_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Snowflake_GetId_FullMethodName = "/Snowflake/GetId"
)

// SnowflakeClient is the client API for Snowflake service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnowflakeClient interface {
	GetId(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IdResponse, error)
}

type snowflakeClient struct {
	cc grpc.ClientConnInterface
}

func NewSnowflakeClient(cc grpc.ClientConnInterface) SnowflakeClient {
	return &snowflakeClient{cc}
}

func (c *snowflakeClient) GetId(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, Snowflake_GetId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnowflakeServer is the server API for Snowflake service.
// All implementations must embed UnimplementedSnowflakeServer
// for forward compatibility
type SnowflakeServer interface {
	GetId(context.Context, *emptypb.Empty) (*IdResponse, error)
	mustEmbedUnimplementedSnowflakeServer()
}

// UnimplementedSnowflakeServer must be embedded to have forward compatible implementations.
type UnimplementedSnowflakeServer struct {
}

func (UnimplementedSnowflakeServer) GetId(context.Context, *emptypb.Empty) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetId not implemented")
}
func (UnimplementedSnowflakeServer) mustEmbedUnimplementedSnowflakeServer() {}

// UnsafeSnowflakeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnowflakeServer will
// result in compilation errors.
type UnsafeSnowflakeServer interface {
	mustEmbedUnimplementedSnowflakeServer()
}

func RegisterSnowflakeServer(s grpc.ServiceRegistrar, srv SnowflakeServer) {
	s.RegisterService(&Snowflake_ServiceDesc, srv)
}

func _Snowflake_GetId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnowflakeServer).GetId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Snowflake_GetId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnowflakeServer).GetId(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Snowflake_ServiceDesc is the grpc.ServiceDesc for Snowflake service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Snowflake_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Snowflake",
	HandlerType: (*SnowflakeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetId",
			Handler:    _Snowflake_GetId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "snowflake.proto",
}
