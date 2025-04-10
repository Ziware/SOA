package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	post "messenger/wall-service/post"
	"messenger/wall-service/utils"

	"google.golang.org/grpc"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal("Not get SERVER_PORT env variable")
	}
	postgresPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatal("Not get POSTGRES_PORT env variable")
	}
	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		log.Fatal("Not get POSTGRES_USER env variable")
	}
	postgresPass := os.Getenv("POSTGRES_PASSWORD")
	if postgresPass == "" {
		log.Fatal("Not get POSTGRES_PASSWORD env variable")
	}
	postgresDB := os.Getenv("POSTGRES_DB")
	if postgresDB == "" {
		log.Fatal("Not get POSTGRES_DB env variable")
	}
	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		log.Fatal("Not get POSTGRES_HOST env variable")
	}
	var dbConf utils.TDBConfig
	dbConf.Host = postgresHost
	dbConf.Port = postgresPort
	dbConf.DBName = postgresDB
	dbConf.User = postgresUser
	dbConf.Password = postgresPass
	dbConf.SSLMode = "disable"

	err = utils.NewClients(&dbConf)
	if err != nil {
		log.Fatal(err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	post.RegisterWallServiceServer(s, &utils.WallService{})

	log.Printf("Staring main user server on port %d", port)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err.Error())
	}
}
