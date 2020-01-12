// Code generated by protoc-gen-go. DO NOT EDIT.
// source: search.proto

package search

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type Result struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Result) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Result)(nil), "Result")
}

func init() { proto.RegisterFile("search.proto", fileDescriptor_453745cff914010e) }

var fileDescriptor_453745cff914010e = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe7, 0x62, 0x0f, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0x2c, 0x4d, 0x2d, 0xaa, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x0c, 0xb8, 0xd8, 0x82, 0x52, 0x8b, 0x4b, 0x73, 0xc0, 0xf2,
	0x25, 0x99, 0x25, 0x39, 0xa9, 0x30, 0x79, 0x30, 0x47, 0x48, 0x80, 0x8b, 0xb9, 0xb4, 0x28, 0x47,
	0x82, 0x09, 0x2c, 0x06, 0x62, 0x1a, 0x85, 0x73, 0xf1, 0x06, 0x83, 0xad, 0x08, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x15, 0x92, 0xe5, 0x62, 0x83, 0x08, 0x08, 0x71, 0xe8, 0x41, 0x2d, 0x93, 0x62,
	0xd7, 0x83, 0x98, 0xaa, 0xc4, 0x20, 0xa4, 0xce, 0xc5, 0xe7, 0x5b, 0x9a, 0x53, 0x92, 0x59, 0x90,
	0x93, 0x8a, 0x47, 0x99, 0x01, 0x63, 0x12, 0x1b, 0xd8, 0xc9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x77, 0x2e, 0xc5, 0x2a, 0xc2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	// Search returns a search result for the query.
	Search(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
	MultipleSearch(ctx context.Context, in *Request, opts ...grpc.CallOption) (SearchService_MultipleSearchClient, error)
}

type searchServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchServiceClient(cc *grpc.ClientConn) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) MultipleSearch(ctx context.Context, in *Request, opts ...grpc.CallOption) (SearchService_MultipleSearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SearchService_serviceDesc.Streams[0], "/SearchService/MultipleSearch", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchServiceMultipleSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SearchService_MultipleSearchClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type searchServiceMultipleSearchClient struct {
	grpc.ClientStream
}

func (x *searchServiceMultipleSearchClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	// Search returns a search result for the query.
	Search(context.Context, *Request) (*Result, error)
	MultipleSearch(*Request, SearchService_MultipleSearchServer) error
}

// UnimplementedSearchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (*UnimplementedSearchServiceServer) Search(ctx context.Context, req *Request) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedSearchServiceServer) MultipleSearch(req *Request, srv SearchService_MultipleSearchServer) error {
	return status.Errorf(codes.Unimplemented, "method MultipleSearch not implemented")
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_MultipleSearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SearchServiceServer).MultipleSearch(m, &searchServiceMultipleSearchServer{stream})
}

type SearchService_MultipleSearchServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type searchServiceMultipleSearchServer struct {
	grpc.ServerStream
}

func (x *searchServiceMultipleSearchServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MultipleSearch",
			Handler:       _SearchService_MultipleSearch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "search.proto",
}
