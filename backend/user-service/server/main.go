package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"user-service/server/utils"

	"github.com/gorilla/mux"
)

func main() {
	privateKeyPath := flag.String("private", "", "path to JWT private key file")
	publicKeyPath := flag.String("public", "", "path to JWT public key file")
	flag.Parse()
	if privateKeyPath == nil || *privateKeyPath == "" {
		log.Fatal("Not get private key")
	}
	if publicKeyPath == nil || *publicKeyPath == "" {
		log.Fatal("Not get public key")
	}
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal("Not get SERVER_PORT env variable")
	}
	posgresPort, err := strconv.Atoi(os.Getenv("POSGRES_PORT"))
	if err != nil {
		log.Fatal("Not get POSGRES_PORT env variable")
	}
	var authConf utils.TAuthConfig
	authConf.JwtPrivateStr = *privateKeyPath
	authConf.JwtPublicStr = *publicKeyPath
	var dbConf utils.TDBConfig
	dbConf.Port = posgresPort
	dbConf.Host = "user-service-postgres"
	dbConf.DBName = "userdb"
	dbConf.User = "userservice"
	dbConf.Password = "userservicepassword"
	dbConf.SSLMode = "disable"
	err = utils.NewContext(authConf, dbConf)
	if err != nil {
		log.Fatal(err.Error())
	}
	router := mux.NewRouter()
	router.HandleFunc("/users/register", utils.RegisterHandler).Methods("POST")
	router.HandleFunc("/users/login", utils.LoginHandler).Methods("POST")
	router.HandleFunc("/users/profile", utils.GetProfileHandler).Methods("GET")
	router.HandleFunc("/users/profile", utils.PutProfileHandler).Methods("PUT")

	log.Printf("Staring main user server on port %d", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
