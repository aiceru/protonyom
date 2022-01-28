// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: protonyom_api_account.proto

package gonyom

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

// AccountApiClient is the client API for AccountApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountApiClient interface {
	Get(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountReply, error)
	Update(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountReply, error)
	Delete(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountReply, error)
}

type accountApiClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountApiClient(cc grpc.ClientConnInterface) AccountApiClient {
	return &accountApiClient{cc}
}

func (c *accountApiClient) Get(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountReply, error) {
	out := new(GetAccountReply)
	err := c.cc.Invoke(ctx, "/protonyom.AccountApi/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountApiClient) Update(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountReply, error) {
	out := new(UpdateAccountReply)
	err := c.cc.Invoke(ctx, "/protonyom.AccountApi/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountApiClient) Delete(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountReply, error) {
	out := new(DeleteAccountReply)
	err := c.cc.Invoke(ctx, "/protonyom.AccountApi/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountApiServer is the server API for AccountApi service.
// All implementations must embed UnimplementedAccountApiServer
// for forward compatibility
type AccountApiServer interface {
	Get(context.Context, *GetAccountRequest) (*GetAccountReply, error)
	Update(context.Context, *UpdateAccountRequest) (*UpdateAccountReply, error)
	Delete(context.Context, *DeleteAccountRequest) (*DeleteAccountReply, error)
	mustEmbedUnimplementedAccountApiServer()
}

// UnimplementedAccountApiServer must be embedded to have forward compatible implementations.
type UnimplementedAccountApiServer struct {
}

func (UnimplementedAccountApiServer) Get(context.Context, *GetAccountRequest) (*GetAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAccountApiServer) Update(context.Context, *UpdateAccountRequest) (*UpdateAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAccountApiServer) Delete(context.Context, *DeleteAccountRequest) (*DeleteAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAccountApiServer) mustEmbedUnimplementedAccountApiServer() {}

// UnsafeAccountApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountApiServer will
// result in compilation errors.
type UnsafeAccountApiServer interface {
	mustEmbedUnimplementedAccountApiServer()
}

func RegisterAccountApiServer(s grpc.ServiceRegistrar, srv AccountApiServer) {
	s.RegisterService(&AccountApi_ServiceDesc, srv)
}

func _AccountApi_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountApiServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protonyom.AccountApi/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountApiServer).Get(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountApi_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountApiServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protonyom.AccountApi/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountApiServer).Update(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountApi_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountApiServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protonyom.AccountApi/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountApiServer).Delete(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountApi_ServiceDesc is the grpc.ServiceDesc for AccountApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protonyom.AccountApi",
	HandlerType: (*AccountApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AccountApi_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AccountApi_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccountApi_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protonyom_api_account.proto",
}
