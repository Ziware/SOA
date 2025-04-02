package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	user "messenger/user-service/server/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req user.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	var headers metadata.MD
	resp, err := UserServiceClient.Register(ctx, &req, grpc.Header(&headers))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Get headers, len: %d", headers.Len())
	for key, value := range headers {
		log.Printf("headers key:val = %s:%s", key, value[0])
	}
	err = ValidateClient.SetCookie(headers, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req user.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	var headers metadata.MD
	resp, err := UserServiceClient.Login(ctx, &req, grpc.Header(&headers))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Get headers, len: %d", headers.Len())
	for key, value := range headers {
		log.Printf("headers key:val = %s:%s", key, value[0])
	}
	err = ValidateClient.SetCookie(headers, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	err := ValidateClient.Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userID, err := ValidateClient.GetUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	req := user.GetProfileRequest{UserId: userID}
	resp, err := UserServiceClient.GetProfile(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func PutProfileHandler(w http.ResponseWriter, r *http.Request) {
	err := ValidateClient.Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userID, err := ValidateClient.GetUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req user.PutProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req.User); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req.User.UserId = userID

	resp, err := UserServiceClient.PutProfile(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
