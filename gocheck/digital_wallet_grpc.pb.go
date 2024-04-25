// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: gocheck/digital_wallet.proto

package gocheckgrpc

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

// DigitalWalletClient is the client API for DigitalWallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DigitalWalletClient interface {
	Transfer(ctx context.Context, in *ReqDigitalWalletTransfer, opts ...grpc.CallOption) (*ResDigitalWalletTransfer, error)
}

type digitalWalletClient struct {
	cc grpc.ClientConnInterface
}

func NewDigitalWalletClient(cc grpc.ClientConnInterface) DigitalWalletClient {
	return &digitalWalletClient{cc}
}

func (c *digitalWalletClient) Transfer(ctx context.Context, in *ReqDigitalWalletTransfer, opts ...grpc.CallOption) (*ResDigitalWalletTransfer, error) {
	out := new(ResDigitalWalletTransfer)
	err := c.cc.Invoke(ctx, "/DigitalWallet/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DigitalWalletServer is the server API for DigitalWallet service.
// All implementations must embed UnimplementedDigitalWalletServer
// for forward compatibility
type DigitalWalletServer interface {
	Transfer(context.Context, *ReqDigitalWalletTransfer) (*ResDigitalWalletTransfer, error)
	mustEmbedUnimplementedDigitalWalletServer()
}

// UnimplementedDigitalWalletServer must be embedded to have forward compatible implementations.
type UnimplementedDigitalWalletServer struct {
}

func (UnimplementedDigitalWalletServer) Transfer(context.Context, *ReqDigitalWalletTransfer) (*ResDigitalWalletTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedDigitalWalletServer) mustEmbedUnimplementedDigitalWalletServer() {}

// UnsafeDigitalWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DigitalWalletServer will
// result in compilation errors.
type UnsafeDigitalWalletServer interface {
	mustEmbedUnimplementedDigitalWalletServer()
}

func RegisterDigitalWalletServer(s grpc.ServiceRegistrar, srv DigitalWalletServer) {
	s.RegisterService(&DigitalWallet_ServiceDesc, srv)
}

func _DigitalWallet_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDigitalWalletTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DigitalWalletServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DigitalWallet/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DigitalWalletServer).Transfer(ctx, req.(*ReqDigitalWalletTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

// DigitalWallet_ServiceDesc is the grpc.ServiceDesc for DigitalWallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DigitalWallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DigitalWallet",
	HandlerType: (*DigitalWalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _DigitalWallet_Transfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gocheck/digital_wallet.proto",
}
