package utils

import (
	"crypto/rsa"
	"time"

	user "messenger/user-service/user"

	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
)

//// Configs + Clients

type TAuthConfig struct {
	JwtPrivateStr string
	JwtPublicStr  string
}

type TAuthClient struct {
	JwtPrivate *rsa.PrivateKey
	JwtPublic  *rsa.PublicKey
}

type TDBConfig struct {
	Host string
	Port int
}

type TDatabase struct {
	db *redis.Client
}

type TKafkaConfig struct {
	Brokers []string
	Timeout time.Duration
}

type TPublisher struct {
	writers map[string]*kafka.Writer
	timeout time.Duration
}

type TClients struct {
	authClient *TAuthClient
	database   *TDatabase
	pb         *TPublisher
}

type UserService struct {
	user.UnimplementedUserProfileServiceServer
}

var cls *TClients

func NewClients(authConf TAuthConfig, dbConf TDBConfig, kafkaConf *TKafkaConfig) error {
	cls = &TClients{}
	var err error
	cls.authClient, err = NewAuthClient(authConf.JwtPrivateStr, authConf.JwtPublicStr)
	if err != nil {
		return err
	}
	cls.database, err = NewDatabase(dbConf)
	if err != nil {
		return err
	}
	cls.pb, err = NewPublisher(kafkaConf)
	if err != nil {
		return err
	}
	return nil
}
