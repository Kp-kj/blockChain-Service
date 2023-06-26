// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: block.proto

package block

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
	Block_Ping_FullMethodName = "/block.Block/Ping"
)

// BlockClient is the client API for Block service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type blockClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockClient(cc grpc.ClientConnInterface) BlockClient {
	return &blockClient{cc}
}

func (c *blockClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Block_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockServer is the server API for Block service.
// All implementations must embed UnimplementedBlockServer
// for forward compatibility
type BlockServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedBlockServer()
}

// UnimplementedBlockServer must be embedded to have forward compatible implementations.
type UnimplementedBlockServer struct {
}

func (UnimplementedBlockServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBlockServer) mustEmbedUnimplementedBlockServer() {}

// UnsafeBlockServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockServer will
// result in compilation errors.
type UnsafeBlockServer interface {
	mustEmbedUnimplementedBlockServer()
}

func RegisterBlockServer(s grpc.ServiceRegistrar, srv BlockServer) {
	s.RegisterService(&Block_ServiceDesc, srv)
}

func _Block_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Block_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Block_ServiceDesc is the grpc.ServiceDesc for Block service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Block_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "block.Block",
	HandlerType: (*BlockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Block_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "block.proto",
}
