package utils

import (
	"database/sql"
	"time"

	post "messenger/wall-service/post"

	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
)

type TDBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	SSLMode  string
}

type TDatabase struct {
	db *sql.DB
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
	db *TDatabase
	pb *TPublisher
}

type WallService struct {
	post.UnimplementedWallServiceServer
}

var cls *TClients

func NewClients(dbConf *TDBConfig, kafkaConfig *TKafkaConfig) error {
	cls = &TClients{}
	var err error
	cls.db, err = NewDatabase(dbConf)
	if err != nil {
		return err
	}
	cls.pb, err = NewPublisher(kafkaConfig)
	if err != nil {
		return err
	}
	return nil
}
