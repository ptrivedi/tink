// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package template

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

// TemplateServiceClient is the client API for TemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateServiceClient interface {
	//
	// CreateTemplate stores a template in the Tinkerbell server.
	CreateTemplate(ctx context.Context, in *WorkflowTemplate, opts ...grpc.CallOption) (*CreateResponse, error)
	//
	// GetTemplate returns a specific template via its identifier.
	GetTemplate(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
	//
	// DeleteTemplate deletes a template via its identifier.
	DeleteTemplate(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Empty, error)
	//
	// ListTemplates returns all the template stored in Tinkerbell server
	ListTemplates(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (TemplateService_ListTemplatesClient, error)
	//
	// UpdateTemplate updates a template
	// TODO: Read the code and figure out how an update work
	UpdateTemplate(ctx context.Context, in *WorkflowTemplate, opts ...grpc.CallOption) (*Empty, error)
}

type templateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateServiceClient(cc grpc.ClientConnInterface) TemplateServiceClient {
	return &templateServiceClient{cc}
}

func (c *templateServiceClient) CreateTemplate(ctx context.Context, in *WorkflowTemplate, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.template.TemplateService/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) GetTemplate(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.template.TemplateService/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) DeleteTemplate(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.template.TemplateService/DeleteTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateServiceClient) ListTemplates(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (TemplateService_ListTemplatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &TemplateService_ServiceDesc.Streams[0], "/github.com.tinkerbell.tink.protos.template.TemplateService/ListTemplates", opts...)
	if err != nil {
		return nil, err
	}
	x := &templateServiceListTemplatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TemplateService_ListTemplatesClient interface {
	Recv() (*WorkflowTemplate, error)
	grpc.ClientStream
}

type templateServiceListTemplatesClient struct {
	grpc.ClientStream
}

func (x *templateServiceListTemplatesClient) Recv() (*WorkflowTemplate, error) {
	m := new(WorkflowTemplate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *templateServiceClient) UpdateTemplate(ctx context.Context, in *WorkflowTemplate, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/github.com.tinkerbell.tink.protos.template.TemplateService/UpdateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServiceServer is the server API for TemplateService service.
// All implementations should embed UnimplementedTemplateServiceServer
// for forward compatibility
type TemplateServiceServer interface {
	//
	// CreateTemplate stores a template in the Tinkerbell server.
	CreateTemplate(context.Context, *WorkflowTemplate) (*CreateResponse, error)
	//
	// GetTemplate returns a specific template via its identifier.
	GetTemplate(context.Context, *GetRequest) (*WorkflowTemplate, error)
	//
	// DeleteTemplate deletes a template via its identifier.
	DeleteTemplate(context.Context, *GetRequest) (*Empty, error)
	//
	// ListTemplates returns all the template stored in Tinkerbell server
	ListTemplates(*ListRequest, TemplateService_ListTemplatesServer) error
	//
	// UpdateTemplate updates a template
	// TODO: Read the code and figure out how an update work
	UpdateTemplate(context.Context, *WorkflowTemplate) (*Empty, error)
}

// UnimplementedTemplateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTemplateServiceServer struct{}

func (UnimplementedTemplateServiceServer) CreateTemplate(context.Context, *WorkflowTemplate) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}

func (UnimplementedTemplateServiceServer) GetTemplate(context.Context, *GetRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplate not implemented")
}

func (UnimplementedTemplateServiceServer) DeleteTemplate(context.Context, *GetRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplate not implemented")
}

func (UnimplementedTemplateServiceServer) ListTemplates(*ListRequest, TemplateService_ListTemplatesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTemplates not implemented")
}

func (UnimplementedTemplateServiceServer) UpdateTemplate(context.Context, *WorkflowTemplate) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemplate not implemented")
}

// UnsafeTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateServiceServer will
// result in compilation errors.
type UnsafeTemplateServiceServer interface {
	mustEmbedUnimplementedTemplateServiceServer()
}

func RegisterTemplateServiceServer(s grpc.ServiceRegistrar, srv TemplateServiceServer) {
	s.RegisterService(&TemplateService_ServiceDesc, srv)
}

func _TemplateService_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.template.TemplateService/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).CreateTemplate(ctx, req.(*WorkflowTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateService_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.template.TemplateService/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).GetTemplate(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateService_DeleteTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).DeleteTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.template.TemplateService/DeleteTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).DeleteTemplate(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemplateService_ListTemplates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TemplateServiceServer).ListTemplates(m, &templateServiceListTemplatesServer{stream})
}

type TemplateService_ListTemplatesServer interface {
	Send(*WorkflowTemplate) error
	grpc.ServerStream
}

type templateServiceListTemplatesServer struct {
	grpc.ServerStream
}

func (x *templateServiceListTemplatesServer) Send(m *WorkflowTemplate) error {
	return x.ServerStream.SendMsg(m)
}

func _TemplateService_UpdateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).UpdateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.tinkerbell.tink.protos.template.TemplateService/UpdateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).UpdateTemplate(ctx, req.(*WorkflowTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

// TemplateService_ServiceDesc is the grpc.ServiceDesc for TemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.tinkerbell.tink.protos.template.TemplateService",
	HandlerType: (*TemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTemplate",
			Handler:    _TemplateService_CreateTemplate_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _TemplateService_GetTemplate_Handler,
		},
		{
			MethodName: "DeleteTemplate",
			Handler:    _TemplateService_DeleteTemplate_Handler,
		},
		{
			MethodName: "UpdateTemplate",
			Handler:    _TemplateService_UpdateTemplate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTemplates",
			Handler:       _TemplateService_ListTemplates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/template/template.proto",
}
