// Code generated by protoc-gen-go.
// source: service.auth/proto/auth.proto
// DO NOT EDIT!

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	service.auth/proto/auth.proto

It has these top-level messages:
	Args
	Customer
*/
package auth

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Args struct {
	From      string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	AuthToken string `protobuf:"bytes,2,opt,name=authToken" json:"authToken,omitempty"`
}

func (m *Args) Reset()         { *m = Args{} }
func (m *Args) String() string { return proto.CompactTextString(m) }
func (*Args) ProtoMessage()    {}

type Customer struct {
	Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	AuthToken string `protobuf:"bytes,2,opt,name=authToken" json:"authToken,omitempty"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}

// Client API for Auth service

type AuthClient interface {
	VerifyToken(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Customer, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) VerifyToken(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := grpc.Invoke(ctx, "/auth.Auth/VerifyToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	VerifyToken(context.Context, *Args) (*Customer, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_VerifyToken_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Args)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthServer).VerifyToken(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyToken",
			Handler:    _Auth_VerifyToken_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
