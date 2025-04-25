package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	stats "messenger/stats-service/stats"
	user "messenger/user-service/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Read private + public keys
func NewAuthClient(jwtPrivatePath string, jwtPublicPath string) (*TAuthClient, error) {
	var auth TAuthClient
	private, err := os.ReadFile(jwtPrivatePath)
	if err != nil {
		return nil, err
	}
	public, err := os.ReadFile(jwtPublicPath)
	if err != nil {
		return nil, err
	}
	jwtPrivate, err := jwt.ParseRSAPrivateKeyFromPEM(private)
	if err != nil {
		return nil, err
	}
	auth.JwtPrivate = jwtPrivate
	jwtPublic, err := jwt.ParseRSAPublicKeyFromPEM(public)
	if err != nil {
		return nil, err
	}
	auth.JwtPublic = jwtPublic
	log.Println("Successfull init of auth client")
	return &auth, nil
}

func NewValidateClient(jwtPublicPath string) (*TAuthClient, error) {
	var auth TAuthClient
	public, err := os.ReadFile(jwtPublicPath)
	if err != nil {
		return nil, err
	}
	jwtPublic, err := jwt.ParseRSAPublicKeyFromPEM(public)
	if err != nil {
		return nil, err
	}
	auth.JwtPublic = jwtPublic
	log.Println("Successfull init of validate client")
	return &auth, nil
}

func (auth *TAuthClient) GetField(req *http.Request, field string) (string, error) {
	cookie, err := req.Cookie("jwt")
	if err != nil {
		if err == http.ErrNoCookie {
			return "", fmt.Errorf("no cookie provided: %v", err.Error())
		}
		return "", fmt.Errorf("some jwt cookie error when trying to get it: %v", err.Error())
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return auth.JwtPublic, nil
	})
	if err != nil {
		return "", fmt.Errorf("error on token parsing: %v", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error in claims")
	}
	if expStr, ok := claims["expires"]; ok {
		parsedTime, err := time.Parse(time.RFC3339, expStr.(string))
		if err != nil {
			return "", fmt.Errorf("cant parse time in token")
		}
		if time.Now().After(parsedTime) {
			return "", fmt.Errorf("your cookie is expired")
		}
	} else {
		return "", fmt.Errorf("not found expires key in token")
	}
	if value, ok := claims[field]; ok {
		return value.(string), nil
	}
	return "", fmt.Errorf("no such key in token: %v", field)
}

func (auth *TAuthClient) GetLogin(req *http.Request) (string, error) {
	return auth.GetField(req, "login")
}

func (auth *TAuthClient) GetUserId(req *http.Request) (string, error) {
	return auth.GetField(req, "user_id")
}

func (auth *TAuthClient) Validate(req *http.Request) error {
	_, err := auth.GetLogin(req)
	if err != nil {
		return err
	}
	_, err = auth.GetUserId(req)
	if err != nil {
		return err
	}
	return nil
}

func (auth *TAuthClient) GetToken(login string, user_id string) (string, error) {
	if auth.JwtPrivate == nil {
		return "", fmt.Errorf("your client can only validate tokens")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"login":   login,
		"user_id": user_id,
		"expires": time.Now().Add(time.Hour * 24).Format(time.RFC3339),
	})
	return token.SignedString(auth.JwtPrivate)
}

func (auth *TAuthClient) SetMeta(login string, userId string, ctx context.Context) error {
	tokenString, err := auth.GetToken(login, userId)
	if err != nil {
		return fmt.Errorf("failed to sign the jwt token: %v", err.Error())
	}
	md := metadata.Pairs(
		"cookie", tokenString,
		"expires", time.Now().Add(time.Hour*24).Format(time.RFC3339),
	)
	log.Printf("Sending meta data")
	grpc.SendHeader(ctx, md)

	return nil
}

func (auth *TAuthClient) SetCookie(md metadata.MD, writer http.ResponseWriter) error {
	token, exists := md["cookie"]
	if !exists {
		return fmt.Errorf("cookie key not exists")
	}
	expires, exists := md["expires"]
	if !exists {
		return fmt.Errorf("expires key not exists")
	}
	exp, err := time.Parse(time.RFC3339, expires[0])
	if err != nil {
		return err
	}
	http.SetCookie(writer, &http.Cookie{
		Name:    "jwt",
		Value:   token[0],
		Path:    "/",
		Expires: exp,
	})
	return nil
}

func (us *UserService) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	log.Printf("Register user")
	log.Printf("%s %s %s", req.Email, req.Login, req.Password)
	if req.Login == "" || req.Email == "" || req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "check your query, there 3 required fields: login, email, password")
	}
	exists, err := cls.database.CheckUserExistsByLogin(req.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	if exists {
		return nil, status.Errorf(codes.AlreadyExists, "user with this login actually exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	log.Printf("UserData: login - %v, email - %v", req.Login, req.Email)
	var usr user.User
	usr.UserId = uuid.New().String()
	usr.Login = req.Login
	usr.Email = req.Email
	usr.PasswordHash = string(hashedPassword)
	err = cls.database.CreateUser(&usr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	err = cls.authClient.SetMeta(usr.Login, usr.UserId, ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	userCreatedEvent := &stats.UserCreated{
		UserId:    usr.UserId,
		Timestamp: time.Now().Unix(),
	}
	err = cls.pb.PublishUserCreated(userCreatedEvent)
	if err != nil {
		log.Printf("Failed to publish CommentCreated event: %v", err)
	}
	resp := user.RegisterResponse{
		User: &usr,
	}
	return &resp, nil
}

func (us *UserService) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	log.Printf("Login user")
	if req.Login == "" || req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "check your query, there 2 required fields: login, password")
	}
	exists, err := cls.database.CheckUserExistsByLogin(req.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	if !exists {
		return nil, status.Errorf(codes.NotFound, "user with this login doesn't exists")
	}
	usr, err := cls.database.GetUserByLogin(req.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	log.Printf("UserData: login - %v", req.Login)
	err = bcrypt.CompareHashAndPassword([]byte(usr.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "your password is incorrect")
	}
	err = cls.authClient.SetMeta(usr.Login, usr.UserId, ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	resp := user.LoginResponse{
		User: usr,
	}
	return &resp, nil
}
