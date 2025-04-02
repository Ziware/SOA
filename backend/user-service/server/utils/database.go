package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	user "messenger/user-service/server/user"

	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrDuplicateKey = errors.New("user login or email mistmatch")
	ErrInvalidData  = errors.New("incorrect user data")
)

func NewDatabase(config TDBConfig) (*TDatabase, error) {
	ctx := context.Background()
	db := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
	})
	if err := db.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	log.Println("successfull init of redis")
	return &TDatabase{db}, nil
}

func (r *TDatabase) Close() error {
	return r.db.Close()
}

func (r *TDatabase) GetUserByID(id string) (*user.User, error) {
	ctx := context.Background()
	ustr, err := r.db.Get(ctx, fmt.Sprintf("profile:%s", id)).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get user data: %e", err)
	}
	var usr user.User
	if err := proto.Unmarshal([]byte(ustr), &usr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %e", err)
	}

	return &usr, nil
}

func (r *TDatabase) GetUserByLogin(login string) (*user.User, error) {
	ctx := context.Background()
	user_id, err := r.db.Get(ctx, fmt.Sprintf("user_id:%s", login)).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get user data: %w", err)
	}
	return r.GetUserByID(user_id)
}

func (r *TDatabase) CreateUser(user *user.User) error {
	if user.Login == "" || user.Email == "" || user.PasswordHash == "" {
		return ErrInvalidData
	}

	ctx := context.Background()

	user.CreatedAt = time.Now().Format(time.RFC3339)
	user.UpdatedAt = time.Now().Format(time.RFC3339)
	data, err := proto.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %w", err)
	}

	err = r.db.Set(ctx, fmt.Sprintf("profile:%s", user.UserId), data, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set user data: %w", err)
	}

	err = r.db.Set(ctx, fmt.Sprintf("user_id:%s", user.Login), user.UserId, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set user data by login: %w", err)
	}

	return nil
}

func (r *TDatabase) UpdateUser(user *user.User) error {
	ctx := context.Background()

	user.UpdatedAt = time.Now().Format(time.RFC3339)
	data, err := proto.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %w", err)
	}

	exists, err := r.CheckUserExistsById(user.UserId)
	if err != nil {
		return err
	} else if !exists {
		return ErrUserNotFound
	}

	err = r.db.Set(ctx, fmt.Sprintf("profile:%s", user.UserId), data, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set user data: %w", err)
	}

	return nil
}

func (r *TDatabase) DeleteUser(id string) error {
	ctx := context.Background()

	userStr, err := r.db.Get(ctx, fmt.Sprintf("profile:%s", id)).Result()
	if err == redis.Nil {
		return ErrUserNotFound
	} else if err != nil {
		return fmt.Errorf("failed to get user data: %w", err)
	}

	var user user.User
	if err := proto.Unmarshal([]byte(userStr), &user); err != nil {
		return fmt.Errorf("failed to unmarshal user data: %w", err)
	}

	err = r.db.Del(ctx, fmt.Sprintf("profile:%s", id)).Err()
	if err != nil {
		return fmt.Errorf("failed to delete user data by id: %w", err)
	}

	err = r.db.Del(ctx, fmt.Sprintf("user_id:%s", user.Login)).Err()
	if err != nil {
		return fmt.Errorf("failed to delete user data by login: %w", err)
	}

	return nil
}

func (r *TDatabase) CheckUserExistsById(id string) (bool, error) {
	ctx := context.Background()
	exists, err := r.db.Exists(ctx, fmt.Sprintf("profile:%s", id)).Result()
	if err != nil {
		return false, fmt.Errorf("error on user existence check by id: %w", err)
	}

	return exists > 0, nil
}

func (r *TDatabase) CheckUserExistsByLogin(login string) (bool, error) {
	ctx := context.Background()
	exists, err := r.db.Exists(ctx, fmt.Sprintf("user_id:%s", login)).Result()
	if err != nil {
		return false, fmt.Errorf("error on user existence check by login: %w", err)
	}

	return exists > 0, nil
}
