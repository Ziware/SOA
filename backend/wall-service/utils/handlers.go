package utils

import (
	"context"
	"errors"
	"log"
	"time"

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
