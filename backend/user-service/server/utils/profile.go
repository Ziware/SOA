package utils

import (
	"context"
	"log"

	user "messenger/user-service/server/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (us *UserService) GetProfile(ctx context.Context, req *user.GetProfileRequest) (*user.GetProfileResponse, error) {
	log.Printf("Get profile")
	usr, err := cls.database.GetUserByID(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	resp := user.GetProfileResponse{
		User: usr,
	}
	return &resp, nil
}

func parsePutProfileRequest(user *user.User, newParams *user.User) {
	if newParams.BirthDate != "" {
		user.BirthDate = newParams.BirthDate
	}
	if newParams.PhoneNumber != "" {
		user.PhoneNumber = newParams.PhoneNumber
	}
	if newParams.Email != "" {
		user.Email = newParams.Email
	}
	if newParams.Name != "" {
		user.Name = newParams.Name
	}
	if newParams.Surname != "" {
		user.Surname = newParams.Surname
	}
}

func (us *UserService) PutProfile(ctx context.Context, req *user.PutProfileRequest) (*user.PutProfileResponse, error) {
	log.Printf("Change profile")
	usr, err := cls.database.GetUserByID(req.User.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	parsePutProfileRequest(usr, req.User)
	err = cls.database.UpdateUser(usr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	resp := user.PutProfileResponse{
		User: usr,
	}
	return &resp, nil
}
