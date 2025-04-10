package utils

import (
	"database/sql"

	post "messenger/wall-service/post"

	_ "github.com/lib/pq"
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

type TClients struct {
	db *TDatabase
}

type WallService struct {
	post.UnimplementedWallServiceServer
}

var cls *TClients

func NewClients(dbConf *TDBConfig) error {
	cls = &TClients{}
	var err error
	cls.db, err = NewDatabase(dbConf)
	if err != nil {
		return err
	}
	return nil
}
