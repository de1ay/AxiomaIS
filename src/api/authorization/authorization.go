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
	var user src.User
	passwordHash := GetHash(authData.Password)
	err := src.Connection.Connection.Where("login=? AND password=?", authData.Login, passwordHash).First(&user).Error
	if err != nil  {
		return conf.ERROR_LOGIN_301
	}
	authData.UserID = int(user.ID)
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
	err := src.Connection.Connection.Create(&src.Session{UserID: uint(authData.UserID), Token: token}).Error
	if err != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return &conf.ApiResponse{200, "success", ApiResponseAuthorizationSuccess{token}}
}

func (authData *AuthData) validateSessionsOverflow() *conf.ApiResponse {
	var count int
	err := src.Connection.Connection.Model(&src.Session{}).Where("user_id=?", authData.UserID).Count(&count).Error
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
	err := src.Connection.Connection.Where("user_id=?", authData.UserID).Limit(1).Delete(&src.Session{}).Error
	if err != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return nil
}

func ValidateToken(token string) (src.User, *conf.ApiResponse) {
	var user src.User
	var session src.Session
	err := src.Connection.Connection.Where("token=?", token).First(&session).Error; if err != nil {
		return user, conf.ERROR_TOKEN_INVALID_401
	}
	err = src.Connection.Connection.Model(&session).Related(&user).Error; if err != nil {
		return user, conf.ERROR_TOKEN_INVALID_401
	}
	if user.FranchiseID != 0 {
		err = src.Connection.Connection.Model(&user).Related(user.Franchise).Error;
		if err != nil {
			return user, conf.ERROR_TOKEN_INVALID_401
		}
	}
	if user.ClinicID != 0 {
		err = src.Connection.Connection.Model(&user).Related(user.Clinic).Error;
		if err != nil {
			return user, conf.ERROR_TOKEN_INVALID_401
		}
	}
	return user, nil
}

func GetHash(str string) string {
	hasher := sha3.New512()
	passwordBytes := []byte(str)
	passwordHashBytes := hasher.Sum(passwordBytes)
	return base64.URLEncoding.EncodeToString(passwordHashBytes)
}