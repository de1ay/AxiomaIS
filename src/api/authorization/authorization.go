package authorization

import (
	"math/rand"
	"golang.org/x/crypto/sha3"
	"encoding/base64"
	"axioma/conf"
	"crypto/sha512"
	"strconv"
	"time"
	"axioma/src"
	"net/http"
)

type AuthData struct {
	Login string
	Password string
	UserID int
}

type ApiResponseAuthorizationSuccess struct {
	Token string `json:"token"`
}

func (authData *AuthData) Authorize(responseWriter http.ResponseWriter) {
	response := authData.validate()
	response.Execute(responseWriter)
}

func (authData *AuthData) getPasswordHash() string {
	hasher := sha3.New512()
	passwordBytes := []byte(authData.Password)
	passwordHashBytes := hasher.Sum(passwordBytes)
	return base64.URLEncoding.EncodeToString(passwordHashBytes)
}

func (authData *AuthData) validate() *conf.ApiResponse {
	err := authData.validateData(); if err != nil {
		return err
	}
	err = authData.validatePassword(); if err != nil {
		return err
	}
	return authData.createSession()
}

func (authData *AuthData) validateData() *conf.ApiResponse {
	if len(authData.Login) < conf.LOGIN_MIN_LENGTH { return conf.ERROR_LOGIN_SHORT_302 }
	if len(authData.Password) < conf.PASSWORD_MIN_LENGTH { return conf.ERROR_PASSWORD_SHORT_303 }
	return nil
}

func (authData *AuthData) validatePassword() *conf.ApiResponse {
	userID := -1
	passwordHash := authData.getPasswordHash()
	err := src.Connection.Connection.QueryRow("SELECT id FROM users WHERE login=? AND password=?", authData.Login, passwordHash).Scan(&userID)
	if err != nil || userID == -1  {
		return conf.ERROR_LOGIN_301
	}
	authData.UserID = userID
	return nil
}

func (authData *AuthData) generateToken() string {
	hasher := sha512.New()
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	rand.Seed(time.Now().UnixNano())
	tokenRawBytes := []byte(authData.Login + currentTime + strconv.FormatInt(rand.Int63(), 10))
	tokenBytes := hasher.Sum(tokenRawBytes)
	return base64.URLEncoding.EncodeToString(tokenBytes)
}

func (authData *AuthData) createSession() *conf.ApiResponse {
	token := authData.generateToken()
	validation := authData.validateSessionsOverflow(); if validation != nil {
		return validation
	}
	dbRequest, err := src.Connection.Connection.Prepare("INSERT INTO sessions(user_id, token) VALUES (?, ?)")
	if err != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	_, err = dbRequest.Exec(authData.UserID, token); if err != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return &conf.ApiResponse{200, "success", ApiResponseAuthorizationSuccess{token}}
}

func (authData *AuthData) validateSessionsOverflow() *conf.ApiResponse {
	var count int
	err := src.Connection.Connection.QueryRow("SELECT COUNT(token) as count FROM sessions WHERE user_id=?", authData.UserID).Scan(&count)
	if err != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	if count > 4 {
		return authData.deleteOldestSession()
	} else {
		return nil
	}
}

func (authData *AuthData) deleteOldestSession() *conf.ApiResponse {
	dbRequest, err := src.Connection.Connection.Prepare("DELETE FROM sessions WHERE user_id=? ORDER BY start_time ASC LIMIT 1")
	if err != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	_, err = dbRequest.Exec(authData.UserID); if err != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return nil
}

func ValidateToken(token string) (int, *conf.ApiResponse) {
	userID := -1
	err := src.Connection.Connection.QueryRow("SELECT user_id FROM sessions WHERE token=?", token).Scan(&userID)
	if err != nil {
		return userID, conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	if userID == -1 {
		return userID, conf.ERROR_TOKEN_INVALID_401
	} else {
		return userID, nil
	}
}