package utils

import (
	"database/sql"
	"fmt"
	"log"

	post "messenger/wall-service/post"

	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewDatabase(cfg *TDBConfig) (*TDatabase, error) {
	var db TDatabase
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	log.Printf("trying to connect to postgres with: %s", dsn)
	var err error
	db.db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &db, nil
}

func (r *TDatabase) insertPost(post *post.Post) error {
	query := `
		INSERT INTO posts (post_id, title, description, creator_id, created_at, updated_at, is_private, tags)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.Exec(query, post.PostId, post.Title, post.Description, post.CreatorId, post.CreatedAt, post.UpdatedAt, post.IsPrivate, pq.Array(post.Tags))
	if err != nil {
		return status.Errorf(codes.Internal, "%s", err.Error())
	}
	return nil
}

func (r *TDatabase) postExists(postID string) (bool, error) {
	query := `SELECT exists (SELECT 1 FROM posts WHERE post_id = $1)`
	var exists bool
	err := r.db.QueryRow(query, postID).Scan(&exists)
	if err != nil {
		return false, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return exists, nil
}

func (r *TDatabase) getPostByID(postID string) (*post.Post, error) {
	exists, err := r.postExists(postID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Could not found post with id: %s", postID)
	}

	query := `
		SELECT post_id, title, description, creator_id, created_at, updated_at, is_private, tags
		FROM posts
		WHERE post_id = $1
	`
	var post post.Post
	row := r.db.QueryRow(query, postID)
	err = row.Scan(&post.PostId, &post.Title, &post.Description, &post.CreatorId, &post.CreatedAt, &post.UpdatedAt, &post.IsPrivate, pq.Array(&post.Tags))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return &post, nil
}

func (r *TDatabase) updatePost(post *post.Post) error {
	query := `
		UPDATE posts
		SET title = $2, description = $3, creator_id = $4, updated_at = $5, is_private = $6, tags = $7
		WHERE post_id = $1
	`
	_, err := r.db.Exec(query, post.PostId, post.Title, post.Description, post.CreatorId, post.UpdatedAt, post.IsPrivate, pq.Array(post.Tags))
	return err
}

func (r *TDatabase) deletePost(postID string) error {
	exists, err := r.postExists(postID)
	if err != nil {
		return status.Errorf(codes.Internal, "%s", err.Error())
	}
	if !exists {
		return status.Errorf(codes.NotFound, "Could not found post with id: %s", postID)
	}
	query := `DELETE FROM posts WHERE post_id = $1`
	_, err = r.db.Exec(query, postID)
	if err != nil {
		return status.Errorf(codes.Internal, "%s", err.Error())
	}
	return nil
}

func (r *TDatabase) listPosts(limit, offset int, requestor_id string) ([]*post.Post, error) {
	query := `
		SELECT post_id, title, description, creator_id, created_at, updated_at, is_private, tags
		FROM posts
		WHERE is_private = false OR creator_id = $3
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(query, limit, offset, requestor_id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	defer rows.Close()

	var posts []*post.Post
	for rows.Next() {
		var post post.Post
		err := rows.Scan(&post.PostId, &post.Title, &post.Description, &post.CreatorId, &post.CreatedAt, &post.UpdatedAt, &post.IsPrivate, pq.Array(&post.Tags))
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *TDatabase) insertComment(comment *post.Comment) error {
	query := `
		INSERT INTO comments (comment_id, post_id, creator_id, text, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(
		query,
		comment.CommentId,
		comment.PostId,
		comment.CreatorId,
		comment.Text,
		comment.CreatedAt,
		comment.UpdatedAt,
	)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to insert comment: %s", err.Error())
	}

	return nil
}

func (r *TDatabase) listComments(postID string, limit, offset int) ([]*post.Comment, error) {
	query := `
		SELECT comment_id, post_id, creator_id, text, created_at, updated_at
		FROM comments
		WHERE post_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.Query(query, postID, limit, offset)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to query comments: %s", err.Error())
	}
	defer rows.Close()

	var comments []*post.Comment

	for rows.Next() {
		var comment post.Comment
		err := rows.Scan(
			&comment.CommentId,
			&comment.PostId,
			&comment.CreatorId,
			&comment.Text,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to scan comment: %s", err.Error())
		}

		comments = append(comments, &comment)
	}

	if err = rows.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "Error iterating comments: %s", err.Error())
	}

	return comments, nil
}
