package utils

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	user "user-service/server/user"

	"github.com/google/uuid"
)

func (us *UserService) GetProfile(ctx context.Context, req *user.GetProfileRequest) (*user.GetProfileResponse, error) {
	log.Printf("Get profile")

	user, err := cls.database.GetUserByID(req.UserId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	var resp TGetProfileResponse
	resp.Message = "Successful Profile get"
	resp.User = *user
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(data)
}

func parsePutProfileRequest(user *TUser, newParams *TPutProfileRequest) {
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
	var newProfileData TPutProfileRequest
	err := json.NewDecoder(req.Body).Decode(&newProfileData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = cls.authClient.Validate(req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}
	userIdStr, err := cls.authClient.GetUserId(req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := cls.database.GetUserByID(userId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	parsePutProfileRequest(user, &newProfileData)
	err = cls.database.UpdateUser(user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	var resp TPutProfileResponse
	resp.Message = "Successful profile update"
	resp.User = *user
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(data)
}
