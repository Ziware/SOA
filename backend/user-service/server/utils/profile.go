package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func GetProfileHandler(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Get profile")
	err := ctx.authClient.Validate(req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}
	userIdStr, err := ctx.authClient.GetUserId(req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := ctx.database.GetUserByID(userId)
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

func ParsePutProfileRequest(user *TUser, newParams *TPutProfileRequest) {
	if newParams.BirthDate != "" {
		user.BirthDate = newParams.BirthDate
	}
	if newParams.PhoneNumber != "" {
		user.PhoneNumber = newParams.PhoneNumber
	}
	if newParams.Email != "" {
		user.Email = newParams.PhoneNumber
	}
	if newParams.Name != "" {
		user.Name = newParams.Name
	}
	if newParams.Surname != "" {
		user.Surname = newParams.Surname
	}
}

func PutProfileHandler(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Change profile")
	var newProfileData TPutProfileRequest
	err := json.NewDecoder(req.Body).Decode(&newProfileData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = ctx.authClient.Validate(req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}
	userIdStr, err := ctx.authClient.GetUserId(req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := ctx.database.GetUserByID(userId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	ParsePutProfileRequest(user, &newProfileData)
	err = ctx.database.UpdateUser(user)
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
