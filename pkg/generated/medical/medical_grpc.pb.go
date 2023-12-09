// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: medical.proto

package medical

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MedicalServiceClient is the client API for MedicalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedicalServiceClient interface {
	CreateHospital(ctx context.Context, in *CreateHospitalRequest, opts ...grpc.CallOption) (*Hospital, error)
	CreatePatient(ctx context.Context, in *CreatePatientRequest, opts ...grpc.CallOption) (*Patient, error)
	RegisterPatient(ctx context.Context, in *RegisterPatientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DischargePatient(ctx context.Context, in *DiscardPatientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetStatOfPatients(ctx context.Context, in *GetStatOfPatientsRequest, opts ...grpc.CallOption) (*PatientsStats, error)
}

type medicalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMedicalServiceClient(cc grpc.ClientConnInterface) MedicalServiceClient {
	return &medicalServiceClient{cc}
}

func (c *medicalServiceClient) CreateHospital(ctx context.Context, in *CreateHospitalRequest, opts ...grpc.CallOption) (*Hospital, error) {
	out := new(Hospital)
	err := c.cc.Invoke(ctx, "/medical.MedicalService/CreateHospital", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalServiceClient) CreatePatient(ctx context.Context, in *CreatePatientRequest, opts ...grpc.CallOption) (*Patient, error) {
	out := new(Patient)
	err := c.cc.Invoke(ctx, "/medical.MedicalService/CreatePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalServiceClient) RegisterPatient(ctx context.Context, in *RegisterPatientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/medical.MedicalService/RegisterPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalServiceClient) DischargePatient(ctx context.Context, in *DiscardPatientRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/medical.MedicalService/DischargePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalServiceClient) GetStatOfPatients(ctx context.Context, in *GetStatOfPatientsRequest, opts ...grpc.CallOption) (*PatientsStats, error) {
	out := new(PatientsStats)
	err := c.cc.Invoke(ctx, "/medical.MedicalService/GetStatOfPatients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedicalServiceServer is the server API for MedicalService service.
// All implementations must embed UnimplementedMedicalServiceServer
// for forward compatibility
type MedicalServiceServer interface {
	CreateHospital(context.Context, *CreateHospitalRequest) (*Hospital, error)
	CreatePatient(context.Context, *CreatePatientRequest) (*Patient, error)
	RegisterPatient(context.Context, *RegisterPatientRequest) (*emptypb.Empty, error)
	DischargePatient(context.Context, *DiscardPatientRequest) (*emptypb.Empty, error)
	GetStatOfPatients(context.Context, *GetStatOfPatientsRequest) (*PatientsStats, error)
	mustEmbedUnimplementedMedicalServiceServer()
}

// UnimplementedMedicalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMedicalServiceServer struct {
}

func (UnimplementedMedicalServiceServer) CreateHospital(context.Context, *CreateHospitalRequest) (*Hospital, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHospital not implemented")
}
func (UnimplementedMedicalServiceServer) CreatePatient(context.Context, *CreatePatientRequest) (*Patient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatient not implemented")
}
func (UnimplementedMedicalServiceServer) RegisterPatient(context.Context, *RegisterPatientRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPatient not implemented")
}
func (UnimplementedMedicalServiceServer) DischargePatient(context.Context, *DiscardPatientRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DischargePatient not implemented")
}
func (UnimplementedMedicalServiceServer) GetStatOfPatients(context.Context, *GetStatOfPatientsRequest) (*PatientsStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatOfPatients not implemented")
}
func (UnimplementedMedicalServiceServer) mustEmbedUnimplementedMedicalServiceServer() {}

// UnsafeMedicalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedicalServiceServer will
// result in compilation errors.
type UnsafeMedicalServiceServer interface {
	mustEmbedUnimplementedMedicalServiceServer()
}

func RegisterMedicalServiceServer(s grpc.ServiceRegistrar, srv MedicalServiceServer) {
	s.RegisterService(&MedicalService_ServiceDesc, srv)
}

func _MedicalService_CreateHospital_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHospitalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalServiceServer).CreateHospital(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical.MedicalService/CreateHospital",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalServiceServer).CreateHospital(ctx, req.(*CreateHospitalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalService_CreatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalServiceServer).CreatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical.MedicalService/CreatePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalServiceServer).CreatePatient(ctx, req.(*CreatePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalService_RegisterPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalServiceServer).RegisterPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical.MedicalService/RegisterPatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalServiceServer).RegisterPatient(ctx, req.(*RegisterPatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalService_DischargePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscardPatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalServiceServer).DischargePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical.MedicalService/DischargePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalServiceServer).DischargePatient(ctx, req.(*DiscardPatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalService_GetStatOfPatients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatOfPatientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalServiceServer).GetStatOfPatients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical.MedicalService/GetStatOfPatients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalServiceServer).GetStatOfPatients(ctx, req.(*GetStatOfPatientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MedicalService_ServiceDesc is the grpc.ServiceDesc for MedicalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedicalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "medical.MedicalService",
	HandlerType: (*MedicalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHospital",
			Handler:    _MedicalService_CreateHospital_Handler,
		},
		{
			MethodName: "CreatePatient",
			Handler:    _MedicalService_CreatePatient_Handler,
		},
		{
			MethodName: "RegisterPatient",
			Handler:    _MedicalService_RegisterPatient_Handler,
		},
		{
			MethodName: "DischargePatient",
			Handler:    _MedicalService_DischargePatient_Handler,
		},
		{
			MethodName: "GetStatOfPatients",
			Handler:    _MedicalService_GetStatOfPatients_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "medical.proto",
}