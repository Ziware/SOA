package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"messenger/user-service/utils"

	user "messenger/user-service/user"

	"google.golang.org/grpc"
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
	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		log.Fatal("Not get REDIS_PORT env variable")
	}
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		log.Fatal("Not get KAFKA_BROKER env variable")
	}

	var authConf utils.TAuthConfig
	authConf.JwtPrivateStr = *privateKeyPath
	authConf.JwtPublicStr = *publicKeyPath

	var dbConf utils.TDBConfig
	dbConf.Port = redisPort
	dbConf.Host = "user-service-redis"

	var kafkaConf utils.TKafkaConfig
	kafkaConf.Timeout = time.Second * 3
	kafkaConf.Brokers = []string{kafkaBroker}

	err = utils.NewClients(authConf, dbConf, &kafkaConf)
	if err != nil {
		log.Fatal(err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserProfileServiceServer(s, &utils.UserService{})

	log.Printf("Staring main user server on port %d", port)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err.Error())
	}
}
