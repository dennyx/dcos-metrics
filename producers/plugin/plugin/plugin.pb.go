// Code generated by protoc-gen-go.
// source: plugin.proto
// DO NOT EDIT!

/*
Package plugin is a generated protocol buffer package.

It is generated from these files:
	plugin.proto

It has these top-level messages:
	MetricsCollectorType
	MetricsMessage
	Datapoint
	Dimensions
*/
package plugin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Specifies the type of metrics message to recieve. This is based on the
// collectors available. Currently we have `node` and `framework` collectors.
type MetricsCollectorType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *MetricsCollectorType) Reset()                    { *m = MetricsCollectorType{} }
func (m *MetricsCollectorType) String() string            { return proto.CompactTextString(m) }
func (*MetricsCollectorType) ProtoMessage()               {}
func (*MetricsCollectorType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MetricsCollectorType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type MetricsMessage struct {
	Name       string       `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Datapoints []*Datapoint `protobuf:"bytes,2,rep,name=datapoints" json:"datapoints,omitempty"`
	Dimensions *Dimensions  `protobuf:"bytes,3,opt,name=dimensions" json:"dimensions,omitempty"`
	Timestamp  int64        `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *MetricsMessage) Reset()                    { *m = MetricsMessage{} }
func (m *MetricsMessage) String() string            { return proto.CompactTextString(m) }
func (*MetricsMessage) ProtoMessage()               {}
func (*MetricsMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MetricsMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricsMessage) GetDatapoints() []*Datapoint {
	if m != nil {
		return m.Datapoints
	}
	return nil
}

func (m *MetricsMessage) GetDimensions() *Dimensions {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *MetricsMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Datapoint struct {
	Timestamp string `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value     string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *Datapoint) Reset()                    { *m = Datapoint{} }
func (m *Datapoint) String() string            { return proto.CompactTextString(m) }
func (*Datapoint) ProtoMessage()               {}
func (*Datapoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Datapoint) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Datapoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Datapoint) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Dimensions struct {
	MesosID            string            `protobuf:"bytes,1,opt,name=mesosID" json:"mesosID,omitempty"`
	ClusterID          string            `protobuf:"bytes,2,opt,name=clusterID" json:"clusterID,omitempty"`
	ContainerID        string            `protobuf:"bytes,3,opt,name=containerID" json:"containerID,omitempty"`
	ExecutorID         string            `protobuf:"bytes,4,opt,name=executorID" json:"executorID,omitempty"`
	FrameworkName      string            `protobuf:"bytes,5,opt,name=frameworkName" json:"frameworkName,omitempty"`
	FrameworkID        string            `protobuf:"bytes,6,opt,name=frameworkID" json:"frameworkID,omitempty"`
	FrameworkRole      string            `protobuf:"bytes,7,opt,name=frameworkRole" json:"frameworkRole,omitempty"`
	FrameworkPrincipal string            `protobuf:"bytes,8,opt,name=frameworkPrincipal" json:"frameworkPrincipal,omitempty"`
	Hostname           string            `protobuf:"bytes,9,opt,name=hostname" json:"hostname,omitempty"`
	Labels             map[string]string `protobuf:"bytes,10,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Dimensions) Reset()                    { *m = Dimensions{} }
func (m *Dimensions) String() string            { return proto.CompactTextString(m) }
func (*Dimensions) ProtoMessage()               {}
func (*Dimensions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Dimensions) GetMesosID() string {
	if m != nil {
		return m.MesosID
	}
	return ""
}

func (m *Dimensions) GetClusterID() string {
	if m != nil {
		return m.ClusterID
	}
	return ""
}

func (m *Dimensions) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *Dimensions) GetExecutorID() string {
	if m != nil {
		return m.ExecutorID
	}
	return ""
}

func (m *Dimensions) GetFrameworkName() string {
	if m != nil {
		return m.FrameworkName
	}
	return ""
}

func (m *Dimensions) GetFrameworkID() string {
	if m != nil {
		return m.FrameworkID
	}
	return ""
}

func (m *Dimensions) GetFrameworkRole() string {
	if m != nil {
		return m.FrameworkRole
	}
	return ""
}

func (m *Dimensions) GetFrameworkPrincipal() string {
	if m != nil {
		return m.FrameworkPrincipal
	}
	return ""
}

func (m *Dimensions) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Dimensions) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricsCollectorType)(nil), "plugin.MetricsCollectorType")
	proto.RegisterType((*MetricsMessage)(nil), "plugin.MetricsMessage")
	proto.RegisterType((*Datapoint)(nil), "plugin.Datapoint")
	proto.RegisterType((*Dimensions)(nil), "plugin.Dimensions")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Metrics service

type MetricsClient interface {
	AttachOutputStream(ctx context.Context, in *MetricsCollectorType, opts ...grpc.CallOption) (Metrics_AttachOutputStreamClient, error)
}

type metricsClient struct {
	cc *grpc.ClientConn
}

func NewMetricsClient(cc *grpc.ClientConn) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) AttachOutputStream(ctx context.Context, in *MetricsCollectorType, opts ...grpc.CallOption) (Metrics_AttachOutputStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Metrics_serviceDesc.Streams[0], c.cc, "/plugin.Metrics/AttachOutputStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsAttachOutputStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Metrics_AttachOutputStreamClient interface {
	Recv() (*MetricsMessage, error)
	grpc.ClientStream
}

type metricsAttachOutputStreamClient struct {
	grpc.ClientStream
}

func (x *metricsAttachOutputStreamClient) Recv() (*MetricsMessage, error) {
	m := new(MetricsMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Metrics service

type MetricsServer interface {
	AttachOutputStream(*MetricsCollectorType, Metrics_AttachOutputStreamServer) error
}

func RegisterMetricsServer(s *grpc.Server, srv MetricsServer) {
	s.RegisterService(&_Metrics_serviceDesc, srv)
}

func _Metrics_AttachOutputStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MetricsCollectorType)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetricsServer).AttachOutputStream(m, &metricsAttachOutputStreamServer{stream})
}

type Metrics_AttachOutputStreamServer interface {
	Send(*MetricsMessage) error
	grpc.ServerStream
}

type metricsAttachOutputStreamServer struct {
	grpc.ServerStream
}

func (x *metricsAttachOutputStreamServer) Send(m *MetricsMessage) error {
	return x.ServerStream.SendMsg(m)
}

var _Metrics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "plugin.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AttachOutputStream",
			Handler:       _Metrics_AttachOutputStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "plugin.proto",
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0x9c, 0x26, 0xf5, 0x04, 0x10, 0x8c, 0x2a, 0x64, 0x45, 0x55, 0x65, 0x59, 0x1c,
	0x22, 0x0e, 0x11, 0x04, 0x09, 0x01, 0x37, 0x44, 0x38, 0x44, 0xa2, 0x05, 0xb9, 0x5c, 0x38, 0x6e,
	0xdd, 0xa1, 0x5d, 0x75, 0xbd, 0x6b, 0xed, 0x8e, 0x81, 0x3c, 0x12, 0x67, 0x5e, 0x10, 0x79, 0xed,
	0xd8, 0x4e, 0xc9, 0x6d, 0xe7, 0x9f, 0xcf, 0xf3, 0x8f, 0x77, 0x66, 0xe1, 0x61, 0xa9, 0xaa, 0x1b,
	0xa9, 0x97, 0xa5, 0x35, 0x6c, 0x70, 0xd2, 0x44, 0xe9, 0x0b, 0x38, 0x39, 0x27, 0xb6, 0x32, 0x77,
	0x1f, 0x8d, 0x52, 0x94, 0xb3, 0xb1, 0xdf, 0xb6, 0x25, 0x21, 0xc2, 0x98, 0xb7, 0x25, 0xc5, 0x41,
	0x12, 0x2c, 0xa2, 0xcc, 0x9f, 0xd3, 0x3f, 0x01, 0x3c, 0x6e, 0xe1, 0x73, 0x72, 0x4e, 0xdc, 0x78,
	0x4c, 0x8b, 0xa2, 0xc3, 0xea, 0x33, 0xbe, 0x02, 0xb8, 0x16, 0x2c, 0x4a, 0x23, 0x35, 0xbb, 0x78,
	0x94, 0x84, 0x8b, 0xd9, 0xea, 0xe9, 0xb2, 0x75, 0x5f, 0xef, 0x32, 0xd9, 0x00, 0xc2, 0x15, 0xc0,
	0xb5, 0x2c, 0x48, 0x3b, 0x69, 0xb4, 0x8b, 0xc3, 0x24, 0x58, 0xcc, 0x56, 0xd8, 0x7d, 0xd2, 0x65,
	0xb2, 0x01, 0x85, 0xa7, 0x10, 0xb1, 0x2c, 0xc8, 0xb1, 0x28, 0xca, 0x78, 0x9c, 0x04, 0x8b, 0x30,
	0xeb, 0x85, 0xf4, 0x12, 0xa2, 0xce, 0x6a, 0x1f, 0x6d, 0x5a, 0xed, 0x85, 0xee, 0x1f, 0x46, 0x83,
	0x7f, 0x38, 0x81, 0xa3, 0x9f, 0x42, 0x55, 0xe4, 0x7b, 0x89, 0xb2, 0x26, 0x48, 0xff, 0x86, 0x00,
	0x7d, 0x37, 0x18, 0xc3, 0xb4, 0x20, 0x67, 0xdc, 0x66, 0xdd, 0x16, 0xdd, 0x85, 0xb5, 0x61, 0xae,
	0x2a, 0xc7, 0x64, 0x37, 0xeb, 0xb6, 0x6e, 0x2f, 0x60, 0x02, 0xb3, 0xdc, 0x68, 0x16, 0x52, 0xfb,
	0x7c, 0x63, 0x31, 0x94, 0xf0, 0x0c, 0x80, 0x7e, 0x53, 0x5e, 0xb1, 0xa9, 0x81, 0xb1, 0x07, 0x06,
	0x0a, 0x3e, 0x87, 0x47, 0x3f, 0xac, 0x28, 0xe8, 0x97, 0xb1, 0x77, 0x17, 0x75, 0xef, 0x47, 0x1e,
	0xd9, 0x17, 0x6b, 0x9f, 0x4e, 0xd8, 0xac, 0xe3, 0x49, 0xe3, 0x33, 0x90, 0xf6, 0xea, 0x64, 0x46,
	0x51, 0x3c, 0xbd, 0x57, 0xa7, 0x16, 0x71, 0x09, 0xd8, 0x09, 0x5f, 0xad, 0xd4, 0xb9, 0x2c, 0x85,
	0x8a, 0x8f, 0x3d, 0x7a, 0x20, 0x83, 0x73, 0x38, 0xbe, 0x35, 0x8e, 0xfd, 0xa5, 0x46, 0x9e, 0xea,
	0x62, 0x7c, 0x03, 0x13, 0x25, 0xae, 0x48, 0xb9, 0x18, 0xfc, 0x62, 0x9c, 0xfd, 0x3f, 0xe5, 0xe5,
	0x67, 0x0f, 0x7c, 0xd2, 0x6c, 0xb7, 0x59, 0x4b, 0xcf, 0xdf, 0xc1, 0x6c, 0x20, 0xe3, 0x13, 0x08,
	0xef, 0x68, 0xdb, 0x5e, 0x7b, 0x7d, 0xec, 0x27, 0x36, 0x1a, 0x4c, 0xec, 0xfd, 0xe8, 0x6d, 0xb0,
	0xfa, 0x0e, 0xd3, 0x76, 0x6b, 0xf1, 0x02, 0xf0, 0x03, 0xb3, 0xc8, 0x6f, 0xbf, 0x54, 0x5c, 0x56,
	0x7c, 0xc9, 0x96, 0x44, 0x81, 0xa7, 0xbb, 0x1e, 0x0e, 0xbd, 0x84, 0xf9, 0xb3, 0x7b, 0xd9, 0x76,
	0xf5, 0xd3, 0x07, 0x2f, 0x83, 0xab, 0x89, 0x7f, 0x4c, 0xaf, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xdd, 0x79, 0x4d, 0x4f, 0x5c, 0x03, 0x00, 0x00,
}
