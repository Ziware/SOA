package utils

import (
	"crypto/rsa"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TUser struct {
	UserId      uuid.UUID `json:"user_id"`
	Login       string    `json:"login"`
	Email       string    `json:"email"`
	PassHash    string    `json:"password_hash"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	BirthDate   string    `json:"birth_date"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

//// Handlers

type TRegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type TRegisterResponse struct {
	Message string `json:"message"`
	UserId  string `json:"user_id"`
}

type TLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TLoginResponse struct {
	Message string `json:"message"`
	User    TUser  `json:"user"`
}

type TGetProfileRequest struct{}

type TGetProfileResponse struct {
	Message string `json:"message"`
	User    TUser  `json:"user"`
}

type TPutProfileRequest struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	BirthDate   string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
}

type TPutProfileResponse struct {
	Message string `json:"message"`
	User    TUser  `json:"user"`
}

type TErrorResponse struct {
	Error string `json:"error"`
}

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
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type TDatabase struct {
	db *sql.DB
}

type TContext struct {
	authClient *TAuthClient
	database   *TDatabase
}

var ctx *TContext

func NewContext(authConf TAuthConfig, dbConf TDBConfig) error {
	ctx = &TContext{}
	var err error
	ctx.authClient, err = NewAuthClient(authConf.JwtPrivateStr, authConf.JwtPublicStr)
	if err != nil {
		return err
	}
	ctx.database, err = NewDatabase(dbConf)
	if err != nil {
		return err
	}
	return nil
}
