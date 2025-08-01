// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.4
// source: proto/control.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 空消息
type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_control_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_control_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_control_proto_rawDescGZIP(), []int{0}
}

// 通用响应包裹：code + message + data
type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 业务状态码
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 描述信息
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_proto_control_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_control_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_control_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// 具体的转发规则结构体，先放在这占位
type ProxyRule struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	SourceAddr      string                 `protobuf:"bytes,1,opt,name=source_addr,json=sourceAddr,proto3" json:"source_addr,omitempty"`                // 监听源地址
	DestinationAddr string                 `protobuf:"bytes,2,opt,name=destination_addr,json=destinationAddr,proto3" json:"destination_addr,omitempty"` // 转发目标地址
	Option          string                 `protobuf:"bytes,3,opt,name=option,proto3" json:"option,omitempty"`                                          // add or del
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ProxyRule) Reset() {
	*x = ProxyRule{}
	mi := &file_proto_control_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProxyRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRule) ProtoMessage() {}

func (x *ProxyRule) ProtoReflect() protoreflect.Message {
	mi := &file_proto_control_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRule.ProtoReflect.Descriptor instead.
func (*ProxyRule) Descriptor() ([]byte, []int) {
	return file_proto_control_proto_rawDescGZIP(), []int{2}
}

func (x *ProxyRule) GetSourceAddr() string {
	if x != nil {
		return x.SourceAddr
	}
	return ""
}

func (x *ProxyRule) GetDestinationAddr() string {
	if x != nil {
		return x.DestinationAddr
	}
	return ""
}

func (x *ProxyRule) GetOption() string {
	if x != nil {
		return x.Option
	}
	return ""
}

type PushRulesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rules         []*ProxyRule           `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"` //可能下发多条
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PushRulesRequest) Reset() {
	*x = PushRulesRequest{}
	mi := &file_proto_control_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushRulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRulesRequest) ProtoMessage() {}

func (x *PushRulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_control_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRulesRequest.ProtoReflect.Descriptor instead.
func (*PushRulesRequest) Descriptor() ([]byte, []int) {
	return file_proto_control_proto_rawDescGZIP(), []int{3}
}

func (x *PushRulesRequest) GetRules() []*ProxyRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

var File_proto_control_proto protoreflect.FileDescriptor

const file_proto_control_proto_rawDesc = "" +
	"\n" +
	"\x13proto/control.proto\x12\acontrol\"\a\n" +
	"\x05Empty\"8\n" +
	"\bResponse\x12\x12\n" +
	"\x04code\x18\x01 \x01(\x05R\x04code\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"o\n" +
	"\tProxyRule\x12\x1f\n" +
	"\vsource_addr\x18\x01 \x01(\tR\n" +
	"sourceAddr\x12)\n" +
	"\x10destination_addr\x18\x02 \x01(\tR\x0fdestinationAddr\x12\x16\n" +
	"\x06option\x18\x03 \x01(\tR\x06option\"<\n" +
	"\x10PushRulesRequest\x12(\n" +
	"\x05rules\x18\x01 \x03(\v2\x12.control.ProxyRuleR\x05rules2K\n" +
	"\x0eControlService\x129\n" +
	"\tPushRules\x12\x19.control.PushRulesRequest\x1a\x11.control.ResponseB\bZ\x06proto/b\x06proto3"

var (
	file_proto_control_proto_rawDescOnce sync.Once
	file_proto_control_proto_rawDescData []byte
)

func file_proto_control_proto_rawDescGZIP() []byte {
	file_proto_control_proto_rawDescOnce.Do(func() {
		file_proto_control_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_control_proto_rawDesc), len(file_proto_control_proto_rawDesc)))
	})
	return file_proto_control_proto_rawDescData
}

var file_proto_control_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_control_proto_goTypes = []any{
	(*Empty)(nil),            // 0: control.Empty
	(*Response)(nil),         // 1: control.Response
	(*ProxyRule)(nil),        // 2: control.ProxyRule
	(*PushRulesRequest)(nil), // 3: control.PushRulesRequest
}
var file_proto_control_proto_depIdxs = []int32{
	2, // 0: control.PushRulesRequest.rules:type_name -> control.ProxyRule
	3, // 1: control.ControlService.PushRules:input_type -> control.PushRulesRequest
	1, // 2: control.ControlService.PushRules:output_type -> control.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_control_proto_init() }
func file_proto_control_proto_init() {
	if File_proto_control_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_control_proto_rawDesc), len(file_proto_control_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_control_proto_goTypes,
		DependencyIndexes: file_proto_control_proto_depIdxs,
		MessageInfos:      file_proto_control_proto_msgTypes,
	}.Build()
	File_proto_control_proto = out.File
	file_proto_control_proto_goTypes = nil
	file_proto_control_proto_depIdxs = nil
}
