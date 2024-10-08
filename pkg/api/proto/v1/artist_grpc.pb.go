// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: api/proto/v1/artist.proto

package artistv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Artist_Get_FullMethodName     = "/artist.Artist/Get"
	Artist_GetAll_FullMethodName  = "/artist.Artist/GetAll"
	Artist_Create_FullMethodName  = "/artist.Artist/Create"
	Artist_Replace_FullMethodName = "/artist.Artist/Replace"
	Artist_Delete_FullMethodName  = "/artist.Artist/Delete"
)

// ArtistClient is the client API for Artist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtistClient interface {
	Get(ctx context.Context, in *GetArtistReq, opts ...grpc.CallOption) (*GetArtistReply, error)
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllArtistReply, error)
	Create(ctx context.Context, in *CreateArtistReq, opts ...grpc.CallOption) (*CreateArtistReply, error)
	Replace(ctx context.Context, in *ReplaceArtistReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DeleteArtistReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type artistClient struct {
	cc grpc.ClientConnInterface
}

func NewArtistClient(cc grpc.ClientConnInterface) ArtistClient {
	return &artistClient{cc}
}

func (c *artistClient) Get(ctx context.Context, in *GetArtistReq, opts ...grpc.CallOption) (*GetArtistReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArtistReply)
	err := c.cc.Invoke(ctx, Artist_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllArtistReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllArtistReply)
	err := c.cc.Invoke(ctx, Artist_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistClient) Create(ctx context.Context, in *CreateArtistReq, opts ...grpc.CallOption) (*CreateArtistReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateArtistReply)
	err := c.cc.Invoke(ctx, Artist_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistClient) Replace(ctx context.Context, in *ReplaceArtistReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Artist_Replace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistClient) Delete(ctx context.Context, in *DeleteArtistReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Artist_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtistServer is the server API for Artist service.
// All implementations must embed UnimplementedArtistServer
// for forward compatibility.
type ArtistServer interface {
	Get(context.Context, *GetArtistReq) (*GetArtistReply, error)
	GetAll(context.Context, *emptypb.Empty) (*GetAllArtistReply, error)
	Create(context.Context, *CreateArtistReq) (*CreateArtistReply, error)
	Replace(context.Context, *ReplaceArtistReq) (*emptypb.Empty, error)
	Delete(context.Context, *DeleteArtistReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedArtistServer()
}

// UnimplementedArtistServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedArtistServer struct{}

func (UnimplementedArtistServer) Get(context.Context, *GetArtistReq) (*GetArtistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedArtistServer) GetAll(context.Context, *emptypb.Empty) (*GetAllArtistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedArtistServer) Create(context.Context, *CreateArtistReq) (*CreateArtistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedArtistServer) Replace(context.Context, *ReplaceArtistReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replace not implemented")
}
func (UnimplementedArtistServer) Delete(context.Context, *DeleteArtistReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedArtistServer) mustEmbedUnimplementedArtistServer() {}
func (UnimplementedArtistServer) testEmbeddedByValue()                {}

// UnsafeArtistServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtistServer will
// result in compilation errors.
type UnsafeArtistServer interface {
	mustEmbedUnimplementedArtistServer()
}

func RegisterArtistServer(s grpc.ServiceRegistrar, srv ArtistServer) {
	// If the following call pancis, it indicates UnimplementedArtistServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Artist_ServiceDesc, srv)
}

func _Artist_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Artist_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServer).Get(ctx, req.(*GetArtistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Artist_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Artist_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Artist_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArtistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Artist_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServer).Create(ctx, req.(*CreateArtistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Artist_Replace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceArtistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServer).Replace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Artist_Replace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServer).Replace(ctx, req.(*ReplaceArtistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Artist_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Artist_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServer).Delete(ctx, req.(*DeleteArtistReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Artist_ServiceDesc is the grpc.ServiceDesc for Artist service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Artist_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "artist.Artist",
	HandlerType: (*ArtistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Artist_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Artist_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Artist_Create_Handler,
		},
		{
			MethodName: "Replace",
			Handler:    _Artist_Replace_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Artist_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/artist.proto",
}
