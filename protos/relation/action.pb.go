// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.5
// source: action.proto

package relation

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

type DouyinRelationActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                              // 用户鉴权token
	ToUserId   int64  `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`     // 对方用户id
	ActionType int32  `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"` // 1-关注，2-取消关注
}

func (x *DouyinRelationActionRequest) Reset() {
	*x = DouyinRelationActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinRelationActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinRelationActionRequest) ProtoMessage() {}

func (x *DouyinRelationActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinRelationActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinRelationActionRequest) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinRelationActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DouyinRelationActionRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *DouyinRelationActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type DouyinRelationActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
}

func (x *DouyinRelationActionResponse) Reset() {
	*x = DouyinRelationActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinRelationActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinRelationActionResponse) ProtoMessage() {}

func (x *DouyinRelationActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinRelationActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinRelationActionResponse) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinRelationActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinRelationActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

var File_action_proto protoreflect.FileDescriptor

var file_action_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x69, 0x64, 0x6c, 0x22, 0x75, 0x0a, 0x1e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x74,
	0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x61, 0x0a, 0x1f, 0x64, 0x6f,
	0x75, 0x79, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x32, 0x80, 0x01,
	0x0a, 0x1b, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a,
	0x14, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x69, 0x64, 0x6c, 0x2e, 0x64, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x69, 0x64, 0x6c,
	0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_action_proto_rawDescOnce sync.Once
	file_action_proto_rawDescData = file_action_proto_rawDesc
)

func file_action_proto_rawDescGZIP() []byte {
	file_action_proto_rawDescOnce.Do(func() {
		file_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_action_proto_rawDescData)
	})
	return file_action_proto_rawDescData
}

var file_action_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_action_proto_goTypes = []interface{}{
	(*DouyinRelationActionRequest)(nil),  // 0: idl.douyin_relation_action_request
	(*DouyinRelationActionResponse)(nil), // 1: idl.douyin_relation_action_response
}
var file_action_proto_depIdxs = []int32{
	0, // 0: idl.DouyinRelationActionService.DouyinRelationAction:input_type -> idl.douyin_relation_action_request
	1, // 1: idl.DouyinRelationActionService.DouyinRelationAction:output_type -> idl.douyin_relation_action_response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_action_proto_init() }
func file_action_proto_init() {
	if File_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinRelationActionRequest); i {
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
		file_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinRelationActionResponse); i {
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
			RawDescriptor: file_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_action_proto_goTypes,
		DependencyIndexes: file_action_proto_depIdxs,
		MessageInfos:      file_action_proto_msgTypes,
	}.Build()
	File_action_proto = out.File
	file_action_proto_rawDesc = nil
	file_action_proto_goTypes = nil
	file_action_proto_depIdxs = nil
}
