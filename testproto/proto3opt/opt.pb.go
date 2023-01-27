// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: proto3opt/opt.proto

package proto3opt

import (
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

type SimpleEnum int32

const (
	SimpleEnum_ENUM1 SimpleEnum = 0
	SimpleEnum_ENUM2 SimpleEnum = 1
	SimpleEnum_ENUM3 SimpleEnum = 2
)

// Enum value maps for SimpleEnum.
var (
	SimpleEnum_name = map[int32]string{
		0: "ENUM1",
		1: "ENUM2",
		2: "ENUM3",
	}
	SimpleEnum_value = map[string]int32{
		"ENUM1": 0,
		"ENUM2": 1,
		"ENUM3": 2,
	}
)

func (x SimpleEnum) Enum() *SimpleEnum {
	p := new(SimpleEnum)
	*p = x
	return p
}

func (x SimpleEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SimpleEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_proto3opt_opt_proto_enumTypes[0].Descriptor()
}

func (SimpleEnum) Type() protoreflect.EnumType {
	return &file_proto3opt_opt_proto_enumTypes[0]
}

func (x SimpleEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SimpleEnum.Descriptor instead.
func (SimpleEnum) EnumDescriptor() ([]byte, []int) {
	return file_proto3opt_opt_proto_rawDescGZIP(), []int{0}
}

type OptionalFieldInProto3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionalInt32    *int32      `protobuf:"varint,1,opt,name=optional_int32,json=optionalInt32,proto3,oneof" json:"optional_int32,omitempty"`
	OptionalInt64    *int64      `protobuf:"varint,2,opt,name=optional_int64,json=optionalInt64,proto3,oneof" json:"optional_int64,omitempty"`
	OptionalUint32   *uint32     `protobuf:"varint,3,opt,name=optional_uint32,json=optionalUint32,proto3,oneof" json:"optional_uint32,omitempty"`
	OptionalUint64   *uint64     `protobuf:"varint,4,opt,name=optional_uint64,json=optionalUint64,proto3,oneof" json:"optional_uint64,omitempty"`
	OptionalSint32   *int32      `protobuf:"zigzag32,5,opt,name=optional_sint32,json=optionalSint32,proto3,oneof" json:"optional_sint32,omitempty"`
	OptionalSint64   *int64      `protobuf:"zigzag64,6,opt,name=optional_sint64,json=optionalSint64,proto3,oneof" json:"optional_sint64,omitempty"`
	OptionalFixed32  *uint32     `protobuf:"fixed32,7,opt,name=optional_fixed32,json=optionalFixed32,proto3,oneof" json:"optional_fixed32,omitempty"`
	OptionalFixed64  *uint64     `protobuf:"fixed64,8,opt,name=optional_fixed64,json=optionalFixed64,proto3,oneof" json:"optional_fixed64,omitempty"`
	OptionalSfixed32 *int32      `protobuf:"fixed32,9,opt,name=optional_sfixed32,json=optionalSfixed32,proto3,oneof" json:"optional_sfixed32,omitempty"`
	OptionalSfixed64 *int64      `protobuf:"fixed64,10,opt,name=optional_sfixed64,json=optionalSfixed64,proto3,oneof" json:"optional_sfixed64,omitempty"`
	OptionalFloat    *float32    `protobuf:"fixed32,11,opt,name=optional_float,json=optionalFloat,proto3,oneof" json:"optional_float,omitempty"`
	OptionalDouble   *float64    `protobuf:"fixed64,12,opt,name=optional_double,json=optionalDouble,proto3,oneof" json:"optional_double,omitempty"`
	OptionalBool     *bool       `protobuf:"varint,13,opt,name=optional_bool,json=optionalBool,proto3,oneof" json:"optional_bool,omitempty"`
	OptionalString   *string     `protobuf:"bytes,14,opt,name=optional_string,json=optionalString,proto3,oneof" json:"optional_string,omitempty"`
	OptionalBytes    []byte      `protobuf:"bytes,15,opt,name=optional_bytes,json=optionalBytes,proto3,oneof" json:"optional_bytes,omitempty"`
	OptionalEnum     *SimpleEnum `protobuf:"varint,16,opt,name=optional_enum,json=optionalEnum,proto3,enum=SimpleEnum,oneof" json:"optional_enum,omitempty"`
}

func (x *OptionalFieldInProto3) Reset() {
	*x = OptionalFieldInProto3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto3opt_opt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionalFieldInProto3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionalFieldInProto3) ProtoMessage() {}

func (x *OptionalFieldInProto3) ProtoReflect() protoreflect.Message {
	mi := &file_proto3opt_opt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionalFieldInProto3.ProtoReflect.Descriptor instead.
func (*OptionalFieldInProto3) Descriptor() ([]byte, []int) {
	return file_proto3opt_opt_proto_rawDescGZIP(), []int{0}
}

func (x *OptionalFieldInProto3) GetOptionalInt32() int32 {
	if x != nil && x.OptionalInt32 != nil {
		return *x.OptionalInt32
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalInt64() int64 {
	if x != nil && x.OptionalInt64 != nil {
		return *x.OptionalInt64
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalUint32() uint32 {
	if x != nil && x.OptionalUint32 != nil {
		return *x.OptionalUint32
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalUint64() uint64 {
	if x != nil && x.OptionalUint64 != nil {
		return *x.OptionalUint64
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalSint32() int32 {
	if x != nil && x.OptionalSint32 != nil {
		return *x.OptionalSint32
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalSint64() int64 {
	if x != nil && x.OptionalSint64 != nil {
		return *x.OptionalSint64
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalFixed32() uint32 {
	if x != nil && x.OptionalFixed32 != nil {
		return *x.OptionalFixed32
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalFixed64() uint64 {
	if x != nil && x.OptionalFixed64 != nil {
		return *x.OptionalFixed64
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalSfixed32() int32 {
	if x != nil && x.OptionalSfixed32 != nil {
		return *x.OptionalSfixed32
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalSfixed64() int64 {
	if x != nil && x.OptionalSfixed64 != nil {
		return *x.OptionalSfixed64
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalFloat() float32 {
	if x != nil && x.OptionalFloat != nil {
		return *x.OptionalFloat
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalDouble() float64 {
	if x != nil && x.OptionalDouble != nil {
		return *x.OptionalDouble
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalBool() bool {
	if x != nil && x.OptionalBool != nil {
		return *x.OptionalBool
	}
	return false
}

func (x *OptionalFieldInProto3) GetOptionalString() string {
	if x != nil && x.OptionalString != nil {
		return *x.OptionalString
	}
	return ""
}

func (x *OptionalFieldInProto3) GetOptionalBytes() []byte {
	if x != nil {
		return x.OptionalBytes
	}
	return nil
}

func (x *OptionalFieldInProto3) GetOptionalEnum() SimpleEnum {
	if x != nil && x.OptionalEnum != nil {
		return *x.OptionalEnum
	}
	return SimpleEnum_ENUM1
}

var File_proto3opt_opt_proto protoreflect.FileDescriptor

var file_proto3opt_opt_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x6f, 0x70, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x08, 0x0a, 0x15, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x12,
	0x2a, 0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x02, 0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x55, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x03,
	0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11, 0x48, 0x04, 0x52, 0x0e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x88, 0x01,
	0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x12, 0x48, 0x05, 0x52, 0x0e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x88, 0x01, 0x01, 0x12,
	0x2e, 0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x07, 0x48, 0x06, 0x52, 0x0f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x88, 0x01, 0x01, 0x12,
	0x2e, 0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x06, 0x48, 0x07, 0x52, 0x0f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x88, 0x01, 0x01, 0x12,
	0x30, 0x0a, 0x11, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0f, 0x48, 0x08, 0x52, 0x10, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x88, 0x01,
	0x01, 0x12, 0x30, 0x0a, 0x11, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x10, 0x48, 0x09, 0x52, 0x10,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x48, 0x0a, 0x52, 0x0d, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x2c, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x48, 0x0b, 0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a,
	0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x0c, 0x52, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x42, 0x6f, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x0d, 0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x0e, 0x52,
	0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x35, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x0f, 0x52, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x42, 0x12,
	0x0a, 0x10, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x42, 0x13,
	0x0a, 0x11, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x42, 0x14,
	0x0a, 0x12, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x42, 0x12, 0x0a,
	0x10, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2a, 0x2d, 0x0a, 0x0a, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4e, 0x55, 0x4d, 0x31, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x4e, 0x55, 0x4d, 0x32, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4e,
	0x55, 0x4d, 0x33, 0x10, 0x02, 0x42, 0x15, 0x5a, 0x13, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x6f, 0x70, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto3opt_opt_proto_rawDescOnce sync.Once
	file_proto3opt_opt_proto_rawDescData = file_proto3opt_opt_proto_rawDesc
)

func file_proto3opt_opt_proto_rawDescGZIP() []byte {
	file_proto3opt_opt_proto_rawDescOnce.Do(func() {
		file_proto3opt_opt_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto3opt_opt_proto_rawDescData)
	})
	return file_proto3opt_opt_proto_rawDescData
}

var file_proto3opt_opt_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto3opt_opt_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto3opt_opt_proto_goTypes = []interface{}{
	(SimpleEnum)(0),               // 0: SimpleEnum
	(*OptionalFieldInProto3)(nil), // 1: OptionalFieldInProto3
}
var file_proto3opt_opt_proto_depIdxs = []int32{
	0, // 0: OptionalFieldInProto3.optional_enum:type_name -> SimpleEnum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto3opt_opt_proto_init() }
func file_proto3opt_opt_proto_init() {
	if File_proto3opt_opt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto3opt_opt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionalFieldInProto3); i {
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
	file_proto3opt_opt_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto3opt_opt_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto3opt_opt_proto_goTypes,
		DependencyIndexes: file_proto3opt_opt_proto_depIdxs,
		EnumInfos:         file_proto3opt_opt_proto_enumTypes,
		MessageInfos:      file_proto3opt_opt_proto_msgTypes,
	}.Build()
	File_proto3opt_opt_proto = out.File
	file_proto3opt_opt_proto_rawDesc = nil
	file_proto3opt_opt_proto_goTypes = nil
	file_proto3opt_opt_proto_depIdxs = nil
}
