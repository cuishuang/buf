// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.19.1
// source: buf/alpha/registry/v1alpha1/token.proto

package registryv1alpha1

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

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenServiceClient interface {
	// CreateToken creates a new token suitable for machine-to-machine authentication.
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
	// GetToken gets the specific token for the user
	//
	// This method requires authentication.
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
	// ListTokens lists the users active tokens
	//
	// This method requires authentication.
	ListTokens(ctx context.Context, in *ListTokensRequest, opts ...grpc.CallOption) (*ListTokensResponse, error)
	// DeleteToken deletes an existing token.
	//
	// This method requires authentication.
	DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error)
}

type tokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenServiceClient(cc grpc.ClientConnInterface) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.TokenService/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.TokenService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) ListTokens(ctx context.Context, in *ListTokensRequest, opts ...grpc.CallOption) (*ListTokensResponse, error) {
	out := new(ListTokensResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.TokenService/ListTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error) {
	out := new(DeleteTokenResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.TokenService/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServiceServer is the server API for TokenService service.
// All implementations should embed UnimplementedTokenServiceServer
// for forward compatibility
type TokenServiceServer interface {
	// CreateToken creates a new token suitable for machine-to-machine authentication.
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	// GetToken gets the specific token for the user
	//
	// This method requires authentication.
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
	// ListTokens lists the users active tokens
	//
	// This method requires authentication.
	ListTokens(context.Context, *ListTokensRequest) (*ListTokensResponse, error)
	// DeleteToken deletes an existing token.
	//
	// This method requires authentication.
	DeleteToken(context.Context, *DeleteTokenRequest) (*DeleteTokenResponse, error)
}

// UnimplementedTokenServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (UnimplementedTokenServiceServer) CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedTokenServiceServer) GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedTokenServiceServer) ListTokens(context.Context, *ListTokensRequest) (*ListTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTokens not implemented")
}
func (UnimplementedTokenServiceServer) DeleteToken(context.Context, *DeleteTokenRequest) (*DeleteTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}

// UnsafeTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenServiceServer will
// result in compilation errors.
type UnsafeTokenServiceServer interface {
	mustEmbedUnimplementedTokenServiceServer()
}

func RegisterTokenServiceServer(s grpc.ServiceRegistrar, srv TokenServiceServer) {
	s.RegisterService(&TokenService_ServiceDesc, srv)
}

func _TokenService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.TokenService/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.TokenService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_ListTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).ListTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.TokenService/ListTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).ListTokens(ctx, req.(*ListTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.TokenService/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).DeleteToken(ctx, req.(*DeleteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenService_ServiceDesc is the grpc.ServiceDesc for TokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buf.alpha.registry.v1alpha1.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _TokenService_CreateToken_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _TokenService_GetToken_Handler,
		},
		{
			MethodName: "ListTokens",
			Handler:    _TokenService_ListTokens_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _TokenService_DeleteToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buf/alpha/registry/v1alpha1/token.proto",
}
