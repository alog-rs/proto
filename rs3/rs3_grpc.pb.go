// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rs3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RunescapeClient is the client API for Runescape service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RunescapeClient interface {
	GetPlayerProfile(ctx context.Context, in *GetPlayerProfileRequest, opts ...grpc.CallOption) (*PlayerProfile, error)
}

type runescapeClient struct {
	cc grpc.ClientConnInterface
}

func NewRunescapeClient(cc grpc.ClientConnInterface) RunescapeClient {
	return &runescapeClient{cc}
}

func (c *runescapeClient) GetPlayerProfile(ctx context.Context, in *GetPlayerProfileRequest, opts ...grpc.CallOption) (*PlayerProfile, error) {
	out := new(PlayerProfile)
	err := c.cc.Invoke(ctx, "/rs3.Runescape/GetPlayerProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunescapeServer is the server API for Runescape service.
// All implementations must embed UnimplementedRunescapeServer
// for forward compatibility
type RunescapeServer interface {
	GetPlayerProfile(context.Context, *GetPlayerProfileRequest) (*PlayerProfile, error)
	mustEmbedUnimplementedRunescapeServer()
}

// UnimplementedRunescapeServer must be embedded to have forward compatible implementations.
type UnimplementedRunescapeServer struct {
}

func (UnimplementedRunescapeServer) GetPlayerProfile(context.Context, *GetPlayerProfileRequest) (*PlayerProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerProfile not implemented")
}
func (UnimplementedRunescapeServer) mustEmbedUnimplementedRunescapeServer() {}

// UnsafeRunescapeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunescapeServer will
// result in compilation errors.
type UnsafeRunescapeServer interface {
	mustEmbedUnimplementedRunescapeServer()
}

func RegisterRunescapeServer(s grpc.ServiceRegistrar, srv RunescapeServer) {
	s.RegisterService(&_Runescape_serviceDesc, srv)
}

func _Runescape_GetPlayerProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunescapeServer).GetPlayerProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rs3.Runescape/GetPlayerProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunescapeServer).GetPlayerProfile(ctx, req.(*GetPlayerProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Runescape_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rs3.Runescape",
	HandlerType: (*RunescapeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayerProfile",
			Handler:    _Runescape_GetPlayerProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rs3.proto",
}
