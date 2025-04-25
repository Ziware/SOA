package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"messenger/errors"
	user "messenger/user-service/user"
	post "messenger/wall-service/post"

	"github.com/gorilla/mux"
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
		log.Printf("%s", err.Error())
		errors.SendHttpError(w, err)
		return
	}
	err = ValidateClient.SetCookie(headers, w)
	if err != nil {
		log.Printf("%s", err.Error())
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
		log.Printf("%s", err.Error())
		errors.SendHttpError(w, err)
		return
	}
	err = ValidateClient.SetCookie(headers, w)
	if err != nil {
		log.Printf("%s", err.Error())
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
		errors.SendHttpError(w, err)
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
		errors.SendHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
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
	var pst post.Post
	err = json.NewDecoder(r.Body).Decode(&pst)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	pst.CreatorId = userID

	resp, err := WallServiceClient.CreatePost(context.Background(), &post.CreatePostRequest{Post: &pst})
	if err != nil {
		errors.SendHttpError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp.Post)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	id := vars["id"]

	resp, err := WallServiceClient.GetPost(context.Background(), &post.GetPostRequest{Id: id, UserId: userID})
	if err != nil {
		errors.SendHttpError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp.Post)
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	id := vars["id"]

	var pst post.Post
	err = json.NewDecoder(r.Body).Decode(&pst)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	pst.PostId = id

	resp, err := WallServiceClient.UpdatePost(context.Background(), &post.UpdatePostRequest{Post: &pst, UserId: userID})
	if err != nil {
		errors.SendHttpError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp.Post)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	id := vars["id"]

	_, err = WallServiceClient.DeletePost(context.Background(), &post.DeletePostRequest{Id: id, UserId: userID})
	if err != nil {
		log.Printf("%s", err.Error())
		errors.SendHttpError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ListPostsHandler(w http.ResponseWriter, r *http.Request) {
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

	pageNumber := 1
	pageSize := 10

	queryParams := r.URL.Query()

	if pageNumStr := queryParams.Get("page_number"); pageNumStr != "" {
		if num, err := strconv.Atoi(pageNumStr); err == nil && num > 0 {
			pageNumber = num
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if num <= 0 {
			http.Error(w, "page_number should be above zero", http.StatusBadRequest)
		}
	}

	if pageSizeStr := queryParams.Get("page_size"); pageSizeStr != "" {
		if size, err := strconv.Atoi(pageSizeStr); err == nil && size > 0 {
			pageSize = size
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if size <= 0 {
			http.Error(w, "page_size should be above zero", http.StatusBadRequest)
		}
	}

	resp, err := WallServiceClient.ListPosts(context.Background(), &post.ListPostsRequest{PageNumber: int32(pageNumber), PageSize: int32(pageSize), UserId: userID})
	if err != nil {
		errors.SendHttpError(w, err)
		return
	}

	json.NewEncoder(w).Encode(resp.Posts)
}

func ViewPostHandler(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	postID := vars["id"]

	req := &post.ViewPostRequest{
		PostId: postID,
		UserId: userID,
	}

	resp, err := WallServiceClient.ViewPost(context.Background(), req)
	if err != nil {
		log.Printf("Error viewing post: %s", err.Error())
		errors.SendHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	postID := vars["id"]

	req := &post.LikePostRequest{
		PostId: postID,
		UserId: userID,
	}

	resp, err := WallServiceClient.LikePost(context.Background(), req)
	if err != nil {
		log.Printf("Error liking post: %s", err.Error())
		errors.SendHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	postID := vars["id"]

	var requestBody struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req := &post.CreateCommentRequest{
		PostId: postID,
		UserId: userID,
		Text:   requestBody.Text,
	}

	resp, err := WallServiceClient.CreateComment(context.Background(), req)
	if err != nil {
		log.Printf("Error creating comment: %s", err.Error())
		errors.SendHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp.Comment)
}

func ListCommentsHandler(w http.ResponseWriter, r *http.Request) {
	err := ValidateClient.Validate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	postID := vars["id"]

	pageNumber := 1
	pageSize := 10

	queryParams := r.URL.Query()

	if pageNumStr := queryParams.Get("page_number"); pageNumStr != "" {
		if num, err := strconv.Atoi(pageNumStr); err == nil && num > 0 {
			pageNumber = num
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if num <= 0 {
			http.Error(w, "page_number should be above zero", http.StatusBadRequest)
			return
		}
	}

	if pageSizeStr := queryParams.Get("page_size"); pageSizeStr != "" {
		if size, err := strconv.Atoi(pageSizeStr); err == nil && size > 0 {
			pageSize = size
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if size <= 0 {
			http.Error(w, "page_size should be above zero", http.StatusBadRequest)
			return
		}
	}

	req := &post.ListCommentsRequest{
		PostId:     postID,
		PageNumber: int32(pageNumber),
		PageSize:   int32(pageSize),
	}

	resp, err := WallServiceClient.ListComments(context.Background(), req)
	if err != nil {
		log.Printf("Error listing comments: %s", err.Error())
		errors.SendHttpError(w, err)
		return
	}

	response := struct {
		Comments   []*post.Comment `json:"comments"`
		TotalCount int32           `json:"total_count"`
	}{
		Comments:   resp.Comments,
		TotalCount: resp.TotalCount,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
