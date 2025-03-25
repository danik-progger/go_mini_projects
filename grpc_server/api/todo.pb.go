// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: todo.proto

package grpc_todo

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

type NewTodo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Done          bool                   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewTodo) Reset() {
	*x = NewTodo{}
	mi := &file_todo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewTodo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTodo) ProtoMessage() {}

func (x *NewTodo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTodo.ProtoReflect.Descriptor instead.
func (*NewTodo) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

func (x *NewTodo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewTodo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewTodo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type Todo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Done          bool                   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Todo) Reset() {
	*x = Todo{}
	mi := &file_todo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *Todo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Todo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Todo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *Todo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_todo_proto protoreflect.FileDescriptor

const file_todo_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"todo.proto\x12\x05proto\"S\n" +
	"\aNewTodo\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x12\n" +
	"\x04done\x18\x03 \x01(\bR\x04done\"`\n" +
	"\x04Todo\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x12\n" +
	"\x04done\x18\x03 \x01(\bR\x04done\x12\x0e\n" +
	"\x02id\x18\x04 \x01(\tR\x02id2:\n" +
	"\vTodoService\x12+\n" +
	"\n" +
	"CreateTodo\x12\x0e.proto.NewTodo\x1a\v.proto.Todo\"\x00B5Z3github.com/danik-progger/go_mini_projects/grpc_todob\x06proto3"

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData []byte
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_todo_proto_rawDesc), len(file_todo_proto_rawDesc)))
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_todo_proto_goTypes = []any{
	(*NewTodo)(nil), // 0: proto.NewTodo
	(*Todo)(nil),    // 1: proto.Todo
}
var file_todo_proto_depIdxs = []int32{
	0, // 0: proto.TodoService.CreateTodo:input_type -> proto.NewTodo
	1, // 1: proto.TodoService.CreateTodo:output_type -> proto.Todo
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_todo_proto_rawDesc), len(file_todo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}
