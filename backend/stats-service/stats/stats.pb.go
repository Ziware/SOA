// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: stats.proto

package __

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

type PostCommented struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CommentId     string                 `protobuf:"bytes,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	PostId        string                 `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Timestamp     int64                  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostCommented) Reset() {
	*x = PostCommented{}
	mi := &file_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostCommented) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCommented) ProtoMessage() {}

func (x *PostCommented) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCommented.ProtoReflect.Descriptor instead.
func (*PostCommented) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{0}
}

func (x *PostCommented) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *PostCommented) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *PostCommented) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PostCommented) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type PostViewed struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        string                 `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Timestamp     int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostViewed) Reset() {
	*x = PostViewed{}
	mi := &file_stats_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostViewed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostViewed) ProtoMessage() {}

func (x *PostViewed) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostViewed.ProtoReflect.Descriptor instead.
func (*PostViewed) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{1}
}

func (x *PostViewed) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *PostViewed) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PostViewed) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type PostLiked struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        string                 `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Timestamp     int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostLiked) Reset() {
	*x = PostLiked{}
	mi := &file_stats_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostLiked) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostLiked) ProtoMessage() {}

func (x *PostLiked) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostLiked.ProtoReflect.Descriptor instead.
func (*PostLiked) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{2}
}

func (x *PostLiked) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *PostLiked) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PostLiked) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type UserCreated struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Timestamp     int64                  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserCreated) Reset() {
	*x = UserCreated{}
	mi := &file_stats_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreated) ProtoMessage() {}

func (x *UserCreated) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreated.ProtoReflect.Descriptor instead.
func (*UserCreated) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{3}
}

func (x *UserCreated) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserCreated) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_stats_proto protoreflect.FileDescriptor

const file_stats_proto_rawDesc = "" +
	"\n" +
	"\vstats.proto\x12\x05stats\"~\n" +
	"\rPostCommented\x12\x1d\n" +
	"\n" +
	"comment_id\x18\x01 \x01(\tR\tcommentId\x12\x17\n" +
	"\apost_id\x18\x02 \x01(\tR\x06postId\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\tR\x06userId\x12\x1c\n" +
	"\ttimestamp\x18\x04 \x01(\x03R\ttimestamp\"\\\n" +
	"\n" +
	"PostViewed\x12\x17\n" +
	"\apost_id\x18\x01 \x01(\tR\x06postId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\x03R\ttimestamp\"[\n" +
	"\tPostLiked\x12\x17\n" +
	"\apost_id\x18\x01 \x01(\tR\x06postId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\x03R\ttimestamp\"D\n" +
	"\vUserCreated\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1c\n" +
	"\ttimestamp\x18\x02 \x01(\x03R\ttimestampB\x03Z\x01.b\x06proto3"

var (
	file_stats_proto_rawDescOnce sync.Once
	file_stats_proto_rawDescData []byte
)

func file_stats_proto_rawDescGZIP() []byte {
	file_stats_proto_rawDescOnce.Do(func() {
		file_stats_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_stats_proto_rawDesc), len(file_stats_proto_rawDesc)))
	})
	return file_stats_proto_rawDescData
}

var file_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stats_proto_goTypes = []any{
	(*PostCommented)(nil), // 0: stats.PostCommented
	(*PostViewed)(nil),    // 1: stats.PostViewed
	(*PostLiked)(nil),     // 2: stats.PostLiked
	(*UserCreated)(nil),   // 3: stats.UserCreated
}
var file_stats_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stats_proto_init() }
func file_stats_proto_init() {
	if File_stats_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_stats_proto_rawDesc), len(file_stats_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stats_proto_goTypes,
		DependencyIndexes: file_stats_proto_depIdxs,
		MessageInfos:      file_stats_proto_msgTypes,
	}.Build()
	File_stats_proto = out.File
	file_stats_proto_goTypes = nil
	file_stats_proto_depIdxs = nil
}
