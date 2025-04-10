package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"messenger/api-gateway/handlers"
	"messenger/user-service/utils"
)

func main() {
	publicKeyPath := flag.String("public", "", "path to JWT public key file")
	flag.Parse()
	if publicKeyPath == nil || *publicKeyPath == "" {
		log.Fatal("Not get public key")
	}
	userServiceURL := os.Getenv("USER_SERVICE_URL")
	if userServiceURL == "" {
		log.Fatal("Not get userServiceURL env variable")
	}
	wallServiceURL := os.Getenv("WALL_SERVICE_URL")
	if wallServiceURL == "" {
		log.Fatal("Not get wallServiceURL env variable")
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("Not get SERVER_PORT env variable")
	}
	var authConf utils.TAuthConfig
	authConf.JwtPublicStr = *publicKeyPath

	defer handlers.CloseClients()
	err := handlers.InitClients(authConf, userServiceURL, wallServiceURL)
	if err != nil {
		log.Fatal("Failed to start grpc clients: ", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/users/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/users/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/users/profile", handlers.GetProfileHandler).Methods("GET")
	r.HandleFunc("/users/profile", handlers.PutProfileHandler).Methods("PUT")

	r.HandleFunc("/posts/create", handlers.CreatePostHandler).Methods("POST")
	r.HandleFunc("/posts/{id}", handlers.GetPostHandler).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.UpdatePostHandler).Methods("PUT")
	r.HandleFunc("/posts/{id}", handlers.DeletePostHandler).Methods("DELETE")
	r.HandleFunc("/posts", handlers.ListPostsHandler).Methods("GET")

	log.Printf("API listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
