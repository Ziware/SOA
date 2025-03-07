package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"login":   login,
		"user_id": user_id,
		"expires": time.Now().Add(time.Hour * 24).Format(time.RFC3339),
	})
	return token.SignedString(auth.JwtPrivate)
}

func (auth *TAuthClient) SetCookie(login string, userId string, writer http.ResponseWriter) error {
	tokenString, err := auth.GetToken(login, userId)
	if err != nil {
		return fmt.Errorf("failed to sign the jwt token: %v", err.Error())
	}
	http.SetCookie(writer, &http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24),
	})
	return nil
}

func RegisterHandler(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Register user")
	var regData TRegisterRequest
	err := json.NewDecoder(req.Body).Decode(&regData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if regData.Login == "" || regData.Email == "" || regData.Password == "" {
		http.Error(writer, "check your query, there 3 required fields: login, email, password", http.StatusBadRequest)
		return
	}
	exists, err := ctx.database.CheckUserExistsByLogin(regData.Login)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(writer, "user with this login actually exists", http.StatusConflict)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regData.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("UserData: login - %v, email - %v", regData.Login, regData.Email)
	var user TUser
	user.UserId = uuid.New()
	user.Login = regData.Login
	user.Email = regData.Email
	user.PassHash = string(hashedPassword)
	user.CreatedAt = time.Now()
	err = ctx.database.CreateUser(&user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = ctx.authClient.SetCookie(user.Login, user.UserId.String(), writer)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	var resp TRegisterResponse
	resp.Message = "Successful register"
	resp.UserId = user.UserId.String()
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Write(data)
}

func LoginHandler(writer http.ResponseWriter, req *http.Request) {
	log.Printf("Login user")
	var loginData TLoginRequest
	err := json.NewDecoder(req.Body).Decode(&loginData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if loginData.Login == "" || loginData.Password == "" {
		http.Error(writer, "check your query, there 2 required fields: login, password", http.StatusBadRequest)
		return
	}
	exists, err := ctx.database.CheckUserExistsByLogin(loginData.Login)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(writer, "user with this login doesn't exists", http.StatusNotFound)
		return
	}
	user, err := ctx.database.GetUserByLogin(loginData.Login)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("UserData: login - %v", loginData.Login)
	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(loginData.Password))
	if err != nil {
		http.Error(writer, "your password is incorrect", http.StatusUnauthorized)
		return
	}
	err = ctx.authClient.SetCookie(user.Login, user.UserId.String(), writer)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	var resp TLoginResponse
	resp.Message = "Successful login"
	resp.User = *user
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Write(data)
}
