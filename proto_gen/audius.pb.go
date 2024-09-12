// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: audius.proto

package proto_gen

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

type ManageEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Types that are assignable to Message:
	//
	//	*ManageEntity_UserCreate
	//	*ManageEntity_TrackCreate
	Message isManageEntity_Message `protobuf_oneof:"Message"`
}

func (x *ManageEntity) Reset() {
	*x = ManageEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audius_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManageEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageEntity) ProtoMessage() {}

func (x *ManageEntity) ProtoReflect() protoreflect.Message {
	mi := &file_audius_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageEntity.ProtoReflect.Descriptor instead.
func (*ManageEntity) Descriptor() ([]byte, []int) {
	return file_audius_proto_rawDescGZIP(), []int{0}
}

func (x *ManageEntity) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (m *ManageEntity) GetMessage() isManageEntity_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ManageEntity) GetUserCreate() *UserCreate {
	if x, ok := x.GetMessage().(*ManageEntity_UserCreate); ok {
		return x.UserCreate
	}
	return nil
}

func (x *ManageEntity) GetTrackCreate() *TrackCreate {
	if x, ok := x.GetMessage().(*ManageEntity_TrackCreate); ok {
		return x.TrackCreate
	}
	return nil
}

type isManageEntity_Message interface {
	isManageEntity_Message()
}

type ManageEntity_UserCreate struct {
	UserCreate *UserCreate `protobuf:"bytes,100,opt,name=userCreate,proto3,oneof"`
}

type ManageEntity_TrackCreate struct {
	TrackCreate *TrackCreate `protobuf:"bytes,101,opt,name=trackCreate,proto3,oneof"`
}

func (*ManageEntity_UserCreate) isManageEntity_Message() {}

func (*ManageEntity_TrackCreate) isManageEntity_Message() {}

type UserCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Handle string `protobuf:"bytes,2,opt,name=handle,proto3" json:"handle,omitempty"`
}

func (x *UserCreate) Reset() {
	*x = UserCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audius_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreate) ProtoMessage() {}

func (x *UserCreate) ProtoReflect() protoreflect.Message {
	mi := &file_audius_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreate.ProtoReflect.Descriptor instead.
func (*UserCreate) Descriptor() ([]byte, []int) {
	return file_audius_proto_rawDescGZIP(), []int{1}
}

func (x *UserCreate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserCreate) GetHandle() string {
	if x != nil {
		return x.Handle
	}
	return ""
}

type TrackCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	StreamUrl string `protobuf:"bytes,3,opt,name=stream_url,json=streamUrl,proto3" json:"stream_url,omitempty"`
	UserId    string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *TrackCreate) Reset() {
	*x = TrackCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audius_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackCreate) ProtoMessage() {}

func (x *TrackCreate) ProtoReflect() protoreflect.Message {
	mi := &file_audius_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackCreate.ProtoReflect.Descriptor instead.
func (*TrackCreate) Descriptor() ([]byte, []int) {
	return file_audius_proto_rawDescGZIP(), []int{2}
}

func (x *TrackCreate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TrackCreate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TrackCreate) GetStreamUrl() string {
	if x != nil {
		return x.StreamUrl
	}
	return ""
}

func (x *TrackCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_audius_proto protoreflect.FileDescriptor

var file_audius_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x61, 0x75, 0x64, 0x69, 0x75, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x75, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x75, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x34, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x6b, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_audius_proto_rawDescOnce sync.Once
	file_audius_proto_rawDescData = file_audius_proto_rawDesc
)

func file_audius_proto_rawDescGZIP() []byte {
	file_audius_proto_rawDescOnce.Do(func() {
		file_audius_proto_rawDescData = protoimpl.X.CompressGZIP(file_audius_proto_rawDescData)
	})
	return file_audius_proto_rawDescData
}

var file_audius_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_audius_proto_goTypes = []any{
	(*ManageEntity)(nil), // 0: audius.ManageEntity
	(*UserCreate)(nil),   // 1: audius.UserCreate
	(*TrackCreate)(nil),  // 2: audius.TrackCreate
}
var file_audius_proto_depIdxs = []int32{
	1, // 0: audius.ManageEntity.userCreate:type_name -> audius.UserCreate
	2, // 1: audius.ManageEntity.trackCreate:type_name -> audius.TrackCreate
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_audius_proto_init() }
func file_audius_proto_init() {
	if File_audius_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_audius_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ManageEntity); i {
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
		file_audius_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserCreate); i {
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
		file_audius_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TrackCreate); i {
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
	file_audius_proto_msgTypes[0].OneofWrappers = []any{
		(*ManageEntity_UserCreate)(nil),
		(*ManageEntity_TrackCreate)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_audius_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_audius_proto_goTypes,
		DependencyIndexes: file_audius_proto_depIdxs,
		MessageInfos:      file_audius_proto_msgTypes,
	}.Build()
	File_audius_proto = out.File
	file_audius_proto_rawDesc = nil
	file_audius_proto_goTypes = nil
	file_audius_proto_depIdxs = nil
}
