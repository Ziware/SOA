// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: post.proto

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

type Post struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        string                 `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId     string                 `protobuf:"bytes,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsPrivate     bool                   `protobuf:"varint,7,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
	Tags          []string               `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Post) Reset() {
	*x = Post{}
	mi := &file_post_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Post) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Post) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Post) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Post) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

func (x *Post) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CommentId     string                 `protobuf:"bytes,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	PostId        string                 `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	CreatorId     string                 `protobuf:"bytes,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Text          string                 `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Comment) Reset() {
	*x = Comment{}
	mi := &file_post_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{1}
}

func (x *Comment) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *Comment) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *Comment) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Comment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Comment) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreatePostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Post          *Post                  `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	mi := &file_post_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePostRequest) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type CreatePostResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Post          *Post                  `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePostResponse) Reset() {
	*x = CreatePostResponse{}
	mi := &file_post_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostResponse) ProtoMessage() {}

func (x *CreatePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostResponse.ProtoReflect.Descriptor instead.
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type DeletePostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	mi := &file_post_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeletePostRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeletePostResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePostResponse) Reset() {
	*x = DeletePostResponse{}
	mi := &file_post_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostResponse) ProtoMessage() {}

func (x *DeletePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostResponse.ProtoReflect.Descriptor instead.
func (*DeletePostResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePostResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UpdatePostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Post          *Post                  `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePostRequest) Reset() {
	*x = UpdatePostRequest{}
	mi := &file_post_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostRequest) ProtoMessage() {}

func (x *UpdatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePostRequest) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *UpdatePostRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdatePostResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Post          *Post                  `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePostResponse) Reset() {
	*x = UpdatePostResponse{}
	mi := &file_post_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostResponse) ProtoMessage() {}

func (x *UpdatePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostResponse.ProtoReflect.Descriptor instead.
func (*UpdatePostResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type GetPostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	mi := &file_post_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{8}
}

func (x *GetPostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetPostRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetPostResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Post          *Post                  `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostResponse) Reset() {
	*x = GetPostResponse{}
	mi := &file_post_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResponse) ProtoMessage() {}

func (x *GetPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResponse.ProtoReflect.Descriptor instead.
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{9}
}

func (x *GetPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type ListPostsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PageNumber    int32                  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize      int32                  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPostsRequest) Reset() {
	*x = ListPostsRequest{}
	mi := &file_post_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsRequest) ProtoMessage() {}

func (x *ListPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsRequest.ProtoReflect.Descriptor instead.
func (*ListPostsRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{10}
}

func (x *ListPostsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListPostsRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListPostsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListPostsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Posts         []*Post                `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPostsResponse) Reset() {
	*x = ListPostsResponse{}
	mi := &file_post_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsResponse) ProtoMessage() {}

func (x *ListPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsResponse.ProtoReflect.Descriptor instead.
func (*ListPostsResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{11}
}

func (x *ListPostsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *ListPostsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type ViewPostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        string                 `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ViewPostRequest) Reset() {
	*x = ViewPostRequest{}
	mi := &file_post_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ViewPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewPostRequest) ProtoMessage() {}

func (x *ViewPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewPostRequest.ProtoReflect.Descriptor instead.
func (*ViewPostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{12}
}

func (x *ViewPostRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *ViewPostRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ViewPostResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ViewPostResponse) Reset() {
	*x = ViewPostResponse{}
	mi := &file_post_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ViewPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewPostResponse) ProtoMessage() {}

func (x *ViewPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewPostResponse.ProtoReflect.Descriptor instead.
func (*ViewPostResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{13}
}

func (x *ViewPostResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type LikePostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        string                 `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LikePostRequest) Reset() {
	*x = LikePostRequest{}
	mi := &file_post_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LikePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikePostRequest) ProtoMessage() {}

func (x *LikePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikePostRequest.ProtoReflect.Descriptor instead.
func (*LikePostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{14}
}

func (x *LikePostRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *LikePostRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LikePostResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LikePostResponse) Reset() {
	*x = LikePostResponse{}
	mi := &file_post_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LikePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikePostResponse) ProtoMessage() {}

func (x *LikePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikePostResponse.ProtoReflect.Descriptor instead.
func (*LikePostResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{15}
}

func (x *LikePostResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CreateCommentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        string                 `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Text          string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	mi := &file_post_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{16}
}

func (x *CreateCommentRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *CreateCommentRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateCommentRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CreateCommentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Comment       *Comment               `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCommentResponse) Reset() {
	*x = CreateCommentResponse{}
	mi := &file_post_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentResponse) ProtoMessage() {}

func (x *CreateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentResponse.ProtoReflect.Descriptor instead.
func (*CreateCommentResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{17}
}

func (x *CreateCommentResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type ListCommentsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        string                 `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PageNumber    int32                  `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize      int32                  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCommentsRequest) Reset() {
	*x = ListCommentsRequest{}
	mi := &file_post_proto_msgTypes[18]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentsRequest) ProtoMessage() {}

func (x *ListCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[18]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentsRequest.ProtoReflect.Descriptor instead.
func (*ListCommentsRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{18}
}

func (x *ListCommentsRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *ListCommentsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListCommentsRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListCommentsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListCommentsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Comments      []*Comment             `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCommentsResponse) Reset() {
	*x = ListCommentsResponse{}
	mi := &file_post_proto_msgTypes[19]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentsResponse) ProtoMessage() {}

func (x *ListCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[19]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentsResponse.ProtoReflect.Descriptor instead.
func (*ListCommentsResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{19}
}

func (x *ListCommentsResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *ListCommentsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_post_proto protoreflect.FileDescriptor

const file_post_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"post.proto\x12\x04post\"\xe7\x01\n" +
	"\x04Post\x12\x17\n" +
	"\apost_id\x18\x01 \x01(\tR\x06postId\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1d\n" +
	"\n" +
	"creator_id\x18\x04 \x01(\tR\tcreatorId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x05 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\tR\tupdatedAt\x12\x1d\n" +
	"\n" +
	"is_private\x18\a \x01(\bR\tisPrivate\x12\x12\n" +
	"\x04tags\x18\b \x03(\tR\x04tags\"\xb2\x01\n" +
	"\aComment\x12\x1d\n" +
	"\n" +
	"comment_id\x18\x01 \x01(\tR\tcommentId\x12\x17\n" +
	"\apost_id\x18\x02 \x01(\tR\x06postId\x12\x1d\n" +
	"\n" +
	"creator_id\x18\x03 \x01(\tR\tcreatorId\x12\x12\n" +
	"\x04text\x18\x04 \x01(\tR\x04text\x12\x1d\n" +
	"\n" +
	"created_at\x18\x05 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\tR\tupdatedAt\"3\n" +
	"\x11CreatePostRequest\x12\x1e\n" +
	"\x04post\x18\x01 \x01(\v2\n" +
	".post.PostR\x04post\"4\n" +
	"\x12CreatePostResponse\x12\x1e\n" +
	"\x04post\x18\x01 \x01(\v2\n" +
	".post.PostR\x04post\"<\n" +
	"\x11DeletePostRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\".\n" +
	"\x12DeletePostResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"L\n" +
	"\x11UpdatePostRequest\x12\x1e\n" +
	"\x04post\x18\x01 \x01(\v2\n" +
	".post.PostR\x04post\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\"4\n" +
	"\x12UpdatePostResponse\x12\x1e\n" +
	"\x04post\x18\x01 \x01(\v2\n" +
	".post.PostR\x04post\"9\n" +
	"\x0eGetPostRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\"1\n" +
	"\x0fGetPostResponse\x12\x1e\n" +
	"\x04post\x18\x01 \x01(\v2\n" +
	".post.PostR\x04post\"i\n" +
	"\x10ListPostsRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1f\n" +
	"\vpage_number\x18\x02 \x01(\x05R\n" +
	"pageNumber\x12\x1b\n" +
	"\tpage_size\x18\x03 \x01(\x05R\bpageSize\"V\n" +
	"\x11ListPostsResponse\x12 \n" +
	"\x05posts\x18\x01 \x03(\v2\n" +
	".post.PostR\x05posts\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"C\n" +
	"\x0fViewPostRequest\x12\x17\n" +
	"\apost_id\x18\x01 \x01(\tR\x06postId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\",\n" +
	"\x10ViewPostResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"C\n" +
	"\x0fLikePostRequest\x12\x17\n" +
	"\apost_id\x18\x01 \x01(\tR\x06postId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\",\n" +
	"\x10LikePostResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"\\\n" +
	"\x14CreateCommentRequest\x12\x17\n" +
	"\apost_id\x18\x01 \x01(\tR\x06postId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x12\n" +
	"\x04text\x18\x03 \x01(\tR\x04text\"@\n" +
	"\x15CreateCommentResponse\x12'\n" +
	"\acomment\x18\x01 \x01(\v2\r.post.CommentR\acomment\"\x85\x01\n" +
	"\x13ListCommentsRequest\x12\x17\n" +
	"\apost_id\x18\x01 \x01(\tR\x06postId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x1f\n" +
	"\vpage_number\x18\x03 \x01(\x05R\n" +
	"pageNumber\x12\x1b\n" +
	"\tpage_size\x18\x04 \x01(\x05R\bpageSize\"b\n" +
	"\x14ListCommentsResponse\x12)\n" +
	"\bcomments\x18\x01 \x03(\v2\r.post.CommentR\bcomments\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount2\xcd\x04\n" +
	"\vWallService\x12?\n" +
	"\n" +
	"CreatePost\x12\x17.post.CreatePostRequest\x1a\x18.post.CreatePostResponse\x12?\n" +
	"\n" +
	"DeletePost\x12\x17.post.DeletePostRequest\x1a\x18.post.DeletePostResponse\x12?\n" +
	"\n" +
	"UpdatePost\x12\x17.post.UpdatePostRequest\x1a\x18.post.UpdatePostResponse\x126\n" +
	"\aGetPost\x12\x14.post.GetPostRequest\x1a\x15.post.GetPostResponse\x12<\n" +
	"\tListPosts\x12\x16.post.ListPostsRequest\x1a\x17.post.ListPostsResponse\x129\n" +
	"\bViewPost\x12\x15.post.ViewPostRequest\x1a\x16.post.ViewPostResponse\x129\n" +
	"\bLikePost\x12\x15.post.LikePostRequest\x1a\x16.post.LikePostResponse\x12H\n" +
	"\rCreateComment\x12\x1a.post.CreateCommentRequest\x1a\x1b.post.CreateCommentResponse\x12E\n" +
	"\fListComments\x12\x19.post.ListCommentsRequest\x1a\x1a.post.ListCommentsResponseB\x03Z\x01.b\x06proto3"

var (
	file_post_proto_rawDescOnce sync.Once
	file_post_proto_rawDescData []byte
)

func file_post_proto_rawDescGZIP() []byte {
	file_post_proto_rawDescOnce.Do(func() {
		file_post_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_post_proto_rawDesc), len(file_post_proto_rawDesc)))
	})
	return file_post_proto_rawDescData
}

var file_post_proto_msgTypes = make([]protoimpl.MessageInfo, 20)
var file_post_proto_goTypes = []any{
	(*Post)(nil),                  // 0: post.Post
	(*Comment)(nil),               // 1: post.Comment
	(*CreatePostRequest)(nil),     // 2: post.CreatePostRequest
	(*CreatePostResponse)(nil),    // 3: post.CreatePostResponse
	(*DeletePostRequest)(nil),     // 4: post.DeletePostRequest
	(*DeletePostResponse)(nil),    // 5: post.DeletePostResponse
	(*UpdatePostRequest)(nil),     // 6: post.UpdatePostRequest
	(*UpdatePostResponse)(nil),    // 7: post.UpdatePostResponse
	(*GetPostRequest)(nil),        // 8: post.GetPostRequest
	(*GetPostResponse)(nil),       // 9: post.GetPostResponse
	(*ListPostsRequest)(nil),      // 10: post.ListPostsRequest
	(*ListPostsResponse)(nil),     // 11: post.ListPostsResponse
	(*ViewPostRequest)(nil),       // 12: post.ViewPostRequest
	(*ViewPostResponse)(nil),      // 13: post.ViewPostResponse
	(*LikePostRequest)(nil),       // 14: post.LikePostRequest
	(*LikePostResponse)(nil),      // 15: post.LikePostResponse
	(*CreateCommentRequest)(nil),  // 16: post.CreateCommentRequest
	(*CreateCommentResponse)(nil), // 17: post.CreateCommentResponse
	(*ListCommentsRequest)(nil),   // 18: post.ListCommentsRequest
	(*ListCommentsResponse)(nil),  // 19: post.ListCommentsResponse
}
var file_post_proto_depIdxs = []int32{
	0,  // 0: post.CreatePostRequest.post:type_name -> post.Post
	0,  // 1: post.CreatePostResponse.post:type_name -> post.Post
	0,  // 2: post.UpdatePostRequest.post:type_name -> post.Post
	0,  // 3: post.UpdatePostResponse.post:type_name -> post.Post
	0,  // 4: post.GetPostResponse.post:type_name -> post.Post
	0,  // 5: post.ListPostsResponse.posts:type_name -> post.Post
	1,  // 6: post.CreateCommentResponse.comment:type_name -> post.Comment
	1,  // 7: post.ListCommentsResponse.comments:type_name -> post.Comment
	2,  // 8: post.WallService.CreatePost:input_type -> post.CreatePostRequest
	4,  // 9: post.WallService.DeletePost:input_type -> post.DeletePostRequest
	6,  // 10: post.WallService.UpdatePost:input_type -> post.UpdatePostRequest
	8,  // 11: post.WallService.GetPost:input_type -> post.GetPostRequest
	10, // 12: post.WallService.ListPosts:input_type -> post.ListPostsRequest
	12, // 13: post.WallService.ViewPost:input_type -> post.ViewPostRequest
	14, // 14: post.WallService.LikePost:input_type -> post.LikePostRequest
	16, // 15: post.WallService.CreateComment:input_type -> post.CreateCommentRequest
	18, // 16: post.WallService.ListComments:input_type -> post.ListCommentsRequest
	3,  // 17: post.WallService.CreatePost:output_type -> post.CreatePostResponse
	5,  // 18: post.WallService.DeletePost:output_type -> post.DeletePostResponse
	7,  // 19: post.WallService.UpdatePost:output_type -> post.UpdatePostResponse
	9,  // 20: post.WallService.GetPost:output_type -> post.GetPostResponse
	11, // 21: post.WallService.ListPosts:output_type -> post.ListPostsResponse
	13, // 22: post.WallService.ViewPost:output_type -> post.ViewPostResponse
	15, // 23: post.WallService.LikePost:output_type -> post.LikePostResponse
	17, // 24: post.WallService.CreateComment:output_type -> post.CreateCommentResponse
	19, // 25: post.WallService.ListComments:output_type -> post.ListCommentsResponse
	17, // [17:26] is the sub-list for method output_type
	8,  // [8:17] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_post_proto_init() }
func file_post_proto_init() {
	if File_post_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_post_proto_rawDesc), len(file_post_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   20,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_proto_goTypes,
		DependencyIndexes: file_post_proto_depIdxs,
		MessageInfos:      file_post_proto_msgTypes,
	}.Build()
	File_post_proto = out.File
	file_post_proto_goTypes = nil
	file_post_proto_depIdxs = nil
}
