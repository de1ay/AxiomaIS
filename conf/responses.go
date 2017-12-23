package conf

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type ApiResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message interface{} `json:"message"`
}

func (apiResponse *ApiResponse) toJSON() string {
	resp, _ := json.Marshal(apiResponse)
	return string(resp)
}

func (apiResponse *ApiResponse) Execute(responseWriter http.ResponseWriter) {
	fmt.Fprintf(responseWriter, apiResponse.toJSON())
}

// 100
var ERROR_DATABASE_CONNECTION_CREATION_100 = &ApiResponse{100, "error", "Ошибка соединения с базой данных"}
var ERROR_DATABASE_REQUEST_INVALID_101 = &ApiResponse{101, "error", "Ошибка соединения с базой данных"}

// 200
var REQUEST_SUCCESS_200 = &ApiResponse{200, "success", "Запрос успешно выполнен"}

// 300
var ERROR_UNAUTHORIZED_300 = &ApiResponse{300, "error", "Необходима авторизация"}
var ERROR_LOGIN_301 = &ApiResponse{301, "error", "Неверный логин или пароль"}
var ERROR_LOGIN_SHORT_302 = &ApiResponse{302, "error", "Логин не может быть короче четырёх символов"}
var ERROR_PASSWORD_SHORT_303 = &ApiResponse{303, "error", "Пароль не может быть короче шести символов"}

// 400
var ERROR_ACCESS_400 = &ApiResponse{400, "error", "Недостаточно прав"}
var ERROR_TOKEN_INVALID_401 = &ApiResponse{401, "error", "Неверный токен"}
var ERROR_METHOD_NOT_ALLOWED_402 = &ApiResponse{402, "error", "Метод запрещён"}