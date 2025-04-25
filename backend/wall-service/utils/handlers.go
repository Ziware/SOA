package utils

import (
	"context"
	"errors"
	"log"
	"time"

	stats "messenger/stats-service/stats"
	post "messenger/wall-service/post"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"
)

var (
	ErrForbiddenToDelete = errors.New("Forbidden to delete post")
	ErrForbiddenToUpdate = errors.New("Forbidden to update post")
	ErrForbiddenToWatch  = errors.New("Forbidden to watch post")
)

func (ws *WallService) CreatePost(ctx context.Context, req *post.CreatePostRequest) (*post.CreatePostResponse, error) {
	req.Post.PostId = uuid.New().String()
	req.Post.CreatedAt = time.Now().Format(time.RFC3339)
	req.Post.UpdatedAt = time.Now().Format(time.RFC3339)
	err := cls.db.insertPost(req.Post)
	if err != nil {
		return nil, err
	}
	return &post.CreatePostResponse{Post: req.Post}, nil
}

func (ws *WallService) DeletePost(ctx context.Context, req *post.DeletePostRequest) (*post.DeletePostResponse, error) {
	pst, err := cls.db.getPostByID(req.Id)
	if err != nil {
		return nil, err
	}
	log.Printf("pst: %s", prototext.Format(pst))
	log.Printf("req: %s", prototext.Format(req))
	if pst.CreatorId != req.UserId {
		return nil, status.Errorf(codes.PermissionDenied, "%s", ErrForbiddenToDelete)
	}
	err = cls.db.deletePost(req.Id)
	if err != nil {
		return nil, err
	}
	return &post.DeletePostResponse{Success: true}, nil
}

func (ws *WallService) UpdatePost(ctx context.Context, req *post.UpdatePostRequest) (*post.UpdatePostResponse, error) {
	pst, err := cls.db.getPostByID(req.Post.PostId)
	if err != nil {
		return nil, err
	}

	if pst.CreatorId != req.UserId {
		return nil, status.Errorf(codes.PermissionDenied, "%s", ErrForbiddenToUpdate)
	}
	pst.Description = req.Post.Description
	pst.UpdatedAt = time.Now().Format(time.RFC3339)
	pst.IsPrivate = req.Post.IsPrivate
	pst.Tags = req.Post.Tags
	pst.Title = req.Post.Title
	err = cls.db.updatePost(pst)
	if err != nil {
		return nil, err
	}
	return &post.UpdatePostResponse{Post: req.Post}, nil
}

func (ws *WallService) GetPost(ctx context.Context, req *post.GetPostRequest) (*post.GetPostResponse, error) {
	pst, err := cls.db.getPostByID(req.Id)
	if err != nil {
		return nil, err
	}
	if pst.IsPrivate && req.UserId != pst.CreatorId {
		return nil, status.Errorf(codes.PermissionDenied, "%s", ErrForbiddenToWatch)
	}
	return &post.GetPostResponse{Post: pst}, nil
}

func (ws *WallService) ListPosts(ctx context.Context, req *post.ListPostsRequest) (*post.ListPostsResponse, error) {
	offset := (req.PageNumber - 1) * req.PageSize
	posts, err := cls.db.listPosts(int(req.PageSize), int(offset), req.UserId)
	if err != nil {
		return nil, err
	}
	return &post.ListPostsResponse{Posts: posts, TotalCount: int32(len(posts))}, nil
}

func (ws *WallService) ViewPost(ctx context.Context, req *post.ViewPostRequest) (*post.ViewPostResponse, error) {
	pst, err := cls.db.getPostByID(req.PostId)
	if err != nil {
		return nil, err
	}

	if pst.IsPrivate && req.UserId != pst.CreatorId {
		return nil, status.Errorf(codes.PermissionDenied, "%s", ErrForbiddenToWatch)
	}

	postViewedEvent := &stats.PostViewed{
		PostId:    req.PostId,
		UserId:    req.UserId,
		Timestamp: time.Now().Unix(),
	}

	err = cls.pb.PublishPostViewed(postViewedEvent)
	if err != nil {
		log.Printf("Failed to publish PostViewed event: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to register view: %v", err)
	}

	return &post.ViewPostResponse{Success: true}, nil
}

func (ws *WallService) LikePost(ctx context.Context, req *post.LikePostRequest) (*post.LikePostResponse, error) {
	pst, err := cls.db.getPostByID(req.PostId)
	if err != nil {
		return nil, err
	}

	if pst.IsPrivate && req.UserId != pst.CreatorId {
		return nil, status.Errorf(codes.PermissionDenied, "%s", ErrForbiddenToWatch)
	}

	postLikedEvent := &stats.PostLiked{
		PostId:    req.PostId,
		UserId:    req.UserId,
		Timestamp: time.Now().Unix(),
	}

	err = cls.pb.PublishPostLiked(postLikedEvent)
	if err != nil {
		log.Printf("Failed to publish PostLiked event: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to register like: %v", err)
	}

	return &post.LikePostResponse{Success: true}, nil
}

func (ws *WallService) CreateComment(ctx context.Context, req *post.CreateCommentRequest) (*post.CreateCommentResponse, error) {
	pst, err := cls.db.getPostByID(req.PostId)
	if err != nil {
		return nil, err
	}

	if pst.IsPrivate && req.UserId != pst.CreatorId {
		return nil, status.Errorf(codes.PermissionDenied, "%s", ErrForbiddenToWatch)
	}

	now := time.Now()
	timestamp := now.Format(time.RFC3339)
	comment := &post.Comment{
		CommentId: uuid.New().String(),
		PostId:    req.PostId,
		CreatorId: req.UserId,
		Text:      req.Text,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}

	err = cls.db.insertComment(comment)
	if err != nil {
		return nil, err
	}

	commentCreatedEvent := &stats.PostCommented{
		PostId:    req.PostId,
		UserId:    req.UserId,
		Timestamp: now.Unix(),
	}

	err = cls.pb.PublishCommentCreated(commentCreatedEvent)
	if err != nil {
		log.Printf("Failed to publish CommentCreated event: %v", err)
	}

	return &post.CreateCommentResponse{Comment: comment}, nil
}

func (ws *WallService) ListComments(ctx context.Context, req *post.ListCommentsRequest) (*post.ListCommentsResponse, error) {
	pst, err := cls.db.getPostByID(req.PostId)
	if err != nil {
		return nil, err
	}

	if pst.IsPrivate && req.UserId != pst.CreatorId {
		return nil, status.Errorf(codes.PermissionDenied, "%s", ErrForbiddenToWatch)
	}

	offset := (req.PageNumber - 1) * req.PageSize

	comments, err := cls.db.listComments(req.PostId, int(req.PageSize), int(offset))
	if err != nil {
		return nil, err
	}

	return &post.ListCommentsResponse{
		Comments:   comments,
		TotalCount: int32(len(comments)),
	}, nil
}
