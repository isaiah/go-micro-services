// Code generated by protoc-gen-go.
// source: service.rate/proto/rate.proto
// DO NOT EDIT!

/*
Package rate is a generated protocol buffer package.

It is generated from these files:
	service.rate/proto/rate.proto

It has these top-level messages:
	Args
	Reply
	RatePlan
	RoomType
*/
package rate

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
	HotelIds []int32 `protobuf:"varint,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
	InDate   string  `protobuf:"bytes,2,opt,name=inDate" json:"inDate,omitempty"`
	OutDate  string  `protobuf:"bytes,3,opt,name=outDate" json:"outDate,omitempty"`
}

func (m *Args) Reset()         { *m = Args{} }
func (m *Args) String() string { return proto.CompactTextString(m) }
func (*Args) ProtoMessage()    {}

type Reply struct {
	RatePlans []*RatePlan `protobuf:"bytes,1,rep,name=ratePlans" json:"ratePlans,omitempty"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}

func (m *Reply) GetRatePlans() []*RatePlan {
	if m != nil {
		return m.RatePlans
	}
	return nil
}

type RatePlan struct {
	HotelId  int32     `protobuf:"varint,1,opt,name=hotelId" json:"hotelId,omitempty"`
	Code     string    `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	InDate   string    `protobuf:"bytes,3,opt,name=inDate" json:"inDate,omitempty"`
	OutDate  string    `protobuf:"bytes,4,opt,name=outDate" json:"outDate,omitempty"`
	RoomType *RoomType `protobuf:"bytes,5,opt,name=roomType" json:"roomType,omitempty"`
}

func (m *RatePlan) Reset()         { *m = RatePlan{} }
func (m *RatePlan) String() string { return proto.CompactTextString(m) }
func (*RatePlan) ProtoMessage()    {}

func (m *RatePlan) GetRoomType() *RoomType {
	if m != nil {
		return m.RoomType
	}
	return nil
}

type RoomType struct {
	BookableRate       float64 `protobuf:"fixed64,1,opt,name=bookableRate" json:"bookableRate,omitempty"`
	TotalRate          float64 `protobuf:"fixed64,2,opt,name=totalRate" json:"totalRate,omitempty"`
	TotalRateInclusive float64 `protobuf:"fixed64,3,opt" json:"TotalRateInclusive,omitempty"`
	Code               string  `protobuf:"bytes,4,opt,name=code" json:"code,omitempty"`
	Currency           string  `protobuf:"bytes,5,opt,name=currency" json:"currency,omitempty"`
	RoomDescription    string  `protobuf:"bytes,6,opt,name=roomDescription" json:"roomDescription,omitempty"`
}

func (m *RoomType) Reset()         { *m = RoomType{} }
func (m *RoomType) String() string { return proto.CompactTextString(m) }
func (*RoomType) ProtoMessage()    {}

// Client API for Rate service

type RateClient interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Reply, error)
}

type rateClient struct {
	cc *grpc.ClientConn
}

func NewRateClient(cc *grpc.ClientConn) RateClient {
	return &rateClient{cc}
}

func (c *rateClient) GetRates(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/rate.Rate/GetRates", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rate service

type RateServer interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(context.Context, *Args) (*Reply, error)
}

func RegisterRateServer(s *grpc.Server, srv RateServer) {
	s.RegisterService(&_Rate_serviceDesc, srv)
}

func _Rate_GetRates_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Args)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(RateServer).GetRates(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Rate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rate.Rate",
	HandlerType: (*RateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRates",
			Handler:    _Rate_GetRates_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
