// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// DashboardServiceClient is the client API for DashboardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DashboardServiceClient interface {
	// StartLocalDocument starts a Document on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the dashboard/config.yaml
	//   3. all bytes constituting the Document YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalDocument(ctx context.Context, opts ...grpc.CallOption) (DashboardService_StartLocalDocumentClient, error)
	// StartFromPreviousDocument starts a new Document based on a previous one.
	// If the previous Document does not have the can-replay condition set this call will result in an error.
	StartFromPreviousDocument(ctx context.Context, in *StartFromPreviousDocumentRequest, opts ...grpc.CallOption) (*StartDocumentResponse, error)
	// StartDocumentRequest starts a new Document based on its specification.
	StartDocument(ctx context.Context, in *StartDocumentRequest, opts ...grpc.CallOption) (*StartDocumentResponse, error)
	// Searches for Document known to this instance
	ListDocument(ctx context.Context, in *ListDocumentRequest, opts ...grpc.CallOption) (*ListDocumentResponse, error)
	// Subscribe listens to new Document updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (DashboardService_SubscribeClient, error)
	// GetDocument retrieves details of a single Document
	GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error)
	// Listen listens to Document updates and log output of a running Document
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (DashboardService_ListenClient, error)
	// StopDocument stops a currently running Document
	StopDocument(ctx context.Context, in *StopDocumentRequest, opts ...grpc.CallOption) (*StopDocumentResponse, error)
}

type dashboardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDashboardServiceClient(cc grpc.ClientConnInterface) DashboardServiceClient {
	return &dashboardServiceClient{cc}
}

func (c *dashboardServiceClient) StartLocalDocument(ctx context.Context, opts ...grpc.CallOption) (DashboardService_StartLocalDocumentClient, error) {
	stream, err := c.cc.NewStream(ctx, &DashboardService_ServiceDesc.Streams[0], "/v1.DashboardService/StartLocalDocument", opts...)
	if err != nil {
		return nil, err
	}
	x := &dashboardServiceStartLocalDocumentClient{stream}
	return x, nil
}

type DashboardService_StartLocalDocumentClient interface {
	Send(*StartLocalDocumentRequest) error
	CloseAndRecv() (*StartDocumentResponse, error)
	grpc.ClientStream
}

type dashboardServiceStartLocalDocumentClient struct {
	grpc.ClientStream
}

func (x *dashboardServiceStartLocalDocumentClient) Send(m *StartLocalDocumentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dashboardServiceStartLocalDocumentClient) CloseAndRecv() (*StartDocumentResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartDocumentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dashboardServiceClient) StartFromPreviousDocument(ctx context.Context, in *StartFromPreviousDocumentRequest, opts ...grpc.CallOption) (*StartDocumentResponse, error) {
	out := new(StartDocumentResponse)
	err := c.cc.Invoke(ctx, "/v1.DashboardService/StartFromPreviousDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardServiceClient) StartDocument(ctx context.Context, in *StartDocumentRequest, opts ...grpc.CallOption) (*StartDocumentResponse, error) {
	out := new(StartDocumentResponse)
	err := c.cc.Invoke(ctx, "/v1.DashboardService/StartDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardServiceClient) ListDocument(ctx context.Context, in *ListDocumentRequest, opts ...grpc.CallOption) (*ListDocumentResponse, error) {
	out := new(ListDocumentResponse)
	err := c.cc.Invoke(ctx, "/v1.DashboardService/ListDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (DashboardService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &DashboardService_ServiceDesc.Streams[1], "/v1.DashboardService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &dashboardServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DashboardService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type dashboardServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *dashboardServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dashboardServiceClient) GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error) {
	out := new(GetDocumentResponse)
	err := c.cc.Invoke(ctx, "/v1.DashboardService/GetDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (DashboardService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &DashboardService_ServiceDesc.Streams[2], "/v1.DashboardService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &dashboardServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DashboardService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type dashboardServiceListenClient struct {
	grpc.ClientStream
}

func (x *dashboardServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dashboardServiceClient) StopDocument(ctx context.Context, in *StopDocumentRequest, opts ...grpc.CallOption) (*StopDocumentResponse, error) {
	out := new(StopDocumentResponse)
	err := c.cc.Invoke(ctx, "/v1.DashboardService/StopDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardServiceServer is the server API for DashboardService service.
// All implementations must embed UnimplementedDashboardServiceServer
// for forward compatibility
type DashboardServiceServer interface {
	// StartLocalDocument starts a Document on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the dashboard/config.yaml
	//   3. all bytes constituting the Document YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalDocument(DashboardService_StartLocalDocumentServer) error
	// StartFromPreviousDocument starts a new Document based on a previous one.
	// If the previous Document does not have the can-replay condition set this call will result in an error.
	StartFromPreviousDocument(context.Context, *StartFromPreviousDocumentRequest) (*StartDocumentResponse, error)
	// StartDocumentRequest starts a new Document based on its specification.
	StartDocument(context.Context, *StartDocumentRequest) (*StartDocumentResponse, error)
	// Searches for Document known to this instance
	ListDocument(context.Context, *ListDocumentRequest) (*ListDocumentResponse, error)
	// Subscribe listens to new Document updates
	Subscribe(*SubscribeRequest, DashboardService_SubscribeServer) error
	// GetDocument retrieves details of a single Document
	GetDocument(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error)
	// Listen listens to Document updates and log output of a running Document
	Listen(*ListenRequest, DashboardService_ListenServer) error
	// StopDocument stops a currently running Document
	StopDocument(context.Context, *StopDocumentRequest) (*StopDocumentResponse, error)
	mustEmbedUnimplementedDashboardServiceServer()
}

// UnimplementedDashboardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDashboardServiceServer struct {
}

func (UnimplementedDashboardServiceServer) StartLocalDocument(DashboardService_StartLocalDocumentServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalDocument not implemented")
}
func (UnimplementedDashboardServiceServer) StartFromPreviousDocument(context.Context, *StartFromPreviousDocumentRequest) (*StartDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousDocument not implemented")
}
func (UnimplementedDashboardServiceServer) StartDocument(context.Context, *StartDocumentRequest) (*StartDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartDocument not implemented")
}
func (UnimplementedDashboardServiceServer) ListDocument(context.Context, *ListDocumentRequest) (*ListDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDocument not implemented")
}
func (UnimplementedDashboardServiceServer) Subscribe(*SubscribeRequest, DashboardService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedDashboardServiceServer) GetDocument(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDocument not implemented")
}
func (UnimplementedDashboardServiceServer) Listen(*ListenRequest, DashboardService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedDashboardServiceServer) StopDocument(context.Context, *StopDocumentRequest) (*StopDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopDocument not implemented")
}
func (UnimplementedDashboardServiceServer) mustEmbedUnimplementedDashboardServiceServer() {}

// UnsafeDashboardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DashboardServiceServer will
// result in compilation errors.
type UnsafeDashboardServiceServer interface {
	mustEmbedUnimplementedDashboardServiceServer()
}

func RegisterDashboardServiceServer(s grpc.ServiceRegistrar, srv DashboardServiceServer) {
	s.RegisterService(&DashboardService_ServiceDesc, srv)
}

func _DashboardService_StartLocalDocument_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DashboardServiceServer).StartLocalDocument(&dashboardServiceStartLocalDocumentServer{stream})
}

type DashboardService_StartLocalDocumentServer interface {
	SendAndClose(*StartDocumentResponse) error
	Recv() (*StartLocalDocumentRequest, error)
	grpc.ServerStream
}

type dashboardServiceStartLocalDocumentServer struct {
	grpc.ServerStream
}

func (x *dashboardServiceStartLocalDocumentServer) SendAndClose(m *StartDocumentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dashboardServiceStartLocalDocumentServer) Recv() (*StartLocalDocumentRequest, error) {
	m := new(StartLocalDocumentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DashboardService_StartFromPreviousDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).StartFromPreviousDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DashboardService/StartFromPreviousDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).StartFromPreviousDocument(ctx, req.(*StartFromPreviousDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_StartDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).StartDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DashboardService/StartDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).StartDocument(ctx, req.(*StartDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_ListDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).ListDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DashboardService/ListDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).ListDocument(ctx, req.(*ListDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DashboardServiceServer).Subscribe(m, &dashboardServiceSubscribeServer{stream})
}

type DashboardService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type dashboardServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *dashboardServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DashboardService_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DashboardService/GetDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).GetDocument(ctx, req.(*GetDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DashboardServiceServer).Listen(m, &dashboardServiceListenServer{stream})
}

type DashboardService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type dashboardServiceListenServer struct {
	grpc.ServerStream
}

func (x *dashboardServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DashboardService_StopDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).StopDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DashboardService/StopDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).StopDocument(ctx, req.(*StopDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DashboardService_ServiceDesc is the grpc.ServiceDesc for DashboardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DashboardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DashboardService",
	HandlerType: (*DashboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousDocument",
			Handler:    _DashboardService_StartFromPreviousDocument_Handler,
		},
		{
			MethodName: "StartDocument",
			Handler:    _DashboardService_StartDocument_Handler,
		},
		{
			MethodName: "ListDocument",
			Handler:    _DashboardService_ListDocument_Handler,
		},
		{
			MethodName: "GetDocument",
			Handler:    _DashboardService_GetDocument_Handler,
		},
		{
			MethodName: "StopDocument",
			Handler:    _DashboardService_StopDocument_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalDocument",
			Handler:       _DashboardService_StartLocalDocument_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _DashboardService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _DashboardService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dashboard.proto",
}
