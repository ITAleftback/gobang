// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: user.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Surrender struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Surrender) Reset() {
	*x = Surrender{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Surrender) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Surrender) ProtoMessage() {}

func (x *Surrender) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Surrender.ProtoReflect.Descriptor instead.
func (*Surrender) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type Playchess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I int32 `protobuf:"varint,1,opt,name=i,proto3" json:"i,omitempty"`
	J int32 `protobuf:"varint,2,opt,name=j,proto3" json:"j,omitempty"`
}

func (x *Playchess) Reset() {
	*x = Playchess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Playchess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Playchess) ProtoMessage() {}

func (x *Playchess) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Playchess.ProtoReflect.Descriptor instead.
func (*Playchess) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *Playchess) GetI() int32 {
	if x != nil {
		return x.I
	}
	return 0
}

func (x *Playchess) GetJ() int32 {
	if x != nil {
		return x.J
	}
	return 0
}

type Ready struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ready) Reset() {
	*x = Ready{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ready) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ready) ProtoMessage() {}

func (x *Ready) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ready.ProtoReflect.Descriptor instead.
func (*Ready) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

type Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Login) Reset() {
	*x = Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login) ProtoMessage() {}

func (x *Login) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login.ProtoReflect.Descriptor instead.
func (*Login) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *Login) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Login) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Register struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Register) Reset() {
	*x = Register{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register) ProtoMessage() {}

func (x *Register) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register.ProtoReflect.Descriptor instead.
func (*Register) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *Register) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Register) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Reply2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string         `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	J       []*Reply2_Edge `protobuf:"bytes,2,rep,name=j,proto3" json:"j,omitempty"`
}

func (x *Reply2) Reset() {
	*x = Reply2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply2) ProtoMessage() {}

func (x *Reply2) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply2.ProtoReflect.Descriptor instead.
func (*Reply2) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *Reply2) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Reply2) GetJ() []*Reply2_Edge {
	if x != nil {
		return x.J
	}
	return nil
}

type Reply2_Edge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I []int32 `protobuf:"varint,1,rep,packed,name=i,proto3" json:"i,omitempty"`
}

func (x *Reply2_Edge) Reset() {
	*x = Reply2_Edge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply2_Edge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply2_Edge) ProtoMessage() {}

func (x *Reply2_Edge) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply2_Edge.ProtoReflect.Descriptor instead.
func (*Reply2_Edge) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6, 0}
}

func (x *Reply2_Edge) GetI() []int32 {
	if x != nil {
		return x.I
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x73, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x22, 0x27, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x63, 0x68, 0x65, 0x73, 0x73, 0x12, 0x0c, 0x0a,
	0x01, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x69, 0x12, 0x0c, 0x0a, 0x01, 0x6a,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x6a, 0x22, 0x07, 0x0a, 0x05, 0x72, 0x65, 0x61,
	0x64, 0x79, 0x22, 0x3f, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x42, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x21, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5a, 0x0a, 0x06, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20,
	0x0a, 0x01, 0x6a, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x52, 0x01, 0x6a,
	0x1a, 0x14, 0x0a, 0x04, 0x45, 0x64, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x69, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x01, 0x69, 0x32, 0x2e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x25, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x37, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x2b, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32,
	0x3a, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x63, 0x68, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x09,
	0x50, 0x6c, 0x61, 0x79, 0x63, 0x68, 0x65, 0x73, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x63, 0x68, 0x65, 0x73, 0x73, 0x1a, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x2e, 0x0a, 0x05, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x3a, 0x0a, 0x09, 0x53,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x09, 0x53, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_proto_goTypes = []interface{}{
	(*Surrender)(nil),   // 0: proto.surrender
	(*Playchess)(nil),   // 1: proto.playchess
	(*Ready)(nil),       // 2: proto.ready
	(*Login)(nil),       // 3: proto.login
	(*Register)(nil),    // 4: proto.register
	(*Reply)(nil),       // 5: proto.Reply
	(*Reply2)(nil),      // 6: proto.Reply2
	(*Reply2_Edge)(nil), // 7: proto.Reply2.Edge
}
var file_user_proto_depIdxs = []int32{
	7, // 0: proto.Reply2.j:type_name -> proto.Reply2.Edge
	3, // 1: proto.Login.Login:input_type -> proto.login
	4, // 2: proto.Register.Register:input_type -> proto.register
	1, // 3: proto.Playchess.Playchess:input_type -> proto.playchess
	2, // 4: proto.Ready.Ready:input_type -> proto.ready
	0, // 5: proto.Surrender.Surrender:input_type -> proto.surrender
	5, // 6: proto.Login.Login:output_type -> proto.Reply
	5, // 7: proto.Register.Register:output_type -> proto.Reply
	5, // 8: proto.Playchess.Playchess:output_type -> proto.Reply
	5, // 9: proto.Ready.Ready:output_type -> proto.Reply
	5, // 10: proto.Surrender.Surrender:output_type -> proto.Reply
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Surrender); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Playchess); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ready); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Login); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Register); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply2_Edge); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   5,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginClient interface {
	// 方法 (请求消息结构体) returns (返回消息结构体) {}
	Login(ctx context.Context, in *Login, opts ...grpc.CallOption) (*Reply, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Login(ctx context.Context, in *Login, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.Login/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
type LoginServer interface {
	// 方法 (请求消息结构体) returns (返回消息结构体) {}
	Login(context.Context, *Login) (*Reply, error)
}

// UnimplementedLoginServer can be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (*UnimplementedLoginServer) Login(context.Context, *Login) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Login)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Login/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Login(ctx, req.(*Login))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Login_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// RegisterClient is the client API for Register service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegisterClient interface {
	Register(ctx context.Context, in *Register, opts ...grpc.CallOption) (*Reply, error)
}

type registerClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterClient(cc grpc.ClientConnInterface) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) Register(ctx context.Context, in *Register, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.Register/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServer is the server API for Register service.
type RegisterServer interface {
	Register(context.Context, *Register) (*Reply, error)
}

// UnimplementedRegisterServer can be embedded to have forward compatible implementations.
type UnimplementedRegisterServer struct {
}

func (*UnimplementedRegisterServer) Register(context.Context, *Register) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterRegisterServer(s *grpc.Server, srv RegisterServer) {
	s.RegisterService(&_Register_serviceDesc, srv)
}

func _Register_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Register)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Register/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).Register(ctx, req.(*Register))
	}
	return interceptor(ctx, in, info, handler)
}

var _Register_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Register",
	HandlerType: (*RegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Register_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// PlaychessClient is the client API for Playchess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlaychessClient interface {
	Playchess(ctx context.Context, in *Playchess, opts ...grpc.CallOption) (*Reply, error)
}

type playchessClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaychessClient(cc grpc.ClientConnInterface) PlaychessClient {
	return &playchessClient{cc}
}

func (c *playchessClient) Playchess(ctx context.Context, in *Playchess, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.Playchess/Playchess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaychessServer is the server API for Playchess service.
type PlaychessServer interface {
	Playchess(context.Context, *Playchess) (*Reply, error)
}

// UnimplementedPlaychessServer can be embedded to have forward compatible implementations.
type UnimplementedPlaychessServer struct {
}

func (*UnimplementedPlaychessServer) Playchess(context.Context, *Playchess) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Playchess not implemented")
}

func RegisterPlaychessServer(s *grpc.Server, srv PlaychessServer) {
	s.RegisterService(&_Playchess_serviceDesc, srv)
}

func _Playchess_Playchess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Playchess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaychessServer).Playchess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playchess/Playchess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaychessServer).Playchess(ctx, req.(*Playchess))
	}
	return interceptor(ctx, in, info, handler)
}

var _Playchess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Playchess",
	HandlerType: (*PlaychessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Playchess",
			Handler:    _Playchess_Playchess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// ReadyClient is the client API for Ready service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReadyClient interface {
	Ready(ctx context.Context, in *Ready, opts ...grpc.CallOption) (*Reply, error)
}

type readyClient struct {
	cc grpc.ClientConnInterface
}

func NewReadyClient(cc grpc.ClientConnInterface) ReadyClient {
	return &readyClient{cc}
}

func (c *readyClient) Ready(ctx context.Context, in *Ready, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.Ready/Ready", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReadyServer is the server API for Ready service.
type ReadyServer interface {
	Ready(context.Context, *Ready) (*Reply, error)
}

// UnimplementedReadyServer can be embedded to have forward compatible implementations.
type UnimplementedReadyServer struct {
}

func (*UnimplementedReadyServer) Ready(context.Context, *Ready) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ready not implemented")
}

func RegisterReadyServer(s *grpc.Server, srv ReadyServer) {
	s.RegisterService(&_Ready_serviceDesc, srv)
}

func _Ready_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ready)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReadyServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Ready/Ready",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReadyServer).Ready(ctx, req.(*Ready))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ready_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Ready",
	HandlerType: (*ReadyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ready",
			Handler:    _Ready_Ready_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

// SurrenderClient is the client API for Surrender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SurrenderClient interface {
	Surrender(ctx context.Context, in *Surrender, opts ...grpc.CallOption) (*Reply, error)
}

type surrenderClient struct {
	cc grpc.ClientConnInterface
}

func NewSurrenderClient(cc grpc.ClientConnInterface) SurrenderClient {
	return &surrenderClient{cc}
}

func (c *surrenderClient) Surrender(ctx context.Context, in *Surrender, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/proto.Surrender/Surrender", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SurrenderServer is the server API for Surrender service.
type SurrenderServer interface {
	Surrender(context.Context, *Surrender) (*Reply, error)
}

// UnimplementedSurrenderServer can be embedded to have forward compatible implementations.
type UnimplementedSurrenderServer struct {
}

func (*UnimplementedSurrenderServer) Surrender(context.Context, *Surrender) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Surrender not implemented")
}

func RegisterSurrenderServer(s *grpc.Server, srv SurrenderServer) {
	s.RegisterService(&_Surrender_serviceDesc, srv)
}

func _Surrender_Surrender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Surrender)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurrenderServer).Surrender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Surrender/Surrender",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurrenderServer).Surrender(ctx, req.(*Surrender))
	}
	return interceptor(ctx, in, info, handler)
}

var _Surrender_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Surrender",
	HandlerType: (*SurrenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Surrender",
			Handler:    _Surrender_Surrender_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
