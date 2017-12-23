package authorization

import (
	"net/http"
	"strings"
	"github.com/gorilla/mux"
	"axioma/src/api/authorization"
	"axioma/src/handlers"
	"axioma/conf"
)

func getAuthorizationData(r *http.Request) authorization.AuthData{
	userLogin := strings.TrimSpace(r.FormValue("login"))
	userPassword := strings.TrimSpace(r.FormValue("password"))
	return authorization.AuthData{userLogin, userPassword, -1}
}

func loginAndPasswordAuthHandler(responseWriter http.ResponseWriter, r *http.Request){
	handlers.SetHeaders_API_GET(responseWriter)
	if r.Method == http.MethodGet {
		responseWriter.WriteHeader(http.StatusOK)
		authData := getAuthorizationData(r)
		authData.Authorize(responseWriter)
	} else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		conf.ERROR_METHOD_NOT_ALLOWED_402.Execute(responseWriter)
	}
}

func HandleAuthorization(router *mux.Router)  {
	router.HandleFunc("/token.get", loginAndPasswordAuthHandler)
}
